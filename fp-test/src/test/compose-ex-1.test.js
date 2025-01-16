// npx vitest 启动测试
import { describe, expect, test } from 'vitest'

import * as E from "../ex/compose-ex-1.js"

describe("Compose Exercises 1", () => {
    test("Compose1", async () => {
        expect(await E.getIncompleteTaskSummaries("Punam")).toEqual([
            { "id": 108, "priority": "low", "title": "Adjust the bar", "dueDate": "2013-11-15" }
        ]);
    })

    test("Compose2", async () => {
        expect(await E.getIncompleteTaskSummaries("Scott")).toEqual([
            {
                id: 110, title: "Rename everything",
                dueDate: "2013-11-15", priority: "medium"
            },
            {
                id: 104, title: "Do something",
                dueDate: "2013-11-29", priority: "high"
            }
        ])
    })
})