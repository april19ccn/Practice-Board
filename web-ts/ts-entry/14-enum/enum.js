var Days;
(function (Days) {
    Days[Days["Sun"] = 0] = "Sun";
    Days[Days["Mon"] = 1] = "Mon";
    Days[Days["Tue"] = 2] = "Tue";
    Days[Days["Wed"] = 3] = "Wed";
    Days[Days["Thu"] = 4] = "Thu";
    Days[Days["Fri"] = 5] = "Fri";
    Days[Days["Sat"] = 6] = "Sat";
})(Days || (Days = {}));
;
console.log(Days["Sun"] === 0); // true
console.log(Days["Mon"] === 1); // true
console.log(Days["Tue"] === 2); // true
console.log(Days["Sat"] === 6); // true
// 手动赋值的枚举项可以不是数字，此时需要使用类型断言来让 tsc 无视类型检查 (编译出的 js 仍然是可用的)：
var Days2;
(function (Days2) {
    Days2[Days2["Sun"] = 7] = "Sun";
    Days2[Days2["Mon"] = 10] = "Mon";
    Days2[Days2["Tue"] = 11] = "Tue";
    Days2[Days2["Wed"] = 12] = "Wed";
    Days2[Days2["Thu"] = 13] = "Thu";
    Days2[Days2["Fri"] = 14] = "Fri";
    Days2[Days2["Sat"] = "S"] = "Sat";
})(Days2 || (Days2 = {}));
;
var directions = [0 /* Directions.Up */, 1 /* Directions.Down */, 2 /* Directions.Left */, 3 /* Directions.Right */];
var directions1 = [0 /* Directions.Up */, 1 /* Directions.Down */, 2 /* Directions.Left */, 3 /* Directions.Right */];
