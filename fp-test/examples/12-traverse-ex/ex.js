import { always,Map, curry, IO, Maybe, Either, Identity, left, compose, chain, map, sequence, traverse } from "../../utils/support.js";
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

export const startGame = compose(traverse(Either.of, map(always('game started!'))), map(validate));

traverse(Either.of, validate)([albert, theresa])