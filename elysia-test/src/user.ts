import { Elysia, t } from 'elysia'

export const user = new Elysia({ prefix: '/user' })
    .state({
        user: {} as Record<string, string>,
        session: {} as Record<number, string>
    })

    // 注册
    .put(
        '/sign-up',
        async ({ body: { username, password }, store, error }) => {
            if (store.user[username])
                return error(400, {
                    success: false,
                    message: 'User already exists'
                })

            // insert a username and hashed password with argon2id
            store.user[username] = await Bun.password.hash(password)

            return {
                success: true,
                message: 'User created'
            }
        },
        {
            body: t.Object({
                username: t.String({ minLength: 1 }),
                password: t.String({ minLength: 8 })
            })
        }
    )

    // 登录
    .post(
        '/sign-in',
        async ({
            store: { user, session },
            error,
            body: { username, password },
            cookie: { token }
        }) => {
            if (
                !user[username] ||
                !(await Bun.password.verify(password, user[username]))
            )
                return error(400, {
                    success: false,
                    message: 'Invalid username or password'
                })

            const key = crypto.getRandomValues(new Uint32Array(1))[0]
            session[key] = username
            token.value = key

            return {
                success: true,
                message: `Signed in as ${username}`
            }
        },
        {
            body: t.Object({
                username: t.String({ minLength: 1 }),
                password: t.String({ minLength: 8 })
            }),
            cookie: t.Cookie(
                {
                    token: t.Number()
                },
                {
                    secrets: 'seia'
                }
            )
        }
    )