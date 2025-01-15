import { expect, test } from 'vitest'

import * as E from "./my-ex.js"

test('Exercise 1', () => {
    expect(E.words("Jingle bells Batman smells")).toEqual(['Jingle', 'bells', 'Batman', 'smells']);
});