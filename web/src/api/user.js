// src/api/user.js

import axios from '~/axios';

export function uploadAvatar(data) {
  return axios.post('/api/v1/admin/upload', data, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
}

export function updateUserInfo(data) {
  return axios.post('/api/v1/admin/updateuserinfo', data);
}


// 获取用户列表
export function apiGetUserList(data) {
  return axios.post('/api/v1/user/list', data);
}

// apiAddUser
// 添加用户
export function apiAddUser(data) {
  return axios.post('/api/v1/user/add', data);
}

// apiUpdateUser
// 更新用户
export function apiUpdateUser(data) {
  return axios.put('/api/v1/user/update', data);
}

// apiDeleteUser
// 删除用户
export function apiDeleteUser(id) {
  return axios.delete(`/api/v1/user/delete/${id}`);
}
