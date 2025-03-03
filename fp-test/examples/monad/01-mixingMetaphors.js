import { Either, IO, Maybe, compose, map, curry, toUpperCase } from "../../utils/support.js";
import Task from "data.task";

import fs from "fs";

// readFile :: String -> IO String
const readFile = filename => new IO(() => {
    console.log("**********")
    return fs.readFileSync(filename, 'utf-8')
});

// print :: String -> IO String
const print = x => new IO(() => {
    console.log("print: ", x);
    return x;
});

// cat :: String -> IO (IO String)
const cat = compose(map(print), readFile);

const path = 'd:/Star_Code/A_Study/unocss-practice/fp-test/examples/monad/test.txt'

console.log(cat(path));
// => // IO(IO('[core]\nrepositoryformatversion = 0\n')) 基于下面的展开，我的理解是 IO('[core]\nrepositoryformatversion = 0\n') 先被打开计算， 再然后是 IO(...)
// compose(map(print), readFile)(path)
// compose(map(print), IO(() => fs.readFileSync(path)))
// new IO(R.compose(print, () => fs.readFileSync(path)))
// IO(R.compose(print, () => fs.readFileSync(path)))


console.log(cat(path).unsafePerformIO());
// IO(R.compose(print, () => fs.readFileSync(path))).unsafePerformIO()
// R.compose(print, () => fs.readFileSync(path))() // 这里就需要计算 () => fs.readFileSync(path) 了
// print(fs.readFileSync(path))
// new IO(() => {
//     console.log("print: ", fs.readFileSync(path));
//     return fs.readFileSync(path);
// })

console.log(cat(path).unsafePerformIO().unsafePerformIO());
// () => {
//     console.log("print: ", fs.readFileSync(path));
//     return fs.readFileSync(path);
// }()

console.log(cat(path).join().unsafePerformIO())
// // IO(R.compose(print, () => fs.readFileSync(path))) => join() 有点像刮皮刀 => R.compose(print, () => fs.readFileSync(path))()

///////////////////////////////////////////

//  safeProp :: Key -> {Key: a} -> Maybe a
var safeProp = curry(function (x, obj) {
    return new Maybe(obj[x]);
});

//  safeHead :: [a] -> Maybe a
var safeHead = safeProp(0);

//  firstAddressStreet :: User -> Maybe (Maybe (Maybe Street) )
var firstAddressStreet = compose(
    map(map(safeProp('street'))), map(safeHead), safeProp('addresses')
);

console.log(
    firstAddressStreet(
        { addresses: [{ street: { name: 'Mulburry', number: 8402 }, postcode: "WC2N" }] }
    )
)
// Maybe(Maybe(Maybe({name: 'Mulburry', number: 8402})))

// compose(
//     map(map(safeProp('street'))), map(safeHead), safeProp('addresses')
// )(
//     { addresses: [{ street: { name: 'Mulburry', number: 8402 }, postcode: "WC2N" }] }
// )

// compose(
//     map(map(safeProp('street'))), map(safeHead), safeProp('addresses')
// )(
//     { addresses: [{ street: { name: 'Mulburry', number: 8402 }, postcode: "WC2N" }] }
// )

// compose(
//     map(map(safeProp('street'))), map(safeHead), Maybe([{ street: { name: 'Mulburry', number: 8402 }, postcode: "WC2N" }])
// )

// compose(
//     map(map(safeProp('street'))), Maybe.of(safeHead([{ street: { name: 'Mulburry', number: 8402 }, postcode: "WC2N" }]))
// )

// compose(
//     map(map(safeProp('street'))), Maybe(Maybe({ street: { name: 'Mulburry', number: 8402 }}))
// )

// compose(
//     Maybe.of(
//         map(safeProp('street'))(Maybe({ street: { name: 'Mulburry', number: 8402 }}))
//     )
// )

// compose(
//     Maybe.of(
//         Maybe.of(
//             Maybe({ name: 'Mulburry', number: 8402 })
//         )
//     )
// )

// Maybe(Maybe(Maybe({name: 'Mulburry', number: 8402})))


/////////////
//  join :: Monad m => m (m a) -> m a
var join = function (mma) { return mma.join(); }

//  firstAddressStreet :: User -> Maybe Street
var firstAddressStreet = compose(
    join, map(safeProp('street')), join, map(safeHead), safeProp('addresses')
);

console.log(
    firstAddressStreet(
        { addresses: [{ street: { name: 'Mulburry', number: 8402 }, postcode: "WC2N" }] }
    )
)
// Maybe({name: 'Mulburry', number: 8402})


//////////
// //  log :: a -> IO a
// var log = function (x) {
//     return new IO(function () { console.log(x); return x; });
// }

// //  setStyle :: Selector -> CSSProps -> IO DOM
// var setStyle = curry(function (sel, props) {
//     return new IO(function () { return jQuery(sel).css(props); });
// });

// //  getItem :: String -> IO String
// var getItem = function (key) {
//     return new IO(function () { return localStorage.getItem(key); });
// };

// //  applyPreferences :: String -> IO DOM
// var applyPreferences = compose(
//     join, map(setStyle('#main')), join, map(log), map(JSON.parse), getItem
// );


// applyPreferences('preferences').unsafePerformIO();
// // Object {backgroundColor: "green"}
// // <div style="background-color: 'green'"/>