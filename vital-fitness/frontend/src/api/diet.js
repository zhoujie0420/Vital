import { get, post, del } from '../utils/request'

export const getFoods = (params) => get('/foods/', params)
export const createFood = (data) => post('/foods/', data)
export const createDietRecord = (data) => post('/diets/', data)
export const getDietRecords = (params) => get('/diets/', params)
export const deleteDietRecord = (id) => del(`/diets/${id}`)
