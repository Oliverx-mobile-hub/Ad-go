<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { deleteImageHandler, getAllImagesHandler } from '../../api/display.js'

const router = useRouter()

// 响应式数据
const imageId = ref('')
const loading = ref(false)
const error = ref('')
const success = ref('')
const allImages = ref([])
const totalCount = ref(0)

// 删除图片函数
const deleteImage = async () => {
  if (!imageId.value.trim()) {
    ElMessage.error('请输入图片ID')
    return
  }

  try {
    // 显示确认对话框
    await ElMessageBox.confirm(
      `确定要删除ID为 ${imageId.value} 的图片吗?,删除后，不可恢复。(不建议删除id:1,因为它是默认图片，其他图片可以删除）`,
      '确认删除',
      {
        confirmButtonText: '确定删除',
        cancelButtonText: '取消',
        type: 'warning',
        dangerouslyUseHTMLString: true
      }
    )

    loading.value = true
    error.value = ''
    success.value = ''

    // 调用删除API
    const response = await deleteImageHandler(imageId.value)
    
    if (response.data.success) {
      success.value = `图片ID ${imageId.value} 删除成功！`
      ElMessage.success(success.value)
      // 清空输入框
      imageId.value = ''
      // 重新获取所有图片信息以更新总数
      await fetchAllImages()
    } else {
      throw new Error(response.data.message || '删除失败')
    }
  } catch (err) {
    if (err === 'cancel' || err === 'close') {
      // 用户取消删除
      ElMessage.info('已取消删除操作')
      return
    }
    
    error.value = err.response?.data?.message || err.message || '删除失败'
    ElMessage.error(error.value)
  } finally {
    loading.value = false
  }
}

// 重置表单
const resetForm = () => {
  imageId.value = ''
  error.value = ''
  success.value = ''
}

// 返回图片列表
const goToList = () => {
  router.push('/display/getall')
}

// 获取所有图片信息
const fetchAllImages = async () => {
  try {
    const response = await getAllImagesHandler()
    if (response.data && response.data.status === 200) {
      allImages.value = response.data.data || []
      totalCount.value = allImages.value.length
    } else {
      console.error('获取图片列表失败:', response.data?.message)
    }
  } catch (err) {
    console.error('获取图片列表失败:', err)
  }
}

// 组件挂载时检查路由参数并获取所有图片
onMounted(() => {
  const route = router.currentRoute.value
  if (route.params.id) {
    imageId.value = route.params.id
  }
  fetchAllImages()
})
</script>

<template>
  <div class="delete-container">
    <div class="header">
      <h1>删除图片</h1>
      <div class="actions">
        <button @click="goToList" class="back-btn">返回列表</button>
      </div>
    </div>

    <!-- 成功提示 -->
    <div v-if="success" class="success-message">
      {{ success }}
    </div>

    <!-- 错误提示 -->
    <div v-if="error" class="error-message">
      {{ error }}
    </div>

    <!-- 删除表单 -->
    <div class="delete-form">
      <div class="form-group">
        <label class="form-label">图片ID</label>
        <input 
          v-model="imageId" 
          type="text" 
          class="form-input" 
          placeholder="请输入要删除的图片ID"
          :disabled="loading"
        />
      </div>

      <div class="form-actions">
        <button 
          @click="deleteImage" 
          class="delete-btn" 
          :disabled="loading || !imageId.trim()"
        >
          {{ loading ? '删除中...' : '删除图片' }}
        </button>
        <button @click="resetForm" class="reset-btn" :disabled="loading">
          重置
        </button>
      </div>

      <!-- 所有图片信息 -->
      <div class="images-info">
        <h3>当前图片信息（共 {{ totalCount }} 张图片）</h3>
        <div v-if="allImages.length > 0" class="images-list">
          <div 
            v-for="image in allImages" 
            :key="image.id" 
            class="image-item"
            :class="{ 'current-selected': image.id === imageId }"
          >
            <span class="image-id">ID: {{ image.id }}</span>
            <span class="image-url">{{ image.image_url }}</span>
            <span class="image-desc">{{ image.desc || '-' }}</span>
          </div>
        </div>
        <div v-else class="no-images">
          暂无图片数据
        </div>
      </div>

      <!-- 操作说明 -->
      <div class="operation-info">
        <h3>操作说明：</h3>
        <ul>
          <li>当前共有 <strong>{{ totalCount }}</strong> 张图片</li>
          <li>输入要删除的图片ID</li>
          <li>点击"删除图片"按钮</li>
          <li>系统会弹出确认对话框</li>
          <li>确认后将同时删除数据库记录和uploads文件夹中的文件</li>
          <li>删除后图片总数将变为 <strong>{{ totalCount - 1 }}</strong> 张</li>
          <li>此操作不可恢复，请谨慎操作</li>
        </ul>
      </div>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="loading">
      正在删除图片，请稍候...
    </div>
  </div>
</template>

<style scoped>
.delete-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
  border-bottom: 1px solid #eaeaea;
  padding-bottom: 20px;
}

.header h1 {
  color: #333;
  font-size: 28px;
  font-weight: 600;
  margin: 0;
}

.actions {
  display: flex;
  gap: 10px;
}

.back-btn {
  background-color: #6c757d;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.back-btn:hover {
  background-color: #5a6268;
}

.success-message {
  background-color: #d4edda;
  color: #155724;
  padding: 12px;
  border-radius: 4px;
  margin-bottom: 20px;
  border: 1px solid #c3e6cb;
}

.error-message {
  background-color: #f8d7da;
  color: #721c24;
  padding: 12px;
  border-radius: 4px;
  margin-bottom: 20px;
  border: 1px solid #f5c6cb;
}

.delete-form {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-bottom: 30px;
}

.form-group {
  margin-bottom: 25px;
}

.form-label {
  display: block;
  margin-bottom: 8px;
  font-weight: 600;
  color: #333;
  font-size: 16px;
}

.form-input {
  width: 100%;
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
  box-sizing: border-box;
  transition: border-color 0.3s;
}

.form-input:focus {
  outline: none;
  border-color: #dc3545;
  box-shadow: 0 0 0 2px rgba(220, 53, 69, 0.25);
}

.form-input:disabled {
  background-color: #f8f9fa;
  cursor: not-allowed;
}

.form-actions {
  display: flex;
  gap: 15px;
  margin-bottom: 30px;
}

.delete-btn, .reset-btn {
  padding: 12px 24px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
  font-weight: 500;
  transition: all 0.3s;
}

.delete-btn {
  background-color: #dc3545;
  color: white;
}

.delete-btn:hover:not(:disabled) {
  background-color: #c82333;
  transform: translateY(-1px);
}

.delete-btn:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
  transform: none;
}

.reset-btn {
  background-color: #6c757d;
  color: white;
}

.reset-btn:hover:not(:disabled) {
  background-color: #5a6268;
}

.reset-btn:disabled {
  background-color: #adb5bd;
  cursor: not-allowed;
}

/* 图片信息展示样式 */
.images-info {
  background-color: #f8f9fa;
  padding: 20px;
  border-radius: 6px;
  border-left: 4px solid #007bff;
  margin-bottom: 20px;
}

.images-info h3 {
  color: #333;
  margin-top: 0;
  margin-bottom: 15px;
  font-size: 18px;
  font-weight: 600;
}

.images-list {
  max-height: 200px;
  overflow-y: auto;
}

.image-item {
  display: flex;
  align-items: center;
  padding: 10px;
  margin-bottom: 8px;
  background-color: white;
  border-radius: 4px;
  border: 1px solid #e9ecef;
  transition: all 0.3s;
}

.image-item:hover {
  border-color: #007bff;
  box-shadow: 0 2px 4px rgba(0, 123, 255, 0.1);
}

.image-item.current-selected {
  border-color: #dc3545;
  background-color: #fff5f5;
  box-shadow: 0 2px 4px rgba(220, 53, 69, 0.1);
}

.image-id {
  font-weight: 600;
  color: #007bff;
  min-width: 60px;
  margin-right: 15px;
  font-family: 'Courier New', monospace;
}

.image-url {
  flex: 1;
  color: #666;
  font-family: 'Courier New', monospace;
  font-size: 14px;
  word-break: break-all;
  margin-right: 15px;
}

.image-desc {
  color: #999;
  font-size: 14px;
  min-width: 80px;
  text-align: right;
}

.no-images {
  text-align: center;
  color: #999;
  padding: 20px;
  font-style: italic;
}

.operation-info {
  background-color: #f8f9fa;
  padding: 20px;
  border-radius: 6px;
  border-left: 4px solid #dc3545;
}

.operation-info h3 {
  color: #333;
  margin-top: 0;
  margin-bottom: 15px;
  font-size: 18px;
}

.operation-info ul {
  margin: 0;
  padding-left: 20px;
}

.operation-info li {
  margin-bottom: 8px;
  color: #666;
  line-height: 1.5;
}

.operation-info strong {
  color: #dc3545;
  font-weight: 600;
}

.loading {
  text-align: center;
  padding: 20px;
  color: #666;
  font-size: 16px;
  background-color: #f8f9fa;
  border-radius: 4px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .delete-container {
    padding: 10px;
  }
  
  .header {
    flex-direction: column;
    gap: 15px;
    text-align: center;
  }
  
  .header h1 {
    font-size: 24px;
  }
  
  .delete-form {
    padding: 20px;
  }
  
  .form-actions {
    flex-direction: column;
  }
  
  .delete-btn, .reset-btn {
    width: 100%;
  }
}
</style>