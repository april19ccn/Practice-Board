import * as _ from "ramda"
import { Either, IO, Maybe, Identity, compose, chain, map, curry, toUpperCase, Right, Left, either, left, liftA2 } from "../../utils/support.js";
import Task from "data.task";

var tOfM = compose(Task.of, Maybe.of);

// liftA2(_.concat, tOfM('Rainy Days and Mondays'), tOfM(' always get me down')).fork(console.log, console.log);
// Task(Maybe(Rainy Days and Mondays always get me down))

liftA2(liftA2(_.concat), tOfM('Rainy Days and Mondays'), tOfM(' always get me down')).fork(console.log, console.log);
// Task(Maybe(Rainy Days and Mondays always get me down))


// 同一律
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


// 同态
// A.of(f).ap(A.of(x)) == A.of(f(x))

// Maybe.of(f).ap(Maybe.of(x))
// = Maybe.of(x).map(f)
// = Maybe.of(f(x))

Either.of(_.toUpper).ap(Either.of("oreos")) == Either.of(_.toUpper("oreos"))
console.log(Either.of(_.toUpper).ap(Either.of("oreos")))
