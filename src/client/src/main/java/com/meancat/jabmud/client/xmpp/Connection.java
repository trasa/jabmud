package com.meancat.jabmud.client.xmpp;

import org.jivesoftware.smack.ConnectionConfiguration;
import org.jivesoftware.smack.SmackException;
import org.jivesoftware.smack.StanzaListener;
import org.jivesoftware.smack.XMPPException;
import org.jivesoftware.smack.chat.Chat;
import org.jivesoftware.smack.chat.ChatManager;
import org.jivesoftware.smack.chat.ChatMessageListener;
import org.jivesoftware.smack.packet.IQ;
import org.jivesoftware.smack.packet.Message;
import org.jivesoftware.smack.packet.Stanza;
import org.jivesoftware.smack.provider.ProviderManager;
import org.jivesoftware.smack.tcp.XMPPTCPConnection;
import org.jivesoftware.smack.tcp.XMPPTCPConnectionConfiguration;
import org.jivesoftware.smackx.bytestreams.ibb.provider.DataPacketProvider;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;

import java.io.IOException;

/**
 *
 */
@Component
public class Connection {
    private static final Logger logger = LoggerFactory.getLogger(Connection.class);

    private XMPPTCPConnection conn;

    Chat chat;
    ChatManager cm;


    public void connect() throws IOException, XMPPException, SmackException {

        // hack in a provider
        ProviderManager.addIQProvider("command", "jabmud:iq:command", new CommandIQProvider());


        XMPPTCPConnectionConfiguration config = XMPPTCPConnectionConfiguration.builder()
                .setDebuggerEnabled(true)
                .setServiceName("bw-mbp-trasa.glu.com")
                .setHost("localhost")
                .setSecurityMode(ConnectionConfiguration.SecurityMode.disabled)
                .setCompressionEnabled(false)
                .setUsernameAndPassword("tony.rasa", "password")
                .build();
        conn = new XMPPTCPConnection(config);
        conn.connect();

        conn.login();


/*
        // send a test chat message
        cm = ChatManager.getInstanceFor(conn);

        chat = cm.createChat("jabmud.localhost", new ChatMessageListener() {
            @Override
            public void processMessage(Chat chat, Message message) {
                logger.info("Received Message: {}", message);
            }
        });
        chat.sendMessage("wha");
*/

        // send a test IQ command
        IQ iq = new CommandIQ("bleh");
        iq.setTo("jabmud.localhost");
        conn.sendIqWithResponseCallback(iq, new StanzaListener() {
            @Override
            public void processPacket(Stanza packet) throws SmackException.NotConnectedException {
                logger.info("callback: {}", packet);
            }
        });
    }
}
