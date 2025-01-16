import * as fp from "../index.js";

// 练习 1
//==============
// 通过局部调用（partial apply）移除所有参数

// var words = function(str) {
//     return split(' ', str);
//   };

const words = fp.split(' ');

// 练习 1a
//==============
// 使用 `map` 创建一个新的 `words` 函数，使之能够操作字符串数组

// var sentences = undefined;

const sentences = fp.map(words);

// 练习 2
//==============
// 通过局部调用（partial apply）移除所有参数

// var filterQs = function(xs) {
//     return filter(function(x){ return match(/q/i, x);  }, xs);
// };

const filterQs = fp.filter(fp.match(/q/i));

// 练习 3
//==============
// 使用帮助函数 `_keepHighest` 重构 `max` 使之成为 curry 函数

// 无须改动:
const _keepHighest = function (x, y) { return x >= y ? x : y; };

// 重构这段代码:
// var max = function(xs) {
//   return reduce(function(acc, x){
//     return _keepHighest(acc, x);
//   }, -Infinity, xs);
// };

const max = fp.reduce(_keepHighest, -Infinity);


// 彩蛋 1:
// ============
// 包裹数组的 `slice` 函数使之成为 curry 函数
// //[1,2,3].slice(0, 2)
const slice = fp.curry((start, end, arr) => {
    return arr.slice(start, end);
});


// 彩蛋 2:
// ============
// 借助 `slice` 定义一个 `take` curry 函数，该函数调用后可以取出字符串的前 n 个字符。
var take = slice(0);


export { words, sentences, filterQs, max, slice, take };