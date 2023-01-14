<script setup lang="ts">
import { onBeforeMount, reactive, ref } from "vue";
import { getArticleListApi } from "../api/search";
import router from "../router";
import { formatDate } from "../uitls/formatDate"
import Paginate from "./common/Paginate.vue";

const props = defineProps(['title']);

const total = ref(0);
const isLoad = ref(false);
const articleList = reactive<any>({
  list: [],
});

const pageCount = ref(2);
const pageInfo = reactive({
  title: props.title,
  pagenum: 1,
  pagesize: 5,
});

const getArticleList = async () => {
  const { data: res } = await getArticleListApi(`/user/article/?title=${pageInfo.title}&pagenum=${pageInfo.pagenum}&pagesize=${pageInfo.pagesize}`);
  if (res.total === 0) {
    isLoad.value = true;
  }
  total.value = res.total;
  pageCount.value = res.total / pageInfo.pagesize;
  articleList.list = res.data;
};

onBeforeMount(() => {
  getArticleList();
});

const clickHandler = (currentPage: number) => {
  // console.log(page);
  pageInfo.pagenum = currentPage;
  getArticleList();
};

const handleToDetails = (id: number) => {
  router.push(`/article/details/${id}`);
}

</script>
<template>
  <div class="search-list">
    <div v-if="total === 0 && isLoad" class="search-tip">
      <div class="search-tip__content">
        抱歉，你搜索的文章标题不存在！
      </div>
    </div>
    <div class="article-card">
      <div v-for="(item,idx) in articleList.list" :key="idx" class="article-card__item">
        <div class="article-card__image">
          <img src="" alt="">
        </div>
        <div class="article-card__desc">
          <div class="top-title">
            <span>{{item.Category.name}}</span>
            <h5 @click="handleToDetails(item.ID)">{{item.title}}</h5>
          </div>
          <div class="center-content">
            {{item.desc}}
          </div>
          <div class="bottom-date">
            create:{{formatDate(item.CreatedAt)}} - update:{{formatDate(item.UpdatedAt)}}
            <!-- <span>read:{{item.read_count}}</span> -->
          </div>
        </div>
      </div>
    </div>
    <div class="paginate">
      <Paginate :clickHandler="clickHandler" :pageCount="pageCount" :pageRange="3" :marginPages="2" :hidePrevNext="true"
        :firstLastButton="false" pageLinkClass="page-link" prevLinkClass="prev-link" nextLinkClass="next-link">
      </Paginate>
    </div>
  </div>
</template>
  
<style scoped>
.search-list {
  /* margin-left: 400px; */
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  width: 880px;
  height: auto;
  padding: 10px;
  /* background-color: #fff; */
  background-color: hsla(210, 30%, 8%, 1);
  color: #fff;
  border: 1px solid #535455;
  border-radius: 8px;
}

.search-list>.search-tip {
  width: 550px;
  height: 400px;
  padding: 15px;
  background-color: #fff;
  box-shadow: 0px 3px 8px rgba(0, 0, 0, 0.24);
}

.article-card__item {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 840px;
  height: 140px;
  margin-bottom: 10px;
  border-radius: 8px;
  border: 1px solid #535455;
  /* box-shadow: rgba(9, 30, 66, 0.25) 0px 4px 8px -2px, rgba(9, 30, 66, 0.08) 0px 0px 0px 1px; */
  /* box-shadow: rgba(240, 46, 170, 0.4) -5px 5px, rgba(240, 46, 170, 0.3) -10px 10px, rgba(240, 46, 170, 0.2) -15px 15px, rgba(240, 46, 170, 0.1) -20px 20px, rgba(240, 46, 170, 0.05) -25px 25px; */
  /* box-shadow: rgb(38, 57, 77) 0px 20px 30px -10px; */
  box-shadow: rgba(3, 102, 214, 0.3) 0px 0px 0px 3px;
  /* box-shadow: rgba(67, 71, 85, 0.27) 0px 0px 0.25em, rgba(90, 125, 188, 0.05) 0px 0.25em 1em; */
}

.article-card__item>.article-card__image {
  width: 128px;
  height: 128px;
  padding: 10px;
}

.article-card__item .article-card__desc {
  width: calc(100% - 128px);
  height: 100%;
  padding: 10px;
}

.article-card__desc>.top-title {
  display: flex;
  height: 30px;
  color: #fff;
}

.article-card__desc>.top-title span {
  display: inline-block;
  width: 80px;
  margin-right: 15px;
  background-color: rgb(255, 116, 106);
  text-align: center;
  line-height: 30px;
  border-radius: 4px;
}

.article-card__desc>.top-title>h5 {
  color: hsla(230, 100%, 69%, 1);
  font-size: 18px;
  text-align: left;
  line-height: 30px;
  text-overflow: ellipsis;
}

.article-card__desc>.top-title>h5:hover {
  color: #546ff7;
  cursor: pointer;
  border-bottom: 2px solid #546ff7;
}

.article-card__desc>.center-content {
  display: flex;
  justify-content: flex-start;
  /* align-items: center; */
  height: 60px;
  margin-top: 5px;
  /* padding-left: 15px; */
  border-bottom: 1px solid #ccc;
}

.article-card__desc .bottom-date {
  padding-left: 5px;
  padding-top: 4px;
  text-align: left;
  color: #535455;
  font-size: 14px;
}

.paginate {
  width: 840px;
  margin-top: 10px;
}
</style>