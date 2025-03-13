import { describe, expect, test } from "vitest";
import * as _ from "ramda"
import * as E from "./ex.js";
import { Identity, Maybe, Either, left, either, unsafePerformIO } from "../../utils/support.js";
import Task from "data.task";

describe("Natural Transformations", function () {
    test('Exercise 1', function () {
        const just = E.eitherToMaybe(Either.of('one eyed willy'));
        const noth = E.eitherToMaybe(left('some error'));


        // 验证 just 是 Maybe 的实例
        expect(just).toBeInstanceOf(Maybe);
        // 验证是 Just 类型而非 Nothing
        expect(just.isJust).toBe(true);
        // 验证内部值是否正确
        expect(just.$value).toBe('one eyed willy');


        // 验证 noth 是 Maybe 的实例
        expect(noth).toBeInstanceOf(Maybe);
        // 验证 isNothing
        expect(noth.isNothing).toBe(true);



        // expect(just instanceof Maybe && just.isJust && just.$value === 'one eyed willy').toBeTruthy()

        // expect(noth instanceof Maybe && noth.isNothing).toBeTruthy()

        // assert(
        //     just instanceof Maybe && just.isJust && just.$value === 'one eyed willy',
        //     'The function maps the `Right()` side incorrectly; hint: `Right(14)` should be mapped to `Just(14)`',
        // );

        // assert(
        //     noth instanceof Maybe && noth.isNothing,
        //     'The function maps the `Left()` side incorrectly; hint: `Left(\'error\')` should be mapped to `Nothing`',
        // );
    });

    test("Exercise 2 resolve", function () {
        const res = E.findNameById(1);
        expect(res).toBeInstanceOf(Task);
        return new Promise((done) => {
            res.fork(console.log, function (val) {
                expect(val instanceof Task).toBe(false);
                expect(val instanceof Either).toBe(false);
                expect(val).toEqual('Albert');
                done();
            });
        })
    })

    test("Exercise 2 reject", function () {
        const rej = E.findNameById(999);
        expect(rej).toBeInstanceOf(Task);
        return new Promise((done) => {
            rej.fork(function (val) {
                expect(val instanceof Task).toBe(false);
                expect(val instanceof Either).toBe(false);
                expect(val).toEqual('not found');
                done();
            }, console.log);
        })
    })

    test("Exercise 3", function () {
        const sortLetters = _.compose(E.listToStr, _.sortBy(_.identity), E.strToList);

        expect(sortLetters('sortme')).toEqual("emorst"); // expect放要测值，还是预期值
    })
});