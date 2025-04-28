import { test, expect, vi, type Mocked } from 'vitest'

// 1. 定义正确的模拟结构类型
type AxiosMockType = {
  default: {
    get: typeof vi.fn
  }
}

// 2. 使用 hoisted 创建共享模拟实例
const mockedAxios = vi.hoisted(() => ({
  default: {
    get: vi.fn()
  }
}))

// 3. 创建类型安全的模拟对象
type MockedAxios = Mocked<typeof mockedAxios.default>

test("many api", async () => {
  const mockListData = [{ id: 100 }, { id: 200 }]
  const mockGoodsData = { id: 100, name: "测试商品" }

  // 4. 动态模拟模块
  vi.doMock('axios', () => ({
    default: {
      get: mockedAxios.default.get
        .mockResolvedValueOnce({ data: mockListData })
        .mockResolvedValueOnce({ data: mockGoodsData })
    }
  }))

  // 5. 动态导入被测试模块
  const { getGoods } = await import('./index')
  
  // 6. 获取类型安全的模拟对象
  const axiosMock = mockedAxios.default as MockedAxios

  // 7. 执行测试
  const result = await getGoods(999)

  // 8. 验证结果
  expect(result).toEqual(mockGoodsData)
  
  // 9. 验证调用顺序
  expect(axiosMock.get).toHaveBeenNthCalledWith(1, 'url/goods-list')
  expect(axiosMock.get).toHaveBeenNthCalledWith(2, 'url/goods/100')
})