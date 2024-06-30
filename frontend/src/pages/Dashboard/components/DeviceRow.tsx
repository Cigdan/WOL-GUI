import {Button, Group, Table, Tooltip} from "@mantine/core";
import style from "../dashboard.module.css";
import {Pencil, Power } from "lucide-react";
import {Device} from "../../../misc/Types.ts";
import {useQuery, useMutation} from "@tanstack/react-query";
import {checkDeviceStatus, wakeDevice} from "../../../misc/api.ts";
import DeviceStatus from "./DeviceStatus.tsx";
import Toast from "react-hot-toast";

type DeviceRowProps = {
  device: Device;
  setDeviceToEdit: (device: Device) => void;
}

function DeviceRow(props : DeviceRowProps) {
  const {device} = props;

  const { data, isFetching, isError } = useQuery({
    queryKey: ["deviceStatus" + device.id],
    queryFn: () => checkDeviceStatus(device.id),
    refetchOnMount: false,
    refetchOnWindowFocus: false,
    retry: false,
    refetchInterval: 20000,
  });

  const packetMutation = useMutation({
    mutationFn: () => wakeDevice(device.id),
    onSuccess: () => {
      Toast.success("Successfully sent packet to device");
    },
    onError: (error) => {
      if (error.response) {
        Toast.error(error.response?.data.message);
      } else {
        Toast.error(error.message);
      }
    },
  });

  return (
      <Table.Tr key={device.id}>
        <Table.Td>
          <DeviceStatus status={data ? data.status : -1} isFetching={isFetching} isError={isError} />
        </Table.Td>
        <Table.Td>{device.name}</Table.Td>
        <Table.Td className={style.hiddenInfo}>{device.mac_address}</Table.Td>
        <Table.Td className={style.hiddenInfo}>
          {device.last_online || "Never"}
        </Table.Td>
        <Table.Td>
          <Group gap={0}>
            <Tooltip label={data.status === 1 ? "Device is already on" : "Turn on device"}>
            <Button disabled={data.status === 1} onClick={packetMutation.mutate} variant={"subtle"}>
              <Power size={20}/>
            </Button>
            </Tooltip>
            <Tooltip label={"Edit device"}>
              <Button onClick={() => {
                props.setDeviceToEdit(device);
              }} variant={"subtle"}>
                <Pencil size={20}/>
              </Button>
            </Tooltip>


          </Group>
        </Table.Td>
      </Table.Tr>
  );
}

export default DeviceRow;