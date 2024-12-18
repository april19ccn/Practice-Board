import { Elysia } from "elysia";

export const date = new Elysia()
    .decorate('getDate', () => Date.now())
    .get('/date', ({ getDate }) => getDate())