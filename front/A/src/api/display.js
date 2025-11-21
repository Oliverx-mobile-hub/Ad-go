// 导入配置和请求工具
import { API_CONFIG } from "../config/index.js";
import request from './index.js';



// //获取图片列表
// export const getImageListHandler = (id, image_url, desc) => {
//     return request(API_CONFIG.getImageApi, { "id": id, "image_url": image_url, "desc": desc }, "get", 5000);
// };

// 根据ID获取单张图片
export const getImageByIdHandler = (id) => {
    return request(`${API_CONFIG.getImageApi}/${id}`, {}, "get", 5000);
};

