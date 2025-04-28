import { describe, test,expect } from 'vitest'
import { calculateThePrice } from './index'

describe('Test calulate the price', () => {
    test('Test can return expect price', () => {
        // 創建一個產品物件提供測試
        const shoppingCart = [
            { name: 'milk', price: 39, count: 2 },
            { name: 'apple', price: 25, count: 3 },
        ]
        // 確認期望是否正確
        expect(calculateThePrice(shoppingCart)).toBe(153)
    })
})