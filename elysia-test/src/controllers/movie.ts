import { Elysia, t } from 'elysia'
import { getUserId, userService } from './user'

import { getTitleMovie } from "@/services/movie"

export const movie = new Elysia({ prefix: '/movie' })
    .get(
        "/queryTitle", 
        ({ query: { title } }) => getTitleMovie(title),
        {
            query: t.Object({ title: t.String() })
        }
    )