import * as _ from "ramda"
import { Either, IO, Maybe, compose, chain, map, curry, toUpperCase, Right, Left, either, left, liftA2 } from "../../utils/support.js";
import Task from "data.task";

// 这样是行不通的，因为 2 和 3 都藏在瓶子里。
console.log(_.add(Maybe.of(2), Maybe.of(3)))
// NaN

// 使用可靠的 map 函数试试
const container_of_add_2 = map(_.add, Maybe.of(2));
console.log(container_of_add_2)
// Maybe(add(2))

console.log(
    map((two) => {
        return Maybe.of(3).map(_.add(two))
    }, Maybe.of(2)).join()
)

console.log(
    Maybe.of(2).chain(function (two) {
        return Maybe.of(3).map(_.add(two));
    })
)


/////////////////////// ap
console.log(Maybe.of(_.add(2)).ap(Maybe.of(3)))

console.log(Maybe.of(2).map(_.add).ap(Maybe.of(3)))


/////////////////////// F.of(x).map(f) == F.of(f).ap(F.of(x))
// F.of(x).map(f) == F.of(f).ap(F.of(x))
// F.of(f(x)) == F.of(x).map(f)
// F.of(f(x)) == F.of(f(x))

console.log(Maybe.of(2).map(_.add)) // Maybe(add(2))
console.log(Maybe.of(_.add).ap(Maybe.of(2))) // Maybe(add(2))

console.log(Maybe.of(_.add).ap(Maybe.of(2)).ap(Maybe.of(3)))
// Maybe(5)

Task.of(_.add).ap(Task.of(2)).ap(Task.of(3)).fork(console.log, console.log)
// Task(5)


///////////////////////
// // Http.get :: String -> Task Error HTML

// var renderPage = curry(function(destinations, events) { /* render page */  });

// Task.of(renderPage).ap(Http.get('/destinations')).ap(Http.get('/events'))
// // Task("<div>some page with dest and events</div>")


///////////////////////

// 从 chain 衍生出的 map
X.prototype.map = function (f) {
    var m = this;
    return m.chain(function (a) {
        return m.constructor.of(f(a));
    });
}

// 从 chain/map 衍生出的 ap
X.prototype.ap = function (other) {
    return this.chain(function (f) {
        return other.map(f);
    });
};


