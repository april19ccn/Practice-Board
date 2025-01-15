import * as fp from "./my-fp.js"

// ==== 1
// var words = function(str) {
//     return split(' ', str);
// };

var words = fp.split(' '); 

console.log(words("hello world"))

export { words }