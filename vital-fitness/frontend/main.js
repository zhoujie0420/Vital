import Vue from 'vue'
import App from './App'
import store from './store'

// 引入uView UI
import uView from 'uview-ui'

// 引入全局工具函数
import utils from './utils/index.js'

Vue.config.productionTip = false

Vue.use(uView)

Vue.prototype.$utils = utils

App.mpType = 'app'

const app = new Vue({
    store,
    ...App
})
app.$mount()
