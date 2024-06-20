import {Paper, Stack, Container, TextInput, PasswordInput, Group, Button, Anchor} from '@mantine/core';
import {useForm, SubmitHandler} from "react-hook-form";

type User = {
  username: string;
  password: string;
}

function Register() {
  const {
    register,
    handleSubmit,
    formState: {errors},
  } = useForm<User>();

  const onSubmit: SubmitHandler<User> = data => {
    console.log(data);
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
                  <Anchor href="login">Already created a user?</Anchor>
                  <Button type="submit">Create User</Button>
                </Group>
              </Stack>
            </form>
          </Paper>
        </Container>
      </>
  );
}

export default Register;
