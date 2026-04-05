import { post, get } from '../utils/request'

// 微信登录
export const wxLogin = (code) => post('/auth/wx-login', { code })

// 用户登出
export const logout = () => post('/auth/logout')

// 获取用户资料
export const getProfile = () => get('/users/profile')
