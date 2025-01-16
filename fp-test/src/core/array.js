

import { curry, compose } from "./function.js";

/* ARRAY */
export const map = curry((fp, arr) => {
    return arr.map(fp);
})

export const reduce = curry((fp, init, arr) => {
    return arr.reduce(fp, init);
})

export const filter = curry((fp, arr) => {
    return arr.filter(fp);
})