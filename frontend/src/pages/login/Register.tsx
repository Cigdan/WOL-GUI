import {Paper, Stack, Container, TextInput, PasswordInput, Group, Button, Text} from '@mantine/core';
import {useForm, SubmitHandler} from "react-hook-form";
import {User} from '../../misc/types.ts'
import {
  useMutation,
} from '@tanstack/react-query'
import {createUser} from "../../misc/api.ts";
import {Link, useNavigate} from "@tanstack/react-router";

function Register() {
  const navigate = useNavigate()
  const {
    register,
    handleSubmit,
    formState: {errors},
  } = useForm<User>();

  const registerMutation = useMutation({
    mutationFn: (user) => createUser(user),
    onSuccess: () => {
      navigate({to: "/login"})
    }
  })

  const onSubmit: SubmitHandler<User> = data => {
    registerMutation.mutate(data)
  };

  return (
      <>
        <Container size={"xs"} mt="xl">
          <Paper withBorder shadow="xs" radius="md" p="xl" mt="xl">
            <form onSubmit={handleSubmit(onSubmit)}>
              <Stack>
                <h2>Create User</h2>
                <TextInput
                    radius="md"
                    label="Username"
                    placeholder="Enter your username"
                    {...register("username", {
                      required: "Username is required",
                      minLength: {value: 3, message: "Username is too short"},
                      maxLength: {value: 64, message: "Username is too long"},
                    })}
                    error={errors.username && errors.username.message}
                />
                <PasswordInput
                    radius="md"
                    label="Password"
                    placeholder="Enter your password"
                    {...register("password", {
                      required: "Password is required",
                      minLength: {value: 8, message: "Password has to be at least 8 Characters long"}
                    })}
                    error={errors.password && errors.password.message}
                />
                <Group justify="space-between">
                  <Link  to="/login">Already created a user?</Link>
                  <Button type="submit" disabled={registerMutation.isPending || registerMutation.isSuccess}>Create User</Button>
                </Group>
                {registerMutation.isError && <Text c="red" ta="right">{registerMutation.error.response?.data.message || registerMutation.error.message}</Text>}
              </Stack>
            </form>
          </Paper>
        </Container>
      </>
  );
}

export default Register;
