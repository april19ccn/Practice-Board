// getValue :: Selector -> Task Error (Maybe String)
// postComment :: String -> Task Error Comment
// validate :: String -> Either ValidationError String

// saveComment :: () -> Task Error (Maybe (Either ValidationError (Task Error Comment)))
const saveComment = compose(
    map(map(map(postComment))),
    map(map(validate)), // Task(Maybe(str)).map(map(validate)) => Task(Maybe(validate(str)))
    getValue('#comment'), // 如果成功 Task(Maybe(str)), 如果失败 Task(Error)
);



///////////////////
// promiseToTask :: Promise a b -> Task a b
const promiseToTask = x => new Task((reject, resolve) => x.then(resolve).catch(reject));

// taskToPromise :: Task a b -> Promise a b
const taskToPromise = x => new Promise((resolve, reject) => x.fork(reject, resolve));

const x = Promise.resolve('ring'); // <=> new Promise((resolve, reject) => resolve('ring'))
taskToPromise(promiseToTask(x)) === x;
// promiseToTask(Promise.resolve('ring')) => newTask((reject, resolve) => Promise.resolve('ring').then(resolve).catch(reject));
// taskToPromise() => new Promise((resolve, reject) => Promise.resolve('ring').then(resolve).catch(reject)) // 这里的 resolve 直接将 x 的结果传递给 convertedPromise，无中间处理逻辑，因此二者行为完全一致。

const y = Task.of('rabbit');
promiseToTask(taskToPromise(y)) === y;




// getValue :: Selector -> Task Error (Maybe String)
// postComment :: String -> Task Error Comment
// validate :: String -> Either ValidationError String

// saveComment1 :: () -> Task Error Comment
const saveComment1 = compose(
    chain(postComment), // Task(str2)
    chain(eitherToTask), // Task(str1)
    map(validate), //  Task(Either(str))
    chain(maybeToTask), // Task(Maybe(str)).map(maybeToTask).join() => Task(str)
    getValue('#comment'), // 如果成功 Task(Maybe(str)), 如果失败 Task(Error)
);