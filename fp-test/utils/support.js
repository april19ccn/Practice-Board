/* LICENSE MIT-2.0 - @MostlyAdequate */
/* eslint-disable no-use-before-define, max-len, class-methods-use-this */

// const util = require('util');

// always :: a -> b -> a
const always = curry((a, b) => a);


// compose :: ((a -> b), (b -> c),  ..., (y -> z)) -> a -> z
// const compose = (...fns) => (...args) => fns.reduceRight((res, fn) => [fn.call(null, ...res)], args)[0];

// NOTE A slightly pumped up version of `compose` which also keeps track of the chain
// of callees. In the end, a function created with `compose` holds a `callees` variable
// with the list of all the callees' names.
// This is useful to provide insights during validation of exercises
function compose(...fns) {
    const n = fns.length;
  
    return function $compose(...args) {
      $compose.callees = [];
  
      let $args = args;
  
      for (let i = n - 1; i >= 0; i -= 1) {
        const fn = fns[i];
  
        // assert(
        //   typeof fn === 'function',
        //   `Invalid Composition: ${ordinal(n - i)} element in a composition isn't a function`,
        // );
  
        $compose.callees.push(fn.name);
        $args = [fn.call(null, ...$args)];
      }
  
      return $args[0];
    };
}

// curry :: ((a, b, ...) -> c) -> a -> b -> ... -> c
function curry(fn) {
  const arity = fn.length;

  return function $curry(...args) {
    if (args.length < arity) {
      return $curry.bind(null, ...args);
    }

    return fn.call(null, ...args);
  };
}


// either :: (a -> c) -> (b -> c) -> Either a b -> c
const either = curry((f, g, e) => {
  if (e.isLeft) {
    return f(e.$value);
  }

  return g(e.$value);
});


// identity :: x -> x
const identity = x => x;


// inspect :: a -> String
const inspect = (x) => {
  if (x && typeof x.inspect === 'function') {
    return x.inspect();
  }

  function inspectFn(f) {
    return f.name ? f.name : f.toString();
  }

  function inspectTerm(t) {
    switch (typeof t) {
      case 'string':
        return `'${t}'`;
      case 'object': {
        const ts = Object.keys(t).map(k => [k, inspect(t[k])]);
        return `{${ts.map(kv => kv.join(': ')).join(', ')}}`;
      }
      default:
        return String(t);
    }
  }

  function inspectArgs(args) {
    return Array.isArray(args) ? `[${args.map(inspect).join(', ')}]` : inspectTerm(args);
  }

  return (typeof x === 'function') ? inspectFn(x) : inspectArgs(x);
};


// left :: a -> Either a b
const left = a => new Left(a);


// liftA2 :: (Applicative f) => (a1 -> a2 -> b) -> f a1 -> f a2 -> f b
const liftA2 = curry((fn, a1, a2) => a1.map(fn).ap(a2));


// liftA3 :: (Applicative f) => (a1 -> a2 -> a3 -> b) -> f a1 -> f a2 -> f a3 -> f b
const liftA3 = curry((fn, a1, a2, a3) => a1.map(fn).ap(a2).ap(a3));


// maybe :: b -> (a -> b) -> Maybe a -> b
const maybe = curry((v, f, m) => {
  if (m.isNothing) {
    return v;
  }

  return f(m.$value);
});


const createCompose = curry((F, G) => class Compose {
  constructor(x) {
    this.$value = x;
  }

//   [util.inspect.custom]() {
//     return `Compose(${inspect(this.$value)})`;
//   }

  // ----- Pointed (Compose F G)
  static of(x) {
    return new Compose(F(G(x)));
  }

  // ----- Functor (Compose F G)
  map(fn) {
    return new Compose(this.$value.map(x => x.map(fn)));
  }

  // ----- Applicative (Compose F G)
  ap(f) {
    return f.map(this.$value);
  }
});


class Either {
  constructor(x) {
    this.$value = x;
  }

  // ----- Pointed (Either a)
  static of(x) {
    return new Right(x);
  }
}


class Left extends Either {
  get isLeft() {
    return true;
  }

  get isRight() {
    return false;
  }

  static of(x) {
    throw new Error('`of` called on class Left (value) instead of Either (type)');
  }

//   [util.inspect.custom]() {
//     return `Left(${inspect(this.$value)})`;
//   }

  // ----- Functor (Either a)
  map() {
    return this;
  }

  // ----- Applicative (Either a)
  ap() {
    return this;
  }

  // ----- Monad (Either a)
  chain() {
    return this;
  }

  join() {
    return this;
  }

  // ----- Traversable (Either a)
  sequence(of) {
    return of(this);
  }

  traverse(of, fn) {
    return of(this);
  }
}


class Right extends Either {
  get isLeft() {
    return false;
  }

  get isRight() {
    return true;
  }

  static of(x) {
    throw new Error('`of` called on class Right (value) instead of Either (type)');
  }

//   [util.inspect.custom]() {
//     return `Right(${inspect(this.$value)})`;
//   }

  // ----- Functor (Either a)
  map(fn) {
    return Either.of(fn(this.$value));
  }

  // ----- Applicative (Either a)
  ap(f) {
    return f.map(this.$value);
  }

  // ----- Monad (Either a)
  chain(fn) {
    return fn(this.$value);
  }

  join() {
    return this.$value;
  }

  // ----- Traversable (Either a)
  sequence(of) {
    return this.traverse(of, identity); // this.$value.map(Either.of);
  }

  traverse(of, fn) {
    return fn(this.$value).map(Either.of);
  }
}


class Identity {
  constructor(x) {
    this.$value = x;
  }

//   [util.inspect.custom]() {
//     return `Identity(${inspect(this.$value)})`;
//   }

  // ----- Pointed Identity
  static of(x) {
    return new Identity(x);
  }

  // ----- Functor Identity
  map(fn) {
    return Identity.of(fn(this.$value));
  }

  // ----- Applicative Identity
  ap(f) {
    return f.map(this.$value);
  }

  // ----- Monad Identity
  chain(fn) {
    return this.map(fn).join();
  }

  join() {
    return this.$value;
  }

  // ----- Traversable Identity
  sequence(of) {
    return this.traverse(of, identity);
  }

  traverse(of, fn) {
    return fn(this.$value).map(Identity.of);
  }
}


class IO {
  constructor(fn) {
    this.unsafePerformIO = fn;
  }

//   [util.inspect.custom]() {
//     return 'IO(?)';
//   }

  // ----- Pointed IO
  static of(x) {
    return new IO(() => x);
  }

  // ----- Functor IO
  map(fn) {
    return new IO(compose(fn, this.unsafePerformIO));
  }

  // ----- Applicative IO
  ap(f) {
    return this.chain(fn => f.map(fn));
  }

  // ----- Monad IO
  chain(fn) {
    return this.map(fn).join();
    // this.map(fn) => IO(compose(fn, this.unsafePerformIO))
    // .join() => new IO ( () => compose(fn, this.unsafePerformIO)().unsafePerformIO())
    // => new IO ( () => fn(this.unsafePerformIO()).unsafePerformIO() ) // fn -> IO
  }

  join() {
    return new IO(() => this.unsafePerformIO().unsafePerformIO());
    // return this.unsafePerformIO()
  }
}


class List {
  constructor(xs) {
    this.$value = xs;
  }

//   [util.inspect.custom]() {
//     return `List(${inspect(this.$value)})`;
//   }

  concat(x) {
    return new List(this.$value.concat(x));
  }

  // ----- Pointed List
  static of(x) {
    return new List([x]);
  }

  // ----- Functor List
  map(fn) {
    return new List(this.$value.map(fn));
  }

  // ----- Traversable List
  sequence(of) {
    return this.traverse(of, identity);
  }

  traverse(of, fn) {
    return this.$value.reduce(
      (f, a) => fn(a).map(b => bs => bs.concat(b)).ap(f),
      of(new List([])),
    );
  }
}


class Map {
  constructor(x) {
    this.$value = x;
  }

//   [util.inspect.custom]() {
//     return `Map(${inspect(this.$value)})`;
//   }

//   insert(k, v) {
//     const singleton = {};
//     singleton[k] = v;
//     return Map.of(Object.assign({}, this.$value, singleton));
//   }
insert(k, v) {
    const singleton = {};
    singleton[k] = v;
    return new Map(Object.assign({}, this.$value, singleton));
  }

  reduceWithKeys(fn, zero) {
    return Object.keys(this.$value)
      .reduce((acc, k) => fn(acc, this.$value[k], k), zero);
  }

  // ----- Functor (Map a)
  map(fn) {
    return this.reduceWithKeys(
      (m, v, k) => m.insert(k, fn(v)),
      new Map({}),
    );
  }

  // ----- Traversable (Map a)
  sequence(of) {
    return this.traverse(of, identity);
  }

  traverse(of, fn) {
    return this.reduceWithKeys(
      (f, a, k) => fn(a).map(b => m => m.insert(k, b)).ap(f),
      of(new Map({})),
    );
  }
}


class Maybe {
  get isNothing() {
    return this.$value === null || this.$value === undefined;
  }

  get isJust() {
    return !this.isNothing;
  }

  constructor(x) {
    this.$value = x;
  }

//   [util.inspect.custom]() {
//     return this.isNothing ? 'Nothing' : `Just(${inspect(this.$value)})`;
//   }

  // ----- Pointed Maybe
  static of(x) {
    return new Maybe(x);
  }

  // ----- Functor Maybe
  map(fn) {
    return this.isNothing ? this : Maybe.of(fn(this.$value));
  }

  // ----- Applicative Maybe
  ap(f) {
    return this.isNothing ? this : f.map(this.$value);
  }

  // ----- Monad Maybe
  chain(fn) {
    return this.map(fn).join();
  }

  join() {
    return this.isNothing ? this : this.$value;
  }

  // ----- Traversable Maybe
  sequence(of) {
    return this.traverse(of, identity);
  }

  traverse(of, fn) {
    return this.isNothing ? of(this) : fn(this.$value).map(Maybe.of);
  }
}


class Task {
  constructor(fork) {
    this.fork = fork;
  }

//   [util.inspect.custom]() {
//     return 'Task(?)';
//   }

  static rejected(x) {
    return new Task((reject_, _) => reject_(x));
  }

  // ----- Pointed (Task a)
  static of(x) {
    debugger
    return new Task((_, resolve) => resolve(x));
  }

  // ----- Functor (Task a)
  map(fn) {
    debugger
    return new Task((reject_, resolve) => this.fork(reject_, compose(resolve, fn)));
  }

  // ----- Applicative (Task a)
  ap(f) {
    return this.chain(fn => f.map(fn));
  }

  // ----- Monad (Task a)
  chain(fn) {
    return new Task((reject_, resolve) => this.fork(reject_, x => fn(x).fork(reject_, resolve)));
  }

  join() {
    return this.chain(identity);
  }
}


// nothing :: Maybe a
const nothing = Maybe.of(null);


// reject :: a -> Task a b
const reject = a => Task.rejected(a);


// flip :: (a -> b -> c) -> b -> a -> c
const flip = curry((fn, a, b) => fn(b, a));


// concat :: String -> String -> String
const concat = curry((a, b) => a.concat(b));


// add :: Number -> Number -> Number
const add = curry((a, b) => a + b);


// append :: String -> String -> String
const append = flip(concat);


// chain :: Monad m => (a -> m b) -> m a -> m b
const chain = curry((fn, m) => m.chain(fn));


// eq :: Eq a => a -> a -> Boolean
const eq = curry((a, b) => a === b);


// filter :: (a -> Boolean) -> [a] -> [a]
const filter = curry((fn, xs) => xs.filter(fn));


// forEach :: (a -> ()) -> [a] -> ()
const forEach = curry((fn, xs) => xs.forEach(fn));


// head :: [a] -> a
const head = xs => xs[0];


// intercalate :: String -> [String] -> String
const intercalate = curry((str, xs) => xs.join(str));


// join :: Monad m => m (m a) -> m a
const join = m => m.join();


// last :: [a] -> a
const last = xs => xs[xs.length - 1];


// map :: Functor f => (a -> b) -> f a -> f b
const map = curry((fn, f) => f.map(fn));


// match :: RegExp -> String -> Boolean
const match = curry((re, str) => re.test(str));


// prop :: String -> Object -> a
const prop = curry((p, obj) => obj[p]);


// reduce :: (b -> a -> b) -> b -> [a] -> b
const reduce = curry((fn, zero, xs) => xs.reduce(fn, zero));


// replace :: RegExp -> String -> String -> String
const replace = curry((re, rpl, str) => str.replace(re, rpl));


// reverse :: [a] -> [a]
const reverse = x => (Array.isArray(x) ? x.reverse() : x.split('').reverse().join(''));


// safeHead :: [a] -> Maybe a
const safeHead = compose(Maybe.of, head);


// safeLast :: [a] -> Maybe a
const safeLast = compose(Maybe.of, last);


// safeProp :: String -> Object -> Maybe a
const safeProp = curry((p, obj) => compose(Maybe.of, prop(p))(obj));


// sequence :: (Applicative f, Traversable t) => (a -> f a) -> t (f a) -> f (t a)
const sequence = curry((of, f) => f.sequence(of));


// sortBy :: Ord b => (a -> b) -> [a] -> [a]
const sortBy = curry((fn, xs) => xs.sort((a, b) => {
  if (fn(a) === fn(b)) {
    return 0;
  }

  return fn(a) > fn(b) ? 1 : -1;
}));


// split :: String -> String -> [String]
const split = curry((sep, str) => str.split(sep));


// take :: Number -> [a] -> [a]
const take = curry((n, xs) => xs.slice(0, n));


// toLowerCase :: String -> String
const toLowerCase = s => s.toLowerCase();


// toString :: a -> String
const toString = String;


// toUpperCase :: String -> String
const toUpperCase = s => s.toUpperCase();


// traverse :: (Applicative f, Traversable t) => (a -> f a) -> (a -> f b) -> t a -> f (t b)
const traverse = curry((of, fn, f) => f.traverse(of, fn));


// unsafePerformIO :: IO a -> a
const unsafePerformIO = io => io.unsafePerformIO();


export { always };
export { compose };
export { curry };
export { either };
export { identity };
export { inspect };
export { left };
export { liftA2 };
export { liftA3 };
export { maybe };
export { nothing };
export { reject };
export { createCompose };
export { Either };
export { Left };
export { Right };
export { Identity };
export { IO };
export { List };
export { Map };
export { Maybe };
export { Task };
export { add };
export { append };
export { chain };
export { concat };
export { eq };
export { filter };
export { flip };
export { forEach };
export { head };
export { intercalate };
export { join };
export { last };
export { map };
export { match };
export { prop };
export { reduce };
export { replace };
export { reverse };
export { safeHead };
export { safeLast };
export { safeProp };
export { sequence };
export { sortBy };
export { split };
export { take };
export { toLowerCase };
export { toString };
export { toUpperCase };
export { traverse };
export { unsafePerformIO };