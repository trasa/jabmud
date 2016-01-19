package com.meancat.jabmud.client.xmpp;

import org.jivesoftware.smack.packet.IQ;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class CommandIQ extends IQ {
    private static final Logger logger = LoggerFactory.getLogger(CommandIQ.class);

    String command;

    public CommandIQ() {
        super("command");
    }

    public CommandIQ(String command) {
        this();
        this.command = command;
    }

    @Override
    protected IQChildElementXmlStringBuilder getIQChildElementBuilder(IQChildElementXmlStringBuilder xml) {
        xml.attribute("cmdName", command);
        xml.xmlnsAttribute("jabmud:iq:command");
        xml.rightAngleBracket();
        return xml;
    }

    public String getCommand() {
        return command;
    }

    public void setCommand(String command) {
        this.command = command;
    }
}
