import { describe, expect, test } from "vitest";
import * as _ from "ramda"
import * as E from "./ex.js";
import { Map, List, Identity, Maybe, Either, left, either, unsafePerformIO } from "../../utils/support.js";
import Task from "data.task";

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

describe("Traverse", function () {
    test('Exercise 1', function () {
        const throwUnexpected = () => {
            throw new Error('The function gives incorrect results; a Task has resolved unexpectedly!');
        };

        const routes = new Map({
            '/': '/',
            '/about': '/about',
        });

        const res = E.getJsons(routes)

        return new Promise((done) => {
            res.fork(throwUnexpected, function (val) {
                expect(val.$value['/']).toEqual('json for /');
                expect(val.$value['/about']).toEqual('json for /about');

                // callees 可能与 support.js 里实现有关，现在测不出来
                // const callees = E.getJsons.callees;

                // console.log(callees)

                // expect(callees[0]).toEqual('map');
                // expect(callees[1]).toEqual('sequence');

                done();
            });
        })
    });

    test('Exercise 2', function () {
        const res = E.startGame(new List([albert, theresa]));
        console.log(res)
        expect(res).toBeInstanceOf(Either)
        expect(res.isRight).toBe(true)
        expect(res.$value).toEqual('game started!')

        const rej = E.startGame(new List([gary, { what: 14 }]));
        expect(rej.isLeft).toBe(false)
        expect(rej.$value).toEqual('must have name')
    });

    test('Exercise 3', function () {

    });
});