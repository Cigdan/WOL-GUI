import {createFileRoute, redirect} from '@tanstack/react-router'
import Login from '../pages/login/Login.tsx'

export const Route = createFileRoute('/login')({
  beforeLoad: async () => {
    if (localStorage.getItem('isLoggedIn') === 'true'){
      throw redirect({
        to: '/dashboard',
        replace: true,
      })
    }
  },
  component: () => <Login />,
})