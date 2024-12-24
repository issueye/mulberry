import axios from '~/axios';

// 获取端口号信息列表
export function apiGetPortList(data) {
    return axios.post('/api/v1/port/list', data);
}

// apiAddUser
// 添加端口号信息
export function apiAddPort(data) {
    return axios.post('/api/v1/port', data);
}

// apiUpdateUser
// 更新端口号信息
export function apiUpdatePort(data) {
    return axios.put('/api/v1/port', data);
}

// apiDeleteUser
// 删除端口号信息
export function apiDeletePort(id) {
    return axios.delete(`/api/v1/port/${id}`);
}