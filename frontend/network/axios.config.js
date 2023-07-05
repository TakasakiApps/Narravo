import axios from 'axios'


//配置基础url、超时时间、post请求头
axios.defaults.baseURL = 'http://xxx.xx.xx.xxx:xxxx';
axios.defaults.timeout = 5000;
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
