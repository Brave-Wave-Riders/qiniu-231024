<template>
  <div class="uploadWrap">
    <div>
        <el-input v-model="input" class='titleWrap' placeholder="请输入标题" />
        <el-select v-model="value" value-key="id" placeholder="请选择内容类型">
        <el-option
            v-for="item in options"
            :key="item.id"
            :label="item.label"
            :value="item"
        />
        </el-select>
    </div>
    <el-upload
        class="upload-video"
        drag
        action=""
        :before-upload="handleUpload"
        :auto-upload="true"
    >
        <el-icon class="el-icon--upload"><upload-filled /></el-icon>
        <div class="el-upload__text ml-3">
        拖拽上传
        </div>
        <el-button class="ml-3" type="success">
        <em>点击上传</em>
        </el-button>
        <template #tip>
        <div class="el-upload__tip">
            仅支持mp4文件格式
        </div>
        </template>
    </el-upload>
  </div>
</template>

<script setup>
import { UploadFilled } from '@element-plus/icons-vue'
import { ElNotification } from 'element-plus'
import * as qiniu from 'qiniu-js'
import { ref } from 'vue'

const observer = {
  next(res) {
    // ...
    // console.log(res)
  },
  error(err) {
    // ...
    console.log(err)
  },
  complete(res) {
    // ...
    console.log(res)
  }
}
const handleUpload = (file) => {
  console.log(file)
  if (file.type !== 'video/mp4') {
    ElNotification({
      title: '格式错误',
      message: '请上传mp4文件',
      type: 'warning'
    })
    return false
  }

  const key = file.name
  const token = ''
  const putExtra = {
    fname: file.name,
    params: {},
    mimeType: 'video/mp4'
  }
  const observable = qiniu.upload(file, key, token, putExtra)

  const subscription = observable.subscribe(observer) // 上传开始
  console.log(subscription)

  return false
}

const input = ref('')
const value = ref('')
const options = ref([
  { id: 1, label: 'Option A', desc: 'Option A - 230506' },
  { id: 2, label: 'Option B', desc: 'Option B - 230506' },
  { id: 3, label: 'Option C', desc: 'Option C - 230506' },
  { id: 4, label: 'Option A', desc: 'Option A - 230507' }
])
</script>

<style lang="scss">
.uploadWrap{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin-top: 3%;
  .titleWrap{
    .el-input__inner {
      background: #16304e;
    }
    .el-input{
        --el-input-background-color:#16304e !important;
    }
  }
}
.upload-video{
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100%;
    margin-top: 1%;
    .el-upload{
        background: #16304e;
    }
    .el-upload-dragger {
        background-color: #16304e;
        width: 640px;
        height: 320px;
    }
    .el-icon{
        --font-size: 50px;
        --color: gray;
        margin-top: 10%;
    }
    .el-button{
        margin: 10%;
    }

}

</style>
