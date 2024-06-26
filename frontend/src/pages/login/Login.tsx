import {Paper, Stack, Container, TextInput, PasswordInput, Group, Button, Loader, Text} from '@mantine/core';
import {
  useMutation,
} from '@tanstack/react-query'
import {Link, useNavigate} from "@tanstack/react-router";
import {useForm, SubmitHandler} from "react-hook-form";
import {login} from '../../misc/api.ts'
import {User} from '../../misc/types.ts'

function Login() {
  const navigate = useNavigate()
  const loginMutation = useMutation({
    mutationFn: (user) => login(user),
    onSuccess: () => {
      localStorage.setItem("isLoggedIn", "true")
      navigate({to: "/dashboard"})
    }
  })

  const {
    register,
    handleSubmit,
    formState: {errors},
  } = useForm<User>();

  const onSubmit: SubmitHandler<User> = data => {
    loginMutation.mutate(data)
  };

  return (
      <>
        <Container size={"xs"} mt="xl">
          <Paper withBorder shadow="xs" radius="md" p="xl" mt="xl">
            <form onSubmit={handleSubmit(onSubmit)}>
              <Stack>
                <h2>Login</h2>
                <TextInput
                    radius="md"
                    label="Username"
                    placeholder="Enter your username"
                    {...register("username", {
                      required: "Username is required",
                    })}
                    error={errors.username && errors.username.message}
                />
                <PasswordInput
                    radius="md"
                    label="Password"
                    placeholder="Enter your password"
                    {...register("password", {
                      required: "Password is required",
                    })}
                    error={errors.password && errors.password.message}
                />
                <Group justify="space-between">
                  <Link to="/register">Create new user</Link>
                  <Button type="submit" disabled={loginMutation.isPending || loginMutation.isSuccess}> {
                    loginMutation.isPending ? <Loader color="white" type="dots" /> : 'Login'
                  }</Button>
                </Group>
                {loginMutation.isError && <Text c="red" ta="right">{loginMutation.error.response ? loginMutation.error.response?.message : loginMutation.error.message}</Text>}
              </Stack>
            </form>
          </Paper>
        </Container>
      </>
  );
}

export default Login;