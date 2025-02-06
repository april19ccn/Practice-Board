// var Container = function (x) {
//     this.__value = x;
// }
// 该写法在node环境里会使 Container.of(3) 输出 { __value: 3 }

import util from 'util'
import * as R from 'ramda'

class Container {
    constructor(value) { this.__value = value; }
    // 自定义 inspect 方法
    [util.inspect.custom]() { return `Container(${this.__value})`; }
}

Container.of = function (x) { return new Container(x); };

// console.log(new Container(3)); // 输出 Container(3) 而非 { __value: 3 }

console.log(Container.of(3))
//=> Container(3)

console.log(Container.of("hotdogs"))
//=> Container("hotdogs")

console.log(Container.of(Container.of({ name: "yoda" })))
//=> Container(Container({name: "yoda" }))


//////////////////////////////////////// 第一个 functor
// (a -> b) -> Container a -> Container b
Container.prototype.map = function (f) {
    return Container.of(f(this.__value))
}

console.log(Container.of(2).map(function (two) { return two + 2 }))
//=> Container(4)

console.log(Container.of("flamethrowers").map(function (s) { return s.toUpperCase() }))
//=> Container("FLAMETHROWERS")

console.log(Container.of("bombs").map(R.concat(' away')).map(R.prop('length')))
//=> Container(10)

try {
    console.log(Container.of(null).map(R.match(/a/ig)))
    //=> Cannot read property 'match' of null
} catch (e) {
    console.log(e)
}


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


console.log(Maybe.of("Malkovich Malkovich").map(R.match(/a/ig)))
//=> Maybe(['a', 'a'])

console.log(Maybe.of(null).map(R.match(/a/ig)))
//=> Maybe(null)

console.log(Maybe.of({name: "Boris"}).map(R.prop("age")).map(R.add(10)))
//=> Maybe(null)

console.log(Maybe.of({name: "Dinah", age: 14}).map(R.prop("age")).map(R.add(10)))
//=> Maybe(24)


//  map :: Functor f => (a -> b) -> f a -> f b
var pointfree_map = R.curry(function(f, any_functor_at_all) {
    return any_functor_at_all.map(f);
    // return any_functor_at_all.of(f(any_functor_at_all.__value)) // 不能这么写 1. Maybe.of 是静态方法，如果写成实例方法占内存 2. 不优雅
});

// 改写成 pointfree 写法
console.log(
    R.compose(
        pointfree_map(R.match(/a/ig))
    )(Maybe.of("Malkovich Malkovich"))
)

console.log(
    R.compose(
        pointfree_map(R.add(10)),
        pointfree_map(R.prop("age"))
    )(Maybe.of({name: "Dinah", age: 14}))
)