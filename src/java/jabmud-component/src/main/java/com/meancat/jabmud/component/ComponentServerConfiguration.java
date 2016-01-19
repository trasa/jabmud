package com.meancat.jabmud.component;

import com.codahale.metrics.MetricRegistry;
import com.codahale.metrics.health.HealthCheckRegistry;
import com.codahale.metrics.servlets.HealthCheckServlet;
import com.codahale.metrics.servlets.MetricsServlet;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.meancat.jabmud.shared.SharedConfiguration;
import com.meancat.usefully.jersey.CompressionInterceptor;
import com.meancat.usefully.jersey.MetricFilter;
import com.meancat.usefully.jersey.RequestBodyConverter;
import com.meancat.usefully.jersey.ResponseBodyConverter;
import org.eclipse.jetty.server.HttpConfiguration;
import org.eclipse.jetty.server.HttpConnectionFactory;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.server.ServerConnector;
import org.eclipse.jetty.servlet.ServletHolder;
import org.eclipse.jetty.webapp.WebAppContext;
import org.glassfish.jersey.server.ResourceConfig;
import org.glassfish.jersey.server.filter.HttpMethodOverrideFilter;
import org.glassfish.jersey.server.filter.RolesAllowedDynamicFeature;
import org.glassfish.jersey.servlet.ServletContainer;
import org.reflections.Reflections;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.beans.factory.config.PropertyPlaceholderConfigurer;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Import;
import org.springframework.core.io.*;
import org.springframework.util.StringUtils;
import org.springframework.util.SystemPropertyUtils;
import org.springframework.web.context.WebApplicationContext;
import org.springframework.web.context.support.AbstractRefreshableWebApplicationContext;

import javax.ws.rs.Path;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Date;

import static com.google.common.collect.Lists.newArrayList;


@Configuration
@ComponentScan(basePackageClasses = {ComponentServerConfiguration.class })
@Import({
        SharedConfiguration.class
})
public class ComponentServerConfiguration {
    private static final Logger logger = LoggerFactory.getLogger(ComponentServerConfiguration.class);


    @Value("${http.bindIp}")
    public String bindIp;

    @Value("${http.bindPort:7130}")
    public int bindPort;

    @Value("${http.contentPath}")
    public Resource contentPath;

    // TODO curator/zookeeper support
//    @Value("${zookeeper.servicepath}")
//    public String servicepath;

    // TODO curator/zookeeper support
//    public static final String CONFIG_FILE_PATTERN = "/jabmud/config/component-server.${app.environment}.properties";
//    public static final String GLOBAL_ENV_FILE_PATTERN = "/jabmud/config/global.${app.environment}.properties";


    // TODO no curator/zookeeper support for now
    @Bean
    public static PropertyPlaceholderConfigurer propertyConfigurer(/*CuratorFramework curatorFramework*/) throws Exception {

        ArrayList<Resource> resources = newArrayList();

        /*
        String globalEnvConfigFile = SystemPropertyUtils.resolvePlaceholders(GLOBAL_ENV_FILE_PATTERN);
        logger.info("attempting to read global properties from zookeeper at {}", globalEnvConfigFile);
        if (curatorFramework.checkExists().forPath(globalEnvConfigFile) != null) {
            byte[] data = GzipUtil.decompress(curatorFramework.getData().forPath(globalEnvConfigFile));
            resources.add(new ByteArrayResource(data));
        } else {
            logger.warn("{} not found in zookeeper!  Not reading in any properties from the global config file!", globalEnvConfigFile);
            // but we continue, to read the local config file...
        }
        */

        /*
        // interpret CONFIG_FILE_PATTERN here...
        String configFile = SystemPropertyUtils.resolvePlaceholders(CONFIG_FILE_PATTERN);
        logger.info("attempting to read properties from zookeeper at {}", configFile);
        if (curatorFramework.checkExists().forPath(configFile) != null) {
            byte[] data = GzipUtil.decompress(curatorFramework.getData().forPath(configFile));
            resources.add(new ByteArrayResource(data));
        } else {
            logger.warn("{} not found in zookeeper!  Not reading in any properties from the shared config file!", configFile);
            // but we continue, to read the local config file...
        }
        */

        // local config.properties
        FileSystemResource resource = new FileSystemResource(SystemPropertyUtils.resolvePlaceholders("${app.home}/config.properties"));
        if(resource.exists()) {
            resources.add(resource);
        } else {
            resources.add(new ClassPathResource(ComponentServer.PACKAGE_PATH + "/conf/config.properties"));
        }

        String homeRc = SystemPropertyUtils.resolvePlaceholders("${user.home}/jabmud-component.properties");
        if(Files.exists(Paths.get(homeRc))) {
            logger.warn("Using properties from user.home!");
            resources.add(new DefaultResourceLoader().getResource("file:" + homeRc));
        }

        PropertyPlaceholderConfigurer bean = new PropertyPlaceholderConfigurer();
        bean.setLocations(resources.toArray(new Resource[resources.size()]));
        bean.setSearchSystemEnvironment(true);
        bean.setSystemPropertiesMode(PropertyPlaceholderConfigurer.SYSTEM_PROPERTIES_MODE_OVERRIDE);
        bean.setNullValue("@null");
        return bean;
    }

    // TODO curator/zookeeper support
    @Bean
    public String applicationVersion(/*CuratorFramework curatorFramework*/) throws Exception {
        StringBuilder builder = new StringBuilder("Version: ");
        builder.append(ComponentServer.class.getPackage().getImplementationVersion());
        builder.append("\nDate: ");
        builder.append(new Date(System.currentTimeMillis()));

        /*
        curatorFramework
                .newNamespaceAwareEnsurePath(servicepath)
                .ensure(curatorFramework.getZookeeperClient());
        String path = servicepath + "/" + ComponentServer.class.getSimpleName() + "_";

        curatorFramework
                .create()
                .withMode(CreateMode.EPHEMERAL_SEQUENTIAL)
                .forPath(path, builder.toString().getBytes());
        */

        return ComponentServer.class.getPackage().getImplementationVersion();
    }

    /**
     * Jersey ResourceConfig
     * @return resourceConfig
     */
    @Bean
    public ResourceConfig resourceConfig(Reflections reflections,
                                         ObjectMapper objectMapper,
                                         MetricRegistry metricRegistry
                                         ) throws IOException {
        ResourceConfig resourceConfig = new ResourceConfig();
        // Load all the handlers
        for(Class<?> clazz : reflections.getTypesAnnotatedWith(Path.class)) {
            logger.info("Register REST Service Handler: {}", clazz);
            resourceConfig.register(clazz);
        }

        // Register json serializer/deserializer
        resourceConfig.register(new RequestBodyConverter(objectMapper, metricRegistry));
        resourceConfig.register(new ResponseBodyConverter(objectMapper, metricRegistry));

        // filters
        resourceConfig.register(new HttpMethodOverrideFilter());
        resourceConfig.register(new MetricFilter(metricRegistry));

        // Compression
        resourceConfig.register(new CompressionInterceptor());

        // enable role based permissions https://jersey.java.net/documentation/latest/security.html
        resourceConfig.register(RolesAllowedDynamicFeature.class);


        return resourceConfig;
    }

    @Bean(initMethod = "start", destroyMethod = "stop")
    public Server webServer(ApplicationContext applicationContext,
                            WebAppContext context
                            ) throws Exception {

        // is this being called from the app starting up or from a unit test?
        if (!(applicationContext instanceof AbstractRefreshableWebApplicationContext)) {
            return null; // do nothing.
        }

        Server bean = new Server();

        // http config
        HttpConfiguration httpConfig = new HttpConfiguration();
        httpConfig.setOutputBufferSize(32768);
        httpConfig.setRequestHeaderSize(8192);
        httpConfig.setResponseHeaderSize(8192);

        ServerConnector httpConnector = new ServerConnector(bean, new HttpConnectionFactory(httpConfig));
        httpConnector.setPort(bindPort);
        if (StringUtils.hasText(bindIp)) {
            httpConnector.setHost(bindIp);
        }
        bean.addConnector(httpConnector);

        AbstractRefreshableWebApplicationContext ctx = (AbstractRefreshableWebApplicationContext)applicationContext;
        bean.setHandler(context);

        ctx.setServletContext(context.getServletContext());

        return bean;
    }


    @Bean
    public WebAppContext webAppContext(ApplicationContext applicationContext,
                                       MetricRegistry metricRegistry,
                                       HealthCheckRegistry healthCheckRegistry,
                                       ResourceConfig resourceConfig) throws IOException {
        WebAppContext context;
        /* TODO
        if (isInDockerContainer()) {
            // this is for the docker based compiled resource
            logger.info("Docker Detected - Configuring WebAppContext with ClassPathResource {}", DeerHunterServer.PACKAGE_PATH + "/web");
            context = buildUberJarWebAppContext();
        } else if (!contentPath.getFile().exists()) {
            logger.info("Content Path doesn't exist, using ClassPathResource {} instead", DeerHunterServer.PACKAGE_PATH + "/web");
            context = buildUberJarWebAppContext();
        } else {
            // local development
            logger.warn("Not in Docker Container - using local contentPath - {}", contentPath.getFile().getAbsolutePath());
            context = buildLocalFileWebAppContext();
        }
        */
        context = buildLocalFileWebAppContext();
        // TODO

        context.setContextPath("/");
        context.setAttribute(WebApplicationContext.ROOT_WEB_APPLICATION_CONTEXT_ATTRIBUTE, applicationContext);
        context.getServletContext().setAttribute(MetricsServlet.METRICS_REGISTRY, metricRegistry);
        context.getServletContext().setAttribute(HealthCheckServlet.HEALTH_CHECK_REGISTRY, healthCheckRegistry);

        // jersey client
        context.addServlet(new ServletHolder(new ServletContainer(resourceConfig)), "/api/*");

        return context;
    }

    private WebAppContext buildUberJarWebAppContext() throws IOException {
        WebAppContext context = new WebAppContext();
        context.setContextPath("/");
        context.setResourceBase(new ClassPathResource(ComponentServer.PACKAGE_PATH + "/web").getURI().toString());
        return context;
    }

    private WebAppContext buildLocalFileWebAppContext() throws IOException {
        return new WebAppContext(contentPath.getFile().getAbsolutePath(), "/");
    }
}
