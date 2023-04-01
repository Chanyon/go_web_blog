<script lang="ts" setup>
import Profile from '../components/Profile.vue';
import TopBar from '../components/TopBar.vue';
import Footer from '../components/Footer.vue';
import { ref, watch } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const transitionNameIn = ref("");
const transitionNameOut = ref("");
const reg = /\/article\/category\/d*/gi;
const reg2 = /\/article\/details\/d*/gi;

watch(() => router.currentRoute.value.path, (newValue, oldValue) => {
  if (reg.test(newValue) || newValue === "/") {
    transitionNameIn.value = "animate__animated animate__fadeInLeft";
    transitionNameOut.value = "animate__animated animate__fadeInLeft";
  }
}, { immediate: true });
</script>

<template>
  <div>
    <TopBar></TopBar>
    <main class="main">
      <div class="left">
        <Profile></Profile>
      </div>
      <div class="right">
        <router-view :key="$route.path" v-slot="{ Component }">
          <Transition name="custom-classes" :enter-active-class="transitionNameIn"
            :leave-active-class="transitionNameOut">
            <KeepAlive>
              <component :is="Component" :key="$route.path" />
            </KeepAlive>
          </Transition>
        </router-view>
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

/* .fade-enter-active,
.fade-leave-active {
  transition: opacity 1s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
} */
</style>