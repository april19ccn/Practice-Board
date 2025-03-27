(function () {
    // --- 联合类型
    var myFavoriteNumber;
    myFavoriteNumber = 'seven';
    myFavoriteNumber = 7;
    // --- 访问联合类型的属性和方法
    // myFavoriteNumber = true; // 不能将类型“boolean”分配给类型“string | number”
    // function getLength(something: string | number): number {
    //     return something.length;
    // }
    // index.ts(2,22): error TS2339: Property 'length' does not exist on type 'string | number'.
    //   Property 'length' does not exist on type 'number'.
    // 只能访问此联合类型的所有类型里共有的属性或方法
    function getString(something) {
        return something.toString();
    }
    // --- 联合类型的变量在被赋值的时候，会根据类型推论的规则推断出一个类型：
    var myFavoriteNumber1;
    myFavoriteNumber1 = 'seven'; // myFavoriteNumber 被推断成了 string，访问它的 length 属性不会报错
    console.log(myFavoriteNumber1.length); // 5
    myFavoriteNumber1 = 7; // 第四行的 myFavoriteNumber 被推断成了 number，访问它的 length 属性时就报错了。
    // console.log(myFavoriteNumber.length); // 编译时报错
    // index.ts(5,30): error TS2339: Property 'length' does not exist on type 'number'.
})();
