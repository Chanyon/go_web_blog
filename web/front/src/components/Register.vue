<script lang="ts" setup>
import { reactive, ref, watch } from 'vue';
import { loginApi, registerApi } from '../api/loginAndReg';
import { setStorage } from '../uitls/setStorage';


const emit = defineEmits(['update:isRegister', 'update:isHasUserInfo']);
type UserInfoI = { username: string, password: string, surePassword: string };

const registerInfo = reactive<UserInfoI>({
  username: "",
  password: "",
  surePassword: "",
});

const errorMessage = ref("");

// setTimeout(() => {
//   emit("update:isRegister", false);
// }, 3000)

watch(registerInfo, () => {
  console.log(registerInfo.surePassword);

})
const handleRegister = async () => {
  // try{} catch{}
  if (registerInfo.password !== registerInfo.surePassword || (registerInfo.password.length < 6 || registerInfo.password.length > 16)) {
    errorMessage.value = "密码不一致或长度不正确";
    return;
  }
  const { data: res } = await registerApi({
    username: registerInfo.username,
    password: registerInfo.password,
  });
  if (res.status !== 200) {
    errorMessage.value = res.message;
    return;
  }
  errorMessage.value = "";
  emit("update:isRegister", false);
  // login
  const { data: res2 } = await loginApi({ username: registerInfo.username, password: registerInfo.password });
  setStorage("userInformation", res2);
  // 判断登录状态
  emit("update:isHasUserInfo", true);
};

const handleCancel = () => {
  emit("update:isRegister", false);
};

</script>
<template>
  <div class="login">
    <!-- 使用slot好一点 -->
    <h3>用户注册</h3>
    <form class="form">
      <div class="up">
        <div>
          <label for="password">用户名:</label>
        </div>
        <div>
          <input v-model="registerInfo.username" type="text" id="username" name="username" placeholder="username">
        </div>
      </div>
      <div class="up">
        <div>
          <label for="password">密码:</label>
        </div>
        <div>
          <input v-model="registerInfo.password" type="password" id="password" name="password"
            placeholder="英文、数字、符号(6-16位)">
        </div>
      </div>
      <!-- 使用slot好一点 -->
      <div class="up">
        <div>
          <label for="password1">确认密码:</label>
        </div>
        <div>
          <input v-model="registerInfo.surePassword" type="password" id="password1" name="password1"
            placeholder="英文、数字、符号(6-16位)">
        </div>
      </div>
    </form>
    <div v-if="errorMessage != ''" class="error">{{ errorMessage }}</div>
    <div class="btn">
      <button @click="handleRegister">注册并登录</button>
      <button @click="handleCancel">取消</button>
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
  width: 80px;
  line-height: 40px;
  text-align: right;
  margin-right: 10px;
}

.up>div:last-child {
  width: 200px;
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
  height: 50px;
}

.btn>button {
  border-radius: 8px;
  border: 1px solid transparent;
  /* padding: 0.3em 1em; */
  padding: 4px;
  font-size: 1em;
  font-weight: 500;
  font-family: inherit;
  cursor: pointer;
  transition: border-color 0.25s;
  outline: none;
}

.btn>button:first-child {
  width: 90px;
  background-color: #546ff7;
}

.btn>button:last-child {
  background-color: #ffd978;
}

.btn>button:hover {
  border-color: #242424;
}

.error {
  height: auto;
  color: red;
  font-size: 14px;
}
</style>