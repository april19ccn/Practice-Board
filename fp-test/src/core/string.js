import { curry, compose } from "./function.js";

/* STRING */

// 检索字符串与正则表达式进行匹配的结果
export const match = curry((what, str) => {
    return str.match(what);
});

// 接受一个模式，通过搜索模式将字符串分割成一个有序的子串列表，将这些子串放入一个数组，并返回该数组。
export const split = curry((what, str) =>{
    return str.split(what);
})
