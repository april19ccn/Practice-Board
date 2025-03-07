import * as _ from "ramda"
import { Either, IO, Maybe, compose, chain, map, curry, toUpperCase, Right, Left, either, left } from "../../utils/support.js";
import Task from "data.task";

// 练习 1
// ==========
// 给定一个 user，使用 safeProp 和 map/join 或 chain 安全地获取 sreet 的 name

const safeProp = _.curry(function (x, o) { return Maybe.of(o[x]); });
const user = {
    id: 2,
    name: "albert",
    address: {
        street: {
            number: 22,
            name: 'Walnut St'
        }
    }
};

// const ex1 = (x) => {
//     return compose(map(map(map(safeProp('name')))), map(map(safeProp('street'))), map(safeProp('address')), Maybe.of)(x).$value.$value.$value
// }

const ex1 = compose(chain(safeProp('name')), chain(safeProp('street')), chain(safeProp('address')), Maybe.of)


// console.log(ex1(user))

// 练习 2
// ==========
// 使用 getFile 获取文件名并删除目录，所以返回值仅仅是文件，然后以纯的方式打印文件

var getFile = function () {
    return new IO(function () { return __filename; });
}

var pureLog = function (x) {
    return new IO(function () {
        console.log(x);
        return 'logged ' + x;
    });
}

var ex2 = compose(chain(compose(pureLog, _.last, _.split('\\'))), getFile);



// 练习 3
// ==========
// 使用 getPost() 然后以 post 的 id 调用 getComments()
const getPost = function (i) {
    return new Task(function (rej, res) {
        setTimeout(function () {
            res({ id: i, title: 'Love them tasks' });
        }, 300);
    });
}

const getComments = function (i) {
    return new Task(function (rej, res) {
        setTimeout(function () {
            res([
                { post_id: i, body: "This book should be illegal" },
                { post_id: i, body: "Monads are like smelly shallots" }
            ]);
        }, 300);
    });
}

const ex3 = compose(chain(compose(getComments, _.prop('id'))), getPost);


// 练习 4
// ==========
// 用 validateEmail、addToMailingList 和 emailBlast 实现 ex4 的类型签名

//  addToMailingList :: Email -> IO([Email])
var addToMailingList = (function (list) {
    return function (email) {
        return new IO(function () {
            list.push(email);
            return list;
        });
    }
})([]);

function emailBlast(list) {
    return new IO(function () {
        return 'emailed: ' + list.join(',');
    });
}

var validateEmail = function (x) {
    return x.match(/\S+@\S+\.\S+/) ? (new Right(x)) : (new Left('invalid email'));
}

//  ex4 :: Email -> Either String (IO String)
var ex4 = compose(map(compose(chain(emailBlast), addToMailingList)), validateEmail);

export { user }
export { ex1, ex2, ex3, ex4 }