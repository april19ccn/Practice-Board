// 将一个联合类型断言为其中一个类型
interface Cat {
    name: string;
    run(): void;
}
interface Fish {
    name: string;
    swim(): void;
}

function isFish(animal: Cat | Fish) {
    if (typeof (animal as Fish).swim === 'function') {
        return true;
    }
    return false;
}

// 将一个父类断言为更加具体的子类
class ApiError extends Error {
    code: number = 0;
}
class HttpError extends Error {
    statusCode: number = 200;
}

function isApiError(error: Error) {
    if (typeof (error as ApiError).code === 'number') {
        return true;
    }
    return false;
}
// 更合适的方式来判断是不是 ApiError，那就是使用 instanceof
// function isApiError(error: Error) {
//     if (error instanceof ApiError) {
//         return true;
//     }
//     return false;
// }

interface ApiError extends Error {
    code: number;
}
interface HttpError extends Error {
    statusCode: number;
}

function isApiError1(error: Error) {
    if (error instanceof ApiError) {
        return true;
    }
    return false;
}


// 将任何一个类型断言为 any
(window as any).foo = 1;

// 补充：如何扩展window
function getLocalApiUrl(url: any) {

    interface extendsWindow extends Window { // 注意大写
        globalConfig?: any
    }

    let _window: extendsWindow = window
    let baseUrl = _window?.globalConfig?.api || ''


    //  let baseUrl = window?.globalConfig?.api || ''
    return `${baseUrl}${url}`
}


interface Animal {
    name: string;
}
interface Cat {
    name: string;
    run(): void;
}

function testAnimal(animal: Animal) {
    return (animal as Cat);
}
function testCat(cat: Cat) {
    return (cat as Animal);
}