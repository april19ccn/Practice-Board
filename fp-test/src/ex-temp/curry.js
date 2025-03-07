const { sum } = require("lodash");

const sum3 = (x, y, z) => {
    return x + y + z;
}

console.log(sum3(1, 2, 3))

const sum3_1 = (x) => {
    return (y) => {
        return (z) => {
            return x + y + z;
        }
    }
}

console.log(sum3_1(1)(2)(3))


///////////////////////

const sub_curry = function (fn) {
    const args = [].slice.call(arguments, 1);
    return function() {
        const newArgs = args.concat([].slice.call(arguments));
        return fn.apply(this, newArgs);
    };
};

const curry = function (fn, length) {
    length = length || fn.length;

    const slice = Array.prototype.slice

    return function () {
        if (arguments.length < length) {
            var combined = [fn].concat(slice.call(arguments));
            return curry(sub_curry.apply(this, combined), length - arguments.length);
        } else {
            return fn.apply(this, arguments);
        }
    }
}

const fn = curry(function(a, b, c){
    return [a, b, c];
})

// fn("a", "b", "c");
fn("a", "b")("c");
fn("a")("b", "c");
fn("a")("b")("c");

