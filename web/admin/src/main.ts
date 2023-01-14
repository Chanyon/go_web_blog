import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router';

import {
  Menu, Layout,Row,Col,Card,message,
  Form,Button,Modal,Table,Popconfirm,
  Input,Select,Upload,Switch,Slider} from 'ant-design-vue'
// import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';

const app = createApp(App);

// app.use(Antd);
app.use(Menu).use(Layout).use(Slider)
.use(Card).use(Row).use(Col)
.use(Button).use(Modal).use(Table)
.use(Popconfirm).use(Form).use(Input)
.use(Select).use(Upload).use(Switch);
message.config({
  top: `60px`,
  duration: 2,
  maxCount: 3
});
// app.component('MyC1',myc1); 全局注册
// app.config.globalProperties.$message = message;
app.use(router);
app.mount('#app');
