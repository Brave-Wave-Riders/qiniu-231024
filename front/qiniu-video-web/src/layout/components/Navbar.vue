<template>
  <div class="navbar">
    <div class="headerWrap">
      <header-search class="right-menu-item hover-effect"></header-search>
    </div>
    <div class="right-menu">
      <!-- 头像 -->
      <el-dropdown class="avatar-container" trigger="click">
        <div class="avatar-wrapper">
          <el-avatar
            shape="circle"
            :size="40"
            :src="require('@/assets/default_img.jpg')"
          ></el-avatar>
        </div>
        <template #dropdown>
          <el-dropdown-menu class="user-dropdown">
            <router-link to="/">
              <el-dropdown-item> 首页 </el-dropdown-item>
            </router-link>
            <router-link to="/video/upload">
              <el-dropdown-item> 创作 </el-dropdown-item>
            </router-link>
            <el-dropdown-item divided>
              登出
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>

    </div>
  </div>
</template>

<script setup>
import HeaderSearch from '@/components/HeaderSearch'
// import {} from 'vue'
// import { useStore } from 'vuex'

// const store = useStore()
// const logout = () => {
//   store.dispatch('user/logout')
// }
</script>

<style lang="scss" scoped>

.navbar {
  height: 70px;
  overflow: hidden;
  position: relative;
  background: #16304e;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
  display: flex;
  .headerWrap{
    display: flex;
    align-items: center;
    width: 422px;
    margin: 10px auto;
  }

  .right-menu {
    display: flex;
    align-items: center;
    float: right;
    padding-right: 16px;

    ::v-deep .right-menu-item {
      display: inline-block;
      padding: 0 18px 0 0;
      font-size: 24px;
      color: #5a5e66;
      vertical-align: text-bottom;

      &.hover-effect {
        cursor: pointer;
        transition: background 0.3s;

        &:hover {
          background: rgba(0, 0, 0, 0.025);
        }
      }
    }

    ::v-deep .avatar-container {
      cursor: pointer;
      .avatar-wrapper {
        margin-top: 5px;
        position: relative;
        .el-avatar {
          --el-avatar-background-color: none;
          margin-right: 12px;
        }
      }
    }
  }
}
</style>
