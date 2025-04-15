// import axios 模組
import axios from 'axios'

// 取得所有產品
export const getAllGoods = () => {
    // 使用 axios 中的 get Function 獲得資料
    return axios.get('url/allGoods').then((resp) => {
        return resp.data
    })
}