package tuple_space;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.jgroups.JChannel;
import org.jgroups.raft.RaftHandle;
import org.jgroups.raft.StateMachine;

public class RaftTupleSpaceServer {
    private Logger logger = LogManager.getLogger(RaftTupleSpaceServer.class);
    private JChannel channel;
    private RaftHandle raftHandle;
    private StateMachine stateMachine; // TODO: should be a tuple space.

    public RaftTupleSpaceServer() throws Exception {
        this.channel = new JChannel("src/main/resources/raft.xml");
        this.raftHandle = new RaftHandle(this.channel, this.stateMachine); // TODO: implement state machine
        this.channel.connect("tuple-space"); // cluster name
    }
}
