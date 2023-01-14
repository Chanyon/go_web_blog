<script lang="ts" setup>

</script>
<template>
  <div class="footer">
    <div class="left"></div>
    <div class="center">
      <p class="center-info">
        <a href="https://github.com/Chanyon" target="_blank">
          <span>My github</span>
          <img src="../assets/github.svg" alt="github">
        </a>
      </p>
      <p>by https://gitee.com/wejectchan/ginblog</p>
      <p>©2022 Example. All rights reserved.</p>
    </div>
    <div class="right"></div>
  </div>
</template>
  
<style scoped>
.footer {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
  height: 160px;
}

.center .center-info>a {
  display: flex;
  justify-content: center;
  align-items: center;
  color: #617bff;
}

.center .center-info>a:active {
  color: #fff;
}

.center .center-info>a span {
  margin-right: 10px;
}
</style>