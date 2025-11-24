<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getImageByIdHandler, updateImageHandler } from '../../api/display.js'

const route = useRoute()
const router = useRouter()

const imageId = ref('')
const imageUrl = ref('')
const imageDesc = ref('')
const loading = ref(false)
const error = ref('')
const success = ref('')

// 根据ID获取图片信息
const fetchImageById = async (id) => {
  loading.value = true
  error.value = ''
  
  try {
    const response = await getImageByIdHandler(id)
    if (response.data && response.data.status === 200) {
      const imageData = response.data.data
      imageUrl.value = imageData.image_url || ''
      imageDesc.value = imageData.desc || ''
    } else {
      error.value = response.data?.message || '获取图片信息失败'
    }
  } catch (err) {
    error.value = '网络请求失败，请检查网络连接'
    console.error('获取图片信息失败:', err)
  } finally {
    loading.value = false
  }
}

// 更新图片信息
const updateImage = async () => {
  if (!imageId.value.trim()) {
    error.value = '请输入图片ID'
    return
  }
  
  if (!imageUrl.value.trim()) {
    error.value = '请输入图片URL'
    return
  }
  
  loading.value = true
  error.value = ''
  success.value = ''
  
  try {
    const updateData = {
      image_url: imageUrl.value.trim(),
      desc: imageDesc.value.trim()
    }
    
    const response = await updateImageHandler(imageId.value, updateData)
    if (response.data && response.data.status === 200) {
      success.value = '图片更新成功'
      // 清空表单
      imageId.value = ''
      imageUrl.value = ''
      imageDesc.value = ''
    } else {
      error.value = response.data?.message || '更新图片失败'
    }
  } catch (err) {
    error.value = '网络请求失败，请检查网络连接'
    console.error('更新图片失败:', err)
  } finally {
    loading.value = false
  }
}

// 加载图片信息
const loadImageInfo = () => {
  if (imageId.value.trim()) {
    fetchImageById(imageId.value.trim())
  }
}

// 重置表单
const resetForm = () => {
  imageId.value = ''
  imageUrl.value = ''
  imageDesc.value = ''
  error.value = ''
  success.value = ''
}

// 返回图片列表
const backToList = () => {
  router.push('/display/getall')
}

// 组件挂载时检查路由参数
onMounted(() => {
  if (route.query.id) {
    imageId.value = route.query.id
    fetchImageById(route.query.id)
  } else {
    // 默认加载id为1的图片
    imageId.value = '1'
    fetchImageById('1')
  }
})
</script>

<template>
  <div class="update-container">
    <div class="header">
      <h1>更新图片信息</h1>
      <button @click="backToList" class="back-btn">返回列表</button>
    </div>

    <!-- 错误提示 -->
    <div v-if="error" class="error-message">
      {{ error }}
    </div>

    <!-- 成功提示 -->
    <div v-if="success" class="success-message">
      {{ success }}
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="loading">
      正在处理中...
    </div>

    <!-- 更新表单 -->
    <div class="update-form">
      <div class="form-group">
        <label for="imageId" class="form-label">图片ID *</label>
        <div class="input-with-button">
          <input
            id="imageId"
            v-model="imageId"
            type="text"
            placeholder="请输入要更新的图片ID"
            class="form-input"
          />
          <button @click="loadImageInfo" class="load-btn" :disabled="!imageId.trim()">
            加载信息
          </button>
        </div>
      </div>

      <div class="form-group">
        <label for="imageUrl" class="form-label">图片URL *</label>
        <input
          id="imageUrl"
          v-model="imageUrl"
          type="text"
          placeholder="请输入新的图片URL"
          class="form-input"
        />
      </div>

      <div class="form-group">
        <label for="imageDesc" class="form-label">图片描述</label>
        <textarea
          id="imageDesc"
          v-model="imageDesc"
          placeholder="请输入图片描述（可选）"
          class="form-textarea"
          rows="3"
        ></textarea>
      </div>

      <div class="form-actions">
        <button @click="updateImage" class="update-btn" :disabled="loading || !imageId.trim() || !imageUrl.trim()">
          {{ loading ? '更新中...' : '更新图片' }}
        </button>
        <button @click="resetForm" class="reset-btn" :disabled="loading">
          重置
        </button>
      </div>
    </div>

    <!-- 操作说明 -->
    <div class="instructions">
      <h3>操作说明：</h3>
      <ol>
        <li>输入要更新的图片ID,默认id为1,前端默认展示id:1的图片信息</li>
        <li>点击"加载信息"按钮获取当前图片信息</li>
        <li>修改图片URL和描述信息,可直接复制、粘贴</li>
        <li>点击"更新图片"按钮保存更改</li>
        <li>在前端页面查看是否更新成功</li>
      </ol>
    </div>
  </div>
</template>

<style scoped>
.update-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Helvetica Neue', Arial, sans-serif;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
  border-bottom: 1px solid #e0e0e0;
  padding-bottom: 20px;
}

.header h1 {
  color: #333;
  margin: 0;
  font-size: 28px;
}

.back-btn {
  padding: 10px 20px;
  background-color: #6c757d;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.back-btn:hover {
  background-color: #5a6268;
}

.error-message {
  background-color: #f8d7da;
  color: #721c24;
  padding: 12px;
  border-radius: 4px;
  margin-bottom: 20px;
  border: 1px solid #f5c6cb;
}

.success-message {
  background-color: #d4edda;
  color: #155724;
  padding: 12px;
  border-radius: 4px;
  margin-bottom: 20px;
  border: 1px solid #c3e6cb;
}

.loading {
  text-align: center;
  padding: 20px;
  color: #666;
  font-size: 16px;
}

.update-form {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-bottom: 30px;
}

.form-group {
  margin-bottom: 20px;
}

.form-label {
  display: block;
  margin-bottom: 8px;
  font-weight: 600;
  color: #333;
}

.form-input, .form-textarea {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
  box-sizing: border-box;
}

.form-input:focus, .form-textarea:focus {
  outline: none;
  border-color: #007bff;
  box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
}

.form-textarea {
  resize: vertical;
  min-height: 80px;
}

.input-with-button {
  display: flex;
  gap: 10px;
}

.input-with-button .form-input {
  flex: 1;
}

.load-btn, .update-btn, .reset-btn {
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  white-space: nowrap;
}

.load-btn {
  background-color: #17a2b8;
  color: white;
}

.load-btn:hover:not(:disabled) {
  background-color: #138496;
}

.load-btn:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
}

.update-btn {
  background-color: #007bff;
  color: white;
}

.update-btn:hover:not(:disabled) {
  background-color: #0056b3;
}

.update-btn:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
}

.reset-btn {
  background-color: #6c757d;
  color: white;
  margin-left: 10px;
}

.reset-btn:hover:not(:disabled) {
  background-color: #5a6268;
}

.form-actions {
  display: flex;
  justify-content: flex-start;
  margin-top: 30px;
}

.instructions {
  background-color: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border-left: 4px solid #007bff;
}

.instructions h3 {
  margin-top: 0;
  margin-bottom: 15px;
  color: #333;
}

.instructions ol {
  margin: 0;
  padding-left: 20px;
}

.instructions li {
  margin-bottom: 8px;
  color: #666;
  line-height: 1.5;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .update-container {
    padding: 15px;
  }
  
  .header {
    flex-direction: column;
    gap: 15px;
    text-align: center;
  }
  
  .input-with-button {
    flex-direction: column;
  }
  
  .form-actions {
    flex-direction: column;
    gap: 10px;
  }
  
  .reset-btn {
    margin-left: 0;
  }
}
</style>