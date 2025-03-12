requirejs.config({
    paths: {
        ramda: 'https://cdnjs.cloudflare.com/ajax/libs/ramda/0.13.0/ramda.min',
        jquery: 'https://ajax.googleapis.com/ajax/libs/jquery/2.1.1/jquery.min'
    }
});

require(
    [
        'ramda',
        'jquery'
    ],
    function (_, $) {
        ////////////////////////////////////////////
        const trace = _.curry(function (tag, x) {
            console.log(tag, x);
            return x;
        });

        const Impure = {
            getJSON: _.curry(function (callback, url) {
                $.getJSON(url, callback);
            }),

            setHtml: _.curry(function (sel, html) {
                console.log(sel, html);
                $(sel).html(html);
            })
        };

        const url = function (term) {
            return 'https://api.flickr.com/services/feeds/photos_public.gne?tags=' + term + '&format=json&jsoncallback=?';
        };

        const img = function (url) {
            return $('<img />', { src: url });
        };

        // 查看数据结构
        // var app = _.compose(Impure.getJSON(trace("response")), url);
        // app("cats");

        // 观察 _.prop 的实现
        // var prop = _.curry(function(property, object){
        //     return object[property];
        // });

        ////////////////////////////////////////////
        // var mediaUrl = _.compose(_.prop('m'), _.prop('media'));

        // var srcs = _.compose(_.map(mediaUrl), _.prop('items'));

        // var images = _.compose(_.map(img), srcs);

        // // var renderImages = _.compose(Impure.setHtml("body"), images);
        // var renderImages = _.compose(trace("renderImages"), images);
        
        // var app = _.compose(Impure.getJSON(renderImages), url);

        // app("cats");

        //////////////////////////////////////////// 重构
        const mediaUrl = _.compose(_.prop('m'), _.prop('media'));
        const mediaToImg = _.compose(img, mediaUrl);
        const images = _.compose(_.map(mediaToImg), _.prop('items'));
        // const renderImages = _.compose(Impure.setHtml("body"), images);
        var renderImages = _.compose(trace("renderImages"), images);
        const app = _.compose(Impure.getJSON(renderImages), url);
        app("cats");
    }
);