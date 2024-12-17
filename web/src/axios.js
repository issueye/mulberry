import axios from "axios"
import { toast } from '~/composables/util'
import { getToken } from '~/composables/auth'
import { useUserStore } from '~/store' // 导入 Pinia store

const service = axios.create({
    // baseURL: 'http://192.168.224.223:6677', // 指向后端服务器
    timeout: 8000,
});

// 添加请求拦截器
service.interceptors.request.use(
    function (config) {
        // 往header头自动添加token
        const token = getToken()
        if (token) {
            config.headers["Authorization"] = token
        }

        return config;
    },
    function (error) {
        // 对请求错误做些什么
        return Promise.reject(error);
    }
);

// 添加响应拦截器
service.interceptors.response.use(
    function (response) {
        // 对响应数据做点什么
        const code = response.data.code;
        if (code != 200) {
            toast(response.data.message, "error")
        }

        return response.data.data;
    },
    function (error) {
        console.log('error', error)
        const msg = error.response.data.message || "请求失败"

        if (msg == "非法token，请先登录！") {
            const userStore = useUserStore() // 使用 Pinia store
            userStore.logoutUser().finally(() => location.reload())
        }

        toast(msg, "error")

        return Promise.reject(error);
    })

export default service