var tOfM = compose(Task.of, Maybe.of);

liftA2(_.concat, tOfM('Rainy Days and Mondays'), tOfM(' always get me down'));
// Task(Maybe(Rainy Days and Mondays always get me down))

liftA2(_.concat, Task(Maybe('Rainy Days and Mondays')), Task(Maybe(' always get me down')));
Task(Maybe('Rainy Days and Mondays')).map(_.concat).ap(Task(Maybe(' always get me down')));

// Task(Maybe(str))
// Maybe(str) 是上次 this.fork 返回的值， 即 (_, resolve) => resolve(Maybe(str)) 
// (rej, res) => this.fork(rej, compose(res, _.concat))
// (rej, res) => ((_, resolve) => resolve(Maybe(str)))(rej, compose(res, _.concat)) // 替换this.fork
// (rej, res) => compose(res, _.concat)(Maybe(str)) // 替换 resolve
// (rej, res) => res(_.concat(Maybe(str)))

// Task(_.concat(Maybe(str1))).ap(Task(Maybe(str2)));
// Task(_.concat(Maybe(str1))).chain(fn => Task(Maybe(str2)).map(fn));
// new Task((rej_, res) => this.fork(rej_, x => fn(x).fork(rej_, res)))
(rej_, res) => ((_, resolve) => resolve(_.concat(Maybe(str1))))
    (rej_, x => (fn => Task(Maybe(str2)).map(fn))(x).fork(rej_, res)
    );

(rej_, res) => ((_, resolve) =>
    (x => (fn => Task(Maybe(str2)).map(fn))(x).fork(rej_, res))
        (_.concat(Maybe(str1)))
);

(rej_, res) => (x => (fn => Task(Maybe(str2)).map(fn))(x).fork(rej_, res))
    (_.concat(Maybe(str1)));

(rej_, res) => (fn => Task(Maybe(str2)).map(fn))(_.concat(Maybe(str1))).fork(rej_, res);

(rej_, res) => (Task(Maybe(str2)).map(_.concat(Maybe(str1)))).fork(rej_, res);

(rej_, res) => (new Task((rej1, res1) => ((reject, resolve) => resolve(Maybe(str2)))(rej1, compose(res1, _.concat(Maybe(str1)))))).fork(rej_, res);

(rej_, res) => (Task(compose(res1, _.concat(Maybe(str1)))(Maybe(str2)))).fork(rej_, res);

(rej_, res) => (Task(_.concat(Maybe(str1))(Maybe(str2)))).fork(rej_, res);

(rej_, res) => res(_.concat(Maybe(str1))(Maybe(str2)));