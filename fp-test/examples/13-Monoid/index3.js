// reduce :: (b -> a -> b) -> b -> [a] -> b
// const reduce = curry((fn, zero, xs) => xs.reduce(fn, zero));
import { reduce, concat, compose, reverse, append, identity } from "../../utils/support.js";

const sort = arr => [...arr].sort(); // 默认按字符串 Unicode 排序

// 将加法抽象化
const Sum = (x) => ({
    x,
    concat: (other) => Sum(x + other.x),
})
Sum.empty = () => Sum(0)

// concat :: Semigroup s => s -> s -> s
// const concat = x => y => x.concat(y)

console.log(Sum(1).concat(Sum(2)))
console.log(concat(Sum(1))(Sum(2)))
console.log(concat(Sum(1), Sum(2)))
console.log([Sum(1), Sum(2)].reduce(concat)) // Sum(3)

// fold :: Monoid m => m -> [m] -> m
const fold = reduce(concat)

console.log(fold(Sum.empty(), [Sum(1), Sum(2)]))



const Endo = run => ({
    run,
    concat: other =>
        Endo(compose(run, other.run))
})

Endo.empty = () => Endo(identity)


// in action

// thingDownFlipAndReverse :: Endo [String] -> [String]
const thingDownFlipAndReverse = fold(Endo.empty(), [Endo(reverse), Endo(sort), Endo(append('thing down'))])

// [Endo(reverse), Endo(sort), Endo(append('thing down'))].reduce(concat, Endo(() => []))
// concat(Endo(() => []), Endo(reverse)) => Endo(() => []).concat(Endo(reverse)) => function (Endo(reverse)) { return Endo(compose(() => [], reverse)) }

console.log(thingDownFlipAndReverse.run(['let me work it', 'is it worth it?']))
// ['thing down', 'let me work it', 'is it worth it?']