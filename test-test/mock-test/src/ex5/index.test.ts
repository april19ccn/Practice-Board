import { describe, test, expect, vi } from "vitest";
import { lottery } from "./index"

describe("Test Lottery", () => {

    test("Test lottery", () => {
        const mockRandom = vi.spyOn(Math, "random") 

        mockRandom
            .mockReturnValueOnce(0.5)
            .mockReturnValue(0.7) // 前一个必须用 Once，否则后者会取代前者

        expect(lottery()).toBe("LOSE")
        expect(lottery()).toBe("WIN")

        console.log(mockRandom.mock.calls)

        expect(mockRandom.mock.calls.length).toBe(2)
    })
})

