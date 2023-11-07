<template>
  <router-view>
  <div class="pa123">
    <div class="show-video">
      <div class="box">
        <video-player id='my-player' class="video-player vjs-custom-skin"
              ref="videoPlayer"
              :playsinline="true"
              :options="playerOptions"
              :loop="true"
              :autoplay="true"
              @mounted="handleMounted"
              :source="[videoSource]"
          >
        </video-player>
      </div>
      <div class="operation">
        <div>
          <el-icon :size="50"><ArrowUpBold @click="onUpBtnClicked"/></el-icon>
        </div>
        <div>
          <el-icon :size="50"><ArrowDownBold @click="onDownBtnClicked"/></el-icon>
        </div>
        <br/>
        <div>
          <el-avatar
            shape="circle"
            :size="50"
            :src="require('@/assets/default_img.jpg')"
          ></el-avatar>
        </div>
      </div>
    </div>
    <div class="video-info">
      <h1 class="video-title">
        <span>{{key}}</span>
      </h1>
      <el-divider class="video-comment" content-position="left">
        <span>评论</span>

      </el-divider>
    </div>
  </div>
  </router-view>
</template>
<script>
import { VideoPlayer } from '@videojs-player/vue'
// import videojs from 'video.js'
import 'video.js/dist/video-js.css'
import 'videojs-contrib-hls'
// import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
import videojs from 'video.js'

// const Base64 = require('js-base64').Base64
export default {
  components: {
    VideoPlayer
  },
  data() {
    return {
      playerOptions: {
        // width: '640px',
        // height: '480px',
        // width: '100%',
        // height: (window.screen.width * 9) / 16,
        autoplay: true,
        muted: false,
        loop: true,
        preload: 'auto',
        language: 'zh-CN',
        aspectRatio: '16:9',
        fluid: true, // 当true时，Video.js player将拥有流体大小。换句话说，它将按比例缩放以适应其容器。
        sources: [{
          type: 'application/x-mpegURL', // mp4格式视频,若为m3u8格式，type需设置为 application/x-mpegURL
          withCredentials: false,
          src: '' // url地址
        }],
        techOrder: ['html5'],
        flash: { hls: { withCredentials: false } },
        html5: { hls: { withCredentials: false } },
        notSupportedMessage: '此视频暂无法播放，请稍后再试', // 允许覆盖Video.js无法播放媒体源时显示的默认信息。
        controls: true, // 是否显示控制栏
        controlBar: {
          timeDivider: true,
          durationDisplay: true,
          remainingTimeDisplay: false,
          fullscreenToggle: true // 是否显示全屏按钮
        }
      },
      videoSource: {
        type: 'application/x-mpegURL', // mp4格式视频,若为m3u8格式，type需设置为 application/x-mpegURL
        withCredentials: false,
        src: '' // url地址
      },
      key: '',
      player: videojs.player,
      showRouterView: true
    }
  },
  created() {
    const route = useRoute()
    this.key = route.params.id
    console.log(this.key)
    const store = useStore()
    // const url = store.getters.base + Base64.encode(this.key)

    const url = store.getters.base + this.key
    console.log(url)
    // let url = 'http://s3hk0nelw.bkt.clouddn.com/video/mp4.v1080'
    // url = 'http://s3hk0nelw.bkt.clouddn.com/%E7%94%B2%E6%96%B9%E8%AF%B4%EF%BC%9A%E4%BD%A0%E6%80%A7%E6%84%9F%E4%B8%80%E7%82%B9%EF%BC%8C%E6%93%A6%E8%BE%B9%E4%B8%80%E7%82%B9%20-%201.%E7%94%B2%E6%96%B9%E8%AF%B4%EF%BC%9A%E4%BD%A0%E6%80%A7%E6%84%9F%E4%B8%80%E7%82%B9%EF%BC%8C%E6%93%A6%E8%BE%B9%E4%B8%80%E7%82%B9(Av917944767,P1)'
    this.playerOptions.sources[0].src = url
  },
  methods: {
    onUpBtnClicked() {
      console.log('up')
      // const route = useRoute()
      const key = this.key
      // const store = useStore()
      // const videos = store.getters.videoData
      const videos = this.$store.getters.videoData
      const len = videos.length
      console.log(key + ' ' + videos + ' ' + len)
      // const router = useRouter()
      for (let i = 0; i < len; i++) {
        if (videos[i].desc === key) {
          console.log('find! ' + i)
          if (i === 0) {
            break
          }
          this.$router.push({
            path: `/video/${videos[i - 1].desc}`,
            query: {
              date: new Date().getTime()
            }
          })
          this.$route.params.id = videos[i - 1].desc

          break
        }
      }
    },
    onDownBtnClicked() {
      console.log('down')
      const key = this.key
      // const store = useStore()
      // const videos = store.getters.videoData
      const videos = this.$store.getters.videoData
      const len = videos.length
      console.log(key + ' ' + videos + ' ' + len)
      // const router = useRouter()
      for (let i = 0; i < len; i++) {
        console.log(videos[i].desc)
        if (videos[i].desc === key) {
          console.log('find! ' + i)
          if (i === len - 1) {
            break
          }

          this.$router.push({
            path: `/video/${videos[i + 1].desc}`,
            query: {
              date: new Date().getTime()
            }
          })
          break
        }
      }
    },
    handleMounted(payload) {
      // this.player.value = payload.player
      this.player = payload.player.player_
      this.id = this.player.id_
      // payload.player.play()
      console.log('Basic player mounted', this.player)
    }
  }

}

</script>

<style lang="scss">
.show-video{
  display: flex;
  flex-direction: row;
  .box{
    width: 85%;
    // height: 737.88px;
    position: relative;
    overflow: auto;
  }
  .operation{
    display: flex;
    flex-direction: column;
    justify-content: center;
    margin-top: 50px;
    margin-left: 50px;
    right: 0px;
    width: 15%;
    // height: 737.88px;
    position: relative;
    overflow: auto;
  }
}
.video-title span{
  color: rgba($color: #fff, $alpha: .9);
  font-family: PingFang SC,DFPKingGothicGB-Regular,sans-serif;
  font-size: 32px;
  font-weight: 400;
  line-height: 64px;
}
.el-divider{
  background-color: #181f2f;
}
.el-divider__text{
  background: #16304e;
}
.video-comment span{
  color: rgba($color: #fff, $alpha: .34);
  background: #16304e;
  font-family: PingFang SC,DFPKingGothicGB-Regular,sans-serif;
  font-size: 25px;
  font-weight: 100;
  line-height: 32px;
}
</style>
