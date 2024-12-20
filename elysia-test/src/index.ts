import { Elysia } from "elysia";
import { opentelemetry } from '@elysiajs/opentelemetry'
import { swagger } from '@elysiajs/swagger'

import { base } from './controllers/base'
import { user } from './controllers/user'
import { note } from './controllers/note'
import { date } from './controllers/date'

import { controller } from "./controllers";

export const app = new Elysia()
    .use(opentelemetry()) 
    .use(swagger({
        scalarConfig: {
            defaultHttpClient: {
                targetKey: 'javascript',
                clientKey: 'axios',
            },
        }
    }))
    .onError(({ error, code }) => { 
        if (code === 'NOT_FOUND') return 'Not Found :('

        console.error(error) 
    }) 
    // .use(base)
    // .use(user)
    // .use(note)
    // .use(date)
    .use(controller)
    .listen(3000);

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`
);
