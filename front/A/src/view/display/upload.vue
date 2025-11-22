<script setup>
import { ref } from 'vue'
import { uploadImageHandler } from '../../api/display.js'
import { ElMessage } from 'element-plus'

// 表单数据
const formData = ref({
  id: '',
  desc: '',
  images: null
})

// 加载状态
const loading = ref(false)

// 文件上传前的验证
const beforeUpload = (file) => {
  const isJPG = file.type === 'image/jpeg' || file.type === 'image/png' || file.type === 'image/gif'
  const isLt2M = file.size / 1024 / 1024 < 2

  if (!isJPG) {
    ElMessage.error('只能上传 JPG/PNG/GIF 格式的图片!')
    return false
  }
  if (!isLt2M) {
    ElMessage.error('图片大小不能超过 2MB!')
    return false
  }
  
  formData.value.images = file
  return false // 手动上传
}

// 处理文件选择变化
const handleFileChange = (file, fileList) => {
  formData.value.images = file.raw
}

// 提交表单
const submitForm = async () => {
  // 表单验证
  if (!formData.value.id) {
    ElMessage.error('请输入图片ID')
    return
  }
  if (!formData.value.desc) {
    ElMessage.error('请输入图片描述')
    return
  }
  if (!formData.value.images) {
    ElMessage.error('请选择要上传的图片')
    return
  }

  loading.value = true
  
  try {
    // 创建FormData对象
    const data = new FormData()
    data.append('id', formData.value.id)
    data.append('desc', formData.value.desc)
    data.append('images', formData.value.images)

    // 调用上传API
    const response = await uploadImageHandler(data)
    
    if (response.status === 200) {
      ElMessage.success('图片上传成功!')
      // 重置表单
      formData.value = {
        id: '',
        desc: '',
        images: null
      }
    } else {
      ElMessage.error(response.message || '上传失败')
    }
  } catch (error) {
    console.error('上传错误:', error)
    ElMessage.error('上传失败，请重试')
  } finally {
    loading.value = false
  }
}

// 重置表单
const resetForm = () => {
  formData.value = {
    id: '',
    desc: '',
    images: null
  }
}
</script>

<template>
  <div class="upload-container">
    <div class="upload-header">
      <h1>图片上传</h1>
      <p>请填写图片信息并选择要上传的图片文件</p>
    </div>
    
    <div class="upload-form">
      <el-form :model="formData" label-width="100px">
        <el-form-item label="图片ID" required>
          <div class="id-tip">
            <el-alert 
              title="注意: ID将作为文件的名称,建议使用数字,如ID:1,则文件名为1.jpg" 
              type="warning" 
              :closable="false"
              show-icon
              style="font-size: 10px; padding: 4px 8px;"
            />
          </div>
          <el-input 
            v-model="formData.id" 
            placeholder="请输入图片ID"
            maxlength="50"
            show-word-limit
          />
        </el-form-item>
        
        <el-form-item label="图片描述" required>
          <el-input 
            v-model="formData.desc" 
            type="textarea"
            :rows="3"
            placeholder="请输入图片描述信息"
            maxlength="200"
            show-word-limit
          />
        </el-form-item>
        
        <el-form-item label="选择图片" required>
          <el-upload
            class="upload-demo"
            :auto-upload="false"
            :before-upload="beforeUpload"
            :on-change="handleFileChange"
            :show-file-list="false"
            accept=".jpg,.jpeg,.png,.gif"
          >
            <el-button type="primary">选择文件</el-button>
            <template #tip>
              <div class="el-upload__tip">
                <p>支持 JPG、PNG、GIF 格式，大小不超过 2MB</p>
                <p v-if="formData.images" class="file-info">
                  已选择: {{ formData.images.name }}
                </p>
              </div>
            </template>
          </el-upload>
        </el-form-item>
        
        <el-form-item>
          <el-button 
            type="primary" 
            :loading="loading"
            @click="submitForm"
          >
            {{ loading ? '上传中...' : '上传图片' }}
          </el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<style scoped>
.upload-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
}

.upload-header {
  text-align: center;
  margin-bottom: 30px;
}

.upload-header h1 {
  color: #333;
  margin-bottom: 10px;
}

.upload-header p {
  color: #666;
  font-size: 14px;
}

.upload-form {
  background: #fff;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.file-info {
  color: #67c23a;
  font-weight: bold;
  margin-top: 5px;
}

.el-upload__tip {
  margin-top: 10px;
  font-size: 12px;
  color: #909399;
}

.el-upload__tip p {
  margin: 5px 0;
}
</style>
