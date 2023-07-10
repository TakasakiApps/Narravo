import axios from 'axios'
import{ Encrypt } from '../encrypt/index'
//配置基础url、超时时间、post请求头

const client = axios.create({ 
    // 配置
    baseURL : 'http://127.0.0.1:8080',
    timeout : 5000,
    headers: {
        'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8',
      },
})
client.interceptors.request.use(function(config) {
  if (config.method === 'post' || config.method === 'put') {
    config.data = Encrypt(config.data);
  } else if (config.method === 'get' || config.method === 'delete') {
    config.params = Encrypt(config.params);
  }
  return config;
});

export {
    client
}