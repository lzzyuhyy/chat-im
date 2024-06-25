import App from './App'

// #ifndef VUE3
import Vue from 'vue'
import './uni.promisify.adaptor'
Vue.config.productionTip = false
App.mpType = 'app'
const app = new Vue({
  ...App
})
app.$mount()
// #endif

// #ifdef VUE3
import { createSSRApp } from 'vue'
export function createApp() {
  const app = createSSRApp(App)
  return {
    app
  }
}
// #endif

// 拦截器
uni.addInterceptor('request', {
  invoke(args) {
    // request 触发前拼接 url 
    args.url = 'http://127.0.0.1:8088/'+args.url
	
	// // 获取token
	// var token = uni.getStorageSync('token');
	// console.log(token);
	// token = ""
	// // token = "123456"
	// if (token == "") {
	// 	alert("登录状态失效，请重新登录")
	// 	uni.navigateTo({
	// 		url:"/pages/index/index"
	// 	})
	// }else {
	// 	// 获取成功设置到请求头
	// 	// args.header = args.header || {}
	// 	// args.header['token'] = token
	// }

	
  }
})