import * as R from "ramda"

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

console.log(Maybe.of({ name: "Boris" }).map(R.prop("age")).map(R.add(10)))
//=> Maybe(null)

console.log(Maybe.of({ name: "Dinah", age: 14 }).map(R.prop("age")).map(R.add(10)))
//=> Maybe(24)


// 本质是拓宽了map的使用范围，之前map是仅供数组这个特定functor使用
//  map :: Functor f => (a -> b) -> f a -> f b
var pointfree_map = R.curry(function (f, any_functor_at_all) {
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
    )(Maybe.of({ name: "Dinah", age: 14 }))
)

// 直接使用R.map
console.log(
    "R.map",
    R.map(R.match(/a/ig))(Maybe.of("Malkovich Malkovich"))
)
console.log(
    "R.map",
    R.compose(
        R.map(R.add(10)),
        R.map(R.prop("age"))
    )(Maybe.of({ name: "Dinah", age: 14 }))
)

//////////////////////////////////////// 用例
//  safeHead :: [a] -> Maybe(a)
var safeHead = function (xs) {
    return Maybe.of(xs[0]);
};

var streetName = R.compose(pointfree_map(R.prop('street')), safeHead, R.prop('addresses'));

console.log("streetName ==> ", streetName);

console.log(streetName({ addresses: [] }));
// Maybe(null)

console.log(streetName({ addresses: [{ street: "Shady Ln.12", number: 420110 }] }));
// Maybe("Shady Ln.12")

// ramda:map 若第二个参数自身存在 map 方法，则调用自身的 map 方法。
var streetName1 = R.compose(R.map(R.prop('street')), safeHead, R.prop('addresses'));

console.log(streetName1({ addresses: [] }));
// Maybe(null)

console.log(streetName1({ addresses: [{ street: "Shady Ln.", number: 4201 }] }));
// Maybe("Shady Ln.")