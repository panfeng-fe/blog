<template>
  <div class="conrainer">
      <!-- header开始 -->
      <header class="header">
        <!-- header左边logo开始 -->
        <div class="header-info">
          <span class="header-info-logo">Pain</span>
          <span class="header-info-txt">Web开发工程师</span>
        </div>

        <!-- header右边菜单栏开始 -->
        <div class="header-menu">
            <keep-alive>
                <nuxt-link to="/" active-class="header-menu-item-active" class="header-menu-item" exact>首页</nuxt-link>
            </keep-alive>
            <!-- <keep-alive>
                <nuxt-link to="/resume" active-class="header-menu-item-active" class="header-menu-item" exact>简历</nuxt-link>
            </keep-alive> -->
            <keep-alive>
                <nuxt-link to="/about" active-class="header-menu-item-active" class="header-menu-item" exact>关于</nuxt-link>
            </keep-alive> 
        </div>
      </header>

      <!-- content开始 -->
      <nuxt class="content" />

      <!-- side开始 -->
      <Side/>

      <!-- footer开始 -->
      <footer class="footer">
          <div>2020 Pain BLOG By 潘峰，本站博客未经授权禁止转载。</div>
          <div>系统有Nuxt.js + iris(Golang)构建</div>
          <div>Copyright © 皖ICP备19002392号-1</div>   
      </footer>

  </div>
</template>

<script>
import Side from '@/components/side'
export default {
    components:{
      Side
    },
}
</script>
