import { Elysia } from "elysia";

export const base = new Elysia()
    .get('/', ({ path }) => path) 
    .post('/hello', 'Do you miss me?') 
    .get('/hi', () => 'hi~')