import * as R from "ramda"
import { Task, Either, left, map, toUpperCase } from "../../utils/support.js";
import Container from "./01-container.js"

Container.prototype.map = function (f) {
    return Container.of(f(this.__value))
}
//////////

// identity
// map(R.identity) === R.identity;

// composition
// compose(map(f), map(g)) === map(compose(f, g));

var idLaw1 = R.map(R.identity);
var idLaw2 = R.identity;

console.log(idLaw1(Container.of(2)));
//=> Container(2)

console.log(idLaw2(Container.of(2)));
//=> Container(2)


var compLaw1 = R.compose(R.map(R.concat(R.__, " world")), R.map(R.concat(R.__, " cruel")));
var compLaw2 = R.map(R.compose(R.concat(R.__, " world"), R.concat(R.__, " cruel")));

console.log(compLaw1(Container.of("Goodbye")));
//=> Container('Goodbye cruel world')

console.log(compLaw2(Container.of("Goodbye")));
//=> Container('Goodbye cruel world')



//  topRoute :: String -> Container(String)
var topRoute = R.compose(Container.of, R.reverse);

//  bottomRoute :: String -> Container(String)
var bottomRoute = R.compose(R.map(R.reverse), Container.of);


console.log(topRoute("hi"));
// Container("ih")

console.log(bottomRoute("hi"));
// Container("ih")



var nested = Task.of([Either.of('pillows'), left('no sleep for you')]);
console.log(map(map(map(toUpperCase)), nested));
// Task([Right("PILLOWS"), Left("no sleep for you")])
// 左1 map: Task.map
// 左2 map: Array.map
// 左3 map: Either.map / Left.map

// [ => 
nested => Task.of([Either.of("pillows"), left("no sleep for you")]);
// => 
Task((_, resolve) => resolve([Either.of("pillows"), left("no sleep for you")]));
this.fork = (_, resolve) => resolve([Either.of("pillows"), left("no sleep for you")]);
// ]

// [ => 
map(map(map(toUpperCase)), nested);
// => 
nested.map(map(map(toUpperCase)));
// =>
new Task((reject_, resolve) => this.fork(reject_, compose(resolve, map(map(toUpperCase)))));
// =>
new Task((reject_, resolve) => compose(resolve, map(map(toUpperCase)))([Either.of("pillows"), left("no sleep for you")]));
// =>
Task((reject_, resolve) => compose(resolve, map(map(toUpperCase)))([Either.of("pillows"), left("no sleep for you")]));
this.fork = (reject_, resolve) => compose(resolve, map(map(toUpperCase)))([Either.of("pillows"), left("no sleep for you")]);
// ]

(reject_, resolve) => compose(resolve, map(map(toUpperCase)))([Either.of("pillows"), left("no sleep for you")]);
//  执行过程
(reject_, resolve) => compose(resolve, [Either.of("pillows"), left("no sleep for you")].map(map(toUpperCase)));

(reject_, resolve) => compose(resolve, [Either.of("pillows").map(toUpperCase), left("no sleep for you").map(toUpperCase)]);

(reject_, resolve) => resolve([Either.of("pillows").map(toUpperCase), left("no sleep for you").map(toUpperCase)]);

(reject_, resolve) => resolve([Right("PILLOWS"), Left("no sleep for you")]);



//////////////////////// 组合 functor 
class Compose {
    constructor(fgx) {
        this.getCompose = fgx;
    }

    static of(fgx) {
        return new Compose(fgx);
    }

    map(fn) {
        return new Compose(map(map(fn), this.getCompose));
    }
}

const tmd = Task.of(Maybe.of('Rock over London'));

const ctmd = Compose.of(tmd);
// Compose(Task(Just('Rock over London')))

const ctmd2 = map(append(', rock on, Chicago'), ctmd);
// ctmd.map(append(', rock on, Chicago'));
// new Compose(map(map(append(', rock on, Chicago')), Task(Just('Rock over London'))))
// Compose(Task(Just('Rock over London, rock on, Chicago')))

ctmd2.getCompose;
// Task(Just('Rock over London, rock on, Chicago'))