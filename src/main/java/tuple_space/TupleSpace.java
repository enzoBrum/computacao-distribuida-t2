package tuple_space;

import java.util.ArrayList;
import java.util.List;

public class TupleSpace {
    private List<Object[]> tuples;

    public TupleSpace() {
        this.tuples = new ArrayList<>();
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
        return new Object[]{};
    }

    public void get(Object[] tuple) {
        Object[] res = read(tuple);
        if (res.length > 0) {
            this.tuples.remove(tuple);
        }
    }

    private boolean matchWithWildcards(Object[] pattern, Object[] candidate) {
        if (pattern.length != candidate.length) return false;
        for (int i = 0; i < pattern.length; i++) {
            if (!(pattern[i].equals(candidate[i]) || pattern[i].equals('*'))) {
                return false;
            }
        }
        return true;
    }
}