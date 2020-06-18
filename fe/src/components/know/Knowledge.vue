<template>
  <div class="f-height">
    <knowledge-tree
      class="c-ktree lr-border f-height"
      @node-selected="nodeSelected"
      :know_id="know_id"
      v-on:mode-switched="modeSwitched"
    ></knowledge-tree>

    <div class="know-opt">
      <el-button
        type="success"
        icon="el-icon-plus"
        circle
        size="mini"
        class="btn-add"
        :disabled="!know_id"
        @click="onAddBtnClick()"
      ></el-button>
      <el-button
        icon="el-icon-video-play"
        size="mini"
        @click="onPlayBtnClick"
        class="btn-play"
        :disabled="editing"
      ></el-button>
    </div>
    <knowledge-points class="c-kpoints" :know_id="know_id" :editing="editing"></knowledge-points>
  </div>
</template>

<script>
import KnowledgeTree from "./Tree.vue";
import KnowledgePoints from "../card/CardList.vue";

export default {
  name: "Knowledge",
  components: {
    KnowledgeTree,
    KnowledgePoints
  },
  data() {
    let know_id = "";
    if (this.$route.params.know_id) {
      know_id = this.$route.params.know_id;
    }
    return {
      know_id: parseInt(know_id),
      editing: false
    };
  },
  methods: {
    nodeSelected(know_id) {
      if (this.know_id == know_id) {
        return;
      }
      this.know_id = know_id;
      this.$router.push({
        name: "know_list",
        params: {
          know_id: this.know_id
        }
      });
    },
    onPlayBtnClick() {
      this.$router.push({
        name: "learning",
        params: {
          know_id: this.know_id
        }
      });
    },
    modeSwitched(editing) {
      this.editing = editing;
    },
    onAddBtnClick() {
      this.$router.push({
        name: "card_edit",
        params: {
          card_id: 0
        },
        query: {
          know_id: this.know_id
        }
      });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.c-ktree {
  width: 303px;
  float: left;
}

.c-kpoints {
  width: 895px;
  float: left;
}

.know-opt {
  float: right;
  margin: 5px;
  width: 875px;
}
.btn-play {
  float: right;
  margin-right: 5px;
}
</style>
