import { createFileRoute } from '@tanstack/react-router'
import Dashboard from "../pages/Dashboard/Dashboard.tsx";

export const Route = createFileRoute('/dashboard')({
  component: () => <Dashboard />
})