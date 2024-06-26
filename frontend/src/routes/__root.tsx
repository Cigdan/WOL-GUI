import {createRootRoute, Outlet} from '@tanstack/react-router'
import Sidebar from "../layouts/Sidebar.tsx";
import {Group} from "@mantine/core";

export const Route = createRootRoute({
  component: () => {
    if (localStorage.getItem('isLoggedIn') === 'true') {
      return (
          <Group grow>
            <Sidebar/>
            <Outlet/>
          </Group>

      )
    }
  }
})
