import {
  Button,
  Group,
  Paper,
  Stack,
  Text,
  Title,
  useMantineColorScheme,
} from "@mantine/core";
import { useMutation } from "@tanstack/react-query";
import {
  House,
  Settings,
  LogOut,
  Sun,
  Moon,
  Power,
  Menu,
} from "lucide-react";
import { logout } from "../misc/api.ts";
import Toast from "react-hot-toast";
import { useState, useEffect, useRef } from "react";
import classes from "./sidebar.module.css";
import { Link, useNavigate, useRouterState } from "@tanstack/react-router";

const links = [
  {
    name: "Dashboard",
    href: "/dashboard",
    icon: <House width={24} />,
  },
  {
    name: "Settings",
    href: "/settings",
    icon: <Settings width={24} />,
  },
];

const themes = [
  {
    name: "light",
    icon: <Sun width={24} />,
  },
  {
    name: "dark",
    icon: <Moon width={24} />,
  },
];

function Sidebar() {
  const theme = useMantineColorScheme();
  const state = useRouterState();
  const [navbarOpen, setNavbarOpen] = useState(false);
  const navigate = useNavigate();
  const navbarRef = useRef(null);

  const logoutMutation = useMutation({
    mutationFn: () => logout(),
    onSuccess: () => {
      Toast.success("Logout successful");
      navigate({ to: "/login" });
    },
    onError: (error) => {
      if (error.response) {
        Toast.error(error.response?.data.message);
      } else {
        Toast.error(error.message);
      }
    },
  });

  const handleClickOutside = (event) => {
    if (navbarRef.current && !navbarRef.current.contains(event.target)) {
      setNavbarOpen(false);
    }
  };

  useEffect(() => {
    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  return (
      <>
        <Button
            variant={"outline"}
            className={`${classes.toggleButton} ${navbarOpen && classes.buttonHidden}`}
            onClick={() => setNavbarOpen(!navbarOpen)}
        >
          <Menu size={24} />
        </Button>
        <Paper
            ref={navbarRef}
            className={`${classes.navbar} ${navbarOpen && classes.open}`}
            withBorder
            radius={0}
        >
          <Stack className={classes.navItems} justify={"space-between"}>
            <Stack gap="sm">
              <Group justify={"center"} my={"sm"}>
                <Power size={24} />
                <Title order={2}>Wake on Lan</Title>
              </Group>
              {links.map((link) => (
                  <Link key={link.href} to={link.href} onClick={() => setNavbarOpen(false)}>
                    <Button
                        size="md"
                        justify={"start"}
                        variant={
                          state.location.pathname.toLowerCase() ===
                          link.href.toLowerCase()
                              ? "gradient"
                              : "subtle"
                        }
                        className={classes.navButton}
                    >
                      <Group>
                        {link.icon}
                        <Text>{link.name}</Text>
                      </Group>
                    </Button>
                  </Link>
              ))}
            </Stack>
            <Stack>
              <Group grow>
                {themes.map((scheme) => (
                    <Button
                        key={scheme.name}
                        variant={
                          theme.colorScheme === scheme.name ? "gradient" : "subtle"
                        }
                        onClick={() => theme.setColorScheme(scheme.name)}
                    >
                      {scheme.icon}
                    </Button>
                ))}
              </Group>
              <Button
                  size="md"
                  variant={"subtle"}
                  onClick={() => logoutMutation.mutate()}
              >
                <Group>
                  <LogOut width={24} />
                  <Text>Logout</Text>
                </Group>
              </Button>
            </Stack>
          </Stack>
        </Paper>
      </>
  );
}

export default Sidebar;
