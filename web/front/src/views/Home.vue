<script lang="ts" setup>
// import ArticleList from '../components/ArticleList.vue';
import Profile from '../components/Profile.vue';
import TopBar from '../components/TopBar.vue';
import Footer from '../components/Footer.vue';
/**
 * 问题1：切换路由 闪屏.
 */
// todo: footer component
//bg:rgb(14, 20, 27) nav:#617bbf fg:#fff #ff7008 
</script>
<template>
  <div>
    <TopBar></TopBar>
    <main class="main">
      <div class="left">
        <Profile></Profile>
      </div>
      <div class="right">
        <router-view :key="$route.path">
        </router-view>
        <!-- <Transition> -->
        <!-- <component :is="ArticleList" /> -->
        <!-- </Transition> -->
      </div>
    </main>
    <footer>
      <Footer></Footer>
    </footer>
  </div>
</template>

<style scoped>
.main {
  width: 100vw;
  margin-top: 20px;
  display: flex;
  justify-content: center;
  /* align-items: center; */
  padding: 20px 0px 10px 0px;
}

/* .v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
} */
</style>