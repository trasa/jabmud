package com.meancat.jabmud.client.ui;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.support.AbstractApplicationContext;

import javax.swing.*;
import java.awt.*;
import java.awt.event.KeyEvent;
import java.awt.event.KeyListener;
import java.awt.event.WindowAdapter;
import java.awt.event.WindowEvent;

public class MainForm {
    private static final Logger logger = LoggerFactory.getLogger(MainForm.class);

    AbstractApplicationContext applicationContext;
    private JTextField inputText;
    private JTextArea ouputTextArea;
    private JList playerList;
    private JPanel panel;

    public MainForm(AbstractApplicationContext applicationContext) {
        this.applicationContext = applicationContext;
        // TODO fill in the @Autowired dependencies of this class ..
        // applicationContext.
    }

    public static void show(final AbstractApplicationContext applicationContext) {
        final JFrame frame = new JFrame("MainForm");
        MainForm mainForm = new MainForm(applicationContext);
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

        mainForm.inputText.addKeyListener(new KeyListener() {
            @Override
            public void keyTyped(KeyEvent e) {
            }

            @Override
            public void keyPressed(KeyEvent e) {
                if (e.getKeyCode() == KeyEvent.VK_ENTER) {
                    logger.info("enter pressed");
                }
            }

            @Override
            public void keyReleased(KeyEvent e) {
            }
        });
        frame.pack();
        frame.setVisible(true);
    }
}
