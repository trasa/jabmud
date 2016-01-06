package com.meancat.jabmud.client;

import com.meancat.jabmud.shared.SharedConfiguration;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.config.PropertyPlaceholderConfigurer;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Import;
import org.springframework.core.io.ClassPathResource;
import org.springframework.core.io.DefaultResourceLoader;
import org.springframework.core.io.FileSystemResource;
import org.springframework.core.io.Resource;
import org.springframework.util.SystemPropertyUtils;

import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;

import static com.google.common.collect.Lists.newArrayList;

/**
 *
 */
@Configuration
@ComponentScan(basePackageClasses = {ClientConfiguration.class})
@Import({
        SharedConfiguration.class
})
public class ClientConfiguration {
        private static final Logger logger = LoggerFactory.getLogger(ClientConfiguration.class);

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
                if (resource.exists()) {
                        resources.add(resource);
                } else {
                        resources.add(new ClassPathResource(Client.PACKAGE_PATH + "/conf/config.properties"));
                }

                String homeRc = SystemPropertyUtils.resolvePlaceholders("${user.home}/jabmud-client.properties");
                if (Files.exists(Paths.get(homeRc))) {
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

}
