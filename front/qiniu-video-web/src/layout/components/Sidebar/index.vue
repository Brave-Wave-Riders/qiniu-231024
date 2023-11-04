<template>
  <div class="">
    <div class="logo-container">
      <el-avatar
        :size="logoHeight"
        :src="require('@/assets/title.png')"
      />
      <h1 class="logo-title" >
        遁形
      </h1>
    </div>
    <el-scrollbar>
      <sidebar-menu></sidebar-menu>
    </el-scrollbar>
  </div>
</template>

<script setup>
import SidebarMenu from './SidebarMenu'
import {} from 'vue'
// import titleImg from '@/assets/title.png'

// shape="square"
const logoHeight = 44
</script>

<style lang="scss" scoped>
.logo-container {
  height: v-bind(logoHeight) + 'px';
  padding: 10px 0 22px 0;
  display: flex;
  align-items: center;
  justify-content: center;
  .logo-title {
    margin-left: 10px;
    color: #fff;
    font-weight: 600;
    line-height: 50px;
    font-size: 20px;
    white-space: nowrap;
  }
}
</style>
