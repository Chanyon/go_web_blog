<script setup lang="ts">
import { computed, ref, onBeforeMount, reactive } from 'vue';
import type { Ref } from 'vue';
import { message } from 'ant-design-vue';
import axios from 'axios';
import { DataItem, columns } from './category'
import { addCategory, deleteCateById, editCategoryById, getCategoryById, getCategoryInfo } from '../../network/category'

const dataSource: Ref<Array<DataItem>> = ref([]);

const addCateVisible = ref(false);
const editCateVisible = ref(false);

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
  // hideOnSinglePage: true,
  showTotal: (total: number) => `共${total}条`,
}));

const handleTableChange = (pag: { current: number, pageSize: number }, filters: any, sorter: any) => {
  queryParam.value.pagenum = pag.current;
  queryParam.value.pagesize = pag.pageSize;
  page.value.current = pag.current;
  page.value.pageSize = pag.pageSize;
  getCategory();
}

const validateMessages = {
  required: '${label} is required!',
};

const formState = reactive({
  cate: {
    name: '',
  },
});
const editState = reactive({
  cate: {
    name: '',
    id: '',
  },
});

// 网络请求
const getCategory = async () => {
  const { data: ret } = await getCategoryInfo(`admin/category`, {
    pagesize: queryParam.value.pagesize,
    pagenum: queryParam.value.pagenum,
  });
  dataSource.value = ret.data;
  total.value = ret.total;
}
onBeforeMount(() => {
  getCategory();
});

const addCateCancel = () => {
  addCateVisible.value = false;
  formState.cate.name = '';
  message.info('新增分类取消')
}

const addCateOk = async () => {
  if (formState.cate.name === '') {
    return message.warning('请填写分类名');
  } else {
    const { data: ret } = await addCategory(`admin/category/add`, {
      name: formState.cate.name,
    });
    if (ret.status !== 200) {
      return message.error(ret.message);
    }
    addCateVisible.value = false;
    formState.cate.name = "";
    getCategory();
    message.success(ret.message);
  }
}
const onEdit = async (id: string) => {
  editCateVisible.value = true;
  const { data: ret } = await getCategoryById(`http://127.0.0.1:3000/api/v1/admin/category/${id}`);
  editState.cate.name = ret.data.name;
  editState.cate.id = id;
}
const editCateOk = async () => {
  if (editState.cate.name === '') {
    return message.warning('请填写分类名');
  } else {
    const { data: ret } = await editCategoryById(`admin/category/${editState.cate.id}`, {
      name: editState.cate.name,
    });
    if (ret.status !== 200) {
      return message.error(ret.message);
    }
    editCateVisible.value = false;
    message.success('修改成功');
    getCategory();
  }
}
const editCateCancel = () => {
  editCateVisible.value = false;
  editState.cate.name = '';
  message.info('修改分类取消')
}
const onDelete = async (id: string) => {
  const { data: ret } = await deleteCateById(`admin/category/${id}`);
  if (ret.status !== 200) {
    return message.error(ret.message);
  }
  message.success('删除成功');
  getCategory();
};
</script>

<template>
  <a-card>
    <a-row :gutter="20">
      <a-col :span="4">
        <a-button type="primary" size="large" @click="addCateVisible = true" style="margin-bottom: 15px">新增分类</a-button>
      </a-col>
    </a-row>
    <a-table rowKey="id" :columns="columns" :data-source="dataSource" bordered :pagination="pagination"
      @change="handleTableChange">
      <template #bodyCell="{ column, record }">
        <template v-if="column.dataIndex === 'operation'">
          <a-popconfirm v-if="dataSource.length" title="是否删除?不可恢复！" ok-text="是" cancel-text="否"
            @confirm="onDelete(record.id)">
            <a-button type="danger">Delete</a-button>
          </a-popconfirm>
          <a-button type="primary" danger style="margin-left: 10px;" @click="onEdit(record.id)">Edit</a-button>
        </template>
      </template>
    </a-table>
  </a-card>
  <!-- add -->
  <a-modal closable title="新增分类" :visible="addCateVisible" width="60%" @cancel="addCateCancel" @ok="addCateOk"
    destroyOnClose>
    <a-form :model="formState" name="" :validate-messages="validateMessages">
      <a-form-item :name="['cate', 'name']" label="分类名称" :rules="[{ required: true }]">
        <a-input v-model:value="formState.cate.name" />
      </a-form-item>
    </a-form>
  </a-modal>
  <!-- 编辑 -->
  <a-modal closable title="编辑分类" :visible="editCateVisible" width="60%" @cancel="editCateCancel" @ok="editCateOk"
    destroyOnClose>
    <a-form :model="editState" name="" :validate-messages="validateMessages">
      <a-form-item :name="['cate', 'name']" label="分类名称" :rules="[{ required: true }]">
        <a-input v-model:value="editState.cate.name" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<style scoped>

</style>
