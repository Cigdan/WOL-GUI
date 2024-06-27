import {createFileRoute} from '@tanstack/react-router'
import Login from '../pages/login/Login.tsx'

export const Route = createFileRoute('/login')({
  component: () => <Login />
})