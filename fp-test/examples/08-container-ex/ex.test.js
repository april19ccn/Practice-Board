import { describe, expect, test } from "vitest";
import * as E from "./ex.js";
import { Identity, Maybe, Either, left } from "../../utils/support.js";

describe("Functor Exercises", () => {
    test('Exercise 1', function () {
        expect(E.ex1(Identity.of(2))).toEqual(Identity.of(3));
    });

    test('Exercise 2', function () {
        const xs = Identity.of(['do', 'ray', 'me', 'fa', 'so', 'la', 'ti', 'do']);
        expect(E.ex2(xs)).toEqual(Identity.of('do'));
    });

    test('Exercise 3', function () {
        const user = { id: 2, name: "Albert" };
        expect(E.ex3(user)).toEqual(Maybe.of('A'));
    });

    test('Exercise 4 - 1', function () {
        expect(E.ex4_1("4")).toEqual(Maybe.of(4));
    });

    test('Exercise 4 - 2', function () {
        expect(E.ex4_2("4")).toEqual(Maybe.of(4));
    });

    // vitest 不支持异步回调， jest支持
    // test('Exercise 5',function (done) {
    //     E.ex5(13).fork(console.log, function (res) {
    //         console.log(123)
    //         expect.toEqual(res).toEqual('LOVE THEM FUTURES');
    //         done()
    //     })
    // });

    test('Exercise 5', async () => {
        return new Promise((done) => {
            E.ex5(13).fork(
                (error) => done.fail(error),
                (res) => {
                    expect(res).toEqual('LOVE THEM FUTURES');
                    done();
                }
            );
        });
    });

    test('Exercise 6', function () {
        expect(E.ex6({ active: false, name: 'Gary' })).toEqual(left('Your account is not active'))
        expect(E.ex6({ active: true, name: 'Theresa' })).toEqual(Either.of('Welcome Theresa'))
    });

    test('Exercise 7', function () {
        expect(E.ex7("fpguy99")).toEqual(Either.of("fpguy99"));
        expect(E.ex7("...")).toEqual(left("You need > 3"));
    });

    test('Exercise 8', function () {
        expect(E.ex8("fpguy99").unsafePerformIO()).toEqual("fpguy99-saved");
        expect(E.ex8("...").unsafePerformIO()).toEqual("You need > 3");
    });

})