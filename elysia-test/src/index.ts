import { Elysia } from "elysia";
import { swagger } from '@elysiajs/swagger'

import { base } from './base'
import { user } from './user'
import { note } from './note'
import { date } from './date'

const app = new Elysia()
    .use(swagger({
        scalarConfig: {
            defaultHttpClient: {
                targetKey: 'javascript',
                clientKey: 'axios',
            },
        }
    }))
    .use(base)
    .use(user)
    .use(note)
    .use(date)
    .listen(3000);

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`
);
