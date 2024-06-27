import { createFileRoute } from '@tanstack/react-router'
import Dashboard from "../pages/Dashboard/Dashboard.tsx";
import PrivateRoute from "../misc/PrivateRoute.tsx";

export const Route = createFileRoute('/dashboard')({
  component: () => (
      <PrivateRoute Component={Dashboard} />
  )
})