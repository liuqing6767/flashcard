<template>
  <div class="form-login">
    <el-form ref="form" :model="form">
      <el-form-item>
        <el-input v-model="form.email" placeholder="注册邮箱"></el-input>
      </el-form-item>
      <el-form-item>
        <el-input v-model="form.password" placeholder="密码" type="password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-row>
          <el-col :span="16">
            <el-button type="primary" size="mini" @click="onLogin">登录</el-button>
          </el-col>
          <el-col :span="4">
            <el-link type="success" @click="onJump('register')">注册</el-link>
          </el-col>
          <el-col :span="4">
            <el-link type="warning" @click="onJump('reset_password')">找回密码</el-link>
          </el-col>
        </el-row>
      </el-form-item>
    </el-form>

    <About></About>
  </div>
</template>

<script>
import About from "../commons/About.vue";
import { SetUser } from "../../plugins/auth";
export default {
  components: { About },
  data() {
    return {
      form: {
        email: "",
        password: ""
      }
    };
  },
  methods: {
    onLogin() {
      if (this.form.email == "") {
        this.$notify.error({
          title: "错误",
          message: "请输入邮箱"
        });
        return;
      }
      if (this.form.password == "") {
        this.$notify.error({
          title: "错误",
          message: "请输入密码"
        });
        return;
      }

      this.$http
        .post("/auth/login", this.form)
        .then(res => {
          if (res.data.errno == 0) {
            SetUser(res.data.data.email);
            window.location.href = "/know";
            return;
          }
          this.$notify.error({
            title: "错误",
            message: res.data.msg
          });
        })
        .catch(err => {
          window.console.log(err);
        });
    },
    onJump(r) {
      this.$router.push({ name: r });
    }
  }
};
</script>

<style scoped>
.form-login {
  width: 32%;
  margin: auto;
  padding-top: 20%;
}
</style>