var BOSH_SERVICE = 'http://localhost:5280/http-bind'
;(function($) {
    var app = $.sammy(function() {
        this.use(Sammy.EJS);

        this.get('#/', function() {
            this.render('templates/index.ejs', function(html) {
                $('#mainContent').html(html);
                var connection = new Strophe.Connection(BOSH_SERVICE);

                function displayMessage(msg) {
                    var display = $('#display');
                    display
                        .append(msg + "\n\n")
                        .stop()
                        .animate({ scrollTop: display[0].scrollHeight}, 800);
                }

                function onConnect(status)
                {
                    if (status == Strophe.Status.CONNECTING) {
                        displayMessage('Strophe is connecting.');
                    } else if (status == Strophe.Status.CONNFAIL) {
                        displayMessage('Strophe failed to connect.');
                        $('#connect').get(0).value = 'connect';
                    } else if (status == Strophe.Status.DISCONNECTING) {
                        displayMessage('Strophe is disconnecting.');
                    } else if (status == Strophe.Status.DISCONNECTED) {
                        displayMessage('Strophe is disconnected.');
                        $('#connect').get(0).value = 'connect';
                    } else if (status == Strophe.Status.CONNECTED) {
                        displayMessage('Strophe is connected.');
                        connection.addHandler(onMessage, null, 'message', null, null, null);
                        connection.addHandler(onIq, null, "iq", null, null, null);
                        connection.addHandler(onPresence, null, "presence", null, null, null);
                        // send presence
                        connection.send($pres().tree());
                    }
                }

                function onMessage(msg) {
                    var to = msg.getAttribute('to');
                    var from = msg.getAttribute('from');
                    var type = msg.getAttribute('type');
                    var bodyElements = msg.getElementsByTagName('body');

                    var bodyStr = $('<div/>').append(bodyElements).html();

                    displayMessage("(" + type + " from " + from + ") " + bodyStr);

                    // we must return true to keep the handler alive.
                    // returning false would remove it after it finishes.
                    return true;
                }

                function onIq(msg) {
                    console.log(msg);
                    var to = msg.getAttribute('to');
                    var from = msg.getAttribute('from');
                    var type = msg.getAttribute('type');

                    // TODO: jquery.text() strips out all the xml leaving just the text, which
                    // isn't really what we want.
                    var bodyElements = $(msg).text();
                    var bodyStr = $('<div/>').append(bodyElements).html();

                    displayMessage("(iq " + type + " from " + from + ") " + bodyStr);

                    // we must return true to keep the handler alive.
                    // returning false would remove it after it finishes.
                    return true;
                }

                function onPresence(msg) {
                    console.log(msg);
                    var to = msg.getAttribute("to");
                    var from = msg.getAttribute("from");
                    var body = $('<div/>').text($(msg).html()).html();

                    displayMessage("(presence from " + from + ") :" + body);

                    // we must return true to keep the handler alive.
                    // returning false would remove it after it finishes.
                    return true;
                }

                function send() {
                    var buf = $('#buf');
                    var msg = buf.val();
                    if (msg === "") {
                        // nothing to do
                        return;
                    }
                    console.log("sending '" + msg + "'");
                    displayMessage("[sending '" + msg + "']");
                    var command = msg.split(" ");

                    var to = "jabmud.localhost";
                    var from = $('#jid').get(0).value;

                    var iqCommand = $iq({ to: to, from: from, type: 'get'}).c('command').attrs({name: command[0]});
                    for (i = 1; i < command.length; i++) {
                        iqCommand = iqCommand.c('arg').t(command[i]).up();
                    }
                    connection.send(iqCommand.tree());

                    // clear the buffer for the next line
                    buf.val('');

                    // testing:
                    // var iqroster = $iq({type: 'get'}).c('query', {xmlns: 'jabber:iq:roster'});
                    // connection.sendIQ(iqroster, rostercallback)
                }

                // testing
                /*
                function rostercallback(iq) {
                    console.log("rostercallback");
                    console.log(iq);
                    $(iq).find('item').each(function(){
                        var jid = $(this).attr('jid'); // The jabber_id of your contact
                        // You can probably put them in a unordered list and and use their jids as ids.
                        displayMessage(jid)
                    });
                }
                */

                $('#connect').click(function(e) {
                    // Uncomment the following lines to spy on the wire traffic.
                    //connection.rawInput = function (data) { displayMessage('RECV: ' + data); };
                    //connection.rawOutput = function (data) { displayMessage('SEND: ' + data); };

                    // Uncomment the following line to see all the debug output.
                    //Strophe.log = function (level, msg) { displayMessage('LOG: ' + msg); };

                    var button = $('#connect').get(0);
                    if (button.value == 'connect') {
                        button.value = 'disconnect';

                        connection.connect($('#jid').get(0).value, $('#pass').get(0).value, onConnect);
                    } else {
                        button.value = 'connect';
                        connection.disconnect();
                    }
                });

                $('#send').click(function(e) {
                    send();
                });

                $('#testing').click(function() {
                    var to = "jabmud.localhost/mynick";
                    var from = $('#jid').get(0).value;

                    var c = $pres({ to: to, from: from })
                        .c('x').attrs({
                            xmlns: "http://jabber.org/protocol/muc"
                        });
                    displayMessage("sending command");
                    connection.send(c.tree());
                });

                $('#buf')
                    .bind("enterKey", function(e) {
                        send();
                    })
                    .keyup(function (e) {
                        if (e.keyCode == 13) {
                            $(this).trigger('enterKey');
                        }
                    });
            });
        });
    });

    $(function() {
        app.run('#/');
    });
})(jQuery);