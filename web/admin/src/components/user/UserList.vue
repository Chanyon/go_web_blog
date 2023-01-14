<script setup lang="ts">
import { message } from 'ant-design-vue';
import axios from 'axios';
import { computed, onBeforeMount, reactive, ref, watch } from 'vue';
import { getSearchData, UpdateUserPass,updateUserInfo,
getUserList,getUserInfoById,deleteUserInfoById,addUserInfo as addUser } from '../../network/user';
import { columns, UserInfo} from './userList';

const searchParam = reactive({
  name: '',
})

const dataSource = ref([]);
const addUserVisible = ref(false);
const editUserVisible = ref(false);
const checkUserPassVisible = ref(false);
const IsAdmin = ref(false);

const queryParam = ref({
  pagesize: 5,
  pagenum: 1,
});
const page = ref({
  current: 1,
  pageSize: 5,
});
const total = ref(20);
const pagination = computed(() => ({
  total: total.value,
  current: page.value.current,
  pageSize: page.value.pageSize,
  showSizeChanger: true,
  showTotal: (total: number) => `共${total}条`,
}));
const handleTableChange = (pag: { current: number, pageSize: number }, filters: any, sorter: any) => {
  queryParam.value.pagenum = pag.current;
  queryParam.value.pagesize = pag.pageSize;
  page.value.current = pag.current;
  page.value.pageSize = pag.pageSize;
  getUserInfo();
}

const onSearch = async () => {
  const { data: ret } = await getSearchData("admin/users", {
    username: searchParam.name,
    pagesize: page.value.pageSize,
    pagenum: page.value.current,
  });
  dataSource.value = ret.data;
  total.value = ret.total;
}

// 
const getUserInfo = async () => {
  try {
    const { data: ret } = await getUserList(`admin/users`, {
      params: {
        username: searchParam.name,
        pagesize: page.value.pageSize,
        pagenum: page.value.current,
      }
    });
    dataSource.value = ret.data;
    total.value = ret.total;
  } catch (error) {
    message.error("获取数据失败");
  }
}

const validateMessages = {
  required: '${label} is required!',
};

const addUserInfo = reactive<UserInfo>({
  username: '',
  password: '',
  checkpass: '',
});
const editUserInfo = reactive<Pick<UserInfo,'username' | 'role' | 'id'|'password'|'checkpass'>>({
  id: undefined,
  username: '',
  role: 2,
  password: '',
  checkpass: '',
});
const addUserInfoOk = async () => {
  try {
    if (addUserInfo.password === "" || addUserInfo.password !== addUserInfo.checkpass) {
      return message.info("请检查密码");
    } else {
      const { data: ret } = await addUser(`admin/users`, addUserInfo);
      if (ret.status !== 200) {
        return message.info("请更换用户名或检查密码长度");
      }
      message.success("添加成功");
      addUserVisible.value = false;
      getUserInfo();
    }
  } catch (error) {

  }
}
const userInfoCancel = () => {
  addUserVisible.value = false;
  addUserInfo.username = "";
  addUserInfo.password = "";
  addUserInfo.checkpass = "";
}

const onDelete = async (id: number) => {
  try {
    const { data: ret } = await deleteUserInfoById(`admin/user/${id}`);
    if (ret.status !== 200) { return message.error("删除失败"); }
    message.success("删除成功");
    getUserInfo();
  } catch (error) {
    message.error("发生错误");
  }
}

const editUserInfoHandler = async (id: number) => {
  editUserVisible.value = true;
  editUserInfo.id = id;
  const { data: ret } = await getUserInfoById(`admin/user/${id}`);
  // 通过id拿到数据
  editUserInfo.username = ret.data.username;
  if(ret.data.role === 1) {
    IsAdmin.value = true;
  }else{
    IsAdmin.value = false;
  }
}

const setAdminFn = (checked: boolean) => {
console.log(checked);

  if(checked) {
    editUserInfo.role = 1;
  }else {
    editUserInfo.role = 2;
  }
}

const editUserInfoCancel = () => {
  editUserVisible.value = false;
  editUserInfo.username = "";
  editUserInfo.role = 2;
}

const editUserInfoOk = async () => {
  const { data: ret } = await updateUserInfo(`admin/user/${editUserInfo.id}`,{
    username: editUserInfo.username,
    role: editUserInfo.role,
  });
  if(ret.status !== 200) {
    return message.error(ret.message);
  }
  message.success("更新成功");
  editUserInfoCancel();
  getUserInfo();
}

// 更新密码
const updateUserPass = reactive<Pick<UserInfo,'password' | 'checkpass'|'id'>>({
  id: undefined,
  password: '',
  checkpass: '',
})

const updateUserPassHandler = (id:number) => {
  checkUserPassVisible.value = true;
  updateUserPass.id = id;
}

const updateUserPassCancel = () => {
  checkUserPassVisible.value = false;
  updateUserPass.checkpass = "";
  updateUserPass.password = "";
}

const updateUserPassOk = async () => {
  try {
  // 密码验证，todo!
    const { data: ret } = await UpdateUserPass(`admin/user/${updateUserPass.id}}`,{
      password: updateUserPass.password,
      checkpass: updateUserPass.checkpass,
    });
    if(ret.status !== 200) {
      return message.error(ret.message);
    }
    message.success(ret.message);
    updateUserPassCancel()
  } catch (error) {
    message.error("更新密码失败");
  }
}

onBeforeMount(() => {
  getUserInfo();
})
</script>

<template>
  <a-card>
    <a-row :gutter="20">
      <a-col :span="6">
        <a-input-search v-model:value="searchParam.name" placeholder="input search text" enter-button
          @search="onSearch" />
      </a-col>
      <a-col :span="4">
        <a-button type="primary" @click="addUserVisible = true">新增</a-button>
      </a-col>
    </a-row>
    <a-row>
      <a-col :span="24">
        <a-table rowKey="ID" :columns="columns" :dataSource="dataSource" bordered :pagination="pagination"
          @change="handleTableChange">
          <template #bodyCell="{ column, record }">
            <template v-if="column.dataIndex === 'role'">
              <span>{{ record.role === 1 ? '管理员' : '订阅者' }}</span>
            </template>
            <template v-if="column.dataIndex === 'operation'">
              <a-button type="primary" danger style="margin-right: 10px;" @click="editUserInfoHandler(record.ID)">编辑</a-button>
              <a-button type="primary" danger style="margin-right: 10px;" @click="updateUserPassHandler">改密</a-button>
              <a-popconfirm title="是否删除?不可恢复！" ok-text="是" cancel-text="否" @confirm="onDelete(record.ID)">
                <a-button type="danger">删除</a-button>
              </a-popconfirm>
            </template>
          </template>
        </a-table>
      </a-col>
    </a-row>
  </a-card>
  <!-- 新增用户 -->
  <a-modal closeable :visible="addUserVisible" title="新增用户" width="40%" @ok="addUserInfoOk" @cancel="userInfoCancel"
    destroyOnClose>
    <a-form :model="addUserInfo" :validate-messages="validateMessages" :labelCol="{ span: 6 }"
      :wrapperCol="{ span: 12 }">
      <a-form-item :name="['username']" label="用户名" :rules="[{ required: true }]">
        <a-input v-model:value="addUserInfo.username" placeholder="用户名长度5-8位" />
      </a-form-item>
      <a-form-item :name="['password']" label="密码" :rules="[{ required: true }]">
        <a-input v-model:value="addUserInfo.password" placeholder="密码长度6-16位" />
      </a-form-item>
      <a-form-item :name="['checkpass']" label="确认密码" :rules="[{ required: true }]">
        <a-input v-model:value="addUserInfo.checkpass" placeholder="密码长度6-16位" />
      </a-form-item>
    </a-form>
  </a-modal>
  <!-- 编辑用户 -->
  <a-modal closeable :visible="editUserVisible" title="编辑用户" width="40%" @ok="editUserInfoOk" @cancel=" editUserInfoCancel
" destroyOnClose>
    <a-form :model="editUserInfo" :validate-messages="validateMessages" :labelCol="{ span: 6 }"
      :wrapperCol="{ span: 12 }">
      <a-form-item :name="['username']" label="用户名" :rules="[{ required: true }]">
        <a-input v-model:value="editUserInfo.username" placeholder="用户名长度5-8位" />
      </a-form-item>
      <a-form-item :name="['role']" label="设为管理员">
        <a-switch v-model:checked="IsAdmin" checked-children="是" un-checked-children="否" @change="setAdminFn">
        </a-switch>
      </a-form-item>
    </a-form>
  </a-modal>
  <!-- 更新密码 -->
    <a-modal closeable :visible="checkUserPassVisible" title="修改密码" width="40%" @ok="updateUserPassOk" @cancel=" updateUserPassCancel
" destroyOnClose>
    <a-form :model="updateUserPass" :validate-messages="validateMessages" :labelCol="{ span: 6 }"
      :wrapperCol="{ span: 12 }">
      <a-form-item :name="['password']" label="密码" :rules="[{ required: true }]">
        <a-input v-model:value="updateUserPass.password" placeholder="密码长度6-16位" />
      </a-form-item>
      <a-form-item :name="['checkpass']" label="确认密码" :rules="[{ required: true }]">
        <a-input v-model:value="updateUserPass.checkpass" placeholder="密码长度6-16位" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<style scoped>
</style>
