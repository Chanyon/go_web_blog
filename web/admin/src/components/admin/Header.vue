<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
const router = useRouter();
// const route = useRoute();
const loginOut = () => {
  window.sessionStorage.clear()
  router.push(`/login`);
}
</script>

<template>
  <a-button type="danger" @click="loginOut">退出</a-button>
</template>

<style scoped>
</style>
