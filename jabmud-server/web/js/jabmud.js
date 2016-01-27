;(function($) {
    var app = $.sammy(function() {
        this.use(Sammy.EJS);

        this.get('#/', function() {
            this.render('templates/index.ejs', function(html) {
                $('#mainContent').html(html);
            });
        });
    });

    $(function() {
        app.run('#/');
    });
})(jQuery);