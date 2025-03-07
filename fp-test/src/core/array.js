import { curry, compose } from "./function.js";

/* ARRAY */

export const map = curry(function(fp, arr) {
    return arr.map(fp);
})

export const reduce = curry(function(fp, init, arr) {
    return arr.reduce(fp, init);
})

export const filter = curry(function(fp, arr) {
    return arr.filter(fp);
})