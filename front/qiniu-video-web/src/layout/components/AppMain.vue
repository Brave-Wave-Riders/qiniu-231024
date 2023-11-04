<template>
  <div class="app-main">
    <router-view></router-view>
  </div>
</template>

<script setup>
import { watch } from 'vue'
import { useRoute } from 'vue-router'
// import { useStore } from 'vuex'

const route = useRoute()
watch(route, (to, from) => {
  console.log(to)
})

</script>

<style lang="scss" scoped>
@import '~@/styles/mixin.scss';
@import '~@/styles/variables.module.scss';

.app-main {
  min-height: calc(100vh - 30px);
  position: relative;
  overflow: hidden;
  padding: 0px;
  box-sizing: border-box;
  background-color: #16304e;
  top: 70px;
}
</style>
