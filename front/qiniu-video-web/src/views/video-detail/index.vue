<template>
  <div class="pa123">
    <div class="show-video">
      <div class="box">
        <videoPlayer  class="video-player vjs-custom-skin"
              ref="VideoPlayer"
              :playsinline="true"
              :options="playerOptions"
          >
        </videoPlayer>
      </div>
      <div class="operation">
        <div>
          <el-icon :size="50"><ArrowUpBold /></el-icon>
        </div>
        <div>
          <el-icon :size="50"><ArrowDownBold /></el-icon>
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
        <span>使用video.js 在网站中搭建视频</span>
      </h1>
      <el-divider class="video-comment" content-position="left">
        <span>评论</span>

      </el-divider>
    </div>
  </div>
</template>
<script>
import { VideoPlayer } from '@videojs-player/vue'
import 'video.js/dist/video-js.css'
import 'videojs-contrib-hls'

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
        autoplay: false,
        muted: false,
        loop: false,
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
      }
    }
  },
  created() {
    let url = 'http://s3hk0nelw.bkt.clouddn.com/video/mp4.v1080'
    url = 'http://s3hk0nelw.bkt.clouddn.com/avthumb_test_target.mp4'
    this.playerOptions.sources[0].src = url
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
