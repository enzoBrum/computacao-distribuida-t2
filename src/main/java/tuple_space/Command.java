package tuple_space;

import java.io.Serializable;

public class Command implements Serializable {
    private static final long serialVersionUID = 1L;

    public enum Header {
        GET,
        READ,
        WRITE,
        GET_ALL
    };

    private Header header;
    private Object[] tuple;

    public Command(Header h, Object[] t) {
        this.tuple = t;
        this.header = h;
    }

    void setHeader(Header h) {
        header = h;
    }

    void setTuple(Object[] t) {
        tuple = t;
    }

    Header getHeader() {
        return header;
    }

    Object[] getTuple() {
        return tuple;
    }
}
