console.log("traverse-----------")

import { IO, Maybe, compose, chain, map, traverse } from "../../utils/support.js";

// getAttribute :: String -> Node -> Maybe String
const getAttribute = attr => node =>
    node.hasAttribute(attr)
        ? Maybe.of(node.getAttribute(attr))
        : Maybe.empty();

// $ :: Selector -> IO Node
const $ = selector => new IO(() => document.querySelector('#' + selector)); // 实际执行时会真正查询 DOM

// getControlNode :: Selector -> IO (Maybe Node)
const getControlNode = compose(
    // (x) => console.log(x.unsafePerformIO().$value.unsafePerformIO()),
    // map(map($)), // IO(Maybe(IO(node)))

    // (x) => console.log(x.unsafePerformIO().$value),
    chain(traverse(IO.of, $)),       // 关键转换步骤
    map(getAttribute('aria-controls')), // IO(Maybe(node))
    $                                // 初始 IO Node
);

// 查找按钮的控制节点
const program = getControlNode('btn');

console.log(program.unsafePerformIO().$value);

// $ =>
// IO(() => document.querySelector('#btn'))

// map(getAttribute('aria-controls')) =>
// IO(() => document.querySelector('#btn').map(getAttribute('aria-controls')))
// new IO(compose(getAttribute('aria-controls'), () => document.querySelector('#btn')))

// (1) map(map($)) =>
// IO(compose(getAttribute('aria-controls'), () => document.querySelector('#btn'))).map(map($))
// IO(compose(map($), compose(getAttribute('aria-controls'), () => document.querySelector('#btn'))))

// .unsafePerformIO()
// compose(map($), compose(getAttribute('aria-controls'), () => document.querySelector('#btn')))()
// ----> (() => document.querySelector('#btn'))() = node1
// ----> getAttribute('aria-controls')(node1) = Maybe.of(node1.getAttribute('aria-controls')) = Maybe.of('content')
// ----> map($)(Maybe.of('content')) = Maybe.of($('content'))
// = Maybe(IO(() => document.querySelector('#content')))

// .$value
// = IO(() => document.querySelector('#content'))

// .unsafePerformIO()
// = document.querySelector('#content')


// (2) chain(traverse(IO.of, $)) =>
// IO(compose(getAttribute('aria-controls'), () => document.querySelector('#btn'))).map(traverse(IO.of, $)).join()
// new IO(compose(traverse(IO.of, $), compose(getAttribute('aria-controls'), () => document.querySelector('#btn')))).join()

// join() = () => this.unsafePerformIO()
// compose(traverse(IO.of, $), compose(getAttribute('aria-controls'), () => document.querySelector('#btn')))()
// ----> (() => document.querySelector('#btn'))() = node1
// ----> getAttribute('aria-controls')(node1) = Maybe.of(node1.getAttribute('aria-controls')) = Maybe.of('content')
// ----> traverse(IO.of, $)(Maybe.of('content')) 
        // = Maybe.of('content').traverse(IO.of, $) 
        // = $('content').map(Maybe.of) 
        // = IO(() => document.querySelector('#content')).map(Maybe.of)
        // = new IO(compose(Maybe.of, () => document.querySelector('#content')))

// .unsafePerformIO()
// = compose(Maybe.of, () => document.querySelector('#content'))()
// = Maybe.of(document.querySelector('#content'))

// .$value
// = document.querySelector('#content')


// 从宏观角度理解
// $ // 初始 IO Node
// map(getAttribute('aria-controls')), // 得到 IO(Maybe(node1))

// 进一步 通过 从 node1 的 aria-controls 属性获取其属性值对应的元素
// map(map($))
// 打开 IO、再打开 Map， 获取 IO(NODE2)

// chain(traverse(IO.of, $))
// 打开 IO、 取出Maybe的值、$函数处理该值、Maybe仅包裹内部值




// 有个问题哈，IO 不属于 app？ 没有旋转？