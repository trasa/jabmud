package com.meancat.jabmud.client.xmpp;

import org.jivesoftware.smack.ConnectionConfiguration;
import org.jivesoftware.smack.SmackException;
import org.jivesoftware.smack.XMPPException;
import org.jivesoftware.smack.chat.Chat;
import org.jivesoftware.smack.chat.ChatManager;
import org.jivesoftware.smack.chat.ChatMessageListener;
import org.jivesoftware.smack.packet.Message;
import org.jivesoftware.smack.tcp.XMPPTCPConnection;
import org.jivesoftware.smack.tcp.XMPPTCPConnectionConfiguration;
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


        cm = ChatManager.getInstanceFor(conn);

        chat = cm.createChat("jabmud.localhost", new ChatMessageListener() {
            @Override
            public void processMessage(Chat chat, Message message) {
                logger.info("Received Message: {}", message);
            }
        });
        chat.sendMessage("wha");



    }
}
