import { get, post, put, del } from '../utils/request'

// 获取动作库
export const getExercises = (category) => get('/exercises/', { category })

// 创建自定义动作
export const createExercise = (data) => post('/exercises/', data)

// 创建训练记录
export const createWorkout = (data) => post('/workouts/', data)

// 获取训练记录列表
export const getWorkouts = (params) => get('/workouts/', params)

// 获取单条训练记录
export const getWorkout = (id) => get(`/workouts/${id}`)

// 更新训练记录
export const updateWorkout = (id, data) => put(`/workouts/${id}`, data)

// 删除训练记录
export const deleteWorkout = (id) => del(`/workouts/${id}`)
