import { createApp } from 'vue';
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';
import App from './App.vue';
import router from './router';
import pinia from './store';
import { permission } from './directives/permission';
import './styles/global.less';

const app = createApp(App);

app.use(pinia);
app.use(router);
app.use(Antd);
app.directive('permission', permission);

app.mount('#app');
