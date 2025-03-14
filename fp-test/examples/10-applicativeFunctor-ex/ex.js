import * as _ from "ramda"
import { Either, IO, Maybe, compose, chain, map, curry, toUpperCase, Right, Left, either, left, liftA2 } from "../../utils/support.js";
import Task from "data.task";

// fib browser for test
var localStorage = {};



// Exercise 1
// ==========
// Write a function that add's two possibly null numbers together using Maybe and ap()

//  ex1 :: Number -> Number -> Maybe Number
var ex1 = function (x, y) {
    return Maybe.of(_.add).ap(Maybe.of(x)).ap(Maybe.of(y));
    // return liftA2(_.add, Maybe.of(x), Maybe.of(y));
};


// Exercise 2
// ==========
// Now write a function that takes 2 Maybe's and adds them. Use liftA2 instead of ap().

//  ex2 :: Maybe Number -> Maybe Number -> Maybe Number
// var ex2 = (x, y) =>liftA2(_.add, x, y);
var ex2 = liftA2(_.add)



// Exercise 3
// ==========
// Run both getPost(n) and getComments(n) then render the page with both. (the n arg is arbitrary)
var makeComments = _.reduce(function (acc, c) { return acc + "<li>" + c + "</li>" }, "");
var render = _.curry(function (p, cs) { return "<div>" + p.title + "</div>" + makeComments(cs); });

//  ex3 :: Task Error HTML
var ex3 = Task.of(render).ap(getPost(1)).ap(getComments(1))



// Exercise 4
// ==========
// Write an IO that gets both player1 and player2 from the cache and starts the game
localStorage.player1 = "toby";
localStorage.player2 = "sally";

var getCache = function (x) {
    return new IO(function () { return localStorage[x]; });
}
var game = _.curry(function (p1, p2) { return p1 + ' vs ' + p2; });

//  ex4 :: IO String
// var ex4 = IO.of(game).ap(getCache("player1")).ap(getCache("player2"));
const ex4 = liftA2(game, getCache("player1"), getCache("player2"))





// TEST HELPERS
// =====================

function getPost(i) {
    return new Task(function (rej, res) {
        setTimeout(function () { res({ id: i, title: 'Love them tasks' }); }, 300);
    });
}

function getComments(i) {
    return new Task(function (rej, res) {
        setTimeout(function () {
            res(["This book should be illegal", "Monads are like space burritos"]);
        }, 300);
    });
}

export { ex1, ex2, ex3, ex4 }