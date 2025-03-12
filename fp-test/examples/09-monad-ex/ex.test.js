import { describe, expect, test } from "vitest";
import * as _ from "ramda"
import * as E from "./ex.js";
import { Identity, Maybe, Either, left, either, unsafePerformIO } from "../../utils/support.js";

describe("Monad Exercises", function () {
    test('Exercise 1', function () {
        expect(E.ex1(E.user)).toEqual(Maybe.of('Walnut St'))
    });

    test('Exercise 2', function () {
        expect(E.ex2(undefined).unsafePerformIO()).toEqual('logged ex.js')
    });

    test('Exercise 3', async () => {
        return new Promise((done) => {
            E.ex3(13).fork(console.log, function (res) {
                expect(res.map(_.prop('post_id'))).toEqual([13, 13]);
                done();
            });
        })
    });

    test('Exercise 4', function () {
        var getResult = either(_.identity, unsafePerformIO)
        expect(getResult(E.ex4('notanemail'))).toEqual('invalid email')
        expect(getResult(E.ex4('sleepy@grandpa.net'))).toEqual('emailed: sleepy@grandpa.net')
    });
});