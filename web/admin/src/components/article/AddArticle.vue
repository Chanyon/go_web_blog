<script setup lang="ts">
import { onBeforeUnmount, onMounted, reactive, ref, shallowRef, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router'
import { message, UploadChangeParam } from 'ant-design-vue';
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import { IEditorConfig } from '@wangeditor/editor'
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import {
  UploadOutlined
} from '@ant-design/icons-vue';
import { CategoryItem } from './addArticle'
import { getCategoryInfo, getArticleInfoById as getArticleInfo, uploadArticleById, addArticleInfo } from '../../network/article';

const router = useRouter();
const route = useRoute();
type InsertFnType = (url: string, alt: string, href: string) => void

const id = ref(Number(route.params.id));

let articleInfo = reactive({
  id: 0,
  title: "",
  desc: "",
  cid: 1001,
  img: "",
  content: "",
});

let categoryList = ref<CategoryItem[]>([{
  id: 1001,
  name: "JavaScript",
}]);
const getCategory = async () => {
  const { data: ret } = await getCategoryInfo(`admin/category`);
  if (ret.status !== 200) {
    return message.error(ret.message);
  }
  categoryList.value = ret.data.map((data: { id: number; name: string; }) => ({ id: data.id, name: data.name }));
};
const cateChange = (value: number) => {
  articleInfo.cid = value;
};

const getArticleInfoById = async (id: number) => {
  if (id >= 0) {
    const { data: ret } = await getArticleInfo(`admin/article/info/${id}`);
    if (ret.status !== 200) {
      message.error("获取数据失败!");
    }
    const { ID, cid, desc, img, title, content } = ret.data;
    articleInfo.cid = cid;
    articleInfo.id = ID;
    articleInfo.title = title;
    articleInfo.content = content;
    articleInfo.img = img;
    articleInfo.desc = desc;
  }
}

//数据请求 
onMounted(() => {
  getCategory();
  getArticleInfoById(id.value);
})

// 文件上传

const headers = {
  Authorization: `Bearer ${sessionStorage.getItem('token')}`,
}

const fileList = ref([]);
const handleChange = (info: UploadChangeParam) => {
  if (info.file.status !== 'uploading') {
    // console.log(info.file, info.fileList);
    articleInfo.img = info.file.response.url;
  }
  if (info.file.status === 'done') {
    message.success(`${info.file.name} file uploaded successfully`);
  } else if (info.file.status === 'error') {
    message.error(`${info.file.name} file upload failed.`);
  }
};

// 编辑器实例，必须用 shallowRef
const editorRef = shallowRef()
const toolbarConfig = {}
// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()
})

const handleCreated = (editor: any) => {
  editorRef.value = editor // 记录 editor 实例，重要！
}

// 初始化 MENU_CONF 属性,自定义上传图片
const editorConfig: Partial<IEditorConfig> = {
  MENU_CONF: {
    uploadImage: {}
  },
  placeholder: '请输入内容...'
}
editorConfig.MENU_CONF!.uploadImage = {
  server: 'http://127.0.0.1:3000/api/v1/admin/upload',
  // form-data fieldName ，默认值 'wangeditor-uploaded-image'
  fieldName: 'file',

  // 单个文件的最大体积限制，默认为 2M
  maxFileSize: 8 * 1024 * 1024, // 1M

  // 最多可上传几个文件，默认为 100
  maxNumberOfFiles: 10,

  // 选择文件时的类型限制，默认为 ['image/*'] 。如不想限制，则设置为 []
  allowedFileTypes: ['image/*'],

  // 自定义上传参数，例如传递验证的 token 等。参数会被添加到 formData 中，一起上传到服务端。
  // meta: {
  // token: sessionStorage.getItem('token'),
  // otherKey: 'yyy'
  // },

  // 将 meta 拼接到 url 参数中，默认 false
  metaWithUrl: false,

  // 自定义增加 http  header
  headers: {
    Accept: '*/*',
    Authorization: `Bearer ${sessionStorage.getItem('token')}`,
  },

  // 跨域是否传递 cookie ，默认为 false
  withCredentials: false,

  // 超时时间，默认为 10 秒
  timeout: 10 * 1000, // 5 秒
  // 自定义插入图片
  customInsert(res: any, insertFn: InsertFnType) {  // TS 语法
    // res 即服务端的返回结果
    const alt = `图片...`;
    const url = `http://127.0.0.1:3000${res.url}`;
    // 从 res 中找到 url alt href ，然后插图图片
    insertFn(url, alt, "");
  },
  // 单个文件上传成功之后
  onSuccess(file: File, res: any) {
    message.info(`${file.name} 上传成功...`);
  },
  // 单个文件上传失败
  onFailed(file: File, res: any) {
    message.info(`${file.name} 上传失败...`);
  },
}

// 上传文章
const addArticle = async () => {
  if (articleInfo.title === "" && articleInfo.img === "" && articleInfo.desc === "") {
    return message.info(`请添加文章标题`);
  } else {
    try {
      if (0 <= id.value) {
        const { data: ret } = await uploadArticleById(`admin/article/${id.value}`, articleInfo);
        if (ret.status !== 200) {
          message.error(`文章上传失败...`);
        } else {
          message.info("文章跟新成功...");
          router.push("/artlist");
        }
      } else {
        const { data: ret } = await addArticleInfo(`admin/article/add`, articleInfo);
        if (ret.status !== 200) {
          console.log(ret);
          message.error(`文章上传失败...`);
        } else {
          message.info("文章上传成功...");
          router.push("/artlist");
        }
      }
    } catch (error) {

    }
  }
}
const addCancel = () => {
  articleInfo.title = "";
  articleInfo.desc = "";
  articleInfo.id = 0;
  articleInfo.img = "";
  articleInfo.content = "";
  router.push("/artlist");
}
</script>

<template>
  <a-card>
    <h3>{{ id ? "编辑文章" : "新增文章" }}</h3>
    <a-form :model="articleInfo">
      <a-row :gutter="24">
        <a-col :span="16">
          <a-form-item label="文章标题" :name="['title']">
            <a-input type="text" v-model:value="articleInfo.title"></a-input>
          </a-form-item>
          <a-form-item label="文章描述" :name="['desc']">
            <a-textarea v-model:value="articleInfo.desc" placeholder="desc..."
              :auto-size="{ minRows: 2, maxRows: 5 }" />
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item label="文章分类" :name="['cid']">
            <a-select v-model:value="articleInfo.cid" placeholder="请选择分类" @change="cateChange">
              <a-select-option v-for="cate in categoryList" :value="cate.id">
                {{ cate.name }}
              </a-select-option>
            </a-select>
          </a-form-item>
          <a-form-item label="文章缩略图" :name="['img']">
            <a-upload v-model:file-list="fileList" name="file" action="http://127.0.0.1:3000/api/v1/admin/upload"
              @change="handleChange" :headers="headers">
              <a-button>
                <upload-outlined></upload-outlined>
                Click to Upload
              </a-button>
            </a-upload>
            <div v-if="id">
              <img :src="`http://127.0.0.1:3000${articleInfo.img}`"
                style="width: 126px; height: 100px;margin-top: 5px;" />
            </div>
          </a-form-item>
        </a-col>
      </a-row>
      <a-form-item label="文章内容" :name="['content']">
        <!-- 编辑器 -->
        <div style="border: 1px solid #ccc">
          <Toolbar style="border-bottom: 1px solid #ccc" :editor="editorRef" :defaultConfig="toolbarConfig"
            mode="simple" />
          <Editor style="height: 500px; overflow-y: hidden;" v-model="articleInfo.content" :defaultConfig="editorConfig"
            mode="simple" @onCreated="handleCreated" />
        </div>
      </a-form-item>
      <a-form-item class="btn">
        <a-button type="danger" style="margin-right: 15px" @click="addArticle">
          {{
          articleInfo.id ? '更新' : '提交'
          }}
        </a-button>
        <a-button type="primary" @click="addCancel">取消</a-button>
      </a-form-item>
    </a-form>
  </a-card>
</template>

<style scoped>
h3 {
  font-size: 18px;
  font-weight: bold;
}

.btn {
  padding: 20px;
  text-align: center;
}
</style>
