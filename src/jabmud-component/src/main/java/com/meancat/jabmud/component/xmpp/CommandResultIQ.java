package com.meancat.jabmud.component.xmpp;

import org.xmpp.packet.IQ;

/**
 *
 */
public class CommandResultIQ extends IQ {


    public CommandResultIQ(IQ getIQ) {
        super(Type.result, getIQ.getID());
        this.setFrom(getIQ.getTo());
        this.setTo(getIQ.getFrom());

        this.setChildElement(docFactory.createElement("command", "jabmud:iq:command"));
    }

}
