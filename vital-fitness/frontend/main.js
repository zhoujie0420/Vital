import App from './App'
import store from './store'

// 引入uView UI
import uView from 'uview-ui'

// 引入全局工具函数
import utils from './utils/index.js'

Vue.config.productionTip = false

// 使用uView UI
Vue.use(uView)

// 挂载全局工具函数
Vue.prototype.$utils = utils
// 挂载 store
Vue.prototype.$store = store

App.mpType = 'app'

const app = new Vue({
    store,
    ...App
})
app.$mount()
