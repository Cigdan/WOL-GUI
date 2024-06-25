import axios from 'axios';
import {User} from "./Types.ts";

const api = axios.create({
  baseURL: 'http://localhost:8080',
  withCredentials: true,
})


async function login(user: User) {
    const response =  await api.post('/auth/login', {
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

export { login, createUser }