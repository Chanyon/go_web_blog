<script lang="ts" setup>
import { onBeforeMount, ref, watchEffect } from 'vue';
import { useRouter } from 'vue-router';
import { getCategoryApi } from '../api/category';
import Login from '../components/Login.vue';
import Register from '../components/Register.vue';
import Message from '../components/common/Message.vue';
import { useBackImgStore } from '../stores/blogback';
import { storeToRefs } from 'pinia';

type CateItemI = {
  id: number,
  name: string,
  ID?: number,
  CreatedAt?: string,
  DeleteAt?: string,
  UpdatedAt?: string,
};
const cates = ref<Array<CateItemI>>([{ id: 1001, name: "JavaScript" }]);
const isLogin = ref<boolean>(false);
const isRegister = ref<boolean>(false);
const isHasUserInfo = ref(false);
const activeItem = ref(false);
const activeCateItem = ref(-1);
const colors = ["red", "orange", "green", "skyblue"];
const randomColor = colors[Math.floor(Math.random() * colors.length)];
const searchTitle = ref("");
const titleMessage = ref(false);

const store = useBackImgStore();
const { getBackImg } = storeToRefs(store);

const getCategory = async () => {
  const { data: res } = await getCategoryApi();
  cates.value = res.data;
};
onBeforeMount(() => {
  getCategory();
});

const loginOut = () => {
  window.sessionStorage.clear(); //???
  isHasUserInfo.value = false;
  isLogin.value = false;
};

const router = useRouter();
const getCateArticleList = (cateId: number, idx: number) => {
  router.push(`/article/category/${cateId}`)
  activeItem.value = false;
  activeCateItem.value = idx;
}

// fixed nav top value
const topIsGt300 = ref(false);

watchEffect(() => {
  if (window.sessionStorage.getItem("userInformation") != null) {
    isHasUserInfo.value = true;
  }
  window.addEventListener("scroll", (e) => {
    e.preventDefault();
    const top = document.documentElement.scrollTop;
    if (top >= 300 && top <= 600) {
      topIsGt300.value = true;
    } else {
      topIsGt300.value = false;
    }
  });

});

const handleHome = (idx: number) => {
  if (idx === 0) activeItem.value = true;
  activeCateItem.value = -1;
  router.push(`/`)
};

const handleSearch = () => {
  if (searchTitle.value === "") {
    titleMessage.value = true;
    setTimeout(() => {
      titleMessage.value = false;
    }, 4000);
    return;
  }
  router.push(`/search/${searchTitle.value}`);
};

</script>
<template>
  <div class="bar"
    :style="{ background: `url(http://127.0.0.1:3000${getBackImg}) no-repeat`, backgroundSize: 'cover' }">
    <div :class="['top-bar', topIsGt300 ? 'fixed-top' : '']">
      <VMenu :triggers="['hover', 'focus']" :hideTriggers="['click']" :distance="6">
        <div class="top logo" :class="[isLogin ? 'login-logo' : '']"
          :style="{ backgroundColor: isHasUserInfo ? randomColor : '' }">{{ isHasUserInfo ? '退出' : '登录' }}
        </div>
        <!-- <button>Click me</button> -->
        <template #popper>
          <div v-if="!isHasUserInfo" class="btn">
            <!-- <a v-close-popper>Close</a> -->
            <button @click="isLogin = !isLogin">登录</button>
            <button @click="isRegister = !isRegister">注册</button>
          </div>
          <div v-if="isHasUserInfo" class="btn">
            <button @click="loginOut">退出</button>
          </div>
        </template>
      </VMenu>
      <div class="top nav">
        <nav>
          <ul>
            <li @click="handleHome(0)" :class="{ 'active-li-item': activeItem }">首页</li>
            <li v-for="(cate, idx) in cates" :key="idx" @click="getCateArticleList(cate.id, idx)"
              :class="{ 'active-li-item': activeCateItem === idx }">{{ cate.name }}</li>
          </ul>
        </nav>
      </div>
      <div class="top search">
        <span>搜索</span>
        <input v-model="searchTitle" @keyup.enter="handleSearch" type="text" placeholder="搜索标题...">
      </div>
    </div>
    <Login v-if="isLogin && !isRegister" @update:isHasUserInfo="(newValue: boolean) => isHasUserInfo = newValue"
      @update:isLogin="(newValue: boolean) => isLogin = newValue"></Login>
    <Register v-if="isRegister && !isLogin" @update:isRegister="(newValue: boolean) => isRegister = newValue"
      @update:isHasUserInfo="(newValue: boolean) => isHasUserInfo = newValue"></Register>
    <Message v-if="titleMessage" text="请输入文章标题" TipsType="error" :autoClose="true"></Message>
  </div>
</template>
  
<style scoped>
.bar {
  /* width: 100%; */
  height: 450px;
  background-color: #fff;
  background-position: center center;
  background-size: cover;
  background-repeat: no-repeat;
  position: relative;
}

.top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px;
  width: 100vw;
  height: 80px;
  color: #fff;
  /* background-color: rgba(0, 0, 0, .4); */
  background-color: hsla(210, 30%, 8%, 1);
  opacity: 0.8;
}

.fixed-top {
  position: fixed;
  top: 0;
}

.top {
  height: 60px;
}

.logo {
  width: 45px;
  height: 45px;
  margin-left: 90px;
  margin-right: 10px;
  border-radius: 50%;
  line-height: 45px;
  text-align: center;
  font-size: 14px;
  border: 2px solid #fff;
  cursor: pointer;
}

.nav {
  display: flex;
  justify-content: center;
  align-items: center;
}

.nav ul {
  display: flex;
  justify-content: space-between;
  align-items: center;
  list-style: none;
}

.nav ul li {
  width: 100px;
  height: 40px;
  margin: 0 5px 0 0px;
  padding: 0 8px 0 8px;
  text-align: center;
  line-height: 40px;
  border: 1px solid #fff;
  border-radius: 6px;
  background-color: hsla(230, 100%, 69%, 1);
  cursor: pointer;
}

.nav ul li.active-li-item {
  border: 1px solid rgb(255, 116, 106);
  background-color: #242424;
}

.nav ul li:hover {
  border: 1px solid rgb(255, 116, 106);
  background-color: #242424;
  transition: background-color 0.3s linear;
}

.search {
  display: flex;
  align-items: center;
  margin-right: 90px;
}

.search input {
  width: 210px;
  height: 40px;
  padding: 10px;
  border-radius: 20px;
  outline: none;
}

.search span {
  width: 40px;
  margin-right: 10px;
}

.btn>button {
  padding: 10px;
}
</style>