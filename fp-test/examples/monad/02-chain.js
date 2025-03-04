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

// Either.of(String).map(chain(httpPost('/uploads')));
// Either.of(chain(httpPost('/uploads'))(String))