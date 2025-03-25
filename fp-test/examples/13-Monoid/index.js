import { IO, Maybe, Either, Identity, left, compose, chain, map, sequence, traverse } from "../../utils/support.js";

// =============================================================================

// 将加法抽象化
const Sum = (x) => ({
    x,
    concat: (other) => Sum(x + other.x),
})

console.log(Sum(1).concat(Sum(3))) // Sum(4)
console.log(Sum(4).concat(Sum(37))) // Sum(41)

// 拓展
const Product = (x) => ({ x, concat: (other) => Product(x * other.x) })

const Min = (x) => ({ x, concat: (other) => Min(x < other.x ? x : other.x) })

const Max = (x) => ({ x, concat: (other) => Max(x > other.x ? x : other.x) })

// 再拓展
const Any = (x) => ({ x, concat: (other) => Any(x || other.x) })
const All = (x) => ({ x, concat: (other) => All(x && other.x) })

console.log(Any(false).concat(Any(true))) // Any(true)
console.log(Any(false).concat(Any(false))) // Any(false)

console.log(All(false).concat(All(true))) // All(false)
console.log(All(true).concat(All(true))) // All(true)

console.log([(1, 2)].concat([3, 4])) // [1,2,3,4]

console.log('miracle grow'.concat('n')) // miracle grown"

// console.log(Map({ day: 'night' }).concat(Map({ white: 'nikes' }))) // Map({day: 'night', white: 'nikes'})


// =============================================================================

Identity.prototype.concat = function (other) {
	return new Identity(this.$value.concat(other.$value))
}

console.log(Identity.of(Sum(4)).concat(Identity.of(Sum(1)))) // Identity(Sum(5))
console.log(Identity.of(4).concat(Identity.of(1))) // TypeError: this.$value.concat is not a function



// formValues :: Selector -> IO (Map String String)
// validate :: Map String String -> Either Error (Map String String)

formValues('#signup').map(validate).concat(formValues('#terms').map(validate)) // IO(Right(Map({username: 'andre3000', accepted: true})))
formValues('#signup').map(validate).concat(formValues('#terms').map(validate)) // IO(Left('one must accept our totalitarian agreement'))

serverA.get('/friends').concat(serverB.get('/friends')) // Task([friend1, friend2])

// loadSetting :: String -> Task Error (Maybe (Map String Boolean))
loadSetting('email').concat(loadSetting('general')) // Task(Maybe(Map({backgroundColor: true, autoSave: false})))