// 判斷該產品是否有折扣
// checkDiscount 被稱為 SUT 的 DOC （ Depended-on Component 依賴組件）
const checkDiscount = (name: string) => {
    if (name === 'milk') {
        return true
    }
    return false
}

// 計算購買產品的總額
// 被測試的函式 calculateThePrice 被称为 SUT （ System Under Test 測試目標）
const calculateThePrice = (goods: { name: string, price: number, count: number }[], checkDiscount: (name: string) => boolean) => {
    let totalPrice = 0
    goods.forEach((item) => {
        // 先計算原價
        let price = Number(item.price) * Number(item.count)

        // 如果有折扣要半價
        if (checkDiscount(item.name)) {
            price *= 0.5
        }

        // 將價格加到總合上
        totalPrice += price
    })
    return totalPrice
}

export { calculateThePrice }