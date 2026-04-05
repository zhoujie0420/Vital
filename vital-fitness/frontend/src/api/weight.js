import { get, post } from '../utils/request'

export const createWeight = (data) => post('/weights/', data)
export const getWeights = (params) => get('/weights/', params)
