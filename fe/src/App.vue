<template>
  <div class="f-height app">
    <Nav></Nav>

    <div class="content-width f-height lr-border">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Nav from "./components/commons/Nav.vue";

export default {
  name: "app",
  components: { Nav }
};
</script>

<style>
html,
body {
  height: 100%;
  margin: 0;
}
.app {
  padding-top: 60px;
  box-sizing: border-box;
}
.f-height {
  height: 100%;
}
.lr-border {
  border-left: 1px solid #dcdfe6;
  border-right: 1px solid #dcdfe6;
}
.b-border {
  border-bottom: 1px solid #dcdfe6;
}
.content-width {
  width: 1200px;
  margin: auto;
}
.btn {
  border-radius: 4px;
  border: 1px solid #dcdfe6;
}
.border-bottom {
  border-bottom: 1px solid #dcdfe6;
  width: 100%;
}

.f-width {
  width: 100%;
}

.align-center {
  text-align: center;
}
.align-right {
  text-align: right;
}

.float-right {
  float: right;
}
</style>