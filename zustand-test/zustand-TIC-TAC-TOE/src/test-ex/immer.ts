import { produce } from "immer"

const baseState = [
    {
        title: "Learn TypeScript",
        done: true
    },
    {
        title: "Try Immer",
        done: false
    }
]

// const nextState = produce(baseState, draftState => {
//     draftState.push({ title: "Tweet about it", done: false })
//     draftState[1].done = true
// })

// 得到一个新值 =>
// nextState = [
//     {
//         "title": "Learn TypeScript",
//         "done": true
//     },
//     {
//         "title": "Try Immer",
//         "done": true
//     },
//     {
//         "title": "Tweet about it",
//         "done": false
//     }
// ]

const nextState = produce(draftState => {
    draftState.push({ title: "Tweet about it", done: false })
    draftState[1].done = true
})

const nextStateValue = nextState(baseState)

// export default nextState


function a1(this: any, a: number, b: number) {
    return a + b
}

console.log(a1(1, 2))

export default nextStateValue