package com.meancat.jabmud.component.xmpp;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.xmpp.component.Component;
import org.xmpp.component.ComponentException;
import org.xmpp.component.ComponentManager;
import org.xmpp.packet.JID;
import org.xmpp.packet.Packet;

/**
 *
 */
public class MudComponent implements Component {
    private static final Logger logger = LoggerFactory.getLogger(MudComponent.class);

    @Override
    public String getName() {
        return "JabMud";
    }

    @Override
    public String getDescription() {
        return "Jabmud Component";
    }

    @Override
    public void processPacket(Packet packet) {
        logger.info("processPacket:\n{}", packet.toXML());
        // TODO!
    }

    @Override
    public void initialize(JID jid, ComponentManager componentManager) throws ComponentException {
        logger.info("Initialize: {}, {}", jid, componentManager);
    }

    @Override
    public void start() {
        logger.info("MudComponent Start!");
    }

    @Override
    public void shutdown() {
        logger.info("MudComponent shutdown!");
    }
}
