// import axios 模組
import axios from 'axios'

// 取得所有產品
export const getAllGoods = () => {
    // 使用 axios 中的 get Function 獲得資料
    return axios.get('url/allGoods').then((resp) => {
        return resp.data
    })
}


// 获取列表中第一个物品的详细信息
export const getGoods = async (id: number) => {
    const {data} = await axios.get('url/goods-list')

    return axios.get(`url/goods/${data[0].id}`).then((resp) => {
        return resp.data
    })
}