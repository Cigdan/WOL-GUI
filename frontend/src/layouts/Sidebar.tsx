import {Button, Group, Paper, Stack, Text} from '@mantine/core';
import {
  useMutation,
} from '@tanstack/react-query'
import classes from './sidebar.module.css';
import {Link, useNavigate, useRouterState} from "@tanstack/react-router";
import { House, Settings,LogOut   } from 'lucide-react';
import {logout} from '../misc/api.ts'

const links = [
  {
    name: "Dashboard",
    href: "/dashboard",
    icon: <House width={24}/>,
  },
  {
    name: "Settings",
    href: "/settings",
    icon: <Settings width={24}/>,
  },
]

function Sidebar() {
  const state = useRouterState()
  const navigate = useNavigate()
  const logoutMutation = useMutation({
    mutationFn: () => logout(),
    onSuccess: () => {
      navigate({to: "/login"})
    }
  })
  return (
        <Paper className={classes.navbar} withBorder radius={false}>
          <Stack className={classes.navItems} justify={"space-between"} >
          <Stack gap="sm">
          {links.map((link) => {
            return (
                <Link to={link.href}>
                  <Button size="md" justify={"start"} variant={state.location.pathname.toLowerCase() === link.href.toLowerCase() ? "filled" : "subtle" } className={classes.navButton}>
                    <Group>
                      {link.icon}
                      <Text>{link.name}</Text>
                    </Group>
                  </Button>
                </Link>
            )
          })}
          </Stack>
            <Button size="md" variant={"subtle"} onClick={() => logoutMutation.mutate()}>
              <Group>
                <LogOut width={24}/>
                <Text>Logout</Text>
              </Group>
            </Button>
          </Stack>
        </Paper>
  );
}

export default Sidebar;