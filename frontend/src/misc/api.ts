import axios from 'axios';
import {User} from "./Types.ts";

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
  withCredentials: true,
})
// Response interceptor for handling 401 errors
api.interceptors.response.use(
    response => response,
    error => {
      if (error.response && error.response.status === 401) {
        if (!window.location.pathname.includes('/login') && !window.location.pathname.includes('/register')) {
          logout().then(() => {
            console.log('Unauthorized! Redirecting to login...');
          })
              .catch(() => {
                console.log('Error logging out');
              })
              .finally(() => {
                window.location.href = '/login';
              })
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

async function addDevice(device) {
  const response = await api.post('/devices', device)
  return response.data
}

async function editDevice(device) {
  const response = await api.put(`/devices/edit/${device.id}`, device)
  return response.data
}

async function deleteDevice(device) {
  const response = await api.delete(`/devices/delete/${device.id}`)
  return response.data
}

async function checkDeviceStatus(id) {
  const response = await api.get(`/devices/status/${id}`)
  return response.data
}

export {login, createUser, logout, getMyDevices, checkAuth, addDevice, editDevice, deleteDevice, checkDeviceStatus}