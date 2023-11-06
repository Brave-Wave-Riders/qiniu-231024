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
        ref="file"
        class="upload-video"
        drag
        :auto-upload="false"
        :http-request="upload"

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

const input = ref('')
const value = ref('')
const options = ref([
  { id: 1, label: 'Option A', desc: 'Option A - 230506' },
  { id: 2, label: 'Option B', desc: 'Option B - 230506' },
  { id: 3, label: 'Option C', desc: 'Option C - 230506' },
  { id: 4, label: 'Option A', desc: 'Option A - 230507' }
])

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
const upload = (fileObj) => {
  console.log(fileObj.file)
  const file = fileObj.file
  console.log(file.type)
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
  if (file.type !== 'video/mp4') {
    ElNotification({
      title: '格式错误',
      message: '请上传mp4文件',
      type: 'warning'
    })
    return false
  }
  console.log('upload...')
  uploadVideo({
    uid: 1,
    name: input.value,
    type: 2,
    img_url: '123'
  }).then(result => {
    console.log(result)
    console.log('upload......')
    if (!result) {
      return false
    }
    // console.log(result)
    const token = result.token
    const base = result.base

    console.log(token)
    console.log(base)

    const key = input.value

    const putExtra = {
      fname: key,
      params: {},
      mimeType: 'video/mp4'
    }
    const observable = qiniu.upload(file, key, token, putExtra)

    const subscription = observable.subscribe(observer) // 上传开始
    console.log(subscription)
  })

  return true
}
const file = ref()
const submitUpload = () => {
  file.value.submit()
}

</script>

<style lang="scss">
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
