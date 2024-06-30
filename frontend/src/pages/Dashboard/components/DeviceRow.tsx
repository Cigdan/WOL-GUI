import {Button, Group, Table} from "@mantine/core";
import style from "../dashboard.module.css";
import {Pencil, Power } from "lucide-react";
import {Device} from "../../../misc/Types.ts";
import {useQuery} from "@tanstack/react-query";
import {checkDeviceStatus} from "../../../misc/api.ts";
import DeviceStatus from "./DeviceStatus.tsx";

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
            <Button variant={"subtle"}>
              <Power size={20}/>
            </Button>
            <Button onClick={() => {
              props.setDeviceToEdit(device);
            }} variant={"subtle"}>
              <Pencil size={20}/>
            </Button>
          </Group>
        </Table.Td>
      </Table.Tr>
  );
}

export default DeviceRow;