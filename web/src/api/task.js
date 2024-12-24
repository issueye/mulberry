// src/api/user.js

import axios from '~/axios';


// 获取用户列表
export function apiGetTaskList(data) {
  return axios.post('/api/v1/task/list', data);
}

// apiAddUser
// 添加用户
export function apiAddTask(data) {
  return axios.post('/api/v1/task', data);
}

// apiUpdateUser
// 更新用户
export function apiUpdateTask(data) {
  return axios.put('/api/v1/task', data);
}

// apiDeleteUser
// 删除用户
export function apiDeleteTask(id) {
  return axios.delete(`/api/v1/task/${id}`);
}


export function apiSaveCode(data) {
  return axios.put(`/api/v1/task/save_code`, data);
}

export function apiUpdateTaskStatus(id) {
  return axios.put(`/api/v1/task/updateStatus/${id}`);
}

export function apiRunTask(id) {
  return axios.get(`/api/v1/task/run/${id}`);
}

export function apiGetTaskLogList(data) {
  return axios({
    url: '/api/v1/task/log/list',
    method: 'post',
    data: { ...data }
  })
}