// 用来放置项目的配置信息，连接到后端的api接口
console.log("获取环境变量:", import.meta.env)
const BASE_URL = import.meta.env.VITE_BASE_URL
export const API_CONFIG = {
    // 定义登录注册接口
    loginApi: `${BASE_URL}/auth/login`,
    logoutApi: `${BASE_URL}/auth/logout`,
    //image-display接口
    uploadApi: `${BASE_URL}/display/upload`,
    getImageApi: `${BASE_URL}/display/get`,
    updateApi: `${BASE_URL}/display/update`,
    deleteApi: `${BASE_URL}/display/delete`,
    getAllApi: `${BASE_URL}/display/getall`,
}
export const CONFIG = {
    TOKEN_NAME: "Authorization"
}

