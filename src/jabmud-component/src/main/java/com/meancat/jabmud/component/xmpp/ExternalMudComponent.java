package com.meancat.jabmud.component.xmpp;

import org.jivesoftware.whack.ExternalComponentManager;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;
import org.xmpp.component.ComponentException;

import javax.annotation.PostConstruct;

/**
 * Connects to the Jabber Server and wires up a MudComponent to receive traffic.
 */
@Component
public class ExternalMudComponent {
    private static final Logger logger = LoggerFactory.getLogger(ExternalMudComponent.class);

    ExternalComponentManager manager;

    @PostConstruct
    public void init() throws ComponentException {
        start();
    }

    public void start() throws ComponentException {
        // TODO args for connection stuff
        manager = new ExternalComponentManager("localhost", 5275);
        // secret key
        manager.setSecretKey("jabmud", "secret");

        // ?
        manager.setMultipleAllowed("jabmud", true);

        manager.addComponent("jabmud", new MudComponent());
    }

}
