import * as R from "ramda"
import { Task } from "../../utils/support.js";

// Node readfile example:
//=======================

import fs from "fs"

//  readFile :: String -> Task(Error, JSON)
var readFile = function (filename) {
    return new Task(function (reject, result) {
        fs.readFile(filename, 'utf-8', function (err, data) {
            err ? reject(err) : result(data);
        });
    });
};

console.log(readFile("06-test").map(R.split('\n')).map(R.head));
// Task("One morning, as Gregor Samsa was waking up from anxious dreams, he discovered that
// in bed he had been changed into a monstrous verminous bug.")

console.log("compose1 ==> ", R.compose(R.map(R.head), R.map(R.split('\n')))(readFile("06-test")));

// 确保所有传递给R.compose的参数都是函数，输入数据应在组合完成后调用时传入
// console.log("compose2 ==> ", R.compose(R.map(R.head), R.map(R.split('\n')), readFile("06-test")));  // throw new Error('First argument to _arity must be a non-negative integer no greater than ten')
console.log("compose2 ==> ", R.compose(R.map(R.head), R.map(R.split('\n')), () => readFile("06-test"))); // 不执行是 R.compopse整体视为一个函数
console.log("compose2 ==> ", R.compose(R.map(R.head), R.map(R.split('\n')), () => readFile("06-test"))()); // 执行则是 Task { fork: [Function (anonymous)] }

console.log("compose3 ==> ", R.compose(R.map(R.head), R.map(R.split('\n')), readFile));
console.log("compose3 ==> ", R.compose(R.map(R.head), R.map(R.split('\n')), readFile)("06-test"));

// jQuery getJSON example:
//========================

//  getJSON :: String -> {} -> Task(Error, JSON)
var getJSON = R.curry(function(url, params) {
  return new Task(function(reject, result) {
    $.getJSON(url, params, result).fail(reject);
  });
});

console.log(getJSON('/video', {id: 10}).map(R.prop('title')));
// Task("Family Matters ep 15")
console.log(R.compose(R.map(R.prop("title")))(getJSON('/video', {id: 10})));

// 传入普通的实际值也没问题
console.log(Task.of(3).map(function(three){ return three + 1 }));
// Task(4)


// Pure application
//=====================
// blogTemplate :: String

//  blogPage :: Posts -> HTML
var blogPage = Handlebars.compile(blogTemplate);

//  renderPage :: Posts -> HTML
var renderPage = compose(blogPage, sortBy('date'));

//  blog :: Params -> Task(Error, HTML)
var blog = compose(map(renderPage), getJSON('/posts'));


// Impure calling code
//=====================
blog({}).fork(
  function(error){ $("#error").html(error.message); },
  function(page){ $("#main").html(page); }
);

$('#spinner').show();