;(function($) {
    var app = $.sammy(function() {
        this.use(Sammy.EJS);

        this.get('#/', function() {
            console.log("this.get #/");
            this.render('templates/index.ejs', function(html) {
                $('#mainContent').html(html);

                function displayMessage(msg) {
                    var display = $('#display');
                    display
                        .append(msg + "\n\n")
                        .stop()
                        .animate({ scrollTop: display[0].scrollHeight}, 800);
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
                    buf.val('');
                }

                $('#send').click(function(e) {
                    send();
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