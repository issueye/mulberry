import axios from '~/axios';


// 获取用户列表
export function apiGetMenuList(data) {
  return axios.post('/api/v1/menu/list', data);
}

// apiAddUser
// 添加用户
export function apiAddMenu(data) {
  return axios.post('/api/v1/menu/add', data);
}

// apiUpdateUser
// 更新用户
export function apiUpdateMenu(data) {
  return axios.put('/api/v1/menu/update', data);
}

// apiDeleteUser
// 删除用户
export function apiDeleteMenu(id) {
  return axios.delete(`/api/v1/menu/delete/${id}`);
}

export function apiGetRoleMenuList(id) {
  return axios.get(`/api/v1/menu/roleMenus/${id}`);
}

export function apiGetCatalog() {
  return axios.get(`/api/v1/menu/catalog`);
}

// roleMenus
export function apiSaveRoleMenus(id, data) {
  return axios.post(`/api/v1/menu/saveRoleMenus/${id}`, data);
}