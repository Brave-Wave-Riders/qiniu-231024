<template>
  <div :class="{ show: isShow }" class="header-search">
    <div>
      <el-select
        ref="headerSearchSelectRef"
        class="header-search-select"
        v-model="search"
        filterable
        default-first-option
        remote
        placeholder="Search"
        :remote-method="querySearch"
        @change="onSelectChange"
      >
        <el-option
          v-for="option in searchOptions"
          :key="option.item.path"
          :label="option.item.title.join(' > ')"
          :value="option.item"
        ></el-option>
      </el-select>
    </div>
    <div class="search-icon">
      <el-icon size="30"><Search /></el-icon>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

// 控制 search 显示
const isShow = ref(true)
// el-select 实例
const headerSearchSelectRef = ref(null)

// 搜索结果
const searchOptions = ref([])

/**
 * 关闭 search 的处理事件
 */
const onClose = () => {
  headerSearchSelectRef.value.blur()
  isShow.value = false
  searchOptions.value = []
}
/**
 * 监听 search 打开，处理 close 事件
 */
watch(isShow, val => {
  if (val) {
    document.body.addEventListener('click', onClose)
  } else {
    document.body.removeEventListener('click', onClose)
  }
})

</script>

<style lang="scss" >
.el-input--suffix .el-input__inner {
    padding-right: 0px;
    background: #16304e;
    padding-left: 0px;
    border: 0px;
}
.el-select .el-input.is-focus .el-input__inner {
  background-color: #16304e !important;
  border-color: #16304e !important;
}
.header-search {
  display: flex;
  font-size: 0 !important;
  justify-content: center;
  margin-top: 10px;
  background-color: #16304e;
  border: 2px solid rgba($color: #fff, $alpha: .7);

  .search-icon {
    cursor: pointer;
    vertical-align: middle;
    justify-content: center;
    color: #fff;
  }
  .header-search-select {
    font-size: 18px;
    width: 0;
    background: transparent;
    border-radius: 0;
    display: inline-block;
    vertical-align: middle;
    ::v-deep .el-input__inner {
      border-radius: 0;
      border: 0;
      padding-left: 0;
      padding-right: 0;
      box-shadow: none !important;
      border-bottom: 1px solid #16304e;//#d9d9d9;
      vertical-align: middle;
    }
  }
  &.show {
    .header-search-select {
      width: 210px;
      margin-left: 20px;
      color: #16304e;
    }
  }
}
</style>
