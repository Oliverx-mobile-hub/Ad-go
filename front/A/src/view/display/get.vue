<script setup>
import { ref, onMounted } from 'vue'
import { getImageByIdHandler } from '../../api/display.js'
import { ElMessage } from 'element-plus'

// 响应式数据
const imageId = ref('1') // 默认显示ID为1的图片
const imageData = ref(null)
const loading = ref(false)
const error = ref(null)

// 获取图片数据
const fetchImage = async () => {
  if (!imageId.value.trim()) {
    ElMessage.warning('请输入图片ID')
    return
  }
  
  loading.value = true
  error.value = null
  
  try {
    const response = await getImageByIdHandler(imageId.value)
    
    if (response.data.status === 200) {
      // 处理图片路径，将相对路径转换为完整URL
      const imageDataRaw = response.data.data
      if (imageDataRaw.image_url && !imageDataRaw.image_url.startsWith('http')) {
        // 如果是相对路径，添加后端基础URL
        imageDataRaw.image_url = `http://localhost:8080${imageDataRaw.image_url}`
      }
      imageData.value = imageDataRaw
      ElMessage.success('图片获取成功')
    } else {
      error.value = response.data.message || '获取图片失败'
      ElMessage.error(error.value)
    }
  } catch (err) {
    error.value = err.response?.data?.message || '网络请求失败'
    ElMessage.error(error.value)
  } finally {
    loading.value = false
  }
}

// 页面加载时自动获取默认图片
onMounted(() => {
  fetchImage()
})
</script>

<template>
  <div class="image-display-container">
    <div class="header">
      <h1>图片展示</h1>
      <div class="search-box">
        <el-input 
          v-model="imageId" 
          placeholder="请输入图片ID" 
          style="width: 200px; margin-right: 10px;"
          @keyup.enter="fetchImage"
        />
        <el-button 
          type="primary" 
          :loading="loading" 
          @click="fetchImage"
        >
          获取图片
        </el-button>
      </div>
    </div>

    <div class="content">
      <!-- 加载状态 -->
      <div v-if="loading" class="loading-container">
        <el-skeleton :rows="3" animated />
      </div>

      <!-- 错误状态 -->
      <div v-else-if="error" class="error-container">
        <el-result 
          icon="error" 
          title="获取图片失败" 
          :sub-title="error"
        >
          <template #extra>
            <el-button type="primary" @click="fetchImage">重试</el-button>
          </template>
        </el-result>
      </div>

      <!-- 图片展示 -->
      <div v-else-if="imageData" class="image-container">
        <div class="image-card">
          <div class="image-wrapper">
            <img 
              :src="imageData.image_url" 
              :alt="imageData.desc" 
              class="display-image"
              @error="handleImageError"
            />
          </div>
          <div class="image-info">
            <h3>图片信息</h3>
            <div class="info-item">
              <span class="label">ID:</span>
              <span class="value">{{ imageData.id }}</span>
            </div>
            <div class="info-item">
              <span class="label">描述:</span>
              <span class="value">{{ imageData.desc }}</span>
            </div>
            <div class="info-item">
              <span class="label">路径:</span>
              <span class="value">{{ imageData.image_url }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- 无数据状态 -->
      <div v-else class="no-data">
        <el-empty description="暂无图片数据" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.image-display-container {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
  padding-bottom: 20px;
  border-bottom: 1px solid #e4e7ed;
}

.header h1 {
  margin: 0;
  color: #303133;
}

.search-box {
  display: flex;
  align-items: center;
}

.content {
  min-height: 400px;
}

.loading-container,
.error-container,
.no-data {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
}

.image-container {
  display: flex;
  justify-content: center;
}

.image-card {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  overflow: hidden;
  max-width: 600px;
  width: 100%;
}

.image-wrapper {
  padding: 20px;
  text-align: center;
  background: #f5f7fa;
}

.display-image {
  max-width: 100%;
  max-height: 400px;
  object-fit: contain;
  border-radius: 4px;
}

.image-info {
  padding: 20px;
}

.image-info h3 {
  margin: 0 0 15px 0;
  color: #303133;
  border-bottom: 1px solid #e4e7ed;
  padding-bottom: 10px;
}

.info-item {
  display: flex;
  margin-bottom: 10px;
}

.label {
  font-weight: 600;
  color: #606266;
  min-width: 60px;
  margin-right: 10px;
}

.value {
  color: #303133;
  word-break: break-all;
}

@media (max-width: 768px) {
  .image-display-container {
    padding: 10px;
  }
  
  .header {
    flex-direction: column;
    gap: 15px;
    align-items: stretch;
  }
  
  .search-box {
    justify-content: center;
  }
}
</style>

<script>
// 处理图片加载错误
export default {
  methods: {
    handleImageError(event) {
      ElMessage.error('图片加载失败，请检查图片路径')
      event.target.style.display = 'none'
    }
  }
}
</script>
