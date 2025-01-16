function compose() {
    var args = arguments;
    var start = args.length - 1;
    return function() {
        var i = start;
        var result = args[start].apply(this, arguments);
        while (i--) result = args[i].call(this, result);
        return result;
    };
};

var split = function(x) { return x.split(','); };
var toUpperCase = function(x) { return x.toUpperCase(); };
var hello = function(x) { return 'HELLO,' + x; };

console.log(compose(split, toUpperCase, hello)('world'));