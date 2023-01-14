<script lang="ts" setup>
import { reactive, ref } from "vue";
import { loginApi } from "../api/loginAndReg";
// defineProps(['isLogin']);
const emit = defineEmits(['update:isLogin', 'update:isHasUserInfo']);

type UserI = { username: string, password: string };
const userInfo = reactive<UserI>({
  username: "",
  password: "",
});
const errorMessage = reactive({
  isError: false,
  message: "用户名或密码错误!"
});

const text = ref("");

const handleLoginClick = async () => {
  // todo 字符验证
  const { data: res } = await loginApi(userInfo);
  if (res.status === 200) {
    // text.value = res.message;
    window.sessionStorage.setItem("userInformation", JSON.stringify(res));
    if (window.sessionStorage.getItem("userInformation") !== null) {
      emit("update:isHasUserInfo", true);
    } else {
      emit("update:isHasUserInfo", false);
    }
    emit("update:isLogin", false);
  } else {
    errorMessage.isError = true;
    // text.value = res.message;
    emit("update:isLogin", true);
  }
};
</script>
<template>
  <div class="login">
    <h3>用户登录</h3>
    <form class="form">
      <div class="up">
        <div>
          <label for="password">用户名:</label>
        </div>
        <div>
          <input type="text" id="username" v-model="userInfo.username" placeholder="username">
        </div>
      </div>
      <div class="up">
        <div>
          <label for="password">密码:</label>
        </div>
        <div>
          <input type="password" id="password" v-model="userInfo.password" placeholder="password">
        </div>
      </div>
    </form>
    <div v-if="errorMessage.isError" class="error-msg">
      {{ errorMessage.message }}
    </div>
    <div class="btn">
      <button @click="handleLoginClick">登录</button>
      <button @click="emit('update:isLogin', false)">取消</button>
    </div>
  </div>
</template>
  
<style scoped>
.login {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  width: 450px;
  height: 280px;
  padding: 20px;
  border-radius: 10px;
  color: #fff;
  background-color: rgba(0, 0, 0, 0.4);
}

.form>div {
  display: flex;
  justify-content: center;
  height: 40px;
  margin-bottom: 10px;
}

.up>div:first-child {
  width: 60px;
  line-height: 40px;
  text-align: right;
  margin-right: 10px;
}

.up>div:last-child {
  width: 180px;
  height: 100%;
}

.up>div:last-child input {
  width: 100%;
  height: 100%;
  padding: 8px;
  outline: none;
  font-size: 16px;
  border: 1px solid #242424;
  border-radius: 6px;
}

.login>h3 {
  font-size: 18px;
  color: #546ff7;
  margin-bottom: 10px;
}

.btn {
  display: flex;
  justify-content: space-around;
  margin-top: 20px;
  width: 40%;
}

.btn>button {
  border-radius: 8px;
  border: 1px solid transparent;
  padding: 0.6em 1.2em;
  font-size: 1em;
  font-weight: 500;
  font-family: inherit;
  cursor: pointer;
  transition: border-color 0.25s;
  outline: none;
}

.btn>button:first-child {
  background-color: #546ff7;
}

.btn>button:last-child {
  background-color: #ffd978;
}

.btn>button:hover {
  border-color: #242424;
}

.error-msg {
  color: red;
  font-size: 14px;
  font-weight: bold;
}
</style>