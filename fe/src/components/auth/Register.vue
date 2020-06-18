<template>
  <div class="form-login">
    <el-form ref="form" :model="form">
      <el-form-item>
        <el-input v-model="form.email" placeholder="邮箱地址"></el-input>
      </el-form-item>
      <el-form-item>
        <el-input v-model="form.password" placeholder="密码" type="password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-input v-model="form.password2" placeholder="密码确认" type="password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-row>
          <el-col :span="22">
            <el-button type="primary" size="mini" @click="onSubmit">注册</el-button>
          </el-col>
          <el-col :span="2">
            <el-link type="warning" @click="onJump('login')">登录</el-link>
          </el-col>
        </el-row>
      </el-form-item>
    </el-form>

    <About></About>
  </div>
</template>

<script>
import About from "../commons/About.vue";
export default {
  components: { About },
  data() {
    return {
      form: {
        email: "",
        password: "",
        password2: ""
      }
    };
  },
  methods: {
    onSubmit() {
      if (this.form.email == "") {
        this.$notify.error({
          title: "错误",
          message: "请输入邮箱"
        });
        return;
      }
      if (this.form.password.length < 6) {
        this.$notify.error({
          title: "错误",
          message: "密码至少六位"
        });
        return;
      }
      if (this.form.password2.length < 6) {
        this.$notify.error({
          title: "错误",
          message: "验证密码至少六位"
        });
        return;
      }
      if (this.form.password != this.form.password2) {
        this.$notify.error({
          title: "错误",
          message: "两次密码需要一致"
        });
        return;
      }

      this.$http
        .post("/auth/register", this.form)
        .then(res => {
          if (res.data.errno == 0) {
            this.$notify({
              title: "成功",
              message:
                "我们已经将一封邮件发送到你的邮箱，请点击链接激活后使用。",
              type: "success",
              duration: 0
            });
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
      this.$router.replace({ name: r });
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