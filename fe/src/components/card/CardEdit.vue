<template v-loading="loading">
  <div class="fc-card-create">
    <bread-crumb
      :breadcrumb="{items:[{name:'配置', path:''}, {name:'卡片', path: '/card' } ,{name:'编辑', path: '' }]}"
    ></bread-crumb>
    <div class="fc-card-create">
      <el-form>
        <el-row>
          <el-col :span="14">
            <el-row>
              <el-col :span="12">
                <el-select
                  v-model="templateID"
                  placeholder="请选择记忆模板"
                  class="f-width"
                  @change="onSelectChange"
                >
                  <el-option
                    v-for="item in templates"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                  ></el-option>
                </el-select>
              </el-col>

              <el-col :span="12">
                <el-select
                  v-model="remeberCurve"
                  placeholder="请选择记忆曲线"
                  class="f-width"
                  @change="onSelectChange"
                >
                  <el-option
                    v-for="item in remeberCurves"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                  ></el-option>
                </el-select>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="12">
                <el-input
                  type="textarea"
                  v-model="cardData"
                  @dblclick.native="onDataClick"
                  placeholder="卡片内容，单词模板可输入单测后双击完成编辑"
                  :rows="30"
                ></el-input>
              </el-col>
              <el-col :span="12">
                <el-input
                  type="textarea"
                  @dblclick.native="onDemoClick"
                  v-model="templateDemoData"
                  placeholder="内容模板"
                  :rows="30"
                  :disabled="true"
                ></el-input>
              </el-col>
            </el-row>
          </el-col>
          <el-col :span="10">
            <el-row>
              <div class="show" id="show-a">
                <div class="show-div">卡片正面</div>
              </div>
            </el-row>
            <el-row class="show" id="show-b">
              <div class="show-div">卡片反面</div>
            </el-row>
          </el-col>
        </el-row>
        <el-row>
          <el-button type="primary" size="small" @click="onSubmit" style="margin-top:20px">保存</el-button>
          <el-button size="small" @click="goBack(1)" style="margin-top:20px">返回</el-button>
        </el-row>
      </el-form>
    </div>
  </div>
</template>>

<script>
import Vue from "vue/dist/vue.js";
import BreadCrumb from "./../commons/BreadCrumb.vue";

export default {
  name: "CardEdit",
  components: { BreadCrumb },

  data() {
    let _this = this;
    this.$http
      .get("/template")
      .then(res => {
        window.console.log(res);
        let templates = [];
        let templateMap = {};
        for (let index = 0; index < res.data.data.length; index++) {
          const template = res.data.data[index];
          templates.push({
            value: template.id,
            label: template.name
          });
          templateMap[template.id] = template;

          // 把模板注册进去
          Vue.component("a" + template.id, {
            props: {
              data: Object
            },
            template: template.content_a
          });
          Vue.component("b" + template.id, {
            props: {
              data: Object
            },
            template: template.content_b
          });
        }
        _this.templates = templates;
        _this.templateMap = templateMap;

        _this.render();
      })
      .catch(err => {
        window.console.log(err);
      });

    let data = {
      templates: [],
      templateMap: {},
      remeberCurves: [
        {
          value: 1,
          label: "艾宾浩斯记忆曲线"
        }
      ],

      templateDemoData: "",

      cardID: 0,
      cardData: "",
      templateID: "",
      remeberCurve: 1,
      knowID: 0,
      loading: false
    };

    if (this.$route.params.card_id) {
      data.cardID = parseInt(this.$route.params.card_id);
    }

    if (this.$route.query.know_id) {
      data.knowID = parseInt(this.$route.query.know_id);
    }

    if (data.cardID != 0) {
      this.$http
        .get("/card/" + data.cardID)
        .then(res => {
          if (res.data.errno != 0) {
            this.$notify.error({
              title: "错误",
              message: res.data.msg
            });
            this.goBack();
            return;
          }

          let data = res.data.data;
          this.cardData = data.data;
          this.templateID = data.template_id;
          this.knowID = data.know_id;
        })
        .catch(err => {
          window.console.log(err);
        });
    }

    return data;
  },
  methods: {
    render() {
      let list = ["a", "b"];
      for (let i = 0; i < list.length; i++) {
        let data = {};
        try {
          data = JSON.parse(this.cardData);
        } catch (error) {
          return;
        }

        let key = list[i];
        let componentA = Vue.component(key + this.templateID);
        const instanceA = new componentA({
          data: data
        });
        // Vue.set(instanceA, "data", data);
        document.getElementById("show-" + key).innerHTML = "<div></div>";
        instanceA.$mount("#show-" + key + " > div");
      }
    },
    onSubmit() {
      if (!this.templateID) {
        this.$notify.error({
          title: "错误",
          message: "必须选择一个模板"
        });
        return;
      }

      if (!this.cardData) {
        this.$notify.error({
          title: "错误",
          message: "必须填写卡片内容"
        });
        return;
      }

      try {
        JSON.parse(this.cardData);
      } catch (error) {
        this.$notify.error({
          title: "错误",
          message: "卡片内容不是合法的JSON格式"
        });
        return;
      }

      let param = {
        id: this.cardID,
        know_id: this.knowID,
        template_id: this.templateID,
        data: this.cardData
      };

      this.$http
        .post("/card", param)
        .then(res => {
          if (res.data.errno != 0) {
            this.$notify.error({
              title: "错误",
              message: res.data.msg
            });
            return;
          }

          this.$message({
            type: "success",
            message: `卡片保存成功`
          });

          if (this.cardID == 0) {
            this.cardData = this.templateMap[this.templateID].data_demo;
            return;
          }
          this.goBack();
        })
        .catch(err => {
          window.console.log(err);
        });
    },
    onSelectChange() {
      let curTemp = this.templateMap[this.templateID];
      this.templateDemoData =
        curTemp.data_demo + "\n\n双击文本框可快速将数据放入卡片内容";
    },
    onDemoClick() {
      this.cardData = this.templateMap[this.templateID].data_demo;
    },
    onDataClick() {
      let data = this.cardData.toLowerCase();
      if (data == "") {
        this.$notify.error({
          title: "错误",
          message: "单测不能为空"
        });
        return;
      }
      if (data.indexOf(" ") != -1) {
        this.$notify.error({
          title: "错误",
          message: "当前只只能记忆单测，不能记忆句子"
        });
        return;
      }
      this.loading = true;
      this.$http
        .get("/card/word/" + data)
        .then(res => {
          this.loading = false;
          if (res.data.errno != 0) {
            this.$notify.error({
              title: "错误",
              message: res.data.msg
            });
            return;
          }

          let obj = {
            Word: data,
            Detail: res.data.data
          };
          this.cardData = JSON.stringify(obj, null, 4);
        })
        .catch(err => {
          window.console.log(err);
        });
    },
    goBack(arg) {
      let timeout = 2000;
      if (arg) {
        timeout = 0;
      }

      setTimeout(() => {
        this.$router.push({
          name: "know_list",
          params: { know_id: this.knowID }
        });
      }, timeout);
    }
  },
  watch: {
    cardData() {
      this.render();
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this comptemplatent only -->
<style scoped>
.fc-card-create {
  margin: 0px 20px;
}

.show {
  margin: 35px 0 35px 35px;
  width: 430px;
  height: 270px;
  border: 1px solid #dcdfe6;
}

.show-div {
  text-align: center;
  color: #606266;
  padding: 20px;
}
</style>
