import * as fp from "../index.js";
// import R from "../../node_modules/ramda"
import * as R from 'ramda'
// const R = require('ramda');

const data = {
    result: "SUCCESS",
    tasks: [
        {id: 104, complete: false,            priority: "high",
                  dueDate: "2013-11-29",      username: "Scott",
                  title: "Do something",      created: "9/22/2013"},
        {id: 105, complete: false,            priority: "medium",
                  dueDate: "2013-11-22",      username: "Lena",
                  title: "Do something else", created: "9/22/2013"},
        {id: 107, complete: true,             priority: "high",
                  dueDate: "2013-11-22",      username: "Mike",
                  title: "Fix the foo",       created: "9/22/2013"},
        {id: 108, complete: false,            priority: "low",
                  dueDate: "2013-11-15",      username: "Punam",
                  title: "Adjust the bar",    created: "9/25/2013"},
        {id: 110, complete: false,            priority: "medium",
                  dueDate: "2013-11-15",      username: "Scott",
                  title: "Rename everything", created: "10/2/2013"},
        {id: 112, complete: true,             priority: "high",
                  dueDate: "2013-11-27",      username: "Lena",
                  title: "Alter all quuxes",  created: "10/5/2013"}
    ]
};

// Q:
// 我们需要写一个名为 getIncompleteTaskSummaries 的函数，接收一个 username 作为参数，从服务器获取数据，
// 然后筛选出这个用户的未完成的任务的 ids、priorities、titles、和 dueDate 数据，并且按照日期升序排序。

// 模拟服务器数据
const fetchData = function() {
    return Promise.resolve(data)
};

// 方案1 命令式编程（test：✔）
// export const getIncompleteTaskSummaries = (username) => {
//     return fetchData().then(data => {
//         let unCompleteTasks = [];
//         data.tasks.map(task => {
//             if (!task.completed && username == task.username) {
//                 unCompleteTasks.push({id:task.id, priority: task.priority, title: task.title, dueDate: task.dueDate});
//             }
//         })
//         unCompleteTasks.sort((a,b) => {
//             return new Date(a.dueDate).getTime() - new Date(b.dueDate).getTime();
//         })
//         return unCompleteTasks
//     })
// }

// 方案2 方案1的优化（test：✔）
// export const getIncompleteTaskSummaries = (username) => {
//     return fetchData()
//         .then(data => data.tasks)
//         .then(tasks => tasks.filter(task => !task.completed))
//         .then(tasks => tasks.filter(task => task.username == username))
//         .then(tasks => tasks.map(task => {
//             delete task.username
//             delete task.complete
//             delete task.created
//             return task
//         }))
//         .then(tasks => tasks.sort((a, b) => new Date(a.dueDate).getTime() - new Date(b.dueDate).getTime()))
// }


// 方案3 函数式（test：✔）
// const getObjectValue = fp.curry(function(key, obj) {
//     return obj[key];
// });
// const filter = fp.curry(function(fp, arr) {
//     return arr.filter(fp)
// })
// const isTrue = fp.curry(function(key, obj) {
//     return !obj[key];
// })
// const isEqual = fp.curry(function(key, value, x) {
//     return x[key] === value;
// })
// const map = fp.curry(function(fp, arr) {
//     return arr.map(fp);
// })
// const deleteValue = fp.curry(function(key, obj) {
//     delete obj[key];
//     return obj;
// })
// const sort = fp.curry(function(fp, arr) {
//     return arr.sort(fp);
// })
// const timeSort = (a, b) => new Date(a.dueDate).getTime() - new Date(b.dueDate).getTime()

// export const getIncompleteTaskSummaries = async (username) => {
//     return fp.compose(
//         sort(timeSort),
//         map(fp.compose(deleteValue('username'), deleteValue('complete'), deleteValue('created'))),
//         filter(isEqual('username', username)),
//         filter(isTrue('complete')),
//         getObjectValue('tasks')
//     )(await fetchData())
// }

// 方案4（用y-fp实现）



// 方案5（用ramda实现）
export const getIncompleteTaskSummaries = async (username) => {
    return R.compose(
        R.sort(R.ascend(R.prop('dueDate'))),
        R.map(R.pick(['id', 'priority', 'title', 'dueDate'])),
        R.filter(R.propEq(username, 'username')),
        R.filter(R.propEq(false, 'complete')),
        R.prop("tasks")
    )(await fetchData())
}


// getIncompleteTaskSummaries("Punam")