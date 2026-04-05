import { get, post } from '../utils/request'

export const createMood = (data) => post('/moods/', data)
export const getMoods = (params) => get('/moods/', params)
