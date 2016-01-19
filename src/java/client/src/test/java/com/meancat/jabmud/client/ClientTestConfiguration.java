package com.meancat.jabmud.client;

import org.apache.commons.io.FilenameUtils;
import org.springframework.beans.factory.config.PropertyPlaceholderConfigurer;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.util.Properties;

@Configuration
public class ClientTestConfiguration {

    @Bean
    public static PropertyPlaceholderConfigurer propertyConfigurer() {
        Properties properties = new Properties();

        String dir = System.getProperty("user.dir");
        // attempt to correct working dir name so that we can find src/main
        if (dir.endsWith("jabmud")) {
            dir = FilenameUtils.concat(dir, "client");
        }
        System.setProperty("app.home", FilenameUtils.concat(dir, "src/main"));

        System.setProperty("app.environment", "dev");

        System.setProperty("zoo.connectString", "someserver:1234");



        PropertyPlaceholderConfigurer bean = new PropertyPlaceholderConfigurer();
        bean.setProperties(properties);
        return bean;
    }
}
