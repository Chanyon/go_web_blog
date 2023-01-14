<script setup lang="ts">
import { message } from "ant-design-vue";
import { computed, onBeforeMount, ref } from "vue";
import { getComment, updateComment, deleteComment as deleteCom } from "../../network/commentList";
import { columns } from "./comment";

const dataSource = ref([]);
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
  getCommentList();
}

const getCommentList = async () => {
  const { data: ret } = await getComment(`admin/comment/list`,{
    pagesize: queryParam.value.pagesize,
    pagenum: queryParam.value.pagenum,
  });
  if(ret.status !== 200) {
    return message.error("获取数据失败");
  }
  dataSource.value = ret.data;
  total.value = ret.total;
  message.success("获取数据ok");
}

const commentCheck = async (record: any) => {
  if( 1 === record.status ) {
    return message.info("已审核通过");
  }
  const {data: ret} = await updateComment(`admin/check/comment/${record.ID}`,{
    status: 1,
  });
  if(ret.status !== 200) {
    message.error("通过失败");
  }
  getCommentList();
}

const commentUnCheck = async (record: any) => {
  if( 2 === record.status ) {
    return message.info("评论已撤下!");
  }
  const {data: ret} = await updateComment(`admin/uncheck/comment/${record.ID}`,{
    status: 2,
  });
  if(ret.status !== 200) {
    message.error("通过失败");
  }
  getCommentList();
}

const deleteComment = async (id: number) => {
  const {data: ret} = await deleteCom(`admin/delete/comment/${id}`);
  if (ret.status !== 200) {
    message.error("删除失败");
  }
  getCommentList();
}

onBeforeMount(() => {
  getCommentList();
})
</script>

<template>
  <a-card>
    <a-row>
      <a-col :span="24">
        <a-table rowKey="ID" :columns="columns" :dataSource="dataSource" bordered :pagination="pagination"
          @change="handleTableChange">
          <template #bodyCell="{ column, record }">
            <template v-if="column.dataIndex === 'status'">
              <span>{{ record.status === 1 ? '审核通过' : '未审核' }}</span>
            </template>
            <template v-if="column.dataIndex === 'operation'">
              <a-button type="primary" danger style="margin-right: 10px;" @click="commentCheck(record)">通过</a-button>
              <a-button type="primary" danger style="margin-right: 10px;" @click="commentUnCheck(record)">撤下</a-button>
              <a-popconfirm title="是否删除?不可恢复！" ok-text="是" cancel-text="否" @confirm="deleteComment(record.ID)">
                <a-button type="danger">删除</a-button>
              </a-popconfirm>
            </template>
          </template>
        </a-table>
      </a-col>
    </a-row>
  </a-card>
</template>

<style scoped>
</style>
