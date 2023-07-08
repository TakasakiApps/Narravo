import axios from 'axios'

//配置基础url、超时时间、post请求头

const client = axios.create({ 
    // 配置
    baseURL : 'http://127.0.0.1:8080',
    timeout : 5000,
    headers: {
        'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8',
      },
})

export {
    client
}