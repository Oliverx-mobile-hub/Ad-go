<template>
  <div class="image-display-container">
    <div class="header">
      <!-- 展示标题，可更改 -->
      <h1>B展示</h1>
      <!-- 展示图片描述，可更改，在currentImage的description中 更改-->
      <p>{{currentImage.description}}</p>
    </div>
    
    <div class="image-section">
      <div class="image-wrapper">
        <a 
          :href="currentImage.link" 
          target="_blank" 
          class="image-link"
        >
          <img 
            :src="currentImage.src" 
            :alt="currentImage.alt"
            class="display-image"
          />
        </a>
      </div>
    </div>
    

  </div>
</template>

<script>
export default {
  name: 'ImageDisplay',
  data() {
    return {
      currentImage: {
        // 展示图片，可更改，部署时，注意路径
        src: './dog.jpg',
        alt: '示例图片',
        title: '示例图片',
        description: '这是一个示例图片，点击可以跳转到目标页面',
        // 展示跳转链接，可更改
        link: 'http://localhost:5175/display/get'
      }
    }
  },
  methods: {
  }
}
</script>

<style scoped>
.image-display-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  font-family: Arial, sans-serif;
  text-align: center;
}

.header {
  text-align: center;
  margin-bottom: 40px;
}

.header h1 {
  color: #333;
  margin-bottom: 10px;
}

.header p {
  color: #666;
  font-size: 16px;
}

.image-section {
  margin-bottom: 40px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.image-wrapper {
  display: inline-block;
}

.image-link {
  display: block;
  text-decoration: none;
  transition: transform 0.3s ease;
}

.image-link:hover {
  transform: scale(1.02);
}

.display-image {
  max-width: 600px;
  max-height: 400px;
  width: auto;
  height: auto;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}



@media (max-width: 768px) {
  .display-image {
    max-width: 100%;
    max-height: 300px;
  }
}
</style>