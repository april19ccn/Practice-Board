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

export default Container