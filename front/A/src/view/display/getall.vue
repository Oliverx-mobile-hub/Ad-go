<script setup>
import { ref, onMounted } from 'vue'
import { getAllImagesHandler } from '../../api/display.js'
import { useRouter } from 'vue-router'

const images = ref([])
const loading = ref(false)
const error = ref('')
const router = useRouter()

// 获取后端基础URL，如果换服务器可以直接改地址
const getBackendBaseUrl = () => {
  // 从环境变量获取，如果没有则使用默认的localhost:8080
  return import.meta.env.VITE_BACKEND_BASE_URL || 'http://localhost:8080'
}

// 处理图片URL，将相对路径转换为完整URL（仅用于缩略图显示）
const processImageUrl = (imageUrl) => {
  if (imageUrl && !imageUrl.startsWith('http')) {
    // 如果是相对路径，添加后端基础URL前缀（仅用于缩略图显示）
    const backendBaseUrl = getBackendBaseUrl()
    return `${backendBaseUrl}${imageUrl}`
  }
  return imageUrl
}

// 获取数据库原始路径（用于URL显示）
const getOriginalImageUrl = (imageUrl) => {
  // 直接返回数据库中的原始路径，不进行拼接
  return imageUrl
}

// 获取所有图片
const fetchAllImages = async () => {
  loading.value = true
  error.value = ''
  
  try {
    const response = await getAllImagesHandler()
    if (response.data && response.data.status === 200) {
      // 处理图片数据，缩略图使用完整URL，URL文本显示原始路径
      images.value = (response.data.data || []).map(image => ({
        ...image,
        // 缩略图显示使用完整URL
        display_url: processImageUrl(image.image_url),
        // URL文本显示使用原始路径
        original_url: getOriginalImageUrl(image.image_url)
      }))
    } else {
      error.value = response.data?.message || '获取图片列表失败'
    }
  } catch (err) {
    error.value = '网络请求失败，请检查网络连接'
    console.error('获取图片列表失败:', err)
  } finally {
    loading.value = false
  }
}


// 组件挂载时获取数据
onMounted(() => {
  fetchAllImages()
})

// 重新加载数据
const reloadImages = () => {
  fetchAllImages()
}

// 复制URL到剪贴板
const copyToClipboard = (text) => {
  navigator.clipboard.writeText(text).then(() => {
    console.log('已复制到剪贴板:', text)
  }).catch(err => {
    console.error('复制失败:', err)
  })
}

// 跳转到更新页面
const goToUpdatePage = (imageId) => {
  router.push(`/display/update?id=${imageId}`)
}
</script>

<template>
  <div class="getall-container">
    <div class="header">
      <h1>图片信息列表</h1>
      <div class="stats">
        <span class="total-count">总共 {{ images.length }} 张图片</span>
        <button @click="reloadImages" class="reload-btn" :disabled="loading">
          {{ loading ? '加载中...' : '刷新' }}
        </button>
      </div>
    </div>

    <!-- 说明文字 -->
    <div class="info-notice" style="text-align: left;">
      默认展示id:1的图片,限定只更新id:1的数据
    </div>

    <!-- 错误提示 -->
    <div v-if="error" class="error-message">
      {{ error }}
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="loading">
      正在加载图片信息...
    </div>

    <!-- 图片信息展示 -->
    <div v-if="!loading && !error" class="info-container">
      <div v-if="images.length === 0" class="empty-state">
        暂无图片数据
      </div>
      
      <div v-else class="info-table">
        <table class="info-table-content">
          <thead>
            <tr>
              <th class="col-id">ID</th>
              <th class="col-thumbnail">缩略图</th>
              <th class="col-url">图片URL</th>
              <th class="col-desc">描述</th>
              <th class="col-actions">操作</th>
            </tr>
          </thead>
          <tbody>
            <tr 
              v-for="image in images" 
              :key="image.id" 
              class="info-row"
            >
              <td class="col-id">{{ image.id }}</td>
              <td class="col-thumbnail">
                <div class="thumbnail-container">
                  <img 
                    :src="image.display_url" 
                    :alt="image.desc || '图片'" 
                    class="thumbnail" 
                    @error="$event.target.style.display='none'"
                  >
                </div>
              </td>
              <td class="col-url">
                <div class="url-container">
                  <span class="url-text">{{ image.original_url }}</span>
                  <button 
                    @click="copyToClipboard(image.original_url)" 
                    class="copy-btn"
                    title="复制URL"
                  >
                    复制
                  </button>
                </div>
              </td>
              <td class="col-desc">{{ image.desc || '-' }}</td>
              <td class="col-actions">
                <!-- 只显示id为1的图片的更新按钮 -->
                <button 
                  v-if="image.id === '1'"
                  @click="goToUpdatePage(image.id)" 
                  class="update-btn"
                  title="更新图片信息"
                >
                  更新
                </button>
                <span v-else class="no-update">-</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 全局样式重置 */
.getall-container {
  max-width: 1200px;
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
}

.stats {
  display: flex;
  align-items: center;
  gap: 15px;
}

.total-count {
  font-size: 16px;
  color: #666;
  font-weight: 500;
}

.reload-btn {
  padding: 8px 16px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.reload-btn:hover:not(:disabled) {
  background-color: #0056b3;
}

.reload-btn:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
}

.error-message {
  background-color: #f8d7da;
  color: #721c24;
  padding: 12px;
  border-radius: 4px;
  margin-bottom: 20px;
  border: 1px solid #f5c6cb;
}

.loading {
  text-align: center;
  padding: 40px;
  color: #666;
  font-size: 16px;
}

.empty-state {
  text-align: center;
  padding: 60px 20px;
  color: #999;
  font-size: 18px;
}

/* 表格样式 - 确保对齐整齐 */
.info-table {
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.info-table-content {
  width: 100%;
  border-collapse: collapse;
  table-layout: fixed;
  font-size: 14px;
}

/* 表头样式 */
.info-table-content th {
  background-color: #f8f9fa;
  padding: 12px;
  text-align: left;
  font-weight: 600;
  color: #333;
  border-bottom: 2px solid #dee2e6;
  vertical-align: middle;
}

/* 数据单元格样式 */
.info-table-content td {
  padding: 12px;
  border-bottom: 1px solid #e9ecef;
  vertical-align: middle;
}

/* 行悬停效果 */
.info-row:hover {
  background-color: #f8f9fa;
}

/* ID列 - 统一样式和对齐 */
.col-id {
  width: 60px;
  font-weight: 500;
  color: #007bff;
  font-family: 'Courier New', monospace;
  text-align: center;
}

/* 缩略图列样式 */
.col-thumbnail {
  width: 80px;
  text-align: center;
}

.thumbnail-container {
  width: 60px;
  height: 60px;
  overflow: hidden;
  border-radius: 4px;
  border: 1px solid #e0e0e0;
  background-color: #f8f9fa;
  display: flex;
  align-items: center;
  justify-content: center;
}

.thumbnail {
  max-width: 100%;
  max-height: 100%;
  object-fit: cover;
  transition: transform 0.2s ease;
}

.thumbnail:hover {
  transform: scale(1.1);
}

/* URL列 - 统一样式和对齐 */
.col-url {
  width: 50%;
  text-align: left;
}

.url-container {
  display: flex;
  align-items: center;
  gap: 10px;
}

.url-text {
  flex: 1;
  font-family: 'Courier New', monospace;
  font-size: 14px;
  color: #666;
  word-break: break-all;
  line-height: 1.5;
}

.copy-btn {
  padding: 6px 10px;
  background-color: #28a745;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 12px;
  white-space: nowrap;
  flex-shrink: 0;
  height: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.copy-btn:hover {
  background-color: #218838;
}

/* 描述列 - 统一样式和对齐 */
.col-desc {
  width: 25%;
  color: #666;
  word-break: break-word;
  line-height: 1.5;
  text-align: left;
}

/* 确保所有列样式一致 */
.info-table-content th.col-id, .info-table-content td.col-id {
  text-align: center;
}

.info-table-content th.col-thumbnail, .info-table-content td.col-thumbnail {
  text-align: center;
}

.info-table-content th.col-url, .info-table-content td.col-url {
  text-align: left;
}

.info-table-content th.col-desc, .info-table-content td.col-desc {
  text-align: left;
}

/* 操作列样式 */
.col-actions {
  width: 100px;
  text-align: center;
}

.update-btn {
  padding: 6px 12px;
  background-color: #ffc107;
  color: #212529;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 12px;
  white-space: nowrap;
}

.update-btn:hover {
  background-color: #e0a800;
}

/* 说明文字样式 */
.info-notice {
  background-color: #e3f2fd;
  border: 1px solid #bbdefb;
  border-radius: 4px;
  padding: 12px 16px;
  margin: 16px 0;
  color: #1976d2;
  font-size: 14px;
  font-weight: 500;
}

/* 无更新操作提示 */
.no-update {
  color: #999;
  font-style: italic;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .header {
    flex-direction: column;
    gap: 15px;
    text-align: center;
  }
  
  .info-table-content {
    font-size: 14px;
    table-layout: auto;
  }
  
  .info-table-content th,
  .info-table-content td {
    padding: 10px 6px;
  }
  
  .col-id {
    width: auto;
  }
  
  .col-thumbnail {
    width: 60px;
  }
  
  .thumbnail-container {
    width: 40px;
    height: 40px;
  }
  
  .col-url {
    width: auto;
  }
  
  .col-desc {
    width: auto;
  }
  
  .url-container {
    flex-direction: column;
    align-items: flex-start;
    gap: 6px;
  }
  
  .copy-btn {
    align-self: flex-end;
    height: auto;
  }
}

</style>