import axios from 'axios';
import {User} from "./Types.ts";

const api = axios.create({
  baseURL: 'http://localhost:8080',
  withCredentials: true,
})

/* api.interceptors.response.use(async (response) => {
  if (response.status === 401) {
    if (localStorage.getItem('isLoggedIn') === 'true') {
      await logout()
    }
  }
}) */


async function login(user: User) {
    const response = await api.post('/auth/login', {
      username: user.username,
      password: user.password
    })
    return response.data
}

async function createUser(user: User) {
  const response = await api.post("/auth/register", {
    username: user.username,
    password: user.password
  })
  return response.data
}

async function logout() {
  await api.post('/auth/logout')
  localStorage.setItem('isLoggedIn', 'false')
}

async function getMyDevices() {
  const response = await api.get('/devices')
  return response.data
}

export { login, createUser, logout, getMyDevices }