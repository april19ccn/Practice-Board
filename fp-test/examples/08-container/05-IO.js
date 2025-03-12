import * as R from "../../node_modules/ramda/es/index.js"
import * as _ from "../../utils/support.js"

//////////////////////////////////////// 薛定谔的 Maybe
var Maybe = function (x) {
    this.__value = x;
}

Maybe.of = function (x) {
    return new Maybe(x);
}

Maybe.prototype.isNothing = function () {
    return (this.__value === null || this.__value === undefined);
}

Maybe.prototype.map = function (f) {
    return this.isNothing() ? Maybe.of(null) : Maybe.of(f(this.__value));
}
///////////////////////////////////////


//  getFromStorage :: String -> (_ -> String)
var getFromStorage = function (key) {
    return function () {
        return localStorage[key];
    }
}

class IO {
    static of(x) {
        return new IO(() => x);
    }

    constructor(f) {
        this.__value = f;
    }

    map(fn) {
        return new IO(R.compose(fn, this.__value));
    }

    inspect() {
        return `IO(${_.inspect(this.__value)})`;
    }
}

//  io_window_ :: IO Window
var io_window = new IO(function () { return window; });

console.log(io_window.map(function (win) { return win.innerWidth }).inspect());
// IO(1430)

console.log(io_window.map(R.prop('location')).map(R.prop('href')).map(R.split('/')).inspect());
// IO(["http:", "", "localhost:8000", "blog", "posts"])


////// 纯代码库: lib/params.js ///////

//  url :: IO String
var url = new IO(function() { return window.location.href; }); // 这里是new
// 等价于 
// var url = IO.of(window.location.href);
// 不等价于这个 
// var url = IO.of(function() { return window.location.href; });
// console.log(url.__value())

//  toPairs =  String -> [[String]]
var toPairs = R.compose(R.map(R.split('=')), R.split('&'));

//  params :: String -> [[String]]
var params = R.compose(toPairs, R.last, R.split('?'));

//  findParam :: String -> IO Maybe [String]
var findParam = function(key) {
  return R.map(R.compose(Maybe.of, R.filter(R.compose(_.eq(key), R.head)), params), url);
};

console.log(findParam("key"))
// new IO(R.compose(R.compose(Maybe.of, R.filter(R.compose(_.eq(key), R.head)), params), window.location.href))

////// 非纯调用代码: main.js ///////

// 调用 __value() 来运行它！
// ?searchTerm=1234&key=189948
console.log(findParam("key").__value());
// Maybe(['searchTerm', 'wafflehouse'])