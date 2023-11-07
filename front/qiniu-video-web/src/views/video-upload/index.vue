<template>
<!-- 进度条 -->
    <div v-if="loading">
        <el-progress :percentage="progressPercent" :stroke-width="15" striped class="progressWrap" />
        <!-- <el-progress
            :percentage="progressPercent" :text-inside="false"
            :stroke-width="50" status="success">
        </el-progress> -->
    </div>
  <div class="uploadWrap" v-else>
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
            ref="uploadFile"
            class="upload-video"
            drag
            :auto-upload="false"
            :http-request="upload"
            v-model:file-list="fileList"
            :on-change="handleChange"
            :limit="1"
            :on-exceed="handleExceed"
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
        <el-button class="upload_btn ml-3" type="success" @click="submitUpload">
        确认上传
        </el-button>
    </div>

</template>

<script setup>
import { UploadFilled } from '@element-plus/icons-vue'
import { ElNotification } from 'element-plus'
import * as qiniu from 'qiniu-js'
import { ref } from 'vue'
import { uploadVideo } from '@/api/video'

// const Base64 = require('js-base64').Base64

const input = ref('')
const value = ref('')
const options = ref([
  { id: 1, label: '知识', desc: 'Option A - 230506' },
  { id: 2, label: '游戏', desc: 'Option B - 230506' },
  { id: 3, label: '娱乐', desc: 'Option C - 230506' },
  { id: 4, label: '音乐', desc: 'Option A - 230507' },
  { id: 5, label: '美食', desc: 'Option A - 230507' },
  { id: 6, label: '体育', desc: 'Option A - 230507' },
  { id: 7, label: '时尚', desc: 'Option A - 230507' }
])
const loading = ref(false)
const progressPercent = ref(0.0)
const uploadFile = ref(File)
const handleExceed = (files) => {
  console.log(uploadFile)
  uploadFile.value.clearFiles()
  const fileNew = files[0]
  console.log(fileNew)
  //   file.value.uid = genFileId()
  uploadFile.value.handleStart(fileNew)
}
const observer = {
  next(res) {
    // ...
    console.log(res)
    progressPercent.value = res.total.percent
    console.log(loading)
  },
  error(err) {
    // ...
    console.log(err)
    loading.value = false
    progressPercent.value = 0
    ElNotification({
      title: '错误',
      message: '文件上传失败',
      type: 'error'
    })
  },
  complete(res) {
    // ...
    console.log(res)
    loading.value = false
    progressPercent.value = 0
    ElNotification({
      title: '通知',
      message: '文件上传成功',
      type: 'success'
    })
    uploadFile.value.clearFiles()
  }
}
const upload = (fileObj) => {
  console.log(fileObj.file)

  const file = fileObj.file
  console.log(input.value)
  console.log(value.value.valueOf().label)
  if (input.value === '' || !value.value.valueOf().label) {
    ElNotification({
      title: '警告',
      message: '请输入完整信息',
      type: 'warning'
    })
    return false
  }
  if (input.value.length > 79) {
    ElNotification({
      title: '警告',
      message: '标题不能超过79字符',
      type: 'warning'
    })
    return false
  }
  if (file.type !== 'video/mp4') {
    ElNotification({
      title: '格式错误',
      message: '请上传mp4文件',
      type: 'warning'
    })
    return false
  }
  console.log('upload...')
  console.log(value.value.valueOf().id)
  //   const name = Base64.encode(input.value)
  const name = input.value
  //   console.log(name)

  uploadVideo({
    uid: 1,
    name: name,
    type: value.value.valueOf().id,
    img_url: '123'
  }).then(result => {
    console.log(result)
    console.log('upload......')
    if (!result) {
      return false
    }
    loading.value = true
    // console.log(result)
    const token = result.token
    const base = result.base

    console.log(token)
    console.log(base)

    const key = name

    const putExtra = {
      fname: key,
      params: {},
      mimeType: 'video/mp4'
    }
    const observable = qiniu.upload(file, key, token, putExtra)

    const subscription = observable.subscribe(observer) // 上传开始
    console.log(subscription)
  }).catch(error => {
    console.error(error) // Stacktrace
  })

  return true
}

const submitUpload = () => {
  uploadFile.value.submit()
}

</script>

<style lang="scss">
.progressWrap{
    margin-top: 100px ;
    margin-bottom: 15px;
    width: 100%;
}
.uploadWrap{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin-top: 3%;

  .upload_btn{
    margin-top: 3%;
  }
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
