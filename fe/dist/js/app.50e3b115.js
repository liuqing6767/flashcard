(function(t){function e(e){for(var a,r,s=e[0],l=e[1],c=e[2],u=0,p=[];u<s.length;u++)r=s[u],Object.prototype.hasOwnProperty.call(i,r)&&i[r]&&p.push(i[r][0]),i[r]=0;for(a in l)Object.prototype.hasOwnProperty.call(l,a)&&(t[a]=l[a]);d&&d(e);while(p.length)p.shift()();return o.push.apply(o,c||[]),n()}function n(){for(var t,e=0;e<o.length;e++){for(var n=o[e],a=!0,s=1;s<n.length;s++){var l=n[s];0!==i[l]&&(a=!1)}a&&(o.splice(e--,1),t=r(r.s=n[0]))}return t}var a={},i={app:0},o=[];function r(e){if(a[e])return a[e].exports;var n=a[e]={i:e,l:!1,exports:{}};return t[e].call(n.exports,n,n.exports,r),n.l=!0,n.exports}r.m=t,r.c=a,r.d=function(t,e,n){r.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:n})},r.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},r.t=function(t,e){if(1&e&&(t=r(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var n=Object.create(null);if(r.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var a in t)r.d(n,a,function(e){return t[e]}.bind(null,a));return n},r.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return r.d(e,"a",e),e},r.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},r.p="/";var s=window["webpackJsonp"]=window["webpackJsonp"]||[],l=s.push.bind(s);s.push=e,s=s.slice();for(var c=0;c<s.length;c++)e(s[c]);var d=l;o.push([0,"chunk-vendors"]),n()})({0:function(t,e,n){t.exports=n("56d7")},"034f":function(t,e,n){"use strict";var a=n("85ec"),i=n.n(a);i.a},"0709":function(t,e,n){"use strict";var a=n("b44a"),i=n.n(a);i.a},"088d":function(t,e,n){},"0953":function(t,e,n){"use strict";var a=n("f0d9"),i=n.n(a);i.a},"1b2c":function(t,e,n){"use strict";var a=n("fdf2"),i=n.n(a);i.a},"1d05":function(t,e,n){},"1f5b":function(t,e,n){"use strict";var a=n("1d05"),i=n.n(a);i.a},2232:function(t,e,n){"use strict";var a=n("42b2"),i=n.n(a);i.a},"22f5":function(t,e,n){},"2e91":function(t,e,n){},3307:function(t,e,n){},4054:function(t,e,n){},"42b2":function(t,e,n){},"4f5f":function(t,e,n){"use strict";var a=n("3307"),i=n.n(a);i.a},"56d7":function(t,e,n){"use strict";n.r(e);n("e260"),n("e6cf"),n("cca6"),n("a79d");var a=n("2b0e"),i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"f-height app"},[n("Nav"),n("div",{staticClass:"content-width f-height lr-border"},[n("router-view")],1)],1)},o=[],r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"nav"},[n("div",{staticClass:"content-width"},[n("el-menu",{staticClass:"fc-menu",attrs:{mode:"horizontal","background-color":"#282830","text-color":"#f2f2f7","active-text-color":"#0089ff",router:!0}},[n("el-menu-item",{attrs:{index:"/know"}},[t._v("速记卡")]),n("el-submenu",{attrs:{index:"2"}},[n("template",{slot:"title"},[t._v("设置")]),n("el-menu-item",{attrs:{index:"/template"}},[t._v("卡片模板")])],2),t.logined?[n("el-submenu",{staticStyle:{float:"right"},attrs:{index:"3"}},[n("template",{slot:"title"},[t._v(t._s(t.email))]),n("el-menu-item",{attrs:{index:"/auth/reset"}},[t._v("修改密码")]),n("el-menu-item",{on:{click:t.logout}},[t._v("退出")])],2)]:[n("el-menu-item",{staticStyle:{float:"right"},attrs:{index:"/mine/login"}},[t._v(t._s(t.email))])]],2)],1)])},s=[];function l(){return window.localStorage.getItem("email")}function c(t){window.localStorage.setItem("email",t)}function d(){window.localStorage.clear("email")}function u(){return l()}var p={name:"Nav",data:function(){var t=!0,e=l();return e||(t=!1,e="请登录"),{logined:t,email:e}},methods:{logout:function(){d(),window.location.href="/mine/login"}}},m=p,f=(n("7aac"),n("2877")),h=Object(f["a"])(m,r,s,!1,null,null,null),v=h.exports,w={name:"app",components:{Nav:v}},g=w,b=(n("034f"),Object(f["a"])(g,i,o,!1,null,null,null)),_=b.exports,k=(n("96cf"),n("1da1")),y=n("bc3a"),$=n.n(y);a["default"].prototype.$http=$.a,$.a.defaults.baseURL="/api",$.a.defaults.withCredentials=!0,$.a.interceptors.response.use((function(t){return 401==t.data.errno&&(window.location.href="/mine/login"),t}));var x={get:function(){var t=Object(k["a"])(regeneratorRuntime.mark((function t(e,n){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return window.console.log(1),t.prev=1,t.next=4,$.a.get(e,n);case 4:return t.abrupt("return",t.sent);case 7:return t.prev=7,t.t0=t["catch"](1),t.abrupt("return",t.t0);case 10:case"end":return t.stop()}}),t,null,[[1,7]])})));function e(e,n){return t.apply(this,arguments)}return e}(),post:function(){var t=Object(k["a"])(regeneratorRuntime.mark((function t(e,n){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return window.console.log(1),t.prev=1,t.next=4,$.a.post(e,n);case 4:return t.abrupt("return",t.sent);case 7:return t.prev=7,t.t0=t["catch"](1),t.abrupt("return",t.t0);case 10:case"end":return t.stop()}}),t,null,[[1,7]])})));function e(e,n){return t.apply(this,arguments)}return e}()},C=n("5c96"),D=n.n(C);n("0fae");function O(t,e){window.Notification&&"denied"!=Notification.permission&&("granted"!=Notification.permission?Notification.requestPermission().then((function(n){"granted"==n&&O(t,e)})):e?setInterval((function(){t()}),e):t())}a["default"].use(D.a);n("b0c0");var B=n("8c4f"),S=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"form-login"},[n("el-form",{ref:"form",attrs:{model:t.form}},[n("el-form-item",[n("el-input",{attrs:{placeholder:"注册邮箱"},model:{value:t.form.email,callback:function(e){t.$set(t.form,"email",e)},expression:"form.email"}})],1),n("el-form-item",[n("el-input",{attrs:{placeholder:"密码",type:"password"},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}})],1),n("el-form-item",[n("el-row",[n("el-col",{attrs:{span:16}},[n("el-button",{attrs:{type:"primary",size:"mini"},on:{click:t.onLogin}},[t._v("登录")])],1),n("el-col",{attrs:{span:4}},[n("el-link",{attrs:{type:"success"},on:{click:function(e){return t.onJump("register")}}},[t._v("注册")])],1),n("el-col",{attrs:{span:4}},[n("el-link",{attrs:{type:"warning"},on:{click:function(e){return t.onJump("reset_password")}}},[t._v("找回密码")])],1)],1)],1)],1),n("About")],1)},E=[],I=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("el-collapse",{attrs:{value:"1"}},[n("el-collapse-item",{attrs:{title:"关于闪记卡",name:"1"}},[n("div",[t._v("重要的知识是需要理解后记忆的：")]),n("div",[t._v("理解一定是前提，我们提供了知识树供你梳理知识之间的关系；")]),n("div",[t._v("记忆是必要条件，我们提供了记忆周期提醒你去回忆知识点。")])])],1)},j=[],N={name:"About"},T=N,z=Object(f["a"])(T,I,j,!1,null,null,null),A=z.exports,J={components:{About:A},data:function(){return{form:{email:"",password:""}}},methods:{onLogin:function(){var t=this;""!=this.form.email?""!=this.form.password?this.$http.post("/auth/login",this.form).then((function(e){0==e.data.errno&&(c(e.data.data.email),window.location.href="/know"),t.$notify.error({title:"错误",message:e.data.msg})})).catch((function(t){window.console.log(t)})):this.$notify.error({title:"错误",message:"请输入密码"}):this.$notify.error({title:"错误",message:"请输入邮箱"})},onJump:function(t){this.$router.push({name:t})}}},L=J,M=(n("0953"),Object(f["a"])(L,S,E,!1,null,"3c74d301",null)),P=M.exports,R=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("el-tabs",{on:{"tab-click":t.handleClick},model:{value:t.activeName,callback:function(e){t.activeName=e},expression:"activeName"}},[n("el-tab-pane",{attrs:{label:"用户管理",name:"first"}},[t._v("用户管理")]),n("el-tab-pane",{attrs:{label:"配置管理",name:"second"}},[t._v("配置管理")]),n("el-tab-pane",{attrs:{label:"角色管理",name:"third"}},[t._v("角色管理")]),n("el-tab-pane",{attrs:{label:"定时任务补偿",name:"fourth"}},[t._v("定时任务补偿")])],1)},q=[],K={data:function(){return{activeName:"second"}},methods:{handleClick:function(t,e){window.console.log(t,e)}}},V=K,H=Object(f["a"])(V,R,q,!1,null,null,null),Q=H.exports,U=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"form-login"},[n("el-form",{ref:"form",attrs:{model:t.form}},[n("el-form-item",[n("el-input",{attrs:{placeholder:"邮箱地址"},model:{value:t.form.email,callback:function(e){t.$set(t.form,"email",e)},expression:"form.email"}})],1),n("el-form-item",[n("el-input",{attrs:{placeholder:"密码",type:"password"},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}})],1),n("el-form-item",[n("el-input",{attrs:{placeholder:"密码确认",type:"password"},model:{value:t.form.password2,callback:function(e){t.$set(t.form,"password2",e)},expression:"form.password2"}})],1),n("el-form-item",[n("el-row",[n("el-col",{attrs:{span:22}},[n("el-button",{attrs:{type:"primary",size:"mini"},on:{click:t.onSubmit}},[t._v("注册")])],1),n("el-col",{attrs:{span:2}},[n("el-link",{attrs:{type:"warning"},on:{click:function(e){return t.onJump("login")}}},[t._v("登录")])],1)],1)],1)],1),n("About")],1)},W=[],F=(n("ac1f"),n("5319"),{components:{About:A},data:function(){return{form:{email:"",password:"",password2:""}}},methods:{onSubmit:function(){var t=this;""!=this.form.email?this.form.password.length<6?this.$notify.error({title:"错误",message:"密码至少六位"}):this.form.password2.length<6?this.$notify.error({title:"错误",message:"验证密码至少六位"}):this.form.password==this.form.password2?this.$http.post("/auth/register",this.form).then((function(e){0!=e.data.errno?t.$notify.error({title:"错误",message:e.data.msg}):t.$notify({title:"成功",message:"我们已经将一封邮件发送到你的邮箱，请点击链接激活后使用。",type:"success",duration:0})})).catch((function(t){window.console.log(t)})):this.$notify.error({title:"错误",message:"两次密码需要一致"}):this.$notify.error({title:"错误",message:"请输入邮箱"})},onJump:function(t){this.$router.replace({name:t})}}}),G=F,X=(n("0709"),Object(f["a"])(G,U,W,!1,null,"18f99b8e",null)),Y=X.exports,Z=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"form-login"},[n("el-form",{ref:"form",attrs:{model:t.form}},[n("el-form-item",[n("el-input",{attrs:{placeholder:"邮箱地址"},model:{value:t.form.email,callback:function(e){t.$set(t.form,"email",e)},expression:"form.email"}})],1),1==t.step?n("el-form-item",[n("el-input",{attrs:{placeholder:"查看邮箱后得到 code"},model:{value:t.form.code,callback:function(e){t.$set(t.form,"code",e)},expression:"form.code"}})],1):t._e(),2==t.step?n("el-form-item",[n("el-input",{attrs:{placeholder:"密码",type:"password"},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}})],1):t._e(),2==t.step?n("el-form-item",[n("el-input",{attrs:{placeholder:"密码确认",type:"password"},model:{value:t.form.password2,callback:function(e){t.$set(t.form,"password2",e)},expression:"form.password2"}})],1):t._e(),n("el-form-item",[n("el-row",[n("el-col",{attrs:{span:22}},[0==t.step?n("el-button",{attrs:{type:"primary",size:"mini"},on:{click:t.onSendEmail}},[t._v("发送验证码")]):t._e(),1==t.step?n("el-button",{attrs:{type:"primary",size:"mini"},on:{click:t.onReset}},[t._v("验证验证码")]):t._e(),2==t.step?n("el-button",{attrs:{type:"primary",size:"mini"},on:{click:t.onReset}},[t._v("修改密码")]):t._e()],1),n("el-col",{attrs:{span:2}},[n("el-link",{attrs:{type:"warning"},on:{click:function(e){return t.onJump("login")}}},[t._v("登录")])],1)],1)],1)],1),n("About")],1)},tt=[],et={data:function(){return{step:0,form:{email:""}}},components:{About:A},methods:{onSubmit:function(){var t=this;""!=this.form.email?this.$http.post("/auth/reset_email",this.form).then((function(e){if(0==e.data.errno)return t.$notify({title:"成功",message:"我们已经将一封邮件发送到你的邮箱，请输入邮箱中的code",type:"success"}),void(t.step=1);t.$notify.error({title:"错误",message:e.data.msg})})).catch((function(t){window.console.log(t)})):this.$notify.error({title:"错误",message:"请输入邮箱"})},onJump:function(t){this.$router.replace({name:t})}}},nt=et,at=(n("f466"),Object(f["a"])(nt,Z,tt,!1,null,"3c2a3bdc",null)),it=at.exports,ot=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"fc-card-create"},[n("bread-crumb",{attrs:{breadcrumb:{items:[{name:"配置",path:""},{name:"卡片",path:"/card"},{name:"编辑",path:""}]}}}),n("div",{staticClass:"fc-card-create"},[n("el-form",[n("el-row",[n("el-col",{attrs:{span:14}},[n("el-row",[n("el-col",{attrs:{span:12}},[n("el-select",{staticClass:"f-width",attrs:{placeholder:"请选择记忆模板"},on:{change:t.onSelectChange},model:{value:t.templateID,callback:function(e){t.templateID=e},expression:"templateID"}},t._l(t.templates,(function(t){return n("el-option",{key:t.value,attrs:{label:t.label,value:t.value}})})),1)],1),n("el-col",{attrs:{span:12}},[n("el-select",{staticClass:"f-width",attrs:{placeholder:"请选择记忆曲线"},on:{change:t.onSelectChange},model:{value:t.remeberCurve,callback:function(e){t.remeberCurve=e},expression:"remeberCurve"}},t._l(t.remeberCurves,(function(t){return n("el-option",{key:t.value,attrs:{label:t.label,value:t.value}})})),1)],1)],1),n("el-row",[n("el-col",{attrs:{span:12}},[n("el-input",{attrs:{type:"textarea",placeholder:"卡片内容，单词模板可输入单测后双击完成编辑",rows:30},nativeOn:{dblclick:function(e){return t.onDataClick(e)}},model:{value:t.cardData,callback:function(e){t.cardData=e},expression:"cardData"}})],1),n("el-col",{attrs:{span:12}},[n("el-input",{attrs:{type:"textarea",placeholder:"内容模板",rows:30,disabled:!0},nativeOn:{dblclick:function(e){return t.onDemoClick(e)}},model:{value:t.templateDemoData,callback:function(e){t.templateDemoData=e},expression:"templateDemoData"}})],1)],1)],1),n("el-col",{attrs:{span:10}},[n("el-row",[n("div",{staticClass:"show",attrs:{id:"show-a"}},[n("div",{staticClass:"show-div"},[t._v("卡片正面")])])]),n("el-row",{staticClass:"show",attrs:{id:"show-b"}},[n("div",{staticClass:"show-div"},[t._v("卡片反面")])])],1)],1),n("el-row",[n("el-button",{staticStyle:{"margin-top":"20px"},attrs:{type:"primary",size:"small"},on:{click:t.onSubmit}},[t._v("保存")]),n("el-button",{staticStyle:{"margin-top":"20px"},attrs:{size:"small"},on:{click:function(e){return t.goBack(1)}}},[t._v("返回")])],1)],1)],1)],1)},rt=[],st=(n("c975"),n("ba4c")),lt=n.n(st),ct=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"fc-breadcrumb"},[n("el-breadcrumb",[t._l(t.breadcrumb.items,(function(e){return n("el-breadcrumb-item",{key:e.name},[n("router-link",{attrs:{to:e.path}},[t._v(t._s(e.name))])],1)})),t._l(t.breadcrumb.btns,(function(e){return n("el-breadcrumb-item",{key:e.name,staticClass:"nav_child"},[n("router-link",{attrs:{to:e.path}},[t._v(t._s(e.name))])],1)}))],2)],1)},dt=[],ut={name:"BreadCrumb",props:{breadcrumb:Object}},pt=ut,mt=(n("4f5f"),Object(f["a"])(pt,ct,dt,!1,null,"6c61e741",null)),ft=mt.exports,ht={name:"CardEdit",components:{BreadCrumb:ft},data:function(){var t=this,e=this;this.$http.get("/template").then((function(t){window.console.log(t);for(var n=[],a={},i=0;i<t.data.data.length;i++){var o=t.data.data[i];n.push({value:o.id,label:o.name}),a[o.id]=o,lt.a.component("a"+o.id,{props:{data:Object},template:o.content_a}),lt.a.component("b"+o.id,{props:{data:Object},template:o.content_b})}e.templates=n,e.templateMap=a,e.render()})).catch((function(t){window.console.log(t)}));var n={templates:[],templateMap:{},remeberCurves:[{value:1,label:"艾宾浩斯记忆曲线"}],templateDemoData:"",cardID:0,cardData:"",templateID:"",remeberCurve:1,knowID:0,loading:!1};return this.$route.params.card_id&&(n.cardID=parseInt(this.$route.params.card_id)),this.$route.query.know_id&&(n.knowID=parseInt(this.$route.query.know_id)),0!=n.cardID&&this.$http.get("/card/"+n.cardID).then((function(e){if(0!=e.data.errno)return t.$notify.error({title:"错误",message:e.data.msg}),void t.goBack();var n=e.data.data;t.cardData=n.data,t.templateID=n.template_id,t.knowID=n.know_id})).catch((function(t){window.console.log(t)})),n},methods:{render:function(){for(var t=["a","b"],e=0;e<t.length;e++){var n={};try{n=JSON.parse(this.cardData)}catch(r){return}var a=t[e],i=lt.a.component(a+this.templateID),o=new i({data:n});document.getElementById("show-"+a).innerHTML="<div></div>",o.$mount("#show-"+a+" > div")}},onSubmit:function(){var t=this;if(this.templateID)if(this.cardData){try{JSON.parse(this.cardData)}catch(n){return void this.$notify.error({title:"错误",message:"卡片内容不是合法的JSON格式"})}var e={id:this.cardID,know_id:this.knowID,template_id:this.templateID,data:this.cardData};this.$http.post("/card",e).then((function(e){0==e.data.errno?(t.$message({type:"success",message:"卡片保存成功"}),0!=t.cardID?t.goBack():t.cardData=t.templateMap[t.templateID].data_demo):t.$notify.error({title:"错误",message:e.data.msg})})).catch((function(t){window.console.log(t)}))}else this.$notify.error({title:"错误",message:"必须填写卡片内容"});else this.$notify.error({title:"错误",message:"必须选择一个模板"})},onSelectChange:function(){var t=this.templateMap[this.templateID];this.templateDemoData=t.data_demo+"\n\n双击文本框可快速将数据放入卡片内容"},onDemoClick:function(){this.cardData=this.templateMap[this.templateID].data_demo},onDataClick:function(){var t=this,e=this.cardData.toLowerCase();""!=e?-1==e.indexOf(" ")?(this.loading=!0,this.$http.get("/card/word/"+e).then((function(n){if(t.loading=!1,0==n.data.errno){var a={Word:e,Detail:n.data.data};t.cardData=JSON.stringify(a,null,4)}else t.$notify.error({title:"错误",message:n.data.msg})})).catch((function(t){window.console.log(t)}))):this.$notify.error({title:"错误",message:"当前只只能记忆单测，不能记忆句子"}):this.$notify.error({title:"错误",message:"单测不能为空"})},goBack:function(t){var e=this,n=2e3;t&&(n=0),setTimeout((function(){e.$router.push({name:"know_list",params:{know_id:e.knowID}})}),n)}},watch:{cardData:function(){this.render()}}},vt=ht,wt=(n("2232"),Object(f["a"])(vt,ot,rt,!1,null,"975b9c0c",null)),gt=wt.exports,bt=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"f-height"},[n("knowledge-tree",{staticClass:"c-ktree lr-border f-height",attrs:{know_id:t.know_id},on:{"node-selected":t.nodeSelected,"mode-switched":t.modeSwitched}}),n("div",{staticClass:"know-opt"},[n("el-button",{staticClass:"btn-add",attrs:{type:"success",icon:"el-icon-plus",circle:"",size:"mini",disabled:!t.know_id},on:{click:function(e){return t.onAddBtnClick()}}}),n("el-button",{staticClass:"btn-play",attrs:{icon:"el-icon-video-play",size:"mini",disabled:t.editing},on:{click:t.onPlayBtnClick}})],1),n("knowledge-points",{staticClass:"c-kpoints",attrs:{know_id:t.know_id,editing:t.editing}})],1)},_t=[],kt=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("el-row",[n("el-col",{attrs:{span:21}},[n("el-input",{attrs:{placeholder:"关键字过滤"},model:{value:t.filterText,callback:function(e){t.filterText=e},expression:"filterText"}})],1),n("el-col",{attrs:{span:3}},[n("div",{staticClass:"btn-opt"},[t.editing?n("el-button",{attrs:{icon:"el-icon-s-opportunity",type:"success",circle:"",size:"mini"},on:{click:function(e){return t.onModeChange(!1)}}}):n("el-button",{attrs:{type:"primary",icon:"el-icon-edit",circle:"",size:"mini"},on:{click:function(e){return t.onModeChange(!0)}}})],1)])],1),n("el-row",[n("el-tree",{ref:"tree",attrs:{data:t.data,"expand-on-click-node":!1,"filter-node-method":t.filterNode,"default-expanded-keys":t.treeExpandData,draggable:"","node-key":"id","current-node-key":t.know_id,"render-content":t.renderContent,"allow-drop":t.allowDrop,"allow-drag":t.allowDrag},on:{"node-drop":t.handleDrop,"node-click":t.handleNodeClick}})],1)],1)},yt=[],$t=(n("4de4"),n("c740"),n("a434"),n("a9e3"),{name:"KnowledgeTree",watch:{filterText:function(t){this.$refs.tree.filter(t)}},props:{know_id:Number},methods:{onModeChange:function(t){this.editing=t,window.console.log("mode",this.editing),this.data=JSON.parse(JSON.stringify(this.data)),this.$emit("mode-switched",this.editing)},filterNode:function(t,e){return!t||-1!==e.label.indexOf(t)},handleNodeClick:function(t){this.$emit("node-selected",t.id)},handleDrop:function(t,e,n){var a=this,i={type:"drag",id:t.data.id,rid:e.data.id,drop_type:n};window.console.log(i),this.$http.post("/know/modify",i).then((function(t){0==t.data.errno||a.$notify.error({title:"错误",message:t.data.msg})})).catch((function(t){window.console.log(t)}))},allowDrop:function(t,e,n){return!(0==e.data.id&&("prev"==n||"inner"==n))},allowDrag:function(t){return t.data.id},edit:function(t){var e=this,n=t.label;this.$prompt("请输入知识点名称","提示",{confirmButtonText:"确定",cancelButtonText:"取消",inputValue:n,inputValidator:function(t){return""!=t&&t!=n||"请输入和旧名字不一样的名字"}}).then((function(n){var a=n.value;e.$http.post("/know/modify",{type:"rename",id:t.id,name:a}).then((function(n){0==n.data.errno?t.label=a:e.$notify.error({title:"错误",message:n.data.msg})})).catch((function(t){window.console.log(t)}))})).catch((function(){}))},append:function(t){var e=this;this.treeExpandData=[t.id],this.$prompt("请输入子知识点名称","提示",{confirmButtonText:"确定",cancelButtonText:"取消",inputValidator:function(t){return!(!t||t.length<2||t.length>10)||"名字长度要求2-10"}}).then((function(n){var a=n.value;e.$http.post("/know",{pid:t.id,name:a}).then((function(n){if(0==n.data.errno){var i={id:n.data.data.id,label:a,children:[]};t.children||e.$set(t,"children",[]),t.children.push(i)}else e.$notify.error({title:"错误",message:n.data.msg})})).catch((function(t){window.console.log(t)}))})).catch((function(){}))},remove:function(t,e){var n=this;this.$confirm("知识点下的卡片将移动到「默认知识点」中，知识点将被删除,是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((function(){n.$http.post("/know/modify",{type:"delete",id:e.id}).then((function(a){if(0==a.data.errno){var i=t.parent,o=i.data.children||i.data,r=o.findIndex((function(t){return t.id===e.id}));o.splice(r,1)}else n.$notify.error({title:"错误",message:a.data.msg})})).catch((function(t){window.console.log(t)}))})).catch((function(){}))},renderContent:function(t,e){var n=this,a=e.node,i=e.data;return this.editing?t("span",{class:"ktree-node"},[t("span",[" ",a.label," "]),t("span",{class:"r-tree-btn"},[t("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(){return n.append(i)}}},[t("i",{class:"el-icon-circle-plus-outline"})]),t("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(){return n.edit(i)}}},[t("i",{class:"el-icon-edit"})]),t("el-button",{attrs:{size:"mini",type:"text",disabled:0==i.id},on:{click:function(){return n.remove(a,i)}}},[t("i",{class:"el-icon-remove-outline"})])])]):t("span",{class:"ktree-node"},[t("span",[" ",a.label," "]),t("span",{class:"r-tree-btn"},[i.learning,"/",i.total])])}},data:function(){var t=this;return this.$http.get("/know").then((function(e){t.data=[e.data.data]})).catch((function(t){window.console.log(t)})),{treeExpandData:[0],filterText:"",editing:!1,data:[],defaultProps:{children:"children",label:"label"}}}}),xt=$t,Ct=(n("1b2c"),Object(f["a"])(xt,kt,yt,!1,null,null,null)),Dt=Ct.exports,Ot=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("el-row",t._l(t.cards,(function(e){return n("el-col",{key:e.id,staticClass:"one-card",attrs:{span:8}},[n("div",{on:{click:function(n){return t.onCardClick(e.id)}}},[n("el-card",{attrs:{shadow:"hover"}},[n("div",{staticClass:"one-card-content"},[t._v(t._s(e.name))]),t.editing?n("div",{staticClass:"btn-opts"},[n("el-button",{attrs:{type:"primary",icon:"el-icon-edit",circle:"",size:"mini"},on:{click:function(n){return t.onEditBtnClick(e.id)}}}),n("el-button",{attrs:{type:"danger",icon:"el-icon-delete",circle:"",size:"mini"},on:{click:function(n){return t.onDelBtnClick(e.id)}}})],1):n("div",{staticClass:"one-card-ops"},[n("el-row",[n("el-col",{attrs:{span:17}},[n("el-progress",{attrs:{percentage:e.progress_rate}})],1),n("el-col",{attrs:{span:7}},[-1==e.next_learn.indexOf("前")?n("el-tag",{attrs:{size:"medium",type:"success"}},[t._v(t._s(e.next_learn))]):n("el-tag",{attrs:{size:"medium",type:"danger"}},[t._v(t._s(e.next_learn))])],1)],1)],1)])],1)])})),1)],1)},Bt=[],St=(n("4160"),n("159b"),{name:"CardList",props:{editing:Boolean,know_id:Number},methods:{onEditBtnClick:function(t){this.$router.push({name:"card_edit",params:{card_id:t}})},onCardClick:function(t){this.$router.push({name:"learning",params:{know_id:this.know_id},query:{cid:t}})},onDelBtnClick:function(t){var e=this,n=this;this.$confirm("卡片删除后不可恢复,是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((function(){e.$http.post("/card/"+t+"/del").then((function(a){if(0==a.data.errno){var i=new Array;n.cards.forEach((function(e){e.id!=t&&i.push(e)})),n.cards=i}else e.$notify.error({title:"错误",message:a.data.msg})})).catch((function(t){window.console.log(t)}))})).catch((function(){}))},updateData:function(t){var e=this;this.$http.get("/card/know/"+t).then((function(t){window.console.log(t),e.cards=t.data.data,e.formateList(),setInterval((function(){return e.formateList()}),6e4)})).catch((function(t){window.console.log(t)}))},formateList:function(){for(var t=Date.parse(new Date)/1e3,e=0;e<this.cards.length;e++){var n=this.cards[e];n.next_learn=this.formateDiff(n.next_learn_time,t,n.progress_rate)}this.cards=JSON.parse(JSON.stringify(this.cards))},formateDiff:function(t,e,n){if(100==n)return"完成";var a=t-e,i="后";if(a<0&&(i="前",a=-a),a<60)return"现在";var o=parseInt(a/3600);if(o>24)return"> 24h "+i;var r=parseInt(a/60)%60;return 0==o?r+"m"+i:(r<10&&(r="0"+r),o+"h"+r+"m"+i)}},data:function(){return this.know_id&&this.updateData(this.know_id),{cards:[],progressColors:[{color:"#90A19A",percentage:20},{color:"#e6a23c",percentage:40},{color:"#1989fa",percentage:60},{color:"#6f7ad3",percentage:80},{color:"#5cb87a",percentage:100}]}},watch:{know_id:function(t){this.updateData(t)}}}),Et=St,It=(n("1f5b"),Object(f["a"])(Et,Ot,Bt,!1,null,"6c1fa23c",null)),jt=It.exports,Nt={name:"Knowledge",components:{KnowledgeTree:Dt,KnowledgePoints:jt},data:function(){var t="";return this.$route.params.know_id&&(t=this.$route.params.know_id),{know_id:parseInt(t),editing:!1}},methods:{nodeSelected:function(t){this.know_id!=t&&(this.know_id=t,this.$router.push({name:"know_list",params:{know_id:this.know_id}}))},onPlayBtnClick:function(){this.$router.push({name:"learning",params:{know_id:this.know_id}})},modeSwitched:function(t){this.editing=t},onAddBtnClick:function(){this.$router.push({name:"card_edit",params:{card_id:0},query:{know_id:this.know_id}})}}},Tt=Nt,zt=(n("deee"),Object(f["a"])(Tt,bt,_t,!1,null,"ce202454",null)),At=zt.exports,Jt=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"l-panel",on:{keyup:[function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.onViewBtnClick(e)},function(e){return e.type.indexOf("key")||72===e.keyCode?t.onLearnBtnClick(0):null},function(e){return e.type.indexOf("key")||77===e.keyCode?t.onLearnBtnClick(1):null},function(e){return e.type.indexOf("key")||76===e.keyCode?t.onLearnBtnClick(2):null}]}},[n("div",{staticClass:"l-box"},[t._m(0),n("div",{staticClass:"l-opt"},[n("el-row",[n("el-col",{staticClass:"l-opt-l",attrs:{span:7}},[n("el-button",{attrs:{type:"primary",plain:"",icon:"el-icon-view",circle:"",size:"mini"},on:{click:t.onViewBtnClick}},[t._v("S")])],1),n("el-col",{staticClass:"align-center",attrs:{span:10}},[n("el-button",{attrs:{type:"danger",round:"",size:"mini"},on:{click:function(e){return t.onLearnBtnClick(2)}}},[t._v("陌生 L")]),n("el-button",{attrs:{type:"primary",round:"",size:"mini"},on:{click:function(e){return t.onLearnBtnClick(1)}}},[t._v("记得 M")]),n("el-button",{attrs:{type:"success",round:"",size:"mini"},on:{click:function(e){return t.onLearnBtnClick(0)}}},[t._v("熟悉 H")])],1),n("el-col",{staticClass:"l-opt-r align-right",attrs:{span:5}},[n("el-progress",{attrs:{"text-inside":!0,"stroke-width":25,status:"warning",percentage:parseInt(t.index/t.total*100),format:t.rateInfo}})],1),n("el-col",{staticClass:"l-opt-r align-right",attrs:{span:2}},[n("el-button",{attrs:{icon:"el-icon-close",circle:"",size:"mini"},on:{click:t.onExitBtnClick}},[t._v("退出")])],1)],1)],1)])])},Lt=[function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"l-content"},[n("div",{attrs:{id:"learning-content"}})])}],Mt={name:"KnowledgeLearning",components:{},methods:{render:function(t){var e="a";if(t||(e="b"),!(this.index>this.total)){var n=this.cards[this.index],a=lt.a.component(e+n.template_id),i=JSON.parse(n.data),o=new a({data:i});document.getElementById("learning-content").innerHTML="<div></div>",o.$mount("#learning-content > div")}},rateInfo:function(){return this.index+" / "+this.total},onExitBtnClick:function(){this.$router.push({name:"know_list",params:{know_id:this.$route.params.know_id}})},onViewBtnClick:function(){this.isA=!this.isA,this.render(this.isA)},onLearnBtnClick:function(t){var e=this;if(this.$http.post("/card/progress",{f:t,cid:this.cards[this.index].id}).catch((function(t){window.console.log(t)})),this.index++,this.index>=this.total){var n=this.$loading({lock:!0,text:"学习完成，跳转中",background:"rgba(0, 0, 0, 0.7)"});setTimeout((function(){n.close(),e.onExitBtnClick()}),2e3)}else this.isA=!0,this.render(this.isA)}},data:function(){var t=this,e=this,n=this.$route.query["cid"];return this.$http.get("/card/learning/"+this.$route.params.know_id+"?cid="+n).then((function(n){if(0!=n.data.errno)return t.$notify.error({title:"错误",message:n.data.msg}),void setTimeout((function(){t.onExitBtnClick()}),2e3);if(!n.data.data.cards||0==n.data.data.cards.length)return t.$notify.error({title:"错误",message:"没有需要学习的内容，即将返回"}),void setTimeout((function(){t.onExitBtnClick()}),2e3);t.cards=n.data.data.cards,t.templates=n.data.data.templates;for(var a=0;a<t.templates.length;a++){var i=t.templates[a];lt.a.component("a"+i.id,{props:{data:Object},template:i.content_a}),lt.a.component("b"+i.id,{props:{data:Object},template:i.content_b})}t.index=0,t.total=t.cards.length,t.isA=!0,e.render(!0,!0)})).catch((function(t){window.console.log(t)})),{index:0,total:1,isA:!0,cards:[],templates:[]}}},Pt=Mt,Rt=(n("a77f"),Object(f["a"])(Pt,Jt,Lt,!1,null,null,null)),qt=Rt.exports,Kt=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"fc-template-create lr-border"},[n("bread-crumb",{attrs:{breadcrumb:{items:[{name:"配置",path:""},{name:"卡片模板",path:"/template"},{name:"编辑",path:""}]}}}),n("div",{staticClass:"fc-temp-create"},[n("el-form",{ref:"template",attrs:{model:t.template,"label-width":"100px"}},[n("el-form-item",{attrs:{label:"模板名称"}},[n("el-input",{model:{value:t.template.name,callback:function(e){t.$set(t.template,"name",e)},expression:"template.name"}})],1),n("el-form-item",{attrs:{label:"正面内容"}},[n("el-row",[n("el-col",{attrs:{span:12}},[n("el-input",{attrs:{type:"textarea",placeholder:"<div class='width:100px;height:100px'>*</div>",rows:13},model:{value:t.template.content_a,callback:function(e){t.$set(t.template,"content_a",e)},expression:"template.content_a"}})],1),n("el-col",{staticClass:"render",attrs:{span:12}},[n("div",{attrs:{id:"pagea"}},[t._v("预览")])])],1)],1),n("el-form-item",{attrs:{label:"背面内容"}},[n("el-row",[n("el-col",{attrs:{span:12}},[n("el-input",{attrs:{type:"textarea",placeholder:"<div class='width:100px;height:100px'>*</div>",rows:13},model:{value:t.template.content_b,callback:function(e){t.$set(t.template,"content_b",e)},expression:"template.content_b"}})],1),n("el-col",{staticClass:"render",attrs:{span:12}},[n("div",{attrs:{id:"pageb"}},[t._v("预览")])])],1)],1),n("el-form-item",{attrs:{label:"测试数据"}},[n("el-row",[n("el-col",{attrs:{span:12}},[n("el-input",{attrs:{type:"textarea",placeholder:"{}",rows:10},model:{value:t.template.data_demo,callback:function(e){t.$set(t.template,"data_demo",e)},expression:"template.data_demo"}})],1),n("el-col",{attrs:{span:12}})],1)],1),n("el-form-item",{attrs:{label:"卡片名称"}},[n("el-row",[n("el-col",{attrs:{span:12}},[n("el-input",{attrs:{placeholder:"为测试数据的Key, 在创建卡片时给卡片取名"},model:{value:t.template.card_name_key,callback:function(e){t.$set(t.template,"card_name_key",e)},expression:"template.card_name_key"}})],1),n("el-col",{attrs:{span:12}})],1)],1),n("el-form-item",[n("el-button",{attrs:{type:"primary",size:"small"},on:{click:t.onSubmit}},[t._v("保存")])],1)],1)],1)],1)},Vt=[],Ht={name:"CardTemplateEdit",components:{BreadCrumb:ft},watch:{template:{handler:function(){this.onContentChange()},deep:!0}},data:function(){var t=this,e=this.$route.params.temp_id;if(0!=e){this.temp_id=e;var n=this;this.$http.get("/template/"+e).then((function(e){0!=e.data.errno||e.data.data.is_official?(t.$notify.error({title:"错误",message:e.data.msg}),setTimeout((function(){t.$router.push({name:"template_list"})}),2e3)):n.template=e.data.data})).catch((function(t){window.console.log(t)}))}return{temp_id:e,template:{name:"",content_a:'<div style="position:relative; width:100%; height:100%;">\n\t<div style="position:absolute; top:50%; left:50%; transform: translate(-50%, -50%);">\n\t\t{{Q}}\n\t</div>\n</div>',content_b:'<div style="position:relative; width:100%; height:100%;">\n\t<div style="position:absolute; top:50%; left:50%; transform: translate(-50%, -50%);">\n\t\t{{A}}\n\t</div>\n</div>',data_demo:'{\n\t"Q": "问题是什么",\n\t"A": "答案是什么"\n}'}}},methods:{onContentChange:function(){var t=this.template.data_demo,e=lt.a.extend({template:this.template.content_a,data:function(){return JSON.parse(t)}});document.getElementById("pagea").innerHTML="<div></div>",(new e).$mount("#pagea > div");var n=lt.a.extend({template:this.template.content_b,data:function(){return JSON.parse(t)}});document.getElementById("pageb").innerHTML="<div></div>",(new n).$mount("#pageb > div")},onSubmit:function(){var t=this,e={};try{e=JSON.parse(this.template.data_demo)}catch(i){return void this.$notify.error({title:"错误",message:"测试内容非合法JSON"})}var n=e[this.template.card_name_key];if(n&&"string"==typeof n){var a="/template";0!=this.temp_id&&(a=a+"/"+this.temp_id),this.$http.post(a,this.template).then((function(e){0!=e.data.errno?t.$notify.error({title:"错误",message:e.data.msg}):t.$router.push({name:"template_list"})})).catch((function(t){window.console.log(t)}))}else this.$notify.error({title:"错误",message:"卡片名称不是测试数据中的有效的key"})}}},Qt=Ht,Ut=(n("6d84"),Object(f["a"])(Qt,Kt,Vt,!1,null,"1d4c2e9a",null)),Wt=Ut.exports,Ft=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("bread-crumb",{attrs:{breadcrumb:{items:[{name:"配置",path:""},{name:"卡片模板",path:""}]}}}),n("el-row",[n("el-col",{staticClass:"t-row",attrs:{span:8}},[n("el-card",{staticClass:"t-add",attrs:{shadow:"hover"}},[n("el-button",{attrs:{type:"success",icon:"el-icon-plus",circle:"",size:"medium"},on:{click:function(e){return t.onEditBtnClick(0)}}})],1)],1),t._l(t.templates,(function(e){return n("el-col",{key:e.id,staticClass:"t-row",attrs:{span:8}},[n("el-card",{attrs:{shadow:"hover"}},[n("div",{staticClass:"t-card"},[t._v(" "+t._s(e.name)+" "),e.is_official?[t._v("· 官方")]:t._e()],2),n("div",{staticClass:"t-opt"},[n("el-button",{attrs:{type:"primary",icon:"el-icon-edit",circle:"",disabled:e.is_official,size:"mini"},on:{click:function(n){return t.onEditBtnClick(e.id)}}}),n("el-button",{attrs:{type:"danger",icon:"el-icon-delete",circle:"",disabled:e.is_official,size:"mini"},on:{click:function(n){return t.onDelBtnClick(e.id)}}})],1)])],1)}))],2)],1)},Gt=[],Xt={name:"CardTemplateList",props:{msg:String},methods:{onDelBtnClick:function(t){var e=this,n=this;this.$confirm("模板删除后将不能修改，也不能被新的卡片使用,是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((function(){e.$http.post("/template/"+t+"/del").then((function(a){if(0==a.data.errno){var i=new Array;n.templates.forEach((function(e){e.id!=t&&i.push(e)})),n.templates=i}else e.$notify.error({title:"错误",message:a.data.msg})})).catch((function(t){window.console.log(t)}))})).catch((function(){}))},onEditBtnClick:function(t){this.$router.push({name:"template_edit",params:{temp_id:t}})}},components:{BreadCrumb:ft},data:function(){var t=this;return this.$http.get("/template").then((function(e){window.console.log(e),t.templates=e.data.data})).catch((function(t){window.console.log(t)})),{templates:[]}}},Yt=Xt,Zt=(n("b4a6"),Object(f["a"])(Yt,Ft,Gt,!1,null,"45859630",null)),te=Zt.exports;a["default"].use(B["a"]);var ee=new B["a"]({mode:"history",routes:[{name:"main",path:"/",component:At},{name:"know",path:"/know",component:At},{name:"know_list",path:"/know/:know_id",component:At},{name:"card_edit",path:"/card/:card_id",component:gt},{name:"learning",path:"/learning/:know_id",component:qt},{name:"template_list",path:"/template",component:te},{name:"template_edit",path:"/template/:temp_id",component:Wt},{path:"/mine",component:Q},{name:"login",path:"/mine/login",component:P},{name:"register",path:"/mine/register",component:Y},{name:"reset_password",path:"/mine/reset",component:it}]});ee.beforeEach((function(t,e,n){"register"==t.name||"login"==t.name||"reset_password"==t.name||u()?n():n("/mine/login")})),a["default"].prototype.$ajax=x,a["default"].config.productionTip=!1,new a["default"]({el:"#app",router:ee,render:function(t){return t(_)}}),O((function(){this.$http.get("/card/progress").then((function(t){t.data.data&&t.data.data.count&&new window.Notification("闪记卡·熙穆",{body:"你有"+t.data.data.count+"张卡片需要学习了"})})).catch((function(t){window.console.log(t)}))}),6e4)},"6d84":function(t,e,n){"use strict";var a=n("2e91"),i=n.n(a);i.a},"7aac":function(t,e,n){"use strict";var a=n("088d"),i=n.n(a);i.a},"85ec":function(t,e,n){},"8b89":function(t,e,n){},a636:function(t,e,n){},a77f:function(t,e,n){"use strict";var a=n("8b89"),i=n.n(a);i.a},b44a:function(t,e,n){},b4a6:function(t,e,n){"use strict";var a=n("4054"),i=n.n(a);i.a},deee:function(t,e,n){"use strict";var a=n("a636"),i=n.n(a);i.a},f0d9:function(t,e,n){},f466:function(t,e,n){"use strict";var a=n("22f5"),i=n.n(a);i.a},fdf2:function(t,e,n){}});
//# sourceMappingURL=app.50e3b115.js.map