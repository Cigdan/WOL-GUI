import {Button, Group, Modal, Stack, TextInput, Popover, Text, Loader} from "@mantine/core";
import {Device} from "../../../misc/Types.ts";
import {useForm} from "react-hook-form";
import {useMutation, useQueryClient} from "@tanstack/react-query";
import {editDevice, deleteDevice} from "../../../misc/api.ts";
import Toast from "react-hot-toast";

type EditDeviceModalProps = {
  opened: boolean;
  close: () => void;
  setDeviceToEdit: (device: Device | null) => void;
  device: Device;
}

function EditDeviceModal(props: EditDeviceModalProps) {
  const {opened, close, device} = props;
  const queryClient = useQueryClient();

  const editDeviceMutation = useMutation({
    mutationFn: (device: Device) => editDevice(device),
    onSuccess: async () => {
      reset()
      closeModal()
      Toast.success('Successfully edited Device')
      await queryClient.invalidateQueries('devices')
    },
    onError: (error) => {
      if (error.response) {
        Toast.error(error.response?.data.message)
      } else {
        Toast.error(error.message)
      }
    }
  })

  const deleteDeviceMutation = useMutation({
    mutationFn: (device: Device) => deleteDevice(device),
    onSuccess: async () => {
      reset()
      closeModal()
      Toast.success('Successfully deleted Device')
      await queryClient.invalidateQueries('devices')
    },
    onError: (error) => {
      if (error.response) {
        Toast.error(error.response?.data.message)
      } else {
        Toast.error(error.message)
      }
    }
  })

  const {
    register,
    handleSubmit,
    reset,
    formState: {errors},
  } = useForm<Device>({
    defaultValues: {
      id: device.id,
      name: device.name,
      mac_address: device.mac_address,
      ip_address: device.ip_address,
    }
  });

  const onSubmit = (data: Device) => {
    editDeviceMutation.mutate(data)
  }

  function closeModal() {
    close();
    props.setDeviceToEdit(null);
  }

  return (
      <Modal opened={opened} onClose={() => {
        closeModal();
      }} title="Edit Device" centered>
        <form onSubmit={handleSubmit(onSubmit)}>
          <Stack gap={"md"}>
            <Stack>
              <TextInput withAsterisk {...register("name", {
                required: "Name is required",
              })} error={errors.name && errors.name.message} label="Device Name"/>
                <TextInput withAsterisk {...register("mac_address", {
                  required: "Mac Address is required",
                  pattern: {value: /^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$/, message: "Invalid Mac Address"},
                })} error={errors.mac_address && errors.mac_address.message} label="Mac Address"/>

              <TextInput {...register("ip_address", {
                pattern: {
                  value: /^([0-9]{1,3}\.){3}[0-9]{1,3}$/,
                  message: "Invalid IP Address"
                },
              })} error={errors.ip_address && errors.ip_address.message}
                         label="IP Address (Optional for checking status)"/>
              <Group grow>
                <Popover>
                  <Popover.Target>
                    <Button variant={"subtle"} color={"red"}>
                      Delete
                    </Button>
                  </Popover.Target>
                  <Popover.Dropdown>
                    <Stack>
                      <Text>Do you want to delete this device?</Text>
                      <Button disabled={deleteDeviceMutation.isPending} onClick={() => deleteDeviceMutation.mutate(device)} color={"red"}>
                        {deleteDeviceMutation.isPending ? <Loader  type="dots" /> : "Delete"}
                      </Button>
                    </Stack>
                  </Popover.Dropdown>
                </Popover>
                <Button variant={"filled"} type="submit">
                  Edit Device
                </Button>
              </Group>
            </Stack>
          </Stack>
        </form>
      </Modal>
  );
}

export default EditDeviceModal;