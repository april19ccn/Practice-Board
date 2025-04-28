// test/index.test.ts
import { describe, expect, it } from 'bun:test'
import { Elysia } from 'elysia'

import { app } from '../index'

describe('Elysia_1', () => {
    it('return a response', async () => {
        // const app = new Elysia().get('/', () => 'hi')

        const response = await app
            .handle(new Request('http://localhost/hi'))
            .then((res) => res.text())

        expect(response).toBe('hi~')
    })
})

describe('Elysia_2', () => {
    it('return a response', async () => {
        // const app = new Elysia().get('/', () => 'hi')

        const response = await app
            .handle(new Request('http://localhost/note', {method: "get"}))
            .then((res) => res.json())

        expect(response).toMatchObject([ 
			{ 
				data: 'Moonhalo', 
				author: 'saltyaom'
			} 
		] )
    })
})