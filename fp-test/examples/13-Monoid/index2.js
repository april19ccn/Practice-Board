// ------------ 模拟一个简易的 Stream 类（类似 RxJS 的 Observable） ------------
class Stream {
    constructor(subscribe) {
        this.subscribe = subscribe;
    }

    static fromEvent(eventName, element) {
        return new Stream((observer) => {
            const handler = (e) => observer.next(e);
            element.addEventListener(eventName, handler);
            return () => element.removeEventListener(eventName, handler);
        });
    }

    filter(predicate) {
        return new Stream((observer) => {
            return this.subscribe({
                next: (value) => {
                    if (predicate(value)) observer.next(value);
                },
                error: (err) => observer.error(err),
                complete: () => observer.complete(),
            });
        });
    }

    concat(otherStream) {
        return new Stream((observer) => {
            let firstCompleted = false;
            const sub1 = this.subscribe({
                next: (value) => observer.next(value),
                error: (err) => observer.error(err),
                complete: () => {
                    firstCompleted = true;
                    sub2 = otherStream.subscribe(observer);
                },
            });

            let sub2 = null;
            return () => {
                sub1.unsubscribe?.();
                sub2?.unsubscribe?.();
            };
        });
    }

    map(transform) {
        return new Stream((observer) => {
            return this.subscribe({
                next: (value) => observer.next(transform(value)),
                error: (err) => observer.error(err),
                complete: () => observer.complete(),
            });
        });
    }
}

// ------------ 模拟 DOM 元素和事件 ------------
// 假设页面中存在以下元素：
// <button id="submit">Submit</button>
// <form id="myForm"><input type="text"></form>
const $ = (selector) => document.querySelector(selector);

// ------------ submitForm 函数的作用 ------------
// 当用户点击提交按钮或按下回车时，收集表单数据并提交到服务器
const submitForm = (event) => {
    // 阻止默认行为（如表单自动提交）
    event.preventDefault();

    // 收集表单数据（假设表单中有一个输入框）
    const formData = {
        username: $("#myForm input").value,
    };

    // 模拟提交到服务器（例如发送 HTTP 请求）
    console.log("提交数据：", formData);
    simulateServerRequest(formData)
        .then((response) => console.log("成功：", response))
        .catch((error) => console.error("失败：", error));
};

// ------------ 模拟服务器请求 ------------
const simulateServerRequest = (data) => {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            if (data.username) {
                resolve({ status: 200, message: "数据已接收" });
            } else {
                reject({ status: 400, message: "用户名不能为空" });
            }
        }, 1000);
    });
};

// ------------ 事件流的组合与使用 ------------
// 创建两个事件流：
// 1. 点击提交按钮的事件流
const submitStream = Stream.fromEvent("click", $("#submit"));
// 2. 在表单中按下回车键的事件流（过滤非回车事件）
const enterStream = Stream.fromEvent("keydown", $("#myForm"))
    .filter((event) => event.key === "Enter");

// 合并两个事件流，并映射到 submitForm 函数
const combinedStream = submitStream
    .concat(enterStream)
    .map(submitForm);

// 订阅合并后的事件流
const subscription = combinedStream.subscribe({
    next: (event) => {
        // 这里可以添加其他副作用操作（例如显示加载状态）
        console.log("事件触发：", event.type);
    },
    error: (err) => console.error("流发生错误：", err),
    complete: () => console.log("流已完成"),
});

// 取消订阅（例如在组件卸载时调用）
// subscription.unsubscribe();