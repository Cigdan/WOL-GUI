import {createFileRoute, redirect} from '@tanstack/react-router'
import Register from '../pages/login/Register.tsx'

export const Route = createFileRoute('/register')({
  beforeLoad: async () => {
    if (localStorage.getItem('isLoggedIn') === 'true'){
      throw redirect({
        to: '/dashboard',
        replace: true,
      })
    }
  },
  component: () => <Register />,
})