import axios from '~/axios'

export function login(username,password){
    return axios.post("/api/v1/auth/login",{
        username,
        password
    })
}

export function getinfo(){
    return axios.get("/api/v1/admin/userInfo")
}

export function logout(){
    return axios.post("/api/v1/auth/logout")
}

export function updatepassword(data){
    return axios.post("/api/v1/admin/updatepassword",data)
}