package com.meancat.jabmud.client;

import ch.qos.logback.classic.Logger;
import ch.qos.logback.classic.LoggerContext;
import ch.qos.logback.classic.joran.JoranConfigurator;
import ch.qos.logback.core.joran.spi.JoranException;
import ch.qos.logback.core.util.StatusPrinter;
import org.kohsuke.args4j.CmdLineException;
import org.kohsuke.args4j.CmdLineParser;
import org.kohsuke.args4j.Option;
import org.slf4j.LoggerFactory;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.core.io.ClassPathResource;
import org.springframework.util.SystemPropertyUtils;

import java.io.File;
import java.io.IOException;
import java.net.URL;

/**
 * Main Class for Client
 */
public class Client {
    public static final String PACKAGE_PATH = Client.class.getPackage().getName().replace('.', '/');


    private static class ApplicationArguments {
        @Option(required = false, name = "--home", aliases = { "-h" }, usage = "Home directory of application.")
        public String home;

        @Option(required = false, name = "--zoo", aliases = { "-z" }, usage = "Zookeeper connection string.")
        public String zooConnectString;

        @Option(required = true, name = "--env", aliases = { "-e" }, usage = "Environment Name [dev,test]")
        public String environment;

    }

    private static AnnotationConfigApplicationContext applicationContext;


    /**
     * Starts the service.
     *
     * @param args
     *            args
     * @throws Exception
     */
    public static void main(String[] args) throws Exception {
        ApplicationArguments appArgs = new ApplicationArguments();

        CmdLineParser parser = new CmdLineParser(appArgs);

        try {
            parser.parseArgument(args);
        } catch (CmdLineException e) {
            System.err.println(e.getMessage());
            System.err.println();
            parser.printUsage(System.err);
            return;
        }

        // set system variables
        System.setProperty("app.environment", appArgs.environment);

        if(appArgs.home != null) {
            System.setProperty("app.home", appArgs.home);
        } else {
            System.setProperty("app.home", System.getProperty("user.dir"));
        }

        if(appArgs.zooConnectString != null) {
            System.setProperty("zoo.connectString", appArgs.zooConnectString);
        }

        configureLogging();


        System.out.println("Starting Service");

        // start spring
        applicationContext = new AnnotationConfigApplicationContext();
        applicationContext.register(ClientConfiguration.class);
        applicationContext.refresh();

        applicationContext.registerShutdownHook();


        System.out.println("Started Client");

        // wait for app to exit
        while (applicationContext.isActive()) {
            Thread.sleep(5000);
        }

        // Stop all appenders (required if AsyncAppender is used)
        ((LoggerContext) LoggerFactory.getILoggerFactory()).stop();

        // exit application
        System.out.println("Stopped Client");
    }

    private static void configureLogging() throws IOException {
        File localLogbackFile = new File(SystemPropertyUtils.resolvePlaceholders("${app.home}/logback.xml"));
        URL logbackUrl = localLogbackFile.toURI().toURL();

        if(!localLogbackFile.exists()) {
            String logbackPath = PACKAGE_PATH + "/conf/logback.xml";
            ClassPathResource logbackResource = new ClassPathResource(logbackPath);
            logbackUrl = logbackResource.getURL();
        }

        // setup logging
        LoggerContext context = (LoggerContext) LoggerFactory.getILoggerFactory();

        try {
            JoranConfigurator configurator = new JoranConfigurator();
            configurator.setContext(context);
            context.reset();
            configurator.doConfigure(logbackUrl);
        } catch (JoranException je) {
            // ignore since StatusPrinter will handle this
        }

        // fix JUL logging
        java.util.logging.Logger rootLogger = java.util.logging.LogManager.getLogManager().getLogger("");
        java.util.logging.Handler[] handlers = rootLogger.getHandlers();
        // noinspection ForLoopReplaceableByForEach
        for (int i = 0; i < handlers.length; i++) {
            rootLogger.removeHandler(handlers[i]);
        }
        org.slf4j.bridge.SLF4JBridgeHandler.install();

        StatusPrinter.printInCaseOfErrorsOrWarnings(context);

        if(!new File("/.dockerinit").exists()) {
            System.out.println("Removing Loggly Appender: Not running in a docker container");
            Logger logger = context.getLogger(Logger.ROOT_LOGGER_NAME);
            logger.detachAppender("asyncLoggly");
        }

        System.out.println("Loaded Logback Configuration: " + logbackUrl.toExternalForm());
    }

    public static ApplicationContext getApplicationContext() {
        return applicationContext;
    }

}
