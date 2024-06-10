package tuple_space;

import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.DataInput;
import java.io.DataOutput;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.util.ArrayList;
import java.util.List;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.jgroups.raft.StateMachine;

public class TupleSpace implements StateMachine {
    private static final Logger logger = LogManager.getLogger(TupleSpace.class);
    private List<Object[]> tuples;

    public TupleSpace() {
        this.tuples = new ArrayList<>();
    }

    public void setTuples(List<Object[]> otherTuples) {
        this.tuples = otherTuples;
    }

    public List<Object[]> getTuples() {
        return tuples;
    }

    public byte[] apply(byte[] data, int offset, int length, boolean serialize_response) throws Exception {
        Object result = null;
        try (var byteStream = new ByteArrayInputStream(data, offset, length);
                var objectStream = new ObjectInputStream(byteStream)) {

            Command cmd = (Command) objectStream.readObject();

            logger.debug("Command received: %s", cmd.getHeader().toString());
            result = switch (cmd.getHeader()) {
                case GET -> this.get(cmd.getTuple());
                case GET_ALL -> this.tuples;
                case READ -> this.read(cmd.getTuple());
                case WRITE -> {
                    this.write(cmd.getTuple());
                    yield null;
                }
                default -> null;
            };

            if (!serialize_response)
                return null;

        }
        try (var byteStream = new ByteArrayOutputStream(); var objectStream = new ObjectOutputStream(byteStream)) {
            objectStream.writeObject(result);
            return byteStream.toByteArray();
        }
    }

    public void readContentFrom(DataInput in) throws Exception {

    }

    public void writeContentTo(DataOutput out) throws Exception {

    }

    public void write(Object[] tuple) throws Exception {
        for (int i = 0; i < tuple.length; ++i)
            if (tuple[i].equals('*'))
                throw new Exception("Invalid character.");

        this.tuples.add(tuple);
    }

    public Object[] read(Object[] tuple) {
        // (*, 'Fulano', *)
        for (Object[] i : this.tuples) {
            if (matchWithWildcards(tuple, i)) {
                return i;
            }
        }
        return new Object[] {};
    }

    public Object[] get(Object[] tuple) {
        Object[] res = read(tuple);
        if (res.length > 0) {
            this.tuples.remove(tuple);
        }

        return res;
    }

    private boolean matchWithWildcards(Object[] pattern, Object[] candidate) {
        if (pattern.length != candidate.length)
            return false;
        for (int i = 0; i < pattern.length; i++) {
            if (!(pattern[i].equals(candidate[i]) || pattern[i].equals('*'))) {
                return false;
            }
        }
        return true;
    }
}
