<template>
  <div>
    <bread-crumb :breadcrumb="{ items:[{name:'配置', path: ''}, {name:'卡片模板', path: '' } ] }"></bread-crumb>

    <el-row>
      <el-col :span="8" class="t-row">
        <el-card shadow="hover" class="t-add">
          <el-button
            type="success"
            icon="el-icon-plus"
            circle
            size="medium"
            @click="onEditBtnClick(0)"
          ></el-button>
        </el-card>
      </el-col>

      <el-col :span="8" v-for="k in templates" :key="k.id" class="t-row">
        <el-card shadow="hover">
          <div class="t-card">
            {{k.name}}
            <template v-if="k.is_official">· 官方</template>
          </div>
          <div class="t-opt">
            <el-button
              type="primary"
              icon="el-icon-edit"
              circle
              :disabled="k.is_official"
              @click="onEditBtnClick(k.id)"
              size="mini"
            ></el-button>
            <el-button
              type="danger"
              icon="el-icon-delete"
              circle
              :disabled="k.is_official"
              @click="onDelBtnClick(k.id)"
              size="mini"
            ></el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import BreadCrumb from "./../commons/BreadCrumb.vue";

export default {
  name: "CardTemplateList",
  props: {
    msg: String
  },
  methods: {
    onDelBtnClick(temp_id) {
      let _this = this;

      this.$confirm(
        "模板删除后将不能修改，也不能被新的卡片使用,是否继续?",
        "提示",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }
      )
        .then(() => {
          this.$http
            .post("/template/" + temp_id + "/del")
            .then(res => {
              if (res.data.errno != 0) {
                this.$notify.error({
                  title: "错误",
                  message: res.data.msg
                });
                return;
              }

              let temp = new Array();
              _this.templates.forEach(element => {
                if (element.id != temp_id) {
                  temp.push(element);
                }
              });
              _this.templates = temp;
            })
            .catch(err => {
              window.console.log(err);
            });
        })
        .catch(() => {});
    },
    onEditBtnClick(temp_id) {
      this.$router.push({
        name: "template_edit",
        params: {
          temp_id: temp_id
        }
      });
    }
  },
  components: { BreadCrumb },
  data() {
    let _this = this;
    this.$http
      .get("/template")
      .then(res => {
        window.console.log(res);
        _this.templates = res.data.data;
      })
      .catch(err => {
        window.console.log(err);
      });

    return {
      templates: []
    };
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.t-row {
  padding: 10px;
}

.t-card {
  height: 40px;
}

.t-add {
  text-align: center;
  padding: 10px;
}

.t-opt {
  float: right;
  padding-bottom: 10px;
}
</style>
