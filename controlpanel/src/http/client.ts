import axios from 'axios'
// @ts-ignore
// import{ Encrypt } from '../encrypt/index'
// import * as Encrypt from '../encrypt/myModule.browser';
//配置基础url、超时时间、post请求头

const client = axios.create({ 
    // 配置
    timeout : 5000,

})
// client.interceptors.request.use(function(config) {
//   if (config.method === 'post' || config.method === 'put') {
//     if(config.url =='/api/v1/assets/upload/image'){
//       console.log(config.params)
//     config.params = Encrypt(config.data);
//     console.log(config.params)
//     }else{
//       config.data = Encrypt(config.data);
//     }
//   } else if (config.method === 'get' || config.method === 'delete') {
//     config.params = Encrypt(config.params);
//   }
//   return config;
// });

export {
    client
}