import { createLazyFileRoute } from '@tanstack/react-router'
import Register from "../pages/login/Register.tsx";

export const Route = createLazyFileRoute('/register')({
  component: () => <Register />
})