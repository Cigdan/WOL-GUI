import {Button, Group, Paper, Stack, Text, useMantineColorScheme} from '@mantine/core';
import {
  useMutation,
} from '@tanstack/react-query'
import classes from './sidebar.module.css';
import {Link, useNavigate, useRouterState} from "@tanstack/react-router";
import { House, Settings,LogOut, Sun, Moon   } from 'lucide-react';
import {logout} from '../misc/api.ts'
import Toast from 'react-hot-toast';

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

const themes = [
  {
    name: "light",
    icon: <Sun width={24}/>,
  },
  {
    name: "dark",
    icon: <Moon width={24}/>,
  },

]

function Sidebar() {
  const theme = useMantineColorScheme()
  const state = useRouterState()
  const navigate = useNavigate()
  const logoutMutation = useMutation({
    mutationFn: () => logout(),
    onSuccess: () => {
      Toast.success('Logout successful')
      navigate({to: "/login"})
    },
    onError: (error) => {
      if (error.response) {
        Toast.error(error.response?.data.message)
      }
      else {
        Toast.error(error.message)
      }
    }
  })

  return (
        <Paper className={classes.navbar} withBorder radius={0}>
          <Stack className={classes.navItems} justify={"space-between"} >
          <Stack gap="sm">
          {links.map((link) => {
            return (
                <Link to={link.href}>
                  <Button size="md" justify={"start"} variant={state.location.pathname.toLowerCase() === link.href.toLowerCase() ? "gradient" : "subtle" } className={classes.navButton}>
                    <Group>
                      {link.icon}
                      <Text>{link.name}</Text>
                    </Group>
                  </Button>
                </Link>
            )
          })}
          </Stack>
            <Stack>
              <Group grow>
                {themes.map((scheme) => {
                  return (
                      <Button variant={theme.colorScheme === scheme.name ? "gradient" : "subtle"} onClick={() => theme.setColorScheme(scheme.name)}>
                        {scheme.icon}
                      </Button>
                  )
                })}
              </Group>
              <Button size="md" variant={"subtle"} onClick={() => logoutMutation.mutate()}>
                <Group>
                  <LogOut width={24}/>
                  <Text>Logout</Text>
                </Group>
              </Button>
            </Stack>
          </Stack>
        </Paper>
  );
}

export default Sidebar;