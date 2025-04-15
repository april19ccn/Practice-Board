// 計算購買產品的總額
export const calculateThePrice = (goods: { name: string, price: number, count: number }[]) => {
    let totalPrice = 0
    goods.forEach((item) => {
        totalPrice += Number(item.price) * Number(item.count)
    })
    return totalPrice
}