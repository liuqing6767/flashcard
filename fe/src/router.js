import Vue from "vue";
import VueRouter from "vue-router";
import Login from './components/auth/Login.vue';
import Mine from './components/auth/Mine.vue';
import Register from './components/auth/Register.vue';
import Reset from './components/auth/Reset.vue';
import CardEdit from './components/card/CardEdit.vue';
import Knowledge from './components/know/Knowledge.vue';
import KownLearning from './components/know/Learning.vue';
import CardTemplateEdit from './components/template/Edit.vue';
import CardTemplateList from './components/template/List.vue';
import {
    CheckUser
} from './plugins/auth';



// 要告诉 vue 使用 vueRouter
Vue.use(VueRouter);

const router = new VueRouter({
    mode: 'history',
    routes: [{
            name: "know",
            path: '/know', // 所有知识点列表
            component: Knowledge,
        },
        {
            name: "know_list",
            path: '/know/:know_id', // 所有知识点列表
            component: Knowledge,
        },
        {
            name: "card_edit",
            path: '/card/:card_id',
            component: CardEdit,
        },
        {
            name: "learning",
            path: '/learning/:know_id',
            component: KownLearning,
        },
        {
            name: "template_list",
            path: '/template',
            component: CardTemplateList
        },
        {
            name: "template_edit",
            path: '/template/:temp_id',
            component: CardTemplateEdit
        },
        {
            path: '/mine',
            component: Mine
        },
        {
            name: "login",
            path: '/mine/login',
            component: Login
        },
        {
            name: "register",
            path: '/mine/register',
            component: Register
        },
        {
            name: "reset_password",
            path: '/mine/reset',
            component: Reset
        },
    ],
})


router.beforeEach((to, from, next) => {
    if (to.name == "register" || to.name == "login" || to.name == "reset_password" || CheckUser()) {
        next()
        return
    }

    next("/mine/login")
})

export {
    router
};