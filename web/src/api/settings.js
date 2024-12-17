import axios from '~/axios';

/**
 * 获取系统设置
 * @returns 
 */
export function apiGetSystemSettings() {
  return axios.get('/api/v1/settings/system');
}

/**
 * 更新系统设置
 * @param {*} data 
 * @returns 
 */
export function apiSetSystemSettings(data) {
  return axios.put('/api/v1/settings/system', data);
}

/**
 * 获取日志设置
 * @returns 
 */
export function apiGetLoggerSettings() {
  return axios.get('/api/v1/settings/logger');
}

/**
 * 更新日志设置
 * @param {*} data
 * @returns
 */
export function apiSetLoggerSettings(data) {
  return axios.put('/api/v1/settings/logger', data);
}