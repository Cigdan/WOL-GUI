import {Button, Group, Paper, Stack, TextInput, Title} from "@mantine/core";
import {useForm, SubmitHandler} from "react-hook-form";
import {Device} from "../../../misc/Types.ts";
import {useMutation, useQueryClient} from "@tanstack/react-query";
import {addDevice} from "../../../misc/api.ts";
import Toast from "react-hot-toast";

function AddDevice() {
  const {
    register,
    handleSubmit,
    reset,
    formState: {errors},
  } = useForm<Device>({
    defaultValues: {
      name: null,
      mac_address: null,
      ip_address: null,
    }
  });
  const queryClient = useQueryClient();
  const addDeviceMutation = useMutation({
    mutationFn: (device: Device) => addDevice(device),
    onSuccess: async () => {
      await queryClient.invalidateQueries('devices')
      reset()
      Toast.success('Successfully added Device')
    },
    onError: (error) => {
      if (error.response) {
        Toast.error(error.response?.data.message)
      } else {
        Toast.error(error.message)
      }
    }
  })

  const onSubmit: SubmitHandler<Device> = data => {
    addDeviceMutation.mutate(data)
  }

  return (
      <Paper withBorder p={"lg"}>
        <Stack>
          <Title order={2}>Add Device</Title>
          <form onSubmit={handleSubmit(onSubmit)}>
            <Stack>
              <TextInput withAsterisk {...register("name", {
                required: "Name is required",
              })} error={errors.name && errors.name.message} label="Device Name"/>
              <Group grow>
                <TextInput withAsterisk {...register("mac_address", {
                  required: "Mac Address is required",
                  pattern: {value: /^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$/, message: "Invalid Mac Address"},
                })} error={errors.mac_address && errors.mac_address.message} label="Mac Address"/>
                <TextInput {...register("ip_address", {
                  pattern: {
                    value: /^([0-9]{1,3}\.){3}[0-9]{1,3}$/,
                    message: "Invalid IP Address"
                  },
                })} error={errors.ip_address && errors.ip_address.message} label="IP Address"/>
              </Group>
              <Button type="submit">
                Add Device
              </Button>
            </Stack>
          </form>
        </Stack>
      </Paper>
  );
}

export default AddDevice;