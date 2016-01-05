package com.meancat.jabmud.component.xmpp;

import org.jivesoftware.whack.ExternalComponentManager;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.BeansException;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationContextAware;
import org.springframework.stereotype.Component;
import org.springframework.web.context.support.AnnotationConfigWebApplicationContext;
import org.xmpp.component.ComponentException;

import javax.annotation.PostConstruct;

/**
 * Connects to the Jabber Server and wires up a MudComponent to receive traffic.
 */
@Component
public class ExternalMudComponent implements ApplicationContextAware {
    private static final Logger logger = LoggerFactory.getLogger(ExternalMudComponent.class);

    AnnotationConfigWebApplicationContext applicationContext;

    @PostConstruct
    public void init() {
        start();
    }

    public void start() {
        // TODO args for connection stuff
        ExternalComponentManager manager = new ExternalComponentManager("localhost", 5275);
        // secret key
        manager.setSecretKey("jabmud", "secret");

        // ?
        manager.setMultipleAllowed("jabmud", true);

        try {
            manager.addComponent("jabmud", new MudComponent());
            while (applicationContext.isActive()) {
                try {
                    Thread.sleep(5000);
                } catch (InterruptedException e) {
                    logger.error("interrupted", e);
                }
            }
        } catch (ComponentException e) {
            logger.error("Failed to add component", e);
        }
    }

    @Override
    public void setApplicationContext(ApplicationContext applicationContext) throws BeansException {
        this.applicationContext = (AnnotationConfigWebApplicationContext) applicationContext;
    }
}
