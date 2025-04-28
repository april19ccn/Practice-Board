// 將 checkDiscount 放到一個 module 中
const productModule = {
    checkDiscount: (name) => {
        console.log(name)
        if (name === 'milk') {
            return true
        }
        return false
    },
}

const calculateThePrice = (goods) => {
    let totalPrice = 0
    goods.forEach((item) => {
        let price = Number(item.price) * Number(item.count)
        if (productModule.checkDiscount(item.name)) {
            price *= 0.5
        }
        totalPrice += price
    })
    return totalPrice
}

export {
    productModule,
    calculateThePrice
}