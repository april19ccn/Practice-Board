import { always, Map, curry, safeHead, IO, Maybe, Either, Identity, left, compose, chain, map, sequence, traverse } from "../../utils/support.js";
import Task from "data.task";

/* ---------- Chapter 12 ---------- */

const httpGet = function httpGet(route) { return Task.of(`json for ${route}`); };

const routes = new Map({
    '/': '/',
    '/about': '/about',
});

const validate = function validate(player) {
    return player.name
        ? Either.of(player)
        : left('must have name');
};

const readdir = function readdir(dir) {
    return Task.of(['file1', 'file2', 'file3']);
};

const readfile = curry(function readfile(encoding, file) {
    return Task.of(`content of ${file} (${encoding})`);
});


// ex1 ---------------------------------------------

// Considering the following elements:
//
//   // httpGet :: Route -> Task Error JSON
//   // routes :: Map Route Route
//   const routes = new Map({ '/': '/', '/about': '/about' });
//
// Use the traversable interface to change the type signature of `getJsons`.
//
//   getJsons :: Map Route Route -> Task Error (Map Route JSON)

// getJsons :: Map Route Route -> Map Route (Task Error JSON)
// export const getJsons = map(httpGet);

// A1:
// export const getJsons = traverse(Task.of, httpGet);

// A2:
export const getJsons = compose(sequence(Task.of), map(httpGet));


// ex2 ---------------------------------------------

// Using traversable, and the `validate` function, update `startGame` (and its signature)
// to only start the game if all players are valid
//
//   // validate :: Player -> Either String Player
//   validate = player => (player.name ? Either.of(player) : left('must have name'));

// startGame :: [Player] -> [Either Error String]
// export const startGame = compose(map(map(always('game started!'))), map(validate));

// Player:: [a]
// map(validate) :: [Player] -> [Either Error Player]

// export const startGame = compose(traverse(Either.of, map(always('game started!'))), map(validate));

export const startGame = compose(map(always('game started!')), traverse(Either.of, validate));
// new List([albert, theresa])
// => Right([albert, theresa])
// => Right(always('game started!')([albert, theresa]))

// new List([gary, { what: 14 }])
// => validate => [Right(gary), Left('must have name')]


// [gary, { what: 14 }].traverse(Either.of, validate)
//     = [gary, { what: 14 }].reduce(
//         (f, a) => validate(a).map(b => bs => bs.concat(b)).ap(f),
//         Either.of(new List([])),
//     );
//     // gary
//     = (Either([]), gary) => validate(gary).map(b => bs => bs.concat(b)).ap(Either([]))
//     = Either(bs => bs.concat(gary)).ap(Either([]))
//     = Either([]).map(bs => bs.concat(gary))
//     = Either([gary])
//     // { what: 14 }
//     = (Either([gary]), { what: 14 }) => validate({ what: 14 }).map(b => bs => bs.concat(b)).ap(Either([gary]))
//     = left('must have name').map(b => bs => bs.concat(b)).ap(Either([gary]))
//     = left('must have name')

// ex3 ---------------------------------------------
// Considering the following functions:
//
//   readfile :: String -> String -> Task Error String
//   readdir :: String -> Task Error [String]
//
// Use traversable to rearrange and flatten the nested Tasks & Maybe

// readFirst :: String -> Task Error (Maybe (Task Error String))
// export const readFirst = compose(map(map(readfile('utf-8'))), map(safeHead), readdir);

export const readFirst = compose(
    chain(traverse(Task.of, readfile('utf-8'))),
    map(safeHead),
    readdir
)


// Maybe (Task Error String)
const readFirst_1 = compose(
    // chain(),
    sequence(Maybe.of),
    map(safeHead),
    readdir,
) 