package com.meancat.jabmud.client.ui;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.support.AbstractApplicationContext;

import javax.swing.*;
import java.awt.*;
import java.awt.event.WindowAdapter;
import java.awt.event.WindowEvent;

public class MainForm {
    private static final Logger logger = LoggerFactory.getLogger(MainForm.class);

    private JTextField inputText;
    private JTextArea ouputTextArea;
    private JList playerList;
    private JPanel panel;

    public static void show(final AbstractApplicationContext applicationContext) {
        final JFrame frame = new JFrame("MainForm");
        MainForm mainForm = new MainForm();
        frame.setContentPane(mainForm.panel);
        logger.info("panel size {}", mainForm.panel.getPreferredSize());
        frame.setPreferredSize(mainForm.panel.getPreferredSize());

        frame.addWindowListener(new WindowAdapter() {
            @Override
            public void windowClosing(WindowEvent e) {
                logger.info("window closing, shutting down app context");
                applicationContext.close();
                logger.info("stopped");
                frame.dispose();
            }
        });

        frame.pack();
        frame.setVisible(true);
    }
}
