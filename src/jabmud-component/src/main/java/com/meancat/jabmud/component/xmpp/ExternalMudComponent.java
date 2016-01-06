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

    @Override
    public void setApplicationContext(ApplicationContext applicationContext) throws BeansException {
        this.applicationContext = (AnnotationConfigWebApplicationContext) applicationContext;
    }
}
