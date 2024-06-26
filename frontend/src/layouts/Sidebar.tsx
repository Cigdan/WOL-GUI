import {Button, Group, Paper, Stack, Text} from '@mantine/core';
import classes from './sidebar.module.css';
import {Link, useRouterState} from "@tanstack/react-router";
import { House, Settings  } from 'lucide-react';

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
  return (
        <Paper className={classes.navbar} withBorder radius={false}>
          <Stack className={classes.navItems} gap="sm">
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
        </Paper>
  );
}

export default Sidebar;