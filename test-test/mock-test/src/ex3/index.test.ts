import { describe, test, expect, vi } from 'vitest'
import { getAllGoods } from './index'
import axios from 'axios'

// 使用 jset.mock 自動模擬整個 axios 模組
vi.mock('axios')

test('should fetch goods', () => {
    const goods = [{ name: 'Milk' }, { name: 'Apple' }]
    const res = { data: goods };

    // 為 axios 中的 get 模擬回傳值為 res
    // axios 类型增强 「1」
    vi.mocked(axios, { deep: true }).get.mockResolvedValue(res)

    /*
      執行並替回傳值進行斷言，
      這時候 axios 已經被 jest.mock 給模擬了，
      所以 getAllGoods 內的 axios.get 其實不會執行，
      只會回傳用 mockResolvedValue 指定的內容而已
    */
    return getAllGoods().then((resp) => {
        // 從回傳結果中做斷言（第一個產品為 Milk）
        expect(resp[0].name).toEqual('Milk')
    })
})

// 「1」
// vi.mocked
// TypeScript 的类型助手。只返回传入的对象。

// 当 partial 为 true 时，它将期望一个 Partial<T> 作为返回值。
// // vi.mocked(axios.get).mockResolvedValue(res)

// 默认情况下，这只会让 TypeScript 认为第一层的值是模拟的。
// 我们可以将 { deep: true } 作为第二个参数传递给 TypeScript，告诉它整个对象都是模拟的（如果实际上是的话）。
// // vi.mocked(axios, { deep: true }).get.mockResolvedValue(res)