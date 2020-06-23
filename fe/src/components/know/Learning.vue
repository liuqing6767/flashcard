<template>
  <div
    class="l-panel"
    @keyup.enter="onViewBtnClick"
    @keyup.72="onLearnBtnClick(0)"
    @keyup.77="onLearnBtnClick(1)"
    @keyup.76="onLearnBtnClick(2)"
  >
    <div class="l-box">
      <div class="l-content">
        <div id="learning-content"></div>
      </div>
      <div class="l-opt">
        <el-row>
          <el-col :span="7" class="l-opt-l">
            <el-button
              type="primary"
              plain
              icon="el-icon-view"
              circle
              size="mini"
              @click="onViewBtnClick"
            >S</el-button>
          </el-col>
          <el-col :span="10" class="align-center">
            <el-button type="danger" round size="mini" @click="onLearnBtnClick(2)">陌生 L</el-button>
            <el-button type="primary" round size="mini" @click="onLearnBtnClick(1)">记得 M</el-button>
            <el-button type="success" round size="mini" @click="onLearnBtnClick(0)">熟悉 H</el-button>
          </el-col>
          <el-col :span="5" class="l-opt-r align-right">
            <el-progress
              :text-inside="true"
              :stroke-width="25"
              status="warning"
              :percentage="parseInt(index/total*100)"
              :format="rateInfo"
            ></el-progress>
          </el-col>
          <el-col :span="2" class="l-opt-r align-right">
            <el-button icon="el-icon-close" circle size="mini" @click="onExitBtnClick">退出</el-button>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from "vue/dist/vue.js";

export default {
  name: "KnowledgeLearning",
  components: {},
  methods: {
    render(isA) {
      let t = "a";
      if (!isA) {
        t = "b";
      }

      if (this.index > this.total) {
        return;
      }

      let card = this.cards[this.index];

      let component = Vue.component(t + card.template_id);

      let data = JSON.parse(card.data);
      const instance = new component({
        data: data
      });

      document.getElementById("learning-content").innerHTML = "<div></div>";

      instance.$mount("#learning-content > div");
    },

    rateInfo() {
      return this.index + " / " + this.total;
    },
    onExitBtnClick() {
      this.$router.push({
        name: "know_list",
        params: {
          know_id: this.$route.params.know_id
        }
      });
    },
    onViewBtnClick() {
      this.isA = !this.isA;
      this.render(this.isA);
    },
    onLearnBtnClick(f) {
      this.$http
        .post("/card/progress", {
          f: f,
          cid: this.cards[this.index].id
        })
        .catch(err => {
          window.console.log(err);
        });

      this.index++;
      if (this.index >= this.total) {
        const loading = this.$loading({
          lock: true,
          text: "学习完成，跳转中",
          background: "rgba(0, 0, 0, 0.7)"
        });
        setTimeout(() => {
          loading.close();
          this.onExitBtnClick();
        }, 2000);

        return;
      }
      this.isA = true;
      this.render(this.isA);
    }
  },
  data() {
    let _this = this;
    let cid = this.$route.query["cid"];
    this.$http
      .get("/card/learning/" + this.$route.params.know_id + "?cid=" + cid)
      .then(res => {
        if (res.data.errno != 0) {
          this.$notify.error({
            title: "错误",
            message: res.data.msg
          });

          setTimeout(() => {
            this.onExitBtnClick();
          }, 2000);
          return;
        }

        if (!res.data.data.cards || res.data.data.cards.length == 0) {
          this.$notify.error({
            title: "错误",
            message: "没有需要学习的内容，即将返回"
          });

          setTimeout(() => {
            this.onExitBtnClick();
          }, 2000);
          return;
        }

        this.cards = res.data.data.cards;
        this.templates = res.data.data.templates;
        for (let i = 0; i < this.templates.length; i++) {
          const template = this.templates[i];
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

        this.index = 0;
        this.total = this.cards.length;
        this.isA = true;

        _this.render(true, true);
      })
      .catch(err => {
        window.console.log(err);
      });

    return {
      index: 0,
      total: 1,
      isA: true,
      cards: [],
      templates: []
    };
  }
};
</script>

<style>
.l-panel {
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0px;
  right: 0px;
  bottom: 0px;
  left: 0px;
  background: #282830;
}

.l-box {
  background: white;
  margin: auto;
  width: 900px;
  height: 600px;

  position: absolute;
  top: 50%;
  left: 50%;

  margin: -300px 0 0 -450px;
}

.l-content {
  width: 900px;
  height: 560px;
  /* background: black; */
}
.l-opt {
  width: 900px;
  height: 40px;
  padding-top: 5px;
  border-top: 1px solid #dcdfe6;
}
.l-opt-l {
  padding-left: 10px;
}
.l-opt-r {
  padding-right: 10px;
}

#learning-content {
  width: 100%;
  height: 100%;
}
</style>