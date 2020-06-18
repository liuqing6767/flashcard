<template>
  <div class="fc-breadcrumb">
    <el-breadcrumb>
      <el-breadcrumb-item v-for="item in breadcrumb.items" :key="item.name">
        <router-link :to="item.path">{{item.name}}</router-link>
      </el-breadcrumb-item>

      <el-breadcrumb-item v-for="item in breadcrumb.btns" class="nav_child" :key="item.name">
        <router-link :to="item.path">{{item.name}}</router-link>
      </el-breadcrumb-item>
    </el-breadcrumb>
  </div>
</template>

<script>
export default {
  name: "BreadCrumb",
  props: {
    breadcrumb: Object
  }
};
</script>

<style scoped>
.fc-breadcrumb {
  border-bottom: 1px solid #dcdfe6;
  padding-top: 10px;
  padding-bottom: 10px;
  padding-left: 10px;
  margin-bottom: 10px;
}
.nav_child {
  float: right;
  margin-right: 10px;
  cursor: pointer;
}
</style>
