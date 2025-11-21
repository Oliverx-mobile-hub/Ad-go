<template>
  <ImageDisplay />
</template>

<script>
import ImageDisplay from './components/ImageDisplay.vue'

export default {
  name: 'App',
  components: {
    ImageDisplay
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: Arial, sans-serif;
  background-color: #f5f5f5;
}
</style>