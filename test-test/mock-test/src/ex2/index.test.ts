import { describe, test, expect, vi } from 'vitest'
import { calculateThePrice } from './index'

describe('Test calulate the price', () => {
    // 創建一個產品物件提供測試
    const shoppingCart = [
        { name: 'milk', price: 39, count: 2 },
        { name: 'apple', price: 25, count: 3 },
    ]

    // 方法1-建立 Mock 取代 CheckDiscount
    // Mock<Procedure> 类型解释 「2」
    const mockCheckDiscount = vi.fn()

    // 1-設定回傳值
    mockCheckDiscount
        .mockReturnValueOnce(true)
        .mockReturnValue(false)

    // 方法2-建立 Mock 取代 CheckDiscount，并设置 mock 函数
    // const mockCheckDiscount = vi.fn((name: string) => {
    //     if (name === 'milk') {
    //         return true
    //     }
    //     return false
    // })


    test('Test can return expect price', () => {
        // 確認期望是否正確
        expect(calculateThePrice(shoppingCart, mockCheckDiscount)).toBe(114) // 「2」为什么 mockCheckDiscount: Mock<Procedure> 可以赋给 (name: string) => boolean 类型的参数
    })

    // 使用 mock.calls 「1」
    test('Test execute several times of checkDiscount', () => {
        // 確認判斷折扣這件事確實執行了兩次
        expect(mockCheckDiscount.mock.calls.length).toBe(2)
    })

    test('Test mock receive goods name is real', () => {
        // 確認 Mock 接收到正確的產品名稱( calls[第幾次執行][第幾個參數] )
        expect(mockCheckDiscount.mock.calls[0][0]).toBe('milk')
    })
})

// 「1」
// mock.calls
// 这是一个数组，包含了每次调用的所有参数。数组中的每个项目都是那次调用的参数。
// const fn = vi.fn()

// fn('arg1', 'arg2')
// fn('arg3')

// fn.mock.calls
// === [
//     ['arg1', 'arg2'], // first call
//     ['arg3'], // second call
// ]

// 「2」
// interface Mock<T extends Procedure = Procedure> extends MockInstance<T> {
//     new (...args: Parameters<T>): ReturnType<T>;
//     (...args: Parameters<T>): ReturnType<T>;
// }
// type Parameters<T extends (...args: any) => any> = T extends (...args: infer P) => any ? P : never;