import axios from 'axios';
import {User} from "./Types.ts";
import {useNavigate} from "@tanstack/react-router";

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
  withCredentials: true,
})
// Response interceptor for handling 401 errors
api.interceptors.response.use(
    response => response,
    error => {
      if (error.response && error.response.status === 401) {
        if (!window.location.pathname.includes('/login')  && !window.location.pathname.includes('/register')) {
          console.log('Unauthorized! Redirecting to login...');
          const navigate = useNavigate()
          navigate({to: "/login"})
        }
      }
      return Promise.reject(error);
    }
);

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
}

async function checkAuth() {
  const response = await api.get('/auth/check')
  return response.data
}

async function getMyDevices() {
  const response = await api.get('/devices')
  return response.data
}

export { login, createUser, logout, getMyDevices, checkAuth }