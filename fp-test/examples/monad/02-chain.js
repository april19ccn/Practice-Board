import { Either, IO, Maybe, compose, map, curry, toUpperCase, Right } from "../../utils/support.js";
import Task from "data.task";


//  safeProp :: Key -> {Key: a} -> Maybe a
var safeProp = curry(function (x, obj) {
    return new Maybe(obj[x]);
});

//  safeHead :: [a] -> Maybe a
var safeHead = safeProp(0);

//  join :: Monad m => m (m a) -> m a
var join = function (mma) { return mma.join(); }

//  chain :: Monad m => (a -> m b) -> m a -> m b
var chain = curry(function (f, m) {
    return m.map(f).join(); // 或者 compose(join, map(f))(m)
});

// chain
var firstAddressStreet = compose(
    chain(safeProp('street')), chain(safeHead), safeProp('addresses')
);

//  firstAddressStreet :: User -> Maybe Street
// var firstAddressStreet = compose(
//     join, map(safeProp('street')), join, map(safeHead), safeProp('addresses')
// );


console.log(
    firstAddressStreet(
        { addresses: [{ street: { name: 'Mulburry', number: 8402 }, postcode: "WC2N" }] }
    )
)
// Maybe({name: 'Mulburry', number: 8402})



////////////////////////////////////////////////////////////

// getJSON :: Url -> Params -> Task JSON
// querySelector :: Selector -> IO DOM


getJSON('/authenticate', { username: 'stale', password: 'crackers' })
    .chain(function (user) {
        return getJSON('/friends', { user_id: user.id });
    });
// Task([{name: 'Seimith', id: 14}, {name: 'Ric', id: 39}]);


querySelector("input.username").chain(function (uname) {
    return querySelector("input.email").chain(function (email) {
        return IO.of(
            "Welcome " + uname.value + " " + "prepare for spam at " + email.value
        );
    });
});
// IO("Welcome Olivia prepare for spam at olivia@tremorcontrol.net");


Maybe.of(3).chain(function (three) {
    return Maybe.of(2).map(add(three));
});
// Maybe(5);
// =>
// Maybe(3).map(function (three) {
//     return Maybe.of(2).map(add(three));
// }).join()
// =>
// Maybe.of((function (three) {
//     return Maybe.of(2).map(add(three));
// })(3)).join()
// =>
// Maybe(Maybe(5)).join()
// => 
// Maybe(5)


Maybe.of(null).chain(safeProp('address')).chain(safeProp('street'));
// Maybe(null);


///////////////////////////////
const path = 'd:/Star_Code/A_Study/unocss-practice/fp-test/examples/monad/test.txt'

const readFile = (filename) => new IO(() => {
    console.log(`Reading ${filename}...`);
    return `Content of ${filename}`;
});

const print = (content) => new IO(() => {
    console.log(`Printed: ${content}`);
    return content;
});

// 组合操作（返回 IO(IO)）
const readAndPrint = readFile(path).map(print); // IO(IO)
// deepseek IO =>
// readFile(path) => new IO(() => {
//     console.log(`Reading ${filename}...`);
//     return `Content of ${filename}`;
// });
// IO(() => {
//     console.log(`Reading ${path}...`);
//     return `Content of ${path}`;
// }).map(print);
// IO(this.unsafePerformIO).chain(x => IO.of(print(x)))
// new IO(() => (x => IO.of(print(x))(this.unsafePerformIO()).unsafePerformIO())
// IO(() => IO(() => print(this.unsafePerformIO())).unsafePerformIO())
// IO(() => print(this.unsafePerformIO()))
// IO(() => compose(print, this.unsafePerformIO)())

// support.js map => new IO(compose(print, this.unsafePerformIO))


// 展平嵌套 IO
const flattened = readAndPrint.join(); // 等价于 readAndPrint.chain(id)
// deepseek IO =>
// => IO(() => print(this.unsafePerformIO())).chain(x => x);
// => new IO(() => (x => x)(this.unsafePerformIO()).unsafePerformIO());
// => IO(() => (x => x)(print(this.unsafePerformIO())).unsafePerformIO())
// => IO(() => print(this.unsafePerformIO()).unsafePerformIO())

// support.js map().join()
// IO(compose(print, this.unsafePerformIO)).join()
// IO(() => compose(print, this.unsafePerformIO)().unsafePerformIO())
// IO(() => print(this.unsafePerformIO()).unsafePerformIO())

flattened.unsafePerformIO();
// 输出:
// Reading test.txt...
// Printed: Content of test.txt




//////////////////////////
// readFile :: Filename -> Either String (Task Error String)
// httpPost :: String -> String -> Task Error JSON
// upload :: Filename -> Either String (Task Error JSON)
const upload = compose(map(chain(httpPost('/uploads'))), readFile);

// Right =>
// Either(Task String).map(chain(httpPost('/uploads')))
// Either.of(chain(httpPost('/uploads'))(Task String))
Either.of(new Task((reject_, resolve) => this.fork(reject_, x => httpPost('/uploads')(x).fork(reject_, resolve))))


/////////////// 公式推导
mcompose(f, M)(x)
    = compose(chain(f), M)(x)
    = chain(f)(M(x))
    = M(x).map(f).join()
    = M(f(x)).join()
    = f(x)

mcompose(f, M)
    = compose(chain(f), M)
    = compose(compose(join, map(f)), M)
    = compose(join, map(f))(M)
    = compose(join, M(f))
    = compose(join, of(f))
    = f

/////////////////
// 基础函数
const double = x => Maybe.of(x * 2);
const add1 = x => Maybe.of(x + 1);
const M = Maybe.of;

// 左恒等律验证
mcompose(M, double)(2) 
  = compose(chain(double), M)(2) 
  = chain(double)(M(2)) 
  = Maybe.of(2).chain(double) 
  = double(2) 
  === double(2);
// 左恒等不应该是
mcompose(M, double)(2)
    = compose(chain(M), double)(2)
    = chain(M)(double(2))
    = Maybe.of(4).chain(M)
    = Maybe(4).map(M).join()
    = Maybe.of(M(4)).join()
    = Maybe(Maybe(4)).join()
    = Maybe(4)
    === double(2) 
    === Maybe(4)


// 结合律验证
const f = double, g = add1, h = x => Maybe.of(x - 3);
const left = mcompose(mcompose(f, g), h)(5);
const right = mcompose(f, mcompose(g, h))(5);
// left 和 right 均返回 Maybe.of(5)


// 基础定律（join 和 of 形式）是 数学核心，直接定义 Monad 的代数结构。

// 函数组合定律（mcompose 形式）是 应用场景的抽象，描述 Monad 如何安全组合函数。

// 那我的理解就是 mcompose 里 chain 的 M(x) 就是 Monad 的代数结构，f(x) 就是 Monad 的代数运算。

// map 打开 ，之后用join（基础定律）