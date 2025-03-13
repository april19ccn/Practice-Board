import * as _ from "ramda"
import { Either, IO, Maybe, compose, chain, map, curry, prop, toUpperCase, Right, Left, either, left, liftA2, split, intercalate } from "../../utils/support.js";
import Task from "data.task";


// Exercise 1
// ==========
// Write a natural transformation that converts `Either b a` to `Maybe a`
// 写一个 natural transformation 将Either b a转换到Maybe a

// eitherToMaybe :: Either b a -> Maybe a
// const eitherToMaybe = (e) => {
//     console.log(e.$value)
//     return either(() => Maybe.of(null), Maybe.of)(e);
// }

const eitherToMaybe = either(_.always(Maybe.of(null)), Maybe.of);


// Exercise 2
// ==========
// Using `eitherToTask`, simplify `findNameById` to remove the nested `Either`.
// 使用eitherToTask, 简化findNameById方法, 去掉嵌套的Either。

const albert = {
    id: 1,
    active: true,
    name: 'Albert',
    address: {
        street: {
            number: 22,
            name: 'Walnut St',
        },
    },
};

const gary = {
    id: 2,
    active: false,
    name: 'Gary',
    address: {
        street: {
            number: 14,
        },
    },
};

const theresa = {
    id: 3,
    active: true,
    name: 'Theresa',
};

const findUserById = function findUserById(id) {
    switch (id) {
        case 1:
            return Task.of(Either.of(albert));

        case 2:
            return Task.of(Either.of(gary));

        case 3:
            return Task.of(Either.of(theresa));

        default:
            return Task.of(left('not found'));
    }
};

// eitherToTask :: Either a b -> Task a b
const eitherToTask = either(Task.rejected, Task.of);

// findNameById :: Number -> Task Error (Either Error User)
// const findNameById = compose(map(map(prop('name'))), findUserById);

const findNameById = compose(
    map(prop('name')),
    chain(eitherToTask),
    findUserById
);




// Exercise 3
// ==========
// Write the isomorphisms between String and [Char]. 写出字符串与 [Char] 之间的同构关系
//
// As a reminder, the following functions are available in the exercise's context:
//
//   split :: String -> String -> [String]
//   intercalate :: String -> [String] -> String

// strToList :: String -> [Char]
const strToList = split('');

// listToStr :: [Char] -> String
const listToStr = intercalate('');

export { eitherToMaybe, findNameById, strToList, listToStr }