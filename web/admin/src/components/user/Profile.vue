<script setup lang="ts">
import { message, UploadChangeParam } from 'ant-design-vue';
import { onBeforeMount, reactive, ref } from 'vue';
import { getProfile, uploadProfileById, uploadProfileInfo } from '../../network/user'

const profileInfo = reactive({
  id: 0,
  name: "",
  desc: "",
  icp_record: "",
  qq_chat: "",
  wei_chat: "",
  wei_bo: "",
  bili: "",
  email: "",
  img: "", //background
  avatar: "",
});

// 文件上传
const fileList = ref([]);
const handleChange = (info: UploadChangeParam) => {
  if (info.file.status !== 'uploading') {
    console.log(info.file, info.fileList);
    if (info.fileList.length === 1) {
      profileInfo.img = info.fileList[0]?.response.url;
    } else {
      profileInfo.avatar = info.fileList[0]?.response.url;
      profileInfo.img = info.fileList[1]?.response.url;
    }
  }
  if (info.file.status === 'done') {
    message.success(`${info.file.name} file uploaded successfully`);
  } else if (info.file.status === 'error') {
    message.error(`${info.file.name} file upload failed.`);
  }
};

const getProfileInfo = async () => {
  try {
    const { data: ret } = await getProfile(`admin/profile/1`);
    if (ret.status !== 200) {
      message.error("获取数据失败");
    }
    message.success("获取数据成功");
    const profileData = ret.data;
    profileInfo.id = profileData.id;
    profileInfo.name = profileData.name;
    profileInfo.email = profileData.email;
    profileInfo.img = profileData.img;
    profileInfo.bili = profileData.bili;
    profileInfo.desc = profileData.desc;
    profileInfo.avatar = profileData.avatar;
    profileInfo.qq_chat = profileData.qq_chat;
    profileInfo.wei_bo = profileData.wei_bo;
    profileInfo.wei_chat = profileData.wei_chat;
    profileInfo.icp_record = profileData.icp_record;
  } catch (error: any) {
    message.error(error?.message);
  }
}

onBeforeMount(() => {
  getProfileInfo();
})

const uploadProfile = async () => {
  try {
    if (profileInfo.id === 0) {
      const { data: ret } = await uploadProfileInfo(`admin/profile/`, profileInfo);
      if (ret.status !== 200) {
        message.error("上传失败");
      }
      message.success("更新成功");
    } else {
      const { data: ret } = await uploadProfileById(`admin/profile/${profileInfo.id}`, profileInfo);
      if (ret.status !== 200) {
        message.error("更新失败");
      }
      message.success("更新成功");
    }
  } catch (error) {
    if (error) {
      message.info("发生错误");
    }
  }
}

const headers = () => {
  const token = window.sessionStorage.getItem("token");
  const Authorization = `Bearer ${token}`;
  return {
    Authorization,
  };
};

</script>

<template>
  <a-card>
    <h2>个人资料设置</h2>
    <a-form :model="profileInfo" labelAlign="left" :labelCol="{ span: 2 }" :wrapperCol="{ span: 12 }">
      <a-form-item label="昵称" :name="['name']">
        <a-input style="width: 300px" v-model:value="profileInfo.name"></a-input>
      </a-form-item>
      <a-form-item label="个人简介" :name="['desc']">
        <a-textarea style="width: 500px" v-model:value="profileInfo.desc" :auto-size="{ minRows: 2, maxRows: 5 }">
        </a-textarea>
      </a-form-item>
      <a-form-item label="网站备案号" :name="['icp_record']">
        <a-input style="width: 300px" v-model:value="profileInfo.icp_record"></a-input>
      </a-form-item>
      <a-form-item label="QQ" :name="['qq_chat']">
        <a-input style="width: 300px" v-model:value="profileInfo.qq_chat"></a-input>
      </a-form-item>
      <a-form-item label="微信" :name="['wei_chat']">
        <a-input style="width: 300px" v-model:value="profileInfo.wei_chat"></a-input>
      </a-form-item>
      <a-form-item label="微博" :name="['wei_bo']">
        <a-input style="width: 300px" v-model:value="profileInfo.wei_bo"></a-input>
      </a-form-item>
      <a-form-item label="B站" :name="['bili']">
        <a-input style="width: 300px" v-model:value="profileInfo.bili"></a-input>
      </a-form-item>
      <a-form-item label="Email" :name="['email']">
        <a-input style="width: 300px" v-model:value="profileInfo.email"></a-input>
      </a-form-item>
      <a-form-item label="头像" :name="['avatar']">
        <a-upload v-model:file-list="fileList" name="file" :headers="headers()"
          action="http://127.0.0.1:3000/api/v1/admin/upload" @change="handleChange">
          <a-button>
            <upload-outlined></upload-outlined>
            Click to Upload
          </a-button>
        </a-upload>
        <div v-if="profileInfo.avatar">
          <img :src="`http://127.0.0.1:3000${profileInfo.avatar}`"
            style="width: 126px; height: 100px;margin-top: 5px;" />
        </div>
      </a-form-item>
      <a-form-item label="头像背景图" :name="['img']">
        <a-upload v-model:file-list="fileList" name="file" :headers="headers()"
          action="http://127.0.0.1:3000/api/v1/admin/upload" @change="handleChange">
          <a-button>
            <upload-outlined></upload-outlined>
            Click to Upload
          </a-button>
        </a-upload>
        <div v-if="profileInfo.img">
          <img :src="`http://127.0.0.1:3000${profileInfo.img}`" style="width: 126px; height: 100px;margin-top: 5px;" />
        </div>
      </a-form-item>
      <a-form-item label="上传更新">
        <a-button type="primary" @click="uploadProfile">更新</a-button>
      </a-form-item>
    </a-form>
  </a-card>
</template>

<style scoped>

</style>
