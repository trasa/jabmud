package com.meancat.jabmud.component;

import org.apache.commons.io.FilenameUtils;
import org.springframework.beans.factory.config.PropertyPlaceholderConfigurer;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.util.Properties;

@Configuration
public class ComponentServerTestConfiguration {

    @Bean
    public static PropertyPlaceholderConfigurer propertyConfigurer() {
        Properties properties = new Properties();




        properties.put("http.bindIp", "");
        properties.put("http.bindPort", "9000");
        properties.put("http.contentPath", "file:${app.home}/web");

//
//        // what to do about zookeeper?
//        properties.put("zookeeper.gamepath", "/jabmud/dev");
//        properties.put("zookeeper.servicepath", "${zookeeper.gamepath}/services");
//        properties.put("zookeeper.datapath", "${zookeeper.gamepath}/data");





        String dir = System.getProperty("user.dir");
        // attempt to correct working dir name so that we can find src/main
        if (dir.endsWith("jabmud")) {
            dir = FilenameUtils.concat(dir, "jabmud-component");
        }
        System.setProperty("app.home", FilenameUtils.concat(dir, "src/main"));

        System.setProperty("app.environment", "dev");

        System.setProperty("zoo.connectString", "someserver:1234");



        PropertyPlaceholderConfigurer bean = new PropertyPlaceholderConfigurer();
        bean.setProperties(properties);
        return bean;
    }
}
