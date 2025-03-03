import { Either, IO, Maybe, compose, map, curry, toUpperCase } from "../../utils/support.js";
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