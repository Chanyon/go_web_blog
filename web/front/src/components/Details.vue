<script lang="ts" setup>
import { onBeforeMount, reactive, ref, watchEffect } from "vue";
import { getArticleInfoApi, getCommentListApi, sendCommentApi } from "../api/article";
import type { CategoryI, ArticleI } from "../api/article"
import { formatDate } from "../uitls/formatDate";
const props = defineProps(['id']);
const userComment = ref("");

const articleInfo = ref<ArticleI>({
  Category: {
    name: "",
    id: 0,
  },
  ID: 0,
  cid: 0,
  desc: "",
  title: "",
  comment_count: 0,
  content: "",
  read_count: 0,
  img: "",
});

const commentList = reactive<{ list: Array<any>, total: number }>({
  list: [],
  total: 0,
});
const getArticleInfo = async () => {
  const { data: res } = await getArticleInfoApi(`user/article/info/${props.id}`);
  articleInfo.value = res.data;
};

const getCommentList = async () => {
  const { data: res } = await getCommentListApi(`user/comment/front/${props.id}?pagesize=100&pagenum=1`);
  if (res.status === 200 && res.total > 0) {
    commentList.list = res.data;
    commentList.total = res.total;
  } else {
    commentList.list = [];
  }
};
onBeforeMount(() => {
  getArticleInfo();
  getCommentList();
});

const userInfo = ref({
  data: "",
  id: 0,
});

const errorMessage = ref("");
watchEffect(() => {
  const userInformation = window.sessionStorage.getItem("userInformation");
  if (userInformation != null) {
    userInfo.value = JSON.parse(userInformation);
  }
});

const pushComment = async () => {
  if (userComment.value.length === 0) {
    errorMessage.value = "请输入内容";
    return;
  }
  if (userInfo.value.id === 0 && userInfo.value.data.length === 0) {
    errorMessage.value = "登录后，回复";
    return;
  }
  const { data: res } = await sendCommentApi({
    user_id: userInfo.value.id,
    article_id: articleInfo.value.ID,
    article_title: articleInfo.value.title,
    content: userComment.value,
    user_username: userInfo.value.data,
  });
};
</script>
<template>
  <div class="details animate__animated animate__fadeInLeft">
    <div class="title">
      <h2>{{ articleInfo.title }}</h2>
      <span>{{ formatDate(articleInfo.CreatedAt as string) }}</span>
    </div>
    <div class="article-img">
      <img :src="`http://127.0.0.1:3000${articleInfo.img}`" alt="image">
    </div>
    <div class="content" v-html="articleInfo.content">
    </div>
    <div class="comment">
      <p class="comment-top">
      <h3>回复评论</h3>
      <span>{{ commentList.total }}</span>
      </p>
      <div class="push-comment">
        <textarea v-model="userComment" rows="8" autofocus></textarea>
        <p>
          <button @click="pushComment">点击评论</button>
          <span v-if="errorMessage.length > 0" class="error-msg">{{ errorMessage }}</span>
        </p>
      </div>
      <div v-if="commentList.total === 0">
        <span>匿名:</span>
        <span>无评论...</span>
      </div>
      <div v-for="(comm, idx) in commentList.list" :key="idx">
        <span>{{ comm.user_username }}:</span>
        <span>{{ comm.content }}</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.details {
  width: 880px;
  height: auto;
  padding: 10px;
  /* background-color: #fff; */
  background-color: hsla(210, 30%, 8%, 1);
  color: #fff;
  border: 1px solid #535455;
  border-radius: 8px;
}

.details .title {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 75px;
  border-bottom: 1px solid #ccc;
}

.details .title h2 {
  /* flex: 0.7; */
  padding: 4px;
  font-size: 30px;
}

.details .title span {
  padding: 4px;
  font-size: 14px;
}

.details .article-img {
  height: 300px;
  /* overflow: hidden; */
}

.article-img img {
  width: 100%;
  height: 100%;
  vertical-align: top;
  font-size: 0;
}

.details .content {
  margin: 25px 0;
}

.details .comment {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
  padding-left: 20px;
  margin: 0 15px 0 15px;
}

.details .comment div {
  width: 100%;
  height: auto;
  margin-top: 15px;
  text-align: left;
  border-radius: 8px;
  padding: 5px;
  background-color: #535455;
  box-shadow: 3px 4px 6px rgba(255, 255, 255, 0.8);
  font-size: 18px;
}

.details .comment div>span:first-child {
  color: rgb(0, 179, 255);
}

.details .comment div>span:last-child {
  /* color: rgb(0, 179, 255); */
  display: inline-block;
  margin-left: 10px;
  /* height: 90px; */
  /* overflow: auto; */
  word-wrap: break-word;
}

.details .comment .push-comment {
  display: flex;
  flex-direction: column;
  height: 165px;
}

.comment .comment-top {
  display: flex;
}

.comment .comment-top span {
  margin-left: 5px;
  padding: 0 8px 0 8px;
  background-color: #556ff7;
  color: #fff;
}

.comment .push-comment button {
  width: 65px;
  padding: 3px;
  cursor: pointer;
  outline: none;
}

.comment .push-comment textarea {
  width: 100%;
  padding-left: 7px;
  outline: none;
}

.comment .push-comment .error-msg {
  color: rgb(235, 25, 148);
  font-size: 16px;
  margin-left: 15px;
}
</style>