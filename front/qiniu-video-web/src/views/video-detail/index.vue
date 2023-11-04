<template>
  <div class="box">
    <videoPlayer  class="video-player vjs-custom-skin"
          ref="VideoPlayer"
          :playsinline="true"
          :options="playerOptions"
      >
    </videoPlayer>
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
        autoplay: true,
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
    const url = 'http://s3hk0nelw.bkt.clouddn.com/video/mp4.v1080'
    this.playerOptions.sources[0].src = url
  }
}
</script>

<style>
.box{
  margin: 10% 20%;
}
.video-js .vjs-big-play-button{
  margin-left:43%;
  margin-top: 25%;
}
</style>
