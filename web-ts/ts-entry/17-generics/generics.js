var createArray;
createArray = function (length, value) {
    var result = [];
    for (var i = 0; i < length; i++) {
        result[i] = value;
    }
    return result;
};
var x17 = createArray(3, 'x'); // ['x', 'x', 'x']
console.log(x17);
// 泛型参数的默认类型
function createArray171(length, value) {
    var result = [];
    for (var i = 0; i < length; i++) {
        result[i] = value;
    }
    return result;
}
