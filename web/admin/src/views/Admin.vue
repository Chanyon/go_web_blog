<script setup lang="ts">
import Nav from '../components/admin/Nav.vue'
import Header from '../components/admin/Header.vue'
import Footer from '../components/admin/Footer.vue'
import { useRouter, useRoute } from 'vue-router'
// const router = useRouter();
const route = useRoute();

</script>

<template>
  <a-layout class="container">
    <a-layout-sider>
      <Nav></Nav>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="headerBtn">
        <Header></Header>
      </a-layout-header>
      <a-layout-content class="content">
        <router-view :key="route.fullPath"></router-view>
      </a-layout-content>
      <a-layout-footer>
        <Footer></Footer>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<style scoped>
.container {
  height: 100vh;
}

.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
.ant-layout-sider{
  height: 100%;
  /* min-height: 1100px; */
}
.ant-layout{
  height:1100px;
}
</style>