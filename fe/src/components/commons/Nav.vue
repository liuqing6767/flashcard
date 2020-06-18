<template>
  <div class="nav">
    <div class="content-width">
      <el-menu
        class="fc-menu"
        mode="horizontal"
        background-color="#282830"
        text-color="#f2f2f7"
        active-text-color="#0089ff"
        :router="true"
      >
        <el-menu-item index="/know">速记卡</el-menu-item>

        <el-submenu index="2">
          <template slot="title">设置</template>
          <el-menu-item index="/template">卡片模板</el-menu-item>
          <!-- <el-menu-item index="/loop">记忆周期</el-menu-item> -->
        </el-submenu>

        <template v-if="logined">
          <el-submenu index="3" style="float:right">
            <template slot="title">{{email}}</template>
            <el-menu-item index="/auth/reset">修改密码</el-menu-item>
            <el-menu-item @click="logout">退出</el-menu-item>
          </el-submenu>
        </template>
        <template v-else>
          <el-menu-item index="/mine/login" style="float:right">{{email}}</el-menu-item>
        </template>
      </el-menu>
    </div>
  </div>
</template>

<script>
import { GetUser, ClearUser } from "../../plugins/auth";
export default {
  name: "Nav",
  data() {
    let logined = true;
    let email = GetUser();
    if (!email) {
      logined = false;
      email = "请登录";
    }

    return {
      logined: logined,
      email: email
    };
  },
  methods: {
    logout() {
      ClearUser();
      window.location.href = "/mine/login";
    }
  }
};
</script>

<style>
.nav {
  background-color: #282830;
  margin-top: -61px;
}

.fc-menu {
  border-bottom: 0px !important;
}

.fc-menu a {
  text-decoration: none;
}
</style>

