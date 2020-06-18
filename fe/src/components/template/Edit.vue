<template>
  <div class="fc-template-create lr-border">
    <bread-crumb
      :breadcrumb="{items:[{name:'配置', path:''}, {name:'卡片模板', path: '/template' } ,{name:'编辑', path: '' }]}"
    ></bread-crumb>
    <div class="fc-temp-create">
      <el-form ref="template" :model="template" label-width="100px">
        <el-form-item label="模板名称">
          <el-input v-model="template.name"></el-input>
        </el-form-item>

        <el-form-item label="正面内容">
          <el-row>
            <el-col :span="12">
              <el-input
                type="textarea"
                v-model="template.content_a"
                placeholder="<div class='width:100px;height:100px'>*</div>"
                :rows="13"
              ></el-input>
            </el-col>
            <el-col :span="12" class="render">
              <div id="pagea">预览</div>
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item label="背面内容">
          <el-row>
            <el-col :span="12">
              <el-input
                type="textarea"
                v-model="template.content_b"
                placeholder="<div class='width:100px;height:100px'>*</div>"
                :rows="13"
              ></el-input>
            </el-col>
            <el-col :span="12" class="render">
              <div id="pageb">预览</div>
            </el-col>
          </el-row>
        </el-form-item>

        <el-form-item label="测试数据">
          <el-row>
            <el-col :span="12">
              <el-input type="textarea" v-model="template.data_demo" placeholder="{}" :rows="10"></el-input>
            </el-col>
            <el-col :span="12"></el-col>
          </el-row>
        </el-form-item>

        <el-form-item label="卡片名称">
          <el-row>
            <el-col :span="12">
              <el-input v-model="template.card_name_key" placeholder="为测试数据的Key, 在创建卡片时给卡片取名"></el-input>
            </el-col>
            <el-col :span="12"></el-col>
          </el-row>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" size="small" @click="onSubmit">保存</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import Vue from "vue/dist/vue.js";
import BreadCrumb from "./../commons/BreadCrumb.vue";

export default {
  name: "CardTemplateEdit",
  components: { BreadCrumb },

  watch: {
    template: {
      handler: function() {
        this.onContentChange();
      },
      deep: true
    }
  },

  data() {
    let temp_id = this.$route.params.temp_id;
    if (temp_id != 0) {
      this.temp_id = temp_id;
      let _this = this;
      this.$http
        .get("/template/" + temp_id)
        .then(res => {
          if (res.data.errno == 0 && !res.data.data.is_official) {
            _this.template = res.data.data;
            return;
          }

          this.$notify.error({
            title: "错误",
            message: res.data.msg
          });

          setTimeout(() => {
            this.$router.push({ name: "template_list" });
          }, 2000);
        })
        .catch(err => {
          window.console.log(err);
        });
    }

    return {
      temp_id: temp_id,
      template: {
        name: "",
        content_a:
          '<div style="position:relative; width:100%; height:100%;">\n\t<div style="position:absolute; top:50%; left:50%; transform: translate(-50%, -50%);">\n\t\t{{Q}}\n\t</div>\n</div>',
        content_b:
          '<div style="position:relative; width:100%; height:100%;">\n\t<div style="position:absolute; top:50%; left:50%; transform: translate(-50%, -50%);">\n\t\t{{A}}\n\t</div>\n</div>',
        data_demo: '{\n\t"Q": "问题是什么",\n\t"A": "答案是什么"\n}'
      }
    };
  },
  methods: {
    onContentChange() {
      let data_demo = this.template.data_demo;
      let pageA = Vue.extend({
        template: this.template.content_a,
        data() {
          return JSON.parse(data_demo);
        }
      });

      document.getElementById("pagea").innerHTML = "<div></div>";
      new pageA().$mount("#pagea > div");

      let pageB = Vue.extend({
        template: this.template.content_b,
        data() {
          return JSON.parse(data_demo);
        }
      });
      document.getElementById("pageb").innerHTML = "<div></div>";
      new pageB().$mount("#pageb > div");
    },
    onSubmit() {
      let dataDemo = {};
      try {
        dataDemo = JSON.parse(this.template.data_demo);
      } catch (error) {
        this.$notify.error({
          title: "错误",
          message: "测试内容非合法JSON"
        });
        return;
      }
      let name = dataDemo[this.template.card_name_key];
      if (!name || typeof name != "string") {
        this.$notify.error({
          title: "错误",
          message: "卡片名称不是测试数据中的有效的key"
        });
        return;
      }

      let url = "/template";
      if (this.temp_id != 0) {
        url = url + "/" + this.temp_id;
      }
      this.$http
        .post(url, this.template)
        .then(res => {
          if (res.data.errno == 0) {
            this.$router.push({ name: "template_list" });
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
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.fc-temp-create {
  margin: 0px 20px;
}

.render {
  width: 450px;
  height: 280px;
  background: #282830;
  border: #282830 solid 1px;
  margin-left: 20px;
}

#pagea,
#pageb {
  text-align: center;
  color: #ffffff;
  height: 280px;
}
</style>
