// 自己测一点东西

import * as R from 'ramda'

console.log(R.compose(Math.abs, R.add(1), R.multiply(2))(-4)) //=> 7

console.log(R.compose(Math.abs)(R.add(1)))


const u = (d) => Math.abs(d)
const v = (j) => R.add(1)(j)
const w = (z) => R.multiply(2)(z)
console.log(R.compose(R.compose, u, v, w)(-4)())

// console.log(Math.abs(R.add(1, R.multiply(2)(-4))))