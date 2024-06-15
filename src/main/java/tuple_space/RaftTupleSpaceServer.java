package tuple_space;

import java.io.ByteArrayInputStream;
import java.io.DataInput;
import java.net.InetAddress;
import java.nio.ByteBuffer;
import java.util.Arrays;
import java.util.Scanner;
import java.util.stream.Collectors;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.jgroups.Address;
import org.jgroups.JChannel;
import org.jgroups.blocks.cs.BaseServer;
import org.jgroups.blocks.cs.Receiver;
import org.jgroups.blocks.cs.TcpServer;
import org.jgroups.jmx.JmxConfigurator;
import org.jgroups.protocols.raft.Role;
import org.jgroups.protocols.raft.RAFT.RoleChange;
import org.jgroups.raft.RaftHandle;
import org.jgroups.raft.StateMachine;
import org.jgroups.util.Util;

public class RaftTupleSpaceServer {
    private static final Logger logger = LogManager.getLogger(RaftTupleSpaceServer.class);
    private JChannel channel;
    private RaftHandle raftHandle;
    private TupleSpace tupleSpace;

    public RaftTupleSpaceServer() {
        try {
            this.channel = new JChannel(RaftTupleSpaceServer.class.getClassLoader().getResource("raft.xml").toString());
            this.tupleSpace = new TupleSpace();
            this.raftHandle = new RaftHandle(this.channel, this.tupleSpace);
            this.channel.connect("tuple-space");

            this.raftHandle.addRoleListener(role -> logger.debug("Changed role to " + role));
        } catch (Exception e) {
            e.printStackTrace();
            logger.error(e);
        }
    }

    public void start() {
        eventLoop();
    }

    private void eventLoop() {
        while (true) {
            try {
                var scan = new Scanner(System.in);
                scan.nextLine();
                if (tupleSpace != null)
                    tupleSpace.getTuples().forEach(obj -> System.out
                            .println("(" + Arrays.stream(obj).map(Object::toString).collect(Collectors.joining(", "))
                                    + ")"));

                System.out.println(raftHandle.isLeader());

                tupleSpace.getTuples().forEach(obj -> System.out
                        .println("(" + Arrays.stream(obj).map(Object::toString).collect(Collectors.joining(", "))
                                + ")"));

                Object[] new_tuple = new Object[] { "A", "V", 1, 1 };
                System.out.println("\n======================================\n");
                var cmd = new Command(Command.Header.WRITE, new_tuple);

                byte[] buf = Util.objectToByteBuffer(cmd);
                raftHandle.set(buf, 0, buf.length);
            } catch (Exception e) {
                e.printStackTrace();
                logger.error(e);
            }
        }
    }

    public static void main(String[] args) {
        new RaftTupleSpaceServer().start();
    }
}
