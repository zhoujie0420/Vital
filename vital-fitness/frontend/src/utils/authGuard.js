/**
 * 登录守卫工具
 * Auth guard helper
 *
 * 用于在用户触发需要登录的功能时检查登录态。
 * 未登录时弹出确认对话框引导用户前往登录页，让用户可以先自由浏览体验功能服务。
 */

/**
 * 是否已登录
 * @returns {boolean}
 */
export const isLoggedIn = () => {
  return !!uni.getStorageSync('token')
}

/**
 * 检查登录态，未登录时弹窗引导登录
 *
 * @param {object} [options]
 * @param {string} [options.title='登录提示']
 * @param {string} [options.content='该功能需要登录后使用，是否前往登录？']
 * @param {string} [options.confirmText='去登录']
 * @param {string} [options.cancelText='再看看']
 * @returns {boolean} 已登录返回 true，未登录返回 false 并弹窗
 */
export const requireLogin = (options = {}) => {
  if (isLoggedIn()) return true

  const {
    title = '登录提示',
    content = '该功能需要登录后使用，是否前往登录？',
    confirmText = '去登录',
    cancelText = '再看看'
  } = options

  uni.showModal({
    title,
    content,
    confirmText,
    cancelText,
    success: (res) => {
      if (res.confirm) {
        uni.navigateTo({ url: '/pages/auth/login' })
      }
    }
  })

  return false
}

export default {
  isLoggedIn,
  requireLogin
}
