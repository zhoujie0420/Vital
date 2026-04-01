/**
 * 用户相关 API
 */
import { post, get } from '../utils/request'

// 用户注册
export const register = (data) => post('/auth/register', data)

// 用户登录
export const login = (data) => post('/auth/login', data)

// 用户登出
export const logout = () => post('/auth/logout')

// 获取用户资料
export const getProfile = () => get('/users/profile')
