<script setup lang="ts">
import { message } from "ant-design-vue";
import { reactive, watch } from "vue"
import {Login} from "../network/login"
import { useRouter } from 'vue-router'
const router = useRouter();

const loginForm = reactive({
  password: "",
  username: ""
})
const validateMessages = {
  required: '${label} is required!',
}

const loginHandler = async () =>{
  try {
    const {data: ret} = await Login(`admin/login`,{
      username: loginForm.username,
      password: loginForm.password,
    })
    if(ret.status !== 200) {
      message.error(ret.message);
    }
    window.sessionStorage.setItem('token',ret.token);
    router.push('/');
  } catch (error:any) {
    console.log(error);
  }
}

const resetForm = () => {
  loginForm.username = "";
  loginForm.password = "";
}
</script>

<template>
  <div class="container">
    <div class="login-box">
      <a-form class="form" :model="loginForm" :validate-messages="validateMessages" :labelCol="{ span: 6 }"
        :wrapperCol="{ span: 12 }">
        <a-form-item :name="['username']" label="用户名" :rules="[{ required: true }]">
          <a-input v-model:value="loginForm.username" placeholder="用户名长度5-8位" />
        </a-form-item>
        <a-form-item :name="['password']" label="密码" :rules="[{ required: true }]">
          <a-input v-model:value="loginForm.password" placeholder="密码长度6-16位" @keyup.enter="loginHandler"/>
        </a-form-item>
        <a-form-item class="login-btn">
          <a-button type="primary" @click="loginHandler">登录</a-button>
          <a-button type="info" @click="resetForm">取消</a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100vh;
  background-color: #282c34;
}

.login-box {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 450px;
  height: 300px;
  background-color: #fff;
}

.login-btn {
  display: flex;
  justify-content: center;
}
.form {
  width: 100%;
}
</style>
