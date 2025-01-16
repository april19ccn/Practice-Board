// npx vitest 启动测试
import { describe, expect, test } from 'vitest'

import * as E from "../ex/curry-ex.js"

describe("Curry Exercises", () => {
    test('Exercise 1', () => {
        expect(E.words("Jingle bells Batman smells")).toEqual(['Jingle', 'bells', 'Batman', 'smells']);
    });

    test('Exercise 1a', () => {
        expect(E.sentences(["Jingle bells Batman smells", "Robin laid an egg"])).toEqual([['Jingle', 'bells', 'Batman', 'smells'], ['Robin', 'laid', 'an', 'egg']]);
    });

    test('Exercise 2', function () {
        expect(E.filterQs(['quick', 'camels', 'quarry', 'over', 'quails'])).toEqual(['quick', 'quarry', 'quails']);
    });

    test('Exercise 3', function () {
        expect(E.max([323, 523, 554, 123, 5234])).toEqual(5234);
    });

    if (E.slice != undefined) {
        test('Curry Bonus 1', function () {
            expect(E.slice(1)(3)(['a', 'b', 'c'])).toEqual(['b', 'c']);
        });
    }

    if (E.take != undefined) {
        test('Curry Bonus 2', function () {
            expect(E.take(2)(['a', 'b', 'c'])).toEqual(['a', 'b']);
        });
    }
})