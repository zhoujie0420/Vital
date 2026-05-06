/**
 * 统一请求封装
 */

// 条件编译：H5 用相对路径（nginx 代理），小程序直接请求服务器
// #ifdef H5
const BASE_URL = '/api/v1'
// #endif
// #ifndef H5
const BASE_URL = 'https://api.bodysynclab.fun/api/v1'
// #endif

const request = (options) => {
  return new Promise((resolve, reject) => {
    const token = uni.getStorageSync('token')

    uni.request({
      url: BASE_URL + options.url,
      method: options.method || 'GET',
      data: options.data || {},
      header: {
        'Content-Type': 'application/json',
        ...(token ? { 'Authorization': 'Bearer ' + token } : {}),
        ...(options.header || {})
      },
      success: (res) => {
        if (res.statusCode === 401) {
          // token 过期，清除登录状态并提示，避免强制跳转打断用户浏览
          uni.removeStorageSync('token')
          uni.removeStorageSync('userInfo')
          uni.showModal({
            title: '登录提示',
            content: '登录已过期，是否重新登录？',
            confirmText: '去登录',
            cancelText: '再看看',
            success: (modalRes) => {
              if (modalRes.confirm) {
                uni.navigateTo({ url: '/pages/auth/login' })
              }
            }
          })
          reject(new Error('登录已过期'))
          return
        }

        if (res.statusCode >= 200 && res.statusCode < 300) {
          resolve(res.data)
        } else {
          const msg = (res.data && res.data.message) || '请求失败'
          uni.showToast({ title: msg, icon: 'none' })
          reject(new Error(msg))
        }
      },
      fail: (err) => {
        uni.showToast({ title: '网络异常，请稍后重试', icon: 'none' })
        reject(err)
      }
    })
  })
}

export const get = (url, data) => request({ url, method: 'GET', data })
export const post = (url, data) => request({ url, method: 'POST', data })
export const put = (url, data) => request({ url, method: 'PUT', data })
export const del = (url, data) => request({ url, method: 'DELETE', data })

export default request
