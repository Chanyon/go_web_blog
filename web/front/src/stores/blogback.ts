import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
export const useBackImgStore = defineStore('topBackImg', () => {
  const topBackImg = ref("");
  const getBackImg = computed(() => topBackImg.value)
  function setBackImg(url: string) {
    topBackImg.value = url;
  }
  return { topBackImg, getBackImg,setBackImg}
});
