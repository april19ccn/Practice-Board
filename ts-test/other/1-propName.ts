// [propName: string]: any;
// 允许有任意的属性

// [propName: string]: string;
// 一旦定义了任意属性，那么确定属性和可选属性的类型都必须是它的类型的子集
// 故：类型“number”的属性“age”不能赋给“string”索引类型“string”

// 一个接口中只能定义一个任意属性。如果接口中有多个类型的属性，则可以在任意属性中使用联合类型
interface PersonOther1 {
    name: string;
    age?: number;
    [propName: string]: string | number;
}

interface PersonOther12 {
    name: string;
    age?: number;
    x: () => void;
    // 1: () => void; // 类型“() => void”的属性“1”不能赋给“number”索引类型“string | number”。
    [propName: number]: string | number;
}

let tomOther1: Person = {
    name: 'Tom',
    age: 25,
    gender: 'male'
};


// 为什么  [propName: string]: string | number; 会限制确定和可选属性，而 [index: number]: number; 不会呢？
// 在 TypeScript 中，字符串索引签名和数字索引签名对属性的约束行为不同，这与 JavaScript 的底层对象模型和类型系统的设计有关
// 关键区别总结
// 索引签名类型	         约束范围	                      与显式属性的关系
// 字符串索引签名	所有字符串键属性（包括显式属性）	显式属性必须兼容索引签名类型
// 数字索引签名	        仅数字键属性	                显式字符串键属性不受影响

// [propName: number] ---------------------------
interface NumberArrayOther1 {
    [index: number]: number;
    length: number;
    callee: Function;
}


function sumOther1() {
    let args: {
        [index: number]: number;
        length: number;
        callee: Function;
    } = arguments;
}

type Person1 =  {
    name: string;
    age?: number;
    // 1: () => void; // 类型“() => void”的属性“1”不能赋给“number”索引类型“string | number”。
    [propName: number]: string | number;
}



// 任意属性 和 映射是不一样的 --------------

type test_text = 'a' | 'b'

type test_text2 = {
    // name: string; 映射的类型可能不声明属性或方法
    [propName in test_text]: string | number;
}



type MoviesByGenre = {
    action: 'Die Hard';
    comedy: 'Groundhog Day';
    sciFi: 'Blade Runner';
    fantasy: 'The Lord of the Rings: The Fellowship of the Ring';
    drama: 'The Shawshank Redemption';
    horror: 'The Shining';
    romance: 'Titanic';
    animation: 'Toy Story';
    thriller: 'The Silence of the Lambs';
  };

type MovieInfoByGenre<T> = {
	[K in keyof T] : {
		name: T[K],
		year: number,
		director:string
	}
};

type Example = MovieInfoByGenre<MoviesByGenre>;