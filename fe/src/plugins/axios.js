import axios from 'axios';
import Vue from 'vue';

Vue.prototype.$http = axios

axios.defaults.baseURL = "http://10.99.202.18:8080/api"
//axios.defaults.baseURL = "http://localhost:8080/api"
axios.defaults.withCredentials = true;

export default {
    async get(path, params) {
        window.console.log(1)
        try {
            return await axios.get(path, params)
        } catch (err) {
            return err
        }
    },


    async post(path, params) {
        window.console.log(1)
        try {
            return await axios.post(path, params)
        } catch (err) {
            return err
        }
    }
}