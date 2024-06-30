import {createRootRoute, Outlet} from '@tanstack/react-router'
import Sidebar from "../layouts/Sidebar.tsx";
import {Container, Group} from "@mantine/core";
import "../index.css"

const nonSidebarRoutes = ["login", "register"]

export const Route = createRootRoute({
  component: () => {
  if (nonSidebarRoutes.includes(window.location.pathname.split("/")[1])) {
          return <Outlet/>
      }
      return (
          <Group grow align={"start"} className={"container"}>
            <Sidebar/>
            <Container className={"content"} pt={"xl"} mt={"lg"} fluid>
              <Outlet/>
            </Container>

          </Group>
      )
    }
})
