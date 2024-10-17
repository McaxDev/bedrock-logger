import axios from 'axios'
import cookie from 'js-cookie'

// 创建通用的 axios 配置
const createAxiosInstance = (baseURL) => {
    const httpx = axios.create({
        baseURL: baseURL,
        timeout: 50000,
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer yourToken'
        }
    })

    httpx.interceptors.request.use(config => {
        const token = cookie.get('X-bedrock-logger-Token')
        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`  // 将 Token 添加到请求头
        }
        return config
    }, error => {
        return Promise.reject(error)
    })

    httpx.interceptors.response.use(response => {
        const newToken = response.headers['x-bedrock-logger-token']
        if (newToken) {
            cookie.set('X-bedrock-logger-Token', newToken)
        }
        return response
    }, error => {
        if (error.response.status === 401) {
            cookie.remove('X-bedrock-logger-Token')
        } else {
            return Promise.reject(error)
        }
    })

    return httpx
}

const rootDomain = {
    target: 'axtl.cn:520',
    // target: 'mcax.cn',
}

const http = createAxiosInstance(`https://api.${rootDomain.target}/bedrock`)

export { http }