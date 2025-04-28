import { describe, test, expect, vi } from 'vitest'
import { productModule, calculateThePrice } from './index'

describe('Test calulate the price', () => {
    const shoppingCart = [
        { name: 'milk', price: 39, count: 2 },
        { name: 'apple', price: 25, count: 3 },
    ]

    // 以 objectName 及 methodName 創建 spy 替身
    const spyCheckDiscount = vi.spyOn(productModule, 'checkDiscount')

    test('check DOC', () => { // 让其通过 spyOn 真正执行 DOC 逻辑，以防止 DOC 被修改
        // 將 spy 送入測試
        expect(calculateThePrice(shoppingCart)).toBe(114)
    })


    test('Test can return expect price', () => {
        // 設定假資料回傳
        // 如果直接為 spyCheckDiscount 設定假的回傳值，就不會有 console.log 被印出，因為它變得只會依照設定的資料回傳給 SUT 而已
        spyCheckDiscount
            .mockReturnValueOnce(false)
            .mockReturnValueOnce(false)

        // 將 spy 送入測試
        expect(calculateThePrice(shoppingCart)).toBe(153)
    })
})