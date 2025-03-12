import * as _ from "ramda"
import { Maybe, IO, either, Either, left, map, toUpperCase } from "../../utils/support.js";
// const Task = require('data.task');
import Task from "data.task"

// Exercise 1
// ==========
// Use _.add(x,y) and _.map(f,x) to make a function that increments a value inside a functor

const ex1 = _.map(_.add(1));



// Exercise 2
// ==========
// Use _.head to get the first element of the list
// var xs = Identity.of(['do', 'ray', 'me', 'fa', 'so', 'la', 'ti', 'do']);

const ex2 = _.map(_.head());



// Exercise 3
// ==========
// Use safeProp and _.head to find the first initial of the user
const safeProp = _.curry(function (x, o) { return Maybe.of(o[x]); });

// const user = { id: 2, name: "Albert" };

const ex3 = _.compose(_.map(_.head()), safeProp('name'));



// Exercise 4
// ==========
// Use Maybe to rewrite ex4 without an if statement

// var ex4 = function (n) {
//   if (n) { return parseInt(n); }
// };

const ex4_1 = (n) => {
    return Maybe.of(parseInt(n));
};

const ex4_2 = _.compose(_.map(parseInt), Maybe.of);


// Exercise 5
// ==========
// Write a function that will getPost then _.toUpper the post's title

// getPost :: Int -> Future({id: Int, title: String})
var getPost = function (i) {
    return new Task(function (rej, res) {
        setTimeout(function () {
            res({ id: i, title: 'Love them futures' })
        }, 300)
    });
};

var ex5 = _.compose(_.map(_.compose(_.toUpper, _.prop('title'))), getPost);

// console.log(ex5(13))
// console.log(ex5(13).fork(console.log, function (res) {
//     console.log(res)
// }))
// ex5(13).fork(console.log, function (res) {
//     console.log(res)
// })

// Exercise 6
// ==========
// Write a function that uses checkActive() and showWelcome() to grant access or return the error

const showWelcome = _.compose(_.concat("Welcome "),  _.prop('name'));

const checkActive = function (user) {
    console.log(Either.of(user));
    return user.active ? Either.of(user) : left('Your account is not active')
};

const ex6 = _.compose(_.map(showWelcome), checkActive);



// Exercise 7
// ==========
// Write a validation function that checks for a length > 3. It should return Right(x) if it is greater than 3 and Left("You need > 3") otherwise

var ex7 = function (x) {
    return x.length > 3 ? Either.of(x) : left("You need > 3");  // <--- write me. (don't be pointfree)
};



// Exercise 8
// ==========
// Use ex7 above and Either as a functor to save the user if they are valid or return the error message string. Remember either's two arguments must return the same type.

var save = function (x) {
    return new IO(function () {
        console.log("SAVED USER!");
        return x + '-saved';
    });
};

// var ex8 = _.compose(_.map(save), ex7);

// Remember either's two arguments must return the same type.
const ex8 = _.compose(either(IO.of, save), ex7); // 因为要输出同样的类型，如果是map， Right -> IO, Left -> String

export { ex1, ex2, ex3, ex4_1, ex4_2, ex5, ex6, ex7, ex8 };