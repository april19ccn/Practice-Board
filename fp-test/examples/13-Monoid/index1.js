// Map模拟，合并对象属性
const Map = (obj) => ({
    ...obj,
    concat(other) {
        return Map({ ...this, ...other });
    }
});

// Either类型（Right/Left）
class Either {
    static Right(value) {
        return new Right(value);
    }
    static Left(error) {
        return new Left(error);
    }
}

class Right extends Either {
    constructor(value) {
        super();
        this.value = value;
    }
    concat(other) {
        return other.fold(
            e => other, // 保留Left错误
            v => new Right(this.value.concat(v)) // 合并Right的值
        );
    }
    fold(_, g) {
        return g(this.value);
    }
}

class Left extends Either {
    constructor(error) {
        super();
        this.error = error;
    }
    concat(other) {
        return this; // Left优先，不合并错误
    }
    fold(f) {
        return f(this.error);
    }
}

// IO类型，执行副作用并合并结果
class IO {
    constructor(effect) {
        this.effect = effect;
    }
    map(f) {
        return new IO(() => f(this.effect()));
    }
    concat(other) {
        return new IO(() => this.effect().concat(other.effect()));
    }
    run() {
        return this.effect();
    }
}

// 模拟formValues，返回IO(Map)
const formValues = (selector) => new IO(() => {
    switch (selector) {
        case '#signup': return Map({ username: 'andre3000' });
        case '#terms': return Map({ accepted: true });
        default: return Map({});
    }
});

// 验证函数，返回Either
const validate = (data) =>
    data.accepted === false
        ? Either.Left('必须接受条款')
        : Either.Right(data);

// 示例1：合并验证结果
const example1 = formValues('#signup')
    .map(validate)
    .concat(formValues('#terms').map(validate));

console.log(example1);
console.log(example1.run());
// Right { value: { username: 'andre3000', accepted: true } }

// 示例2：其中一个验证失败
const formReject = () => new IO(() => Map({ accepted: false }));
const example2 = formReject().map(validate)
    .concat(formValues('#signup').map(validate));

console.log(example2.run());
// Left { error: '必须接受条款' }

// Task类型，合并异步结果
class Task {
    constructor(fork) {
        this.fork = fork;
    }
    concat(other) {
        return new Task((rej, res) =>
            this.fork(
                e => rej(e),
                a => other.fork(
                    e => rej(e),
                    b => res(a.concat(b))
                )
            )
        );
    }
}

// 模拟API请求
const serverA = { get: () => new Task((_, res) => res(['friend1'])) };
const serverB = { get: () => new Task((_, res) => res(['friend2'])) };

serverA.get().concat(serverB.get())
    .fork(console.error, console.log); // ['friend1', 'friend2']

// Maybe类型，合并可能存在的结果
class Maybe {
    static Just(val) { return new Just(val); }
    static Nothing() { return new Nothing(); }
}

class Just extends Maybe {
    constructor(val) {
        super();
        this.value = val;
    }
    concat(other) {
        return other.fold(
            () => this,
            v => new Just(this.value.concat(v))
        );
    }
    fold(_, f) { return f(this.value); }
}

class Nothing extends Maybe {
    concat(other) { return other; }
    fold(f) { return f(); }
}

// 模拟加载设置
const loadSetting = (key) => new Task((_, res) =>
    setTimeout(() => res(
        key === 'email' ? Maybe.Just(Map({ autoSave: false }))
            : Maybe.Just(Map({ background: true }))
    ), 100)
);

loadSetting('email').concat(loadSetting('general'))
    .fork(console.error, maybe =>
        maybe.fold(
            () => console.log('Nothing'),
            map => console.log(map)
        )
    ); // { autoSave: false, background: true }