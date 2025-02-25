import { Elysia, t } from 'elysia'

export const userService = new Elysia({ name: 'user/service' }) 
	.state({ 
        user: {} as Record<string, string>, 
        session: {} as Record<number, string> 
    }) 
    .model({ 
        signIn: t.Object({ 
            username: t.String({ minLength: 1 }), 
            password: t.String({ minLength: 8 }) 
        }), 
        session: t.Cookie( 
            { 
                token: t.Number() 
            }, 
            { 
                secrets: 'seia'
            } 
        ) 
    }) 
    .model((model) => ({ 
        ...model, 
        optionalSession: t.Optional(model.session) 
    }))
    .macro(({ onBeforeHandle }) => ({ 
        isSignIn(enabled: boolean) { 
            if (!enabled) return

            onBeforeHandle( 
                ({ error, cookie: { token }, store: { session } }) => { 
                    if (!token.value) 
                        return error(401, { 
                            success: false, 
                            message: 'Unauthorized'
                        }) 

                    const username = session[token.value as unknown as number] 

                    if (!username) 
                        return error(401, { 
                            success: false, 
                            message: 'Unauthorized'
                        }) 
                } 
            ) 
        } 
    })) 


export const getUserId = new Elysia() 
    .use(userService) 
    .guard({ 
        cookie: 'session'
    }) 
    .resolve(({ store: { session }, cookie: { token } }) => ({ 
        username: session[token.value] 
    }))
    .as('plugin') 


export const user = new Elysia({ prefix: '/user' })
    .use(userService)

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
            body: 'signIn'
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
            if (!user[username] || !(await Bun.password.verify(password, user[username])))
                return error(400, {
                    success: false,
                    message: 'Invalid username or password'
                })

            const key = crypto.getRandomValues(new Uint32Array(1))[0]
            session[key] = username
            token.value = key

            console.log("token", key)

            return {
                success: true,
                message: `Signed in as ${username}`
            }
        },
        {
            body: 'signIn',
            cookie: 'optionalSession',
        }
    )

    // 登出
    .get( 
        '/sign-out', 
        ({ cookie: { token } }) => { 
            token.remove() 

            return { 
                success: true, 
                message: 'Signed out'
            } 
        }, 
        { 
            cookie: 'optionalSession'
        } 
    )

    // 获取信息
    .use(getUserId)
    .get( 
        '/profile', 
        (({ username }) => ({ 
            success: true, 
            username
        }))
    ) 