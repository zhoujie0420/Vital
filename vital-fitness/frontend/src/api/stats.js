import { get } from '../utils/request'

export const getDashboard = () => get('/dashboard')
export const getStats = (period) => get('/stats', { period })
