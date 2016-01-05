;(function($) {

    var app = $.sammy(function() {
        this.use(Sammy.EJS);

        var sammy = this;
        var flags = {};
        var serveAdd = {};

        function setActive(name) {
            $('#topBar li').removeClass('active');
            $('#topBar li.' + name).addClass('active');
        }

        function toISOString(originalDate) {
            var date = moment(originalDate);

            // zero out the seconds so we start on the minute
            date.seconds(0);

            // Readjust the offset since the dateTimePicker adjusted the value but not the timezone
            date.subtract(date.utcOffset(), 'm');

            return date.toISOString();
        }

       // Paths
        this.get('#/', function() {
            this.render('templates/index.ejs', function(html) {
                $('#mainContent').html(html);
                setActive('home');
            });
        });
    });


    $(function () {
        app.run('#/');
    });

})(jQuery);