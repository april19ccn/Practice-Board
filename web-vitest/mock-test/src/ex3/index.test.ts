import { beforeEach, describe, test, expect, vi } from 'vitest'
import { getAllGoods, getGoods } from './index'
import axios from 'axios'

// 使用 vi.mock 自動模擬整個 axios 模組
vi.mock('axios')

beforeEach(() => {
    // 每次测试前重置所有 Mock 状态
    vi.mocked(axios.get).mockReset()
})

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

test("many api", async () => {
    const mockListData = [{ id: 100 }, { id: 200 }]
    const mockGoodsData = { id: 100, name: "测试商品" }

    // 1. 分两次模拟 axios.get 的返回值
    vi.mocked(axios.get)
        .mockResolvedValueOnce({ data: mockListData }) // 第一次调用返回列表
        .mockResolvedValueOnce({ data: mockGoodsData }) // 第二次调用返回商品详情

    // 2. 执行待测函数
    const result = await getGoods(999) // 注意：参数 999 实际未使用

    // 3. 验证结果
    expect(result).toEqual(mockGoodsData)

    // 4. 验证调用顺序和参数
    expect(axios.get).toHaveBeenNthCalledWith(1, 'url/goods-list')
    expect(axios.get).toHaveBeenNthCalledWith(2, 'url/goods/100')
})


// 「1」
// vi.mocked
// TypeScript 的类型助手。只返回传入的对象。

// 当 partial 为 true 时，它将期望一个 Partial<T> 作为返回值。
// // vi.mocked(axios.get).mockResolvedValue(res)

// 默认情况下，这只会让 TypeScript 认为第一层的值是模拟的。
// 我们可以将 { deep: true } 作为第二个参数传递给 TypeScript，告诉它整个对象都是模拟的（如果实际上是的话）。
// // vi.mocked(axios, { deep: true }).get.mockResolvedValue(res)