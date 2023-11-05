<template>
  <el-scrollbar style="height:100%">
    <div class="index-container">
      <div class="t-list">
        <ul>
          <li
            v-for="(item, index) in tableData"
            :key="index"
          >
            <div @click="onShowClick(item.id)">
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
    :page-sizes="[15, 20, 25, 30, 35]"
    :page-size="size"
    layout="total, sizes, prev, pager, next, jumper"
    :total="total"
  >
  </el-pagination>
</template>

<script setup>
import { ref, onActivated } from 'vue'
import { getVideoList } from '@/api/video'
import { useRouter } from 'vue-router'
// // 数据相关
// const tableData = ref([])
const total = ref(0)
const page = ref(1)
const size = ref(15)
const tableData = ref([
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' },
  { url: 'title.png', desc: '测试图', id: 1, author: '遁形' }])

// 获取数据的方法
const getListData = async () => {
  const result = await getVideoList({
    page: page.value,
    size: size.value
  })
  tableData.value = result.list
  total.value = result.total
}
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
      height: 150px;
    }
  }
}
  .pagination {
    margin-top: 20px;
    text-align: center;
  }
</style>
