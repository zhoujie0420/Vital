/**
 * 全局工具函数
 */

// 格式化日期
export const formatDate = (date) => {
  if (!date) return ''
  const d = new Date(date)
  const year = d.getFullYear()
  const month = String(d.getMonth() + 1).padStart(2, '0')
  const day = String(d.getDate()).padStart(2, '0')
  return `${year}-${month}-${day}`
}

// 格式化日期时间
export const formatDateTime = (date) => {
  if (!date) return ''
  const d = new Date(date)
  return `${formatDate(d)} ${String(d.getHours()).padStart(2, '0')}:${String(d.getMinutes()).padStart(2, '0')}`
}

// 检查是否已登录
export const isLoggedIn = () => {
  return !!uni.getStorageSync('token')
}

export default {
  formatDate,
  formatDateTime,
  isLoggedIn
}
