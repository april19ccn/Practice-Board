// 自己测一点东西
import { Either, IO, Maybe, Identity, compose, chain, map, curry, toUpperCase, Right, Left, either, left, liftA2 } from "../../utils/support.js";
import * as R from 'ramda'

console.log(R.compose(Math.abs, R.add(1), R.multiply(2))(-4)) //=> 7

console.log((R.concat(R.__, "& beyond"))("blood bath "))

console.log(R.compose(R.toUpper, R.concat(R.__, "& beyond"))("blood bath "))

console.log(Maybe.of(R.toUpper).ap(Maybe.of(R.concat(R.__, "& beyond")).ap(Maybe.of("blood bath "))))

// 柯里化compose
const compose3 = f => g => x => f(g(x));
console.log(Maybe.of(compose3).ap(Maybe.of(R.toUpper)).ap(Maybe.of(R.concat(R.__, "& beyond"))).ap(Maybe.of("blood bath ")))

// console.log(R.compose(R.toUpper)(R.concat("& beyond"))("blood bath "))

// console.log(R.compose(Math.abs)(R.add(1)))


// const u = (d) => Math.abs(d)
// const v = (j) => R.add(1)(j)
// const w = (z) => R.multiply(2)(z)
// console.log(R.compose(R.compose, u, v, w)(-4)())

// console.log(Math.abs(R.add(1, R.multiply(2)(-4))))