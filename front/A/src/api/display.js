// 导入配置和请求工具
import { API_CONFIG } from "../config/index.js";
import request from './index.js';



// 根据ID获取单张图片
export const getImageByIdHandler = (id) => {
    return request(`${API_CONFIG.getImageApi}/${id}`, {}, "get", 5000);
};

// 获取所有图片
export const getAllImagesHandler = () => {
    return request(API_CONFIG.getAllApi, {}, "get", 5000);
};

//上传图片
export const uploadImageHandler = (data) => {
    return request(API_CONFIG.uploadApi, data, "post", 30000, {
        'Content-Type': 'multipart/form-data'
    });
};

//更新图片
export const updateImageHandler = (id, data) => {
    return request(`${API_CONFIG.updateApi}/${id}`, data, "put", 5000);
};
