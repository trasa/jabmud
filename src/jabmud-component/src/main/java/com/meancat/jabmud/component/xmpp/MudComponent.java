package com.meancat.jabmud.component.xmpp;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.xmpp.component.AbstractComponent;
import org.xmpp.packet.IQ;
import org.xmpp.packet.Message;

/**
 *
 */
public class MudComponent extends AbstractComponent {
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
    public void start() {
        logger.info("MudComponent Start! My JID is {} and domain is {}", getJID(), getDomain());
    }

    @Override
    protected void handleMessage(Message received) {
        logger.info("handleMessage:\n{}", received.toXML());
        if (received.getType().equals(Message.Type.chat)) {
            Message response = new Message();
            response.setFrom(received.getTo());
            response.setTo(received.getFrom());
            response.setType(received.getType());
            response.setThread(received.getThread());

            response.setBody("Hi!");
            send(response);
        }
    }

    @Override
    protected IQ handleIQGet(IQ iq) {
        logger.info("iq get: {}", iq);
        return new CommandResultIQ(iq);
    }

    @Override
    protected IQ handleIQSet(IQ iq) {
        logger.info("iq set: {}", iq);
        return null;
    }


}
