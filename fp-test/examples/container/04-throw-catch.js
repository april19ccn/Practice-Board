import moment from "moment"
import * as R from "ramda"

//////////////////////////////////////// “纯”错误处理
var Left = function (x) {
    this.__value = x;
}

Left.of = function (x) {
    return new Left(x);
}

Left.prototype.map = function (f) {
    return this;
}

var Right = function (x) {
    this.__value = x;
}

Right.of = function (x) {
    return new Right(x);
}

Right.prototype.map = function (f) {
    return Right.of(f(this.__value));
}

console.log(Right.of("rain").map(function (str) { return "b" + str; }));
// Right("brain")

console.log(Left.of("rain").map(function (str) { return "b" + str; }));
// Left("rain")

console.log(Right.of({ host: 'localhost', port: 80 }).map(R.prop('host')));
// Right('localhost')

console.log(Left.of("rolls eyes...").map(R.prop("host")));
// Left('rolls eyes...')


////// 案例
//  getAge :: Date -> User -> Either(String, Number)
var getAge = R.curry(function (now, user) {
    var birthdate = moment(user.birthdate, 'YYYY-MM-DD');
    if (!birthdate.isValid()) return Left.of("Birth date could not be parsed");
    return Right.of(now.diff(birthdate, 'years'));
});

console.log(getAge(moment(), { birthdate: '2005-12-12' }));
// Right(9)

console.log(getAge(moment(), { birthdate: 'balloons!' }));
// Left("Birth date could not be parsed")


//  fortune :: Number -> String
var fortune = R.compose(R.concat("If you survive, you will be "), R.toString(), R.add(1));

//  zoltar :: User -> Either(String, _)
var zoltar = R.compose(R.map(console.log), R.map(fortune), getAge(moment()));

console.log(zoltar({ birthdate: '2005-12-12' }));
// "If you survive, you will be 10"
// Right(undefined) // R.map(console.log) 取出容器值打印之后，没有给容器新值，导致undefined

console.log(zoltar({ birthdate: 'balloons!' }));
// Left("Birth date could not be parsed")



//  either :: (a -> c) -> (b -> c) -> Either a b -> c
var either = R.curry(function (f, g, e) {
    switch (e.constructor) {
        case Left: return f(e.__value);
        case Right: return g(e.__value);
    }
});

//  zoltar :: User -> _
var zoltar = R.compose(console.log, either(R.identity, fortune), getAge(moment()));

console.log(zoltar({ birthdate: '2005-12-12' }));
// "If you survive, you will be 10"
// undefined

console.log(zoltar({ birthdate: 'balloons!' }));
// "Birth date could not be parsed"
// undefined



///// 测试之前Maybe例子
const getArrayHead = (arr) => {
    if (arr.length === 0) { return Left.of("array is empty") }
    else { return Right.of(arr[0]) }
}

var streetName_Either = R.compose(R.map(R.prop('street')), getArrayHead, R.prop('addresses'));

console.log(streetName_Either({ addresses: [] }));
// Maybe(null)

console.log(streetName_Either({ addresses: [{ street: "Shady Ln. Either", number: 420110 }] }));


