sequence(List.of, Maybe.of(['the facts'])); // [Just('the facts')]

sequence(Task.of, new Map({ a: Task.of(1), b: Task.of(2) })); // Task(Map({ a: 1, b: 2 }))

sequence(IO.of, Either.of(IO.of('buckle my shoe'))); // IO(Right('buckle my shoe'))

sequence(Either.of, [Either.of('wing')]); // Right(['wing'])

sequence(Task.of, left('wing')); // Task(Left('wing'))




// fromPredicate :: (a -> Bool) -> a -> Either e a

// Map
// partition :: (a -> Bool) -> [a] -> [Either e a]
const partition = (f) => map(fromPredicate(f));

// Traverse
// validate :: (a -> Bool) -> [a] -> Either e [a]
const validate = (f) => traverse(Either.of, fromPredicate(f));


// traverse(of, fn) {
//     return this.$value.reduce(
//         (f, a) => fn(a).map(b => bs => bs.concat(b)).ap(f),
//         of(new List([])),
//     );
// }


// 类型签名
// fromPredicate :: (a -> Bool) -> a -> Either e a
const fromPredicate = (predicate) => (x) =>
    predicate(x) ? Right(x) : Left(`验证失败: ${x}`);

// validate :: (a -> Bool) -> [a] -> Either e [a]
const validate1 = (f) => traverse(Either.of, fromPredicate(f));

const isEven = x => x % 2 === 0;

// 用例 1：全部偶数
validate1(isEven)([2, 4, 6]);
// 结果: Right([2, 4, 6])

// 用例 2：包含奇数
validate1(isEven)([2, 3, 4]);
// 结果: Left("验证失败: 3")

traverse(Either.of, fromPredicate(isEven))([2, 3, 4])
    = [2, 3, 4].reduce(
        (f, a) => fn(a).map(b => bs => bs.concat(b)).ap(f),
        Either.of(new List([])),
    )

// a = 2 ：

// (Either.of([]), 2) => fromPredicate(isEven)(2).map(b => bs => bs.concat(b)).ap(Either.of([]))

// // fromPredicate(isEven)(2) => Right(2)
// (Either.of([]), 2) => Right(2).map(b => bs => bs.concat(b)).ap(Either.of([]))

// // Right(2).map(b => bs => bs.concat(b))
// (Either.of([]), 2) => Right(bs => bs.concat(2)).ap(Either.of([]))

// // Either.of(x) => new Right(x)
// (Right([]), 2) => Right(bs => bs.concat(2)).ap(Right([]))


// reduce :: [a] -> (f -> a -> f) -> f -> f
// - `[a]`对应`this.$value`（比如例子中的数组`[2, 3, 4]`）

// - `(f -> a -> f)`对应回调函数`(f, a) => fn(a).map(...).ap(f)`

// - 第一个`f`是初始值`of(new List([]))`（比如`Either.of([])`）

// - 第二个`f`是最终返回的结果（比如`Right([2])`或`Left(...)`）



/////////////////////////////////////////////

// getAttribute :: String -> Node -> Maybe String
// $ :: Selector -> IO Node

// getControlNode :: Selector -> IO (Maybe Node)
// const getControlNode = compose(
//     chain(traverse(IO.of, $)),
//     map(getAttribute('aria-controls')),
//     $
// );

// getControlNode("test")

// $("test") -> IO(Node)
// IO(Node) -> IO(Maybe(Node-'aria-controls'))
// IO(Maybe(Node-'aria-controls')).chain(traverse(IO.of, $)) 
// -> IO(Maybe(Node-'aria-controls')).map(traverse(IO.of, $)).join()
// -> IO(traverse(IO.of, $)(Maybe(Node-'aria-controls'))).join()
// -> IO(Maybe($ Node-'aria-controls'))


