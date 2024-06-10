package tuple_space;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.jgroups.*;
import org.jgroups.util.Util;

import java.io.*;
import java.util.List;
import java.util.Arrays;
import java.util.ArrayList;
import java.util.LinkedList;

public class TupleSpaceServer implements Receiver {
    private TupleSpace tupleSpace;
    private static final Logger logger = LogManager.getLogger(TupleSpaceServer.class);

    public void receive(Message msg) {
        String line = msg.getSrc() + ": " + msg.getObject();
        System.out.println(line);
        synchronized (tupleSpace) {
            try {
                tupleSpace.write(msg.getObject());
                System.out.println(tupleSpace.getTuples().size());
            } catch (Exception e) {
                logger.error("Could not insert tuple into tuple space", e);
            }
        }
    }

    public void write(Object[] tuple) throws Exception {
        tupleSpace.write(tuple);
    }

    public Object[] read(Object[] tuple) {
        return tupleSpace.read(tuple);
    }

    public void get(Object[] tuple) {
        tupleSpace.get(tuple);
    }

    public void getState(OutputStream output) throws Exception {
        synchronized (tupleSpace) {
            Util.objectToStream(tupleSpace.getTuples(), new DataOutputStream(output));
        }
    }

    public void setState(InputStream input) throws Exception {
        List<Object[]> tuples = Util.objectFromStream(new DataInputStream(input));
        synchronized (tupleSpace) {
            tupleSpace.getTuples().clear();
            tupleSpace.getTuples().addAll(tuples);
        }
        System.out.println("received state (" + tuples.size() + " itens in tuple space):");
        tuples.forEach(System.out::println);
    }

    private void start() throws Exception {
        try (var channel = new JChannel()) {
            channel.setReceiver(this);
            channel.connect("TupleSpaceCluster", null, 10000);
            eventLoop(channel);
        }
    }

    private void eventLoop(JChannel channel) {
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        while (true) {
            try {
                System.out.print("> ");
                System.out.flush();
                String line = in.readLine().toLowerCase();
                // if (line.startsWith("quit") || line.startsWith("exit")) {
                // break;
                // }
                // line = "[" + user_name + "] " + line;

                // List<Object[]> myList = Arrays.<Object[]>asList(new Object[]{1,2,3}, new
                // Object[]{"AAA", 1, "BBB"});

                tupleSpace.getTuples().forEach(obj -> {
                    String s = "(";
                    for (var o : obj)
                        s += o.toString() + ", ";
                    s += ")";

                    System.out.println(s);
                });
                // tupleSpace.getTuples().forEach(obj ->
                // System.out.println("(".concat(String.join(", ",
                // Arrays.toString(obj)).concat(")"))));

                Object[] myTuple = new Object[] { "A", "B", 1, 2 };
                Message msg = new ObjectMessage(null, myTuple);
                channel.send(msg);
            } catch (Exception e) {
                logger.error(e);
            }
        }
    }

    TupleSpaceServer() {
        this.tupleSpace = new TupleSpace();
    }

    public static void main(String[] args) throws Exception {
        new TupleSpaceServer().start();
    }
}
