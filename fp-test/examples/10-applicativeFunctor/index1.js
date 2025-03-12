import * as _ from "ramda"
import { Either, IO, Maybe, Identity, compose, chain, map, curry, toUpperCase, Right, Left, either, left, liftA2 } from "../../utils/support.js";
import Task from "data.task";

var tOfM = compose(Task.of, Maybe.of);

// liftA2(_.concat, tOfM('Rainy Days and Mondays'), tOfM(' always get me down')).fork(console.log, console.log);
// Task(Maybe(Rainy Days and Mondays always get me down))

liftA2(liftA2(_.concat), tOfM('Rainy Days and Mondays'), tOfM(' always get me down')).fork(console.log, console.log);
// Task(Maybe(Rainy Days and Mondays always get me down))


////////////////////////////////// 同一律
// A.of(id).ap(v) == v

// Maybe.of(id).ap(v) 
// = v.map(id)
// = v

var v = Identity.of("Pillow Pets");
Identity.of(_.identity).ap(v) == v
console.log(Identity.of(_.identity).ap(v))

// 同一律不允许跨 Applicative 类型混用（如 Identity 操作 Maybe）
var z = Maybe.of("混着用也行吗？")
Identity.of(_.identity).ap(z) == z
console.log(Identity.of(_.identity).ap(z))


////////////////////////////////// 同态
// A.of(f).ap(A.of(x)) == A.of(f(x))

// Maybe.of(f).ap(Maybe.of(x))
// = Maybe.of(x).map(f)
// = Maybe.of(f(x))

Either.of(_.toUpper).ap(Either.of("oreos")) == Either.of(_.toUpper("oreos"))
console.log(Either.of(_.toUpper).ap(Either.of("oreos")))

// Either.of(_.toUpper).ap(Either.of("oreos")) // 用ap的前提是调用ap的对象里包裹的是function  （1）
// Right(_.toUpper).ap(Right("oreos"))
// Right("oreos").map(_.toUpper)
// Either.of(_.toUpper("oreos"))


////////////////////////////////// 互换
// v.ap(A.of(x)) == A.of(function(f) { return f(x) }).ap(v)


var v = Task.of(_.reverse);
var x = 'Sparklehorse';

v.ap(Task.of(x)) == Task.of(function(f) { return f(x) }).ap(v)  // 参考（1）点感悟， Task.of(x) 是没法用ap的， [1] 可证明

// v.ap(Task.of(x)) 推理过程
// v.ap(Task.of(x)) = Task(_.reverse).chain(fn => Task.of('Sparklehorse').map(fn))
new Task((reject_, resolve) => ((rej,res) => res(_.reverse))(reject_, x => (fn => Task.of('Sparklehorse').map(fn))(x).fork(reject_, resolve)));
new Task((reject_, resolve) => (x => (fn => Task.of('Sparklehorse').map(fn))(x).fork(reject_, resolve))(_.reverse));
new Task((reject_, resolve) => (fn => Task.of('Sparklehorse').map(fn))(_.reverse).fork(reject_, resolve));
new Task((reject_, resolve) => Task.of('Sparklehorse').map(_.reverse).fork(reject_, resolve));
Task((reject_, resolve) => Task.of('Sparklehorse').map(_.reverse).fork(reject_, resolve));

    // // Task.of('Sparklehorse').map(_.reverse) ==
    // new Task((reject_, resolve) => this.fork(reject_, compose(resolve, _.reverse)));
    // Task((reject_, resolve) => ((rej,res) => res('Sparklehorse'))(reject_, compose(resolve, _.reverse)));
    // Task((reject_, resolve) => compose(resolve, _.reverse)('Sparklehorse'));
    // Task((reject_, resolve) => resolve(_.reverse('Sparklehorse')));

// => Task((reject_, resolve) => Task((reject_1, resolve1) => resolve1(_.reverse('Sparklehorse'))).fork(reject_, resolve));
// => Task((reject_, resolve) => (reject_1, resolve1) => resolve1(_.reverse('Sparklehorse'))(reject_, resolve));
// => Task((reject_, resolve) => resolve(_.reverse('Sparklehorse'));


// Task.of(function(f) { return f(x) }).ap(v) 推理过程
// Task.of(function(f) { return f(x) }).ap(v) = Task(function(f) { return f(x) }).chain(fn => Task.of(_.reverse).map(fn))
new Task((reject_, resolve) => ((rej,res) => res(function(f) { return f(x) }))(reject_, fn => Task.of(_.reverse).map(fn).fork(reject_, resolve)));
new Task((reject_, resolve) => (fn => Task.of(_.reverse).map(fn).fork(reject_, resolve))(function(f) { return f(x) }));
new Task((reject_, resolve) => Task.of(_.reverse).map(function(f) { return f(x) }).fork(reject_, resolve)); // [1] 如果 x 不构造成 function(f) { return f(x) }， 这里map就会报错
new Task((reject_, resolve) => Task.of(_.reverse).map(function(f) { return f(x) }).fork(reject_, resolve));

    // Task.of(_.reverse).map(function(f) { return f(x) })
    // new Task((reject_, resolve) => this.fork(reject_, compose(resolve, function(f) { return f(x) })));
    // Task((reject_, resolve) => ((rej,res) => res(_.reverse))(reject_, compose(resolve, function(f) { return f(x) })));
    // Task((reject_, resolve) => compose(resolve, function(f) { return f(x) })(_.reverse));
    // Task((reject_, resolve) => resolve(_.reverse('Sparklehorse'));

// 后续推理同上

// fn是函数，a是值
// 感悟（2） Task.of(fn).ap(Task.of(a)) => Task.of(fn(a)), 互换：Task.of(function(f) { return f(a) }).ap(Task.of(fn)) => Task.of((function(f) { return f(a) })(fn)) => Task.of(fn(a))
// 其本质是 就是得到 fn(a) 而已


////////////////////////////////// 组合
A.of(compose).ap(u).ap(v).ap(w) == u.ap(v.ap(w));

var u = IO.of(_.toUpper);
var v = IO.of(_.concat("& beyond"));
var w = IO.of("blood bath ");

IO.of(_.compose).ap(u).ap(v).ap(w) == u.ap(v.ap(w))

IO.of(_.compose).ap(u).ap(v).ap(w)
= new IO(() => _.compose).chain(fn => u.map(fn)).ap(v).ap(w)
= IO(() => _.compose).map(fn => u.map(fn)).join().ap(v).ap(w)
= new IO(_.compose(fn => u.map(fn), () => _.compose)).join().ap(v).ap(w)
= compose(fn => u.map(fn), () => _.compose)().ap(v).ap(w) // 这里用 join 是用的 return this.unsafePerformIO()，便于推导
= u.map(_.compose).ap(v).ap(w)
= new IO(compose(_.compose, u.unsafePerformIO)).ap(v).ap(w)

= IO(compose(_.compose, u.unsafePerformIO)).chain(fn => v.map(fn)).ap(w)
= IO(compose(_.compose, u.unsafePerformIO)).map(fn => v.map(fn)).join().ap(w)
= new IO(compose(fn => v.map(fn), compose(_.compose, u.unsafePerformIO))).join().ap(w)
= compose(fn => v.map(fn), compose(_.compose, u.unsafePerformIO))().ap(w)
= v.map(_.compose(u.unsafePerformIO())).ap(w)
= new IO(compose(_.compose(u.unsafePerformIO()), v.unsafePerformIO)).ap(w)

= IO(compose(_.compose(u.unsafePerformIO()), v.unsafePerformIO)).chain(fn => w.map(fn))
= IO(compose(_.compose(u.unsafePerformIO()), v.unsafePerformIO)).map(fn => w.map(fn)).join()
= new IO(compose(fn => w.map(fn), compose(_.compose(u.unsafePerformIO()), v.unsafePerformIO))).join()
= compose(fn => w.map(fn), compose(_.compose(u.unsafePerformIO()), v.unsafePerformIO))()
= w.map(compose(_.compose(u.unsafePerformIO()), v.unsafePerformIO)())
= w.map(compose(_.compose(u.unsafePerformIO(), v.unsafePerformIO())))
= new IO(compose(compose(_.compose(u.unsafePerformIO(), v.unsafePerformIO())), w.unsafePerformIO))

// A.of(compose).ap(u).ap(v).ap(w).unsafePerformIO()
compose(compose(_.compose(u.unsafePerformIO(), v.unsafePerformIO())), w.unsafePerformIO)()
// 1--> compose(_.compose(u.unsafePerformIO(), v.unsafePerformIO()))(w.unsafePerformIO())
// 2--> _.compose(u.unsafePerformIO(), v.unsafePerformIO())(w.unsafePerformIO())



var u = Maybe.of(_.toUpper);
var v = Maybe.of(_.concat("& beyond"));
var w = Maybe.of("blood bath ");

Maybe.of(compose).ap(Maybe.of(_.toUpper)).ap(Maybe.of(_.concat("& beyond"))).ap(Maybe.of("blood bath "))
= Maybe.of(_.toUpper).map(compose).ap(Maybe.of(_.concat("& beyond"))).ap(Maybe.of("blood bath "))
= Maybe.of(compose(_.toUpper)).ap(Maybe.of(_.concat("& beyond"))).ap(Maybe.of("blood bath "))
= Maybe.of(_.concat("& beyond")).map(compose(_.toUpper)).ap(Maybe.of("blood bath "))
= Maybe.of(compose(_.toUpper)(_.concat("& beyond"))).ap(Maybe.of("blood bath "))
= Maybe.of("blood bath ").map(compose(_.toUpper)(_.concat("& beyond")))
= Maybe.of(compose(_.toUpper)(_.concat("& beyond"))("blood bath "))