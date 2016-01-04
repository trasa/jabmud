package com.meancat.jabmud.shared;

import com.codahale.metrics.MetricRegistry;
import com.codahale.metrics.health.HealthCheckRegistry;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.meancat.usefully.messaging.ObjectMapperFactory;
import com.meancat.usefully.util.PackageUrls;
import com.meancat.usefully.util.UtilConfiguration;
import org.apache.curator.framework.CuratorFramework;
import org.apache.curator.framework.CuratorFrameworkFactory;
import org.apache.curator.retry.ExponentialBackoffRetry;
import org.reflections.Reflections;
import org.reflections.scanners.ResourcesScanner;
import org.reflections.scanners.SubTypesScanner;
import org.reflections.scanners.TypeAnnotationsScanner;
import org.reflections.util.ClasspathHelper;
import org.reflections.util.ConfigurationBuilder;
import org.reflections.util.FilterBuilder;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Import;

import javax.naming.ConfigurationException;
import java.net.URL;
import java.util.Set;

@Configuration
@ComponentScan(basePackageClasses = {SharedConfiguration.class})
@Import({
        UtilConfiguration.class
})
public class SharedConfiguration {
    private static final Logger logger = LoggerFactory.getLogger(SharedConfiguration.class);


//    @Bean(initMethod = "start", destroyMethod = "close")
//    public static CuratorFramework curatorFramework() throws Exception {
//        String zooConnectString = System.getProperty("zoo.connectString");
//        logger.info("Connecting to {}", zooConnectString);
//        return CuratorFrameworkFactory.newClient(zooConnectString, 20000, 20000, new ExponentialBackoffRetry(250, 29, 60_000));
//    }


    @Bean
    public MetricRegistry metricRegistry() {
        return new MetricRegistry();
    }

    @Bean
    public HealthCheckRegistry healthCheckRegistry() {
        return new HealthCheckRegistry();
    }

    /**
     * urls and filterBuilder are defined by the containing service, with the urls and
     * package prefixes needed to configure Reflections.
     *
     * @param packageUrls package urls
     * @param reflectionsFilterBuilder filter prefixes by this
     * @return configured Reflections instance
     */
    @Bean
    public Reflections reflections(PackageUrls packageUrls, FilterBuilder reflectionsFilterBuilder) {
        return new Reflections(new ConfigurationBuilder()
                .filterInputsBy(reflectionsFilterBuilder)
                .setUrls(packageUrls)
                .setScanners(new SubTypesScanner(),
                        new TypeAnnotationsScanner(),
                        new ResourcesScanner()));
    }

    @Bean
    public FilterBuilder filterBuilder() {
        return new FilterBuilder()
                .include("com.meancat.jabmud.component")
                .include("com.meancat.jabmud.shared");
    }

    @Bean
    public PackageUrls packageUrls() {
        Set<URL> urls = ClasspathHelper.forPackage("com.meancat.jabmud.component");
        urls.addAll(ClasspathHelper.forPackage("com.meancat.jabmud.shared"));

        return new PackageUrls(urls);
    }

    @Bean
    public ObjectMapper objectMapper(Reflections reflections, PackageUrls packageUrls) throws ConfigurationException {
        return new ObjectMapperFactory(reflections, packageUrls).create();
    }
}
