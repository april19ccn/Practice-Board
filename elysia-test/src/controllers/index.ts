import { Elysia } from "elysia";

import { base } from './base'
import { user } from './user'
import { note } from './note'
import { date } from './date'

export const controller = new Elysia()
    .use(base)
    .use(user)
    .use(note)
    .use(date)