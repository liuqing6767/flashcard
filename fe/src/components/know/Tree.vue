<template>
  <div>
    <el-row>
      <el-col :span="21">
        <el-input placeholder="关键字过滤" v-model="filterText"></el-input>
      </el-col>
      <el-col :span="3">
        <div class="btn-opt">
          <el-button
            v-if="editing"
            icon="el-icon-s-opportunity"
            type="success"
            circle
            size="mini"
            @click="onModeChange(false)"
          ></el-button>
          <el-button
            v-else
            type="primary"
            icon="el-icon-edit"
            circle
            size="mini"
            @click="onModeChange(true)"
          ></el-button>
        </div>
      </el-col>
    </el-row>

    <el-row>
      <el-tree
        ref="tree"
        :data="data"
        :expand-on-click-node="false"
        :filter-node-method="filterNode"
        @node-drop="handleDrop"
        @node-click="handleNodeClick"
        :default-expanded-keys="treeExpandData"
        draggable
        node-key="id"
        :current-node-key="know_id"
        :render-content="renderContent"
        :allow-drop="allowDrop"
        :allow-drag="allowDrag"
      ></el-tree>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "KnowledgeTree",
  watch: {
    filterText(val) {
      this.$refs.tree.filter(val);
    }
  },
  props: {
    know_id: Number
  },
  methods: {
    onModeChange(editing) {
      this.editing = editing;
      window.console.log("mode", this.editing);
      this.data = JSON.parse(JSON.stringify(this.data));
      this.$emit("mode-switched", this.editing);
    },
    filterNode(value, data) {
      if (!value) return true;
      return data.label.indexOf(value) !== -1;
    },
    handleNodeClick(data) {
      this.$emit("node-selected", data.id);
    },
    handleDrop(draggingNode, dropNode, dropType) {
      let data = {
        type: "drag",
        id: draggingNode.data.id,
        rid: dropNode.data.id,
        drop_type: dropType
      };
      window.console.log(data);
      this.$http
        .post("/know/modify", data)
        .then(res => {
          if (res.data.errno != 0) {
            this.$notify.error({
              title: "错误",
              message: res.data.msg
            });
            return;
          }
        })
        .catch(err => {
          window.console.log(err);
        });
    },
    allowDrop(draggingNode, dropNode, type) {
      return !(dropNode.data.id == 0 && (type == "prev" || type == "inner"));
    },
    allowDrag(draggingNode) {
      return draggingNode.data.id;
    },
    edit(data) {
      let oldName = data.label;
      this.$prompt("请输入知识点名称", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        inputValue: oldName,
        inputValidator: function(val) {
          if (val == "" || val == oldName) {
            return "请输入和旧名字不一样的名字";
          }
          return true;
        }
      })
        .then(({ value }) => {
          this.$http
            .post("/know/modify", {
              type: "rename",
              id: data.id,
              name: value
            })
            .then(res => {
              if (res.data.errno != 0) {
                this.$notify.error({
                  title: "错误",
                  message: res.data.msg
                });
                return;
              }

              data.label = value;
            })
            .catch(err => {
              window.console.log(err);
            });
        })
        .catch(() => {});
    },

    append(data) {
      this.treeExpandData = [data.id];
      this.$prompt("请输入子知识点名称", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        inputValidator(data) {
          if (!data || data.length < 2 || data.length > 10) {
            return "名字长度要求2-10";
          }
          return true;
        }
      })
        .then(({ value }) => {
          this.$http
            .post("/know", {
              pid: data.id,
              name: value
            })
            .then(res => {
              if (res.data.errno != 0) {
                this.$notify.error({
                  title: "错误",
                  message: res.data.msg
                });
                return;
              }

              let newChild = {
                id: res.data.data.id,
                label: value,
                children: []
              };
              if (!data.children) {
                this.$set(data, "children", []);
              }
              data.children.push(newChild);
            })
            .catch(err => {
              window.console.log(err);
            });
        })
        .catch(() => {});
    },
    remove(node, data) {
      this.$confirm(
        "知识点下的卡片将移动到「默认知识点」中，知识点将被删除,是否继续?",
        "提示",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }
      )
        .then(() => {
          this.$http
            .post("/know/modify", {
              type: "delete",
              id: data.id
            })
            .then(res => {
              if (res.data.errno != 0) {
                this.$notify.error({
                  title: "错误",
                  message: res.data.msg
                });
                return;
              }

              let parent = node.parent;
              let children = parent.data.children || parent.data;
              let index = children.findIndex(d => d.id === data.id);
              children.splice(index, 1);
            })
            .catch(err => {
              window.console.log(err);
            });
        })
        .catch(() => {});
    },

    renderContent(h, { node, data }) {
      if (!this.editing) {
        return (
          <span class="ktree-node">
            <span> {node.label} </span>
            <span class="r-tree-btn">
              {data.learning}/{data.total}
            </span>
          </span>
        );
      }

      return (
        <span class="ktree-node">
          <span> {node.label} </span>
          <span class="r-tree-btn">
            <el-button
              size="mini"
              type="text"
              on-click={() => this.append(data)}
            >
              <i class="el-icon-circle-plus-outline"></i>
            </el-button>
            <el-button size="mini" type="text" on-click={() => this.edit(data)}>
              <i class="el-icon-edit"></i>
            </el-button>
            <el-button
              size="mini"
              type="text"
              disabled={data.id == 0}
              on-click={() => this.remove(node, data)}
            >
              <i class="el-icon-remove-outline"></i>
            </el-button>
          </span>
        </span>
      );
    }
  },

  data() {
    this.$http
      .get("/know")
      .then(res => {
        this.data = [res.data.data];
      })
      .catch(err => {
        window.console.log(err);
      });

    return {
      treeExpandData: [0],
      filterText: "",
      editing: false,
      data: [],
      defaultProps: {
        children: "children",
        label: "label"
      }
    };
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
.r-tree-btn {
  float: right;
  padding-right: 10px;
}

.ktree-node {
  width: 100%;
  line-height: 200%;
}

.btn-opt {
  padding: 5px;
}
</style>
