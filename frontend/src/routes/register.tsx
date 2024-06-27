import {createFileRoute} from '@tanstack/react-router'
import Register from '../pages/login/Register.tsx'

export const Route = createFileRoute('/register')({
  component: () => <Register />,
})