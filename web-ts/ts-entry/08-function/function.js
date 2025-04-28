// 函数声明
function sum1(x, y) {
    return x + y;
}
// 函数表达式
var mySum = function (x, y) {
    return x + y;
};
var mySearch;
mySearch = function (source, subString) {
    return source.search(subString) !== -1;
};
// 可选参数
// 必选参数不能位于可选参数后。
// function buildName(firstName?: string, lastName: string) {
//     if (firstName) {
//         return firstName + ' ' + lastName;
//     } else {
//         return lastName;
//     }
// }
// let tomcat = buildName('Tom', 'Cat');
// let tom1 = buildName(undefined, 'Tom');
// 参数默认值
function buildName(firstName, lastName) {
    if (firstName === void 0) { firstName = 'Tom'; }
    return firstName + ' ' + lastName;
}
var tomcat = buildName('Tom', 'Cat');
var cat = buildName(undefined, 'Cat');
// 剩余参数
function push(array) {
    var items = [];
    for (var _i = 1; _i < arguments.length; _i++) {
        items[_i - 1] = arguments[_i];
    }
    items.forEach(function (item) {
        array.push(item);
    });
}
var a = [];
push(a, 1, 2, 3);
function reverse(x) {
    if (typeof x === 'number') {
        return Number(x.toString().split('').reverse().join(''));
    }
    else if (typeof x === 'string') {
        return x.split('').reverse().join('');
    }
}
