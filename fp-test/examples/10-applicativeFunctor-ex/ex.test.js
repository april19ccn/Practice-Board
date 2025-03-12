import { describe, expect, test } from "vitest";
import * as _ from "ramda"
import * as E from "./ex.js";
import { Identity, Maybe, Either, left, either, unsafePerformIO } from "../../utils/support.js";

describe("Applicative Exercises", function () {
    test('Exercise 1', function () {
        expect(Maybe.of(5)).toEqual(E.ex1(2, 3))
        expect(Maybe.of(null)).toEqual(E.ex1(null, 3))
    });

    test('Exercise 2', function () {
        expect(Maybe.of(5)).toEqual(E.ex2(Maybe.of(2), Maybe.of(3)))
        expect(Maybe.of(null)).toEqual(E.ex2(Maybe.of(null), Maybe.of(3)))
    });

    test('Exercise 3', function () {
        return new Promise((done) => {
            E.ex3.fork(console.log, function (html) {
                expect(html).toEqual("<div>Love them tasks</div><li>This book should be illegal</li><li>Monads are like space burritos</li>");
                done();
            });
        })
    });

    test('Exercise 4', function () {
        expect(E.ex4.unsafePerformIO()).toEqual("toby vs sally");
    });
});