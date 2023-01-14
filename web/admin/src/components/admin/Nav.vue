<script setup lang="ts">
import { ref } from 'vue'
import type { Ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  BankOutlined,
  BarsOutlined,
  FolderViewOutlined,
  UserOutlined,
  SettingOutlined
} from '@ant-design/icons-vue';
const router = useRouter();
// const route = useRoute();
const collapsed: Ref<boolean> = ref(false);
const goToPage = (item: any) => {
  router.push(`/${item.key}`).catch(err => err)
}
</script>

<template>
  <a-layout-slider breakpoint="lg" v-model="collapsed">
    <div class="log">
      <h2>{{ collapsed ? 'Blog' : 'My Blog' }}</h2>
    </div>
    <a-menu mode="inline" style="width:100%;" theme="dark" @click="goToPage">
      <a-menu-item key="index">
        <template #icon>
          <bank-outlined />
        </template>
        <span>仪表盘</span>
      </a-menu-item>
      <a-sub-menu key="sub1">
        <template #icon>
          <bars-outlined />
        </template>
        <template #title>文章管理</template>
        <a-menu-item key="addart">
          写文章
        </a-menu-item>
        <a-menu-item key="artlist">
          文章列表
        </a-menu-item>
      </a-sub-menu>
      <a-sub-menu key="sub2">
        <template #icon>
          <bars-outlined />
        </template>
        <template #title>评论管理</template>
        <a-menu-item key="commentlist">评论管理</a-menu-item>
      </a-sub-menu>
      <a-menu-item key="catelist">
        <template #icon>
          <folder-view-outlined />
        </template>
        <span>分类列表</span>
      </a-menu-item>
      <a-menu-item key="userlist">
        <template #icon>
          <user-outlined />
        </template>
        <span>用户列表</span>
      </a-menu-item>
      <a-menu-item key="profile">
        <template #icon>
          <setting-outlined />
        </template>
        <span>个人设置</span>
      </a-menu-item>
    </a-menu>
  </a-layout-slider>
</template>

<style scoped>
.log {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 64px;
  /* margin: 16px; */
  font-size: 16px;
  background-color: #fff;
}

h2 {
  margin: 0;
  text-align: center;
}
</style>