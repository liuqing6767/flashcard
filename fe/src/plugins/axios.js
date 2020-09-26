import axios from 'axios';
import Vue from 'vue';

Vue.prototype.$http = axios

axios.defaults.baseURL = "/api"
axios.defaults.withCredentials = true;

axios.interceptors.response.use(rsp => {
    if (rsp.data.errno == 401) {
        window.location.href = "/mine/login"
    }
    return rsp
})

export default {
    async get(path, params) {
        try {
            return await axios.get(path, params)
        } catch (err) {
            return err
        }
    },

    async post(path, params) {
        try {
            return await axios.post(path, params)
        } catch (err) {
            return err
        }
    }
}
