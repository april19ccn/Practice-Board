import * as R from "ramda"
import Container from "./01-container.js"

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
