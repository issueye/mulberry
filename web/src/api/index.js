import axios from '~/axios'


export function apiGetHomeCount() {
    return axios.get("/api/v1/admin/homeCount")
}