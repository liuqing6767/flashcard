import Vue from 'vue';
import App from './App.vue';
import ajax from './plugins/axios.js';
import './plugins/element.js';
import {
  notification
} from './plugins/notification';
import {
  router
} from './router';


Vue.prototype.$ajax = ajax;

Vue.config.productionTip = false;

new Vue({
  el: '#app',
  router, // 注入到根实例中
  render: h => h(App)
});

notification(function () {
  this.$http
    .get("/card/progress")
    .then(res => {
      if (res.data.data && res.data.data.count) {
        new window.Notification("闪记卡·熙穆", {
          body: "你有" + res.data.data.count + "张卡片需要学习了"
        });
      }
    })
    .catch(err => {
      window.console.log(err);
    });

}, 1000 * 60);