import {Paper, Stack, Container, TextInput, PasswordInput, Group, Button, Anchor} from '@mantine/core';

function Login() {
  return (
      <>
        <Container size={"xs"} mt="xl" >
          <Paper withBorder shadow="xs" radius="md" p="xl" mt="xl">
            <Stack>
              <h2>Login</h2>
              <TextInput
                  radius="md"
                  label="Username"
                  placeholder="Enter your username"
              />
              <PasswordInput
                  radius="md"
                  label="Password"
                  placeholder="Enter your password"
              />
              <Group justify="space-between">
                <Anchor href="register">Create a new account</Anchor>
                <Button>Login</Button>
              </Group>
            </Stack>
          </Paper>
        </Container>
      </>
  );
}

export default Login;