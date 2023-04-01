<script lang="ts" setup>
import { onBeforeMount, ref } from 'vue';
import { getProfileApi } from '../api/profile';
import { useBackImgStore } from '../stores/blogback';
const profile = ref({
  data: {
    avatar: "",
    bili: "",
    desc: "",
    email: "",
    icp_record: "",
    id: "",
    img: "",
    name: "",
    qq_chat: "",
    wei_bo: "",
    wei_chat: "",
  }
});

const skills = ["JavaScript", "TypeScript", "HTML", "CSS", "C/C++", "Rust", "Go", "zig"];

const store = useBackImgStore();
const { setBackImg } = store;
const getProfile = async () => {
  const { data: res } = await getProfileApi(`user/profile/1`);
  profile.value.data = res.data;
  setBackImg(res.data.img);
};

onBeforeMount(() => {
  getProfile();
});

</script>
<template>
  <div class="profile animate__animated animate__fadeInLeft">
    <div class="avatar">
      <div class="top">
        <img :src="`http://127.0.0.1:3000${profile.data.avatar}`" alt="" srcset="">
        <p>{{ profile.data.name }}</p>
      </div>
      <div class="bottom">
        <h3>About Me:</h3>
        <p>{{ profile.data.desc ? profile.data.desc : "" }}</p>
      </div>
    </div>
    <div class="account">
      <div class="account-info">
        <img src="../assets/gmail.svg" alt="" srcset="">
        <span>{{ profile.data.qq_chat }}</span>
      </div>
      <div class="account-info">
        <img src="../assets/gmail.svg" alt="" srcset="">
        <span>{{ profile.data.qq_chat }}</span>
      </div>
      <div class="account-info">
        <img src="../assets/gmail.svg" alt="" srcset="">
        <span>{{ profile.data.qq_chat }}</span>
      </div>
      <div class="account-info">
        <img src="../assets/gmail.svg" alt="" srcset="">
        <span>{{ profile.data.qq_chat }}</span>
      </div>
    </div>
    <div class="skill">
      <h3>Learn program:</h3>
      <p v-for="(item, idx) in skills" :key="idx">
        {{ idx + 1 }}.{{ item }}
      </p>
    </div>
  </div>
</template>
  
<style scoped>
.profile {
  width: 265px;
  border: 1px solid #fff;
  background-color: hsla(210, 30%, 8%, 1);
  color: #535455;
  border-radius: 8px;
  padding: 10px;
  margin-right: 50px;
}

.avatar .top {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 180px;
  background-image: linear-gradient(135deg, #72EDF2 10%, #5151E5 100%);
  color: #fff;
  overflow: hidden;
}


.avatar .top>p {
  margin-top: 5px;
}

.avatar img {
  width: 80px;
  height: 80px;
  border: 1px solid;
  border-radius: 50%;
}

.avatar .bottom>h3 {
  text-align: left;
  padding-left: 15px;
}

.avatar .bottom>p {
  font-size: 16px;
  color: #535355;
  padding-left: 15px;
  text-align: left;
}

.account {
  margin-top: 10px;
  border-top: 2px solid #535355;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
}

.account-info {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  padding-left: 15px;
}

.account-info>img {
  width: 16px;
  height: 16px;
  margin-right: 8px;
}

.skill {
  text-align: left;
}

.skill>h3 {
  padding-left: 15px;
}

.skill>p {
  padding-left: 15px;
}
</style>