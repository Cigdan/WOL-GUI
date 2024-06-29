import {createRootRoute, Outlet} from '@tanstack/react-router'
import Sidebar from "../layouts/Sidebar.tsx";
import {Group} from "@mantine/core";

const nonSidebarRoutes = ["login", "register"]

export const Route = createRootRoute({
  component: () => {
  if (nonSidebarRoutes.includes(window.location.pathname.split("/")[1])) {
          return <Outlet/>
      }
      return (
          <Group grow align={"start"}>
            <Sidebar/>
            <Outlet/>
          </Group>
      )
    }
})
