// src/api/user.js

import axios from '~/axios';


// 获取客户端信息列表
export function apiGetClientList(data) {
  return axios.post('/api/v1/client/list', data);
}

// apiAddUser
// 添加客户端信息
export function apiAddClient(data) {
  return axios.post('/api/v1/client', data);
}

// apiUpdateUser
// 更新客户端信息
export function apiUpdateClient(data) {
  return axios.put('/api/v1/client', data);
}

// apiDeleteUser
// 删除客户端信息
export function apiDeleteClient(id) {
  return axios.delete(`/api/v1/client/${id}`);
}

export function apiCloseClient(id) {
  return axios.put(`/api/v1/client/close/${id}`);
}

export function apiOpenClient(id) {
  return axios.put(`/api/v1/client/open/${id}`);
}