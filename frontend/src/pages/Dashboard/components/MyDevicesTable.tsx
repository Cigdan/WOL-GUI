import {useQuery} from "@tanstack/react-query";
import {getMyDevices} from "../../../misc/api.ts";
import {Container,Loader, Table, Title} from "@mantine/core";
import {Device} from "../../../misc/Types.ts";
import style from "../dashboard.module.css";
import DeviceRow from "./DeviceRow.tsx";
import {useDisclosure} from "@mantine/hooks";
import EditDeviceModal from "./EditDeviceModal.tsx";
import {useEffect, useState} from "react";

function MyDevicesTable() {
  const [opened, { open, close }] = useDisclosure(false);
  const [deviceToEdit, setDeviceToEdit] = useState<Device | null>(null);
  const { data, isLoading, isError } = useQuery({
    queryKey: ["devices"],
    queryFn: getMyDevices,
    refetchOnMount: false,
    refetchOnWindowFocus: false,
  });

  useEffect(() => {
    if (deviceToEdit) {
      open();
    }
  }, [deviceToEdit, open]);

  if (isLoading) {
    return (
        <Container fluid>
          <Loader />
        </Container>
    )
  }
  if (isError) {
    return (
        <Container fluid>
          <Title order={4}>Error loading devices</Title>
        </Container>
    )
  }
  if (!data.data || data.data.length === 0) {
    return (
        <Container fluid>
          <Title order={4}>No devices found</Title>
        </Container>
    )
  }
  return (
      <>
      {deviceToEdit && <EditDeviceModal setDeviceToEdit={setDeviceToEdit} device={deviceToEdit} opened={opened} close={close} />}
        <Table align={"left"} verticalSpacing="md">
          <Table.Thead>
            <Table.Tr>
              <Table.Th>
                Status
              </Table.Th>
              <Table.Th>
                Name
              </Table.Th>
              <Table.Th className={style.hiddenInfo}>
                Mac Address
              </Table.Th>
              <Table.Th className={style.hiddenInfo}>
                Last Online
              </Table.Th>
              <Table.Th>
                Actions
              </Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {data.data.map((device: Device) => {
              return (
                  <DeviceRow device={device} setDeviceToEdit={setDeviceToEdit} />
              )
            })}
          </Table.Tbody>
        </Table>
      </>
  );
}

export default MyDevicesTable;