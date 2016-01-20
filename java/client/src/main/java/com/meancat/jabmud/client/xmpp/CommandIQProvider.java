package com.meancat.jabmud.client.xmpp;

import org.jivesoftware.smack.SmackException;
import org.jivesoftware.smack.provider.IQProvider;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.xmlpull.v1.XmlPullParser;
import org.xmlpull.v1.XmlPullParserException;

import java.io.IOException;

public class CommandIQProvider extends IQProvider<CommandIQ> {
    private static final Logger logger = LoggerFactory.getLogger(CommandIQProvider.class);

    @Override
    public CommandIQ parse(XmlPullParser parser, int initialDepth) throws XmlPullParserException, IOException, SmackException {

        String cmdName = ""; // if this is null, then IQ.toString() will throw NullPointerException
        outerloop: while(true) {
            int eventType = parser.next();
            logger.info("eventType is {}", eventType);
            switch(eventType) {
                case XmlPullParser.START_TAG:
                    String elementName = parser.getName();
                    logger.info("elementName is {}", elementName);
                    cmdName = parser.getAttributeValue(null, "cmdName");
                    logger.info("cmdName is {}", cmdName);
                    break;

                case XmlPullParser.END_TAG:
                    if (parser.getDepth() == initialDepth) {
                        break outerloop;
                    }
                    break;
            }
        }
        return new CommandIQ(cmdName);
    }
}
