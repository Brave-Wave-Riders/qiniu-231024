<template>
  <router-view/>
  <el-scrollbar style="height:100%">
    <div class="index-container">
      <div class="t-list">
        <ul>
          <li
            v-for="(item, index) in tableData"
            :key="index"
          >
            <div @click="onShowClick(item.desc)">
              <img :src=item.url>
              <h4 class="titleWrap">{{ item.desc }}</h4>
              <div class="authorWrap">
                <span>@ </span>
                <span>{{item.author}}</span>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </el-scrollbar>
  <el-pagination
    class="pagination"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
    :current-page="page"
    :page-sizes="[10, 15, 20, 25, 30, 35]"
    :page-size="size"
    layout="total, sizes, prev, pager, next, jumper"
    :total="total"
  >
  </el-pagination>
</template>

<script setup>
import { ref, onActivated } from 'vue'
import { getVideoList, getMyVideoList } from '@/api/video'
import { useRouter, useRoute } from 'vue-router'
import { useStore } from 'vuex'

// const Base64 = require('js-base64').Base64

const store = useStore()
// // 数据相关
// const tableData = ref([])
const total = ref(0)
const page = ref(1)
const size = ref(10)
const tableData = ref([
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' },
  { url: 'title.png', desc: '测试图', icon: 'default.jpg', author: '遁形' }])

let id = '0'
// 获取数据的方法
const getListData = async () => {
  const route = useRoute()
  if (route) {
    id = route.params.id
  }
  console.log('id = ', id)
  let result = null
  if (id !== '8') {
    result = await getVideoList(id, {
      page: page.value,
      size: size.value
    })
  } else {
    result = await getMyVideoList({
      page: page.value,
      size: size.value
    })
  }
  // const
  // console.log(result.data)
  // console.log(result.total)
  if (result.total === 0) {
    return
  }
  const res = result.data
  for (let i = 0; i < res.length; i++) {
    // const desc = result.data.desc
    // console.log(typeof result.data)
    // result.data[i].desc = Base64.decode(result.data[i].desc)
    res[i].url = result.base + res[i].desc + '.jpg'
    console.log(result.data[i].url)
  }
  store.commit('video/setVideoData', result.data)
  store.commit('video/setVideoNum', result.total)
  store.commit('video/setBaseUrl', result.base)
  tableData.value = result.data
  total.value = result.total
}
getListData()
// console.log(getListData)
// 处理数据不重新加载的问题
onActivated(getListData)

/**
 * size 改变触发
 */
const handleSizeChange = currentSize => {
  size.value = currentSize
  getListData()
}

/**
 * 页码改变触发
 */
const handleCurrentChange = currentPage => {
  page.value = currentPage
  getListData()
}
/**
 * 查看按钮点击事件
 */
const router = useRouter()
const onShowClick = id => {
  router.push(`/video/${id}`)
}
</script>

<style lang="scss" scoped>
@import '~@/styles/variables.module.scss';

.index-container {
  min-height: calc(100vh - 50px);
  font-family: PingFang SC,DFPKingGothicGB-Regular,sans-serif;
  .t-list{
    min-height: calc(100vh - 50px);
    .titleWrap {
      color: rgba($color: #fff, $alpha: .9);
      line-height: 28px;
      width: calc( (100vw - #{$sideBarWidth}) / 5 - 36px);
    }
    .authorWrap{
      color: rgba($color: #fff, $alpha: .5);
      line-height: 22px;
      font-size: 14px;
    }
    ul {
      display: flex;
      flex-wrap: wrap;
      position: relative;
    }
    li{
      padding: 3px;
      list-style: none;
      margin-right: 15px;
      margin-bottom: 15px;
      border: 1px solid #121111;
      border-radius: 10px;
    }
    img{
      width: calc( (100vw - #{$sideBarWidth}) / 5 - 36px);
      height: 180px;
    }
  }
}
  .pagination {
    margin-top: 20px;
    text-align: center;
  }
</style>
