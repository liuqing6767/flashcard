<template>
  <div>
    <el-row>
      <!-- <el-col :span="8" class="one-card" v-if="editing">
        <el-card shadow="hover" class="card-add">
          <el-button
            type="success"
            icon="el-icon-plus"
            circle
            size="medium"
            @click="onEditBtnClick(0)"
          ></el-button>
        </el-card>
      </el-col>-->

      <el-col :span="8" v-for="k in cards" :key="k.id" class="one-card">
        <div @click="onCardClick(k.id)">
          <el-card shadow="hover">
            <div class="one-card-content">{{k.name}}</div>

            <div class="one-card-ops" v-if="!editing">
              <el-row>
                <el-col :span="17">
                  <el-progress :percentage="k.progress_rate"></el-progress>
                  <!-- <el-progress :percentage="k.progress_rate" :color="progressColors"></el-progress> -->
                </el-col>
                <el-col :span="7">
                  <el-tag
                    size="medium"
                    type="success"
                    v-if="k.next_learn.indexOf('前') == -1"
                  >{{k.next_learn}}</el-tag>
                  <el-tag size="medium" type="danger" v-else>{{k.next_learn}}</el-tag>
                </el-col>
              </el-row>
            </div>
            <div v-else class="btn-opts">
              <el-button
                type="primary"
                icon="el-icon-edit"
                circle
                @click="onEditBtnClick(k.id)"
                size="mini"
              ></el-button>
              <el-button
                type="danger"
                icon="el-icon-delete"
                circle
                @click="onDelBtnClick(k.id)"
                size="mini"
              ></el-button>
            </div>
          </el-card>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "CardList",
  props: {
    editing: Boolean,
    know_id: Number
  },
  methods: {
    onEditBtnClick(card_id) {
      this.$router.push({
        name: "card_edit",
        params: {
          card_id: card_id
        }
      });
    },
    onCardClick(card_id) {
      this.$router.push({
        name: "learning",
        params: {
          know_id: this.know_id
        },
        query: {
          cid: card_id
        }
      });
    },
    onDelBtnClick(card_id) {
      let _this = this;

      this.$confirm("卡片删除后不可恢复,是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          this.$http
            .post("/card/" + card_id + "/del")
            .then(res => {
              if (res.data.errno != 0) {
                this.$notify.error({
                  title: "错误",
                  message: res.data.msg
                });
                return;
              }

              let temp = new Array();
              _this.cards.forEach(element => {
                if (element.id != card_id) {
                  temp.push(element);
                }
              });
              _this.cards = temp;
            })
            .catch(err => {
              window.console.log(err);
            });
        })
        .catch(() => {});
    },
    updateData(know_id) {
      this.$http
        .get("/card/know/" + know_id)
        .then(res => {
          window.console.log(res);
          this.cards = res.data.data;
          this.formateList();
          setInterval(() => this.formateList(), 1000 * 60);
        })
        .catch(err => {
          window.console.log(err);
        });
    },
    formateList() {
      let now = Date.parse(new Date()) / 1000;
      for (let index = 0; index < this.cards.length; index++) {
        const element = this.cards[index];
        element.next_learn = this.formateDiff(
          element.next_learn_time,
          now,
          element.progress_rate
        );
      }
      this.cards = JSON.parse(JSON.stringify(this.cards));
    },
    formateDiff(time, now, progress_rate) {
      if (progress_rate == 100) {
        return "完成";
      }
      let diff = time - now;
      let ab = "后";
      if (diff < 0) {
        ab = "前";
        diff = -diff;
      }
      if (diff < 60) {
        return "现在";
      }

      let hour = parseInt(diff / 3600);
      if (hour > 24) {
        return "> 24h " + ab;
      }

      let min = parseInt(diff / 60) % 60;

      if (hour == 0) {
        return min + "m" + ab;
      }

      if (min < 10) {
        min = "0" + min;
      }
      return hour + "h" + min + "m" + ab;
    }
  },
  data() {
    if (this.know_id) {
      this.updateData(this.know_id);
    }
    return {
      cards: [],
      progressColors: [
        { color: "#90A19A", percentage: 20 },
        { color: "#e6a23c", percentage: 40 },
        { color: "#1989fa", percentage: 60 },
        { color: "#6f7ad3", percentage: 80 },
        { color: "#5cb87a", percentage: 100 }
      ]
    };
  },
  watch: {
    know_id: function(val) {
      this.updateData(val);
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.one-card {
  padding: 10px;
}

.one-card-content {
  height: 65px;
}

.card-add {
  text-align: center;
  padding: 23px;
}

.btn-opts {
  float: right;
  margin: 5px;
}
</style>
