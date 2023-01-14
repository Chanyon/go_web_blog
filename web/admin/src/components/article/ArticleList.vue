<script setup lang="ts">
import { computed, ref, onBeforeMount, reactive } from 'vue';
import type { Ref } from 'vue';
import { useRouter, useRoute } from 'vue-router'
import { message } from 'ant-design-vue';
import { CategoryItem, ArticleItem, columns } from './addArticle';
import { deleteArticleById, getArticleByCid, getArticleByPage, getCategoryInfo } from '../../network/article';

const router = useRouter();


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
  getArticleList();
}

const searchParams = reactive({
  title: '',
});

let categoryList = ref<CategoryItem[]>([{
  id: 0,
  name: "默认",
}]);
const getCategory = async () => {
  const { data: ret } = await getCategoryInfo(`admin/category`);
  if (ret.status !== 200) {
    return message.error(ret.message);
  }
  categoryList.value = ret.data.map((data: { id: number; name: string; }) => ({ id: data.id, name: data.name }));
};

const articleList: Ref<ArticleItem[]> = ref([]);
const getArticleList = async () => {
  const { data: ret } = await getArticleByPage(`admin/article`, {
    pagesize: queryParam.value.pagesize,
    pagenum: queryParam.value.pagenum,
    title: searchParams.title
  });
  if (ret.status !== 200) {
    return message.error(ret.message);
  }
  articleList.value = ret.data;
  total.value = ret.total;
}

//数据请求 
onBeforeMount(() => {
  getCategory()
  getArticleList();
});
const onSearch = () => {
  getArticleList();
}
const onShowAll = () => {
  if (searchParams.title !== "") {
    searchParams.title = ""
  }
  getArticleList();
}

const onEdit = (id: number) => {
  router.push(`/addart/${id}`);
}

const onDelete = async (id: number) => {
  try {
    const { data: ret } = await deleteArticleById(`admin/article/${id}`);
    if (ret.status !== 200) {
      return message.error("删除失败!");
    }
    message.success("删除成功");
    getArticleList();
  } catch (error) {
  }
}
const CateChange = async (cid: number) => {
  try {
    const { data: ret } = await getArticleByCid(`admin/article/category/${cid}`, {
      params: {
        pagesize: queryParam.value.pagesize,
        pagenum: queryParam.value.pagenum,
        title: searchParams.title
      }
    });
    if (ret.status !== 200) {
      return message.error("获取失败!");
    }
    // message.success("获取成功");
    articleList.value = ret.data;
    console.log(articleList.value);

    total.value = ret.total;
  } catch (error) {
  }
}
</script>

<template>
  <a-card>
    <a-row :gutter="20">
      <a-col :span="6">
        <a-input-search v-model:value="searchParams.title" placeholder="input search text" enter-button
          @search="onSearch" />
      </a-col>
      <a-col :span="2">
        <a-button type="primary" @click="router.push('/addart')">新增</a-button>
      </a-col>
      <a-col :span="3">
        <a-select placeholder="请选择分类" @change="CateChange" style="width:200px;">
          <a-select-option v-for="cate in categoryList" :value="cate.id">
            {{ cate.name }}
          </a-select-option>
        </a-select>
      </a-col>
      <a-col :span="1" style="margin-left:35px;">
        <a-button type="info" @click="onShowAll">显示全部</a-button>
      </a-col>
    </a-row>
    <a-table rowKey="ID" :columns="columns" :data-source="articleList" bordered :pagination="pagination"
      @change="handleTableChange" style="margin-top:15px" :scroll="{ x: 1500, y: 1000 }">
      <template #bodyCell="{ column, record }">
        <template v-if="column.dataIndex === 'operation'">
          <a-button type="primary" danger style="margin-right: 10px;" @click="onEdit(record.ID)">编辑</a-button>
          <a-popconfirm v-if="articleList.length" title="是否删除?不可恢复！" ok-text="是" cancel-text="否"
            @confirm="onDelete(record.ID)">
            <a-button type="danger">删除</a-button>
          </a-popconfirm>
        </template>
        <template v-if="column.dataIndex === 'name'">
          <span>{{record.Category.name}}</span>
        </template>
        <template v-if="column.dataIndex === 'img'">
          <img :src="`http://localhost:3000${record.img}`" width="50" height="50" />
        </template>
      </template>
    </a-table>
  </a-card>
</template>

<style scoped>

</style>
