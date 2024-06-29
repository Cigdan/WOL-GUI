import {useQuery} from "@tanstack/react-query";
import {getMyDevices} from "../../../misc/api.ts";
import {Button, ColorSwatch, Container, Group, Loader, Table, Title} from "@mantine/core";
import {Device} from "../../../misc/Types.ts";
import {Pencil, Power} from "lucide-react";
import style from "../dashboard.module.css";

function MyDevicesTable() {
  const { data, isLoading, isError } = useQuery({
    queryKey: ["devices"],
    queryFn: getMyDevices,
    refetchOnMount: false,
    refetchOnWindowFocus: false,
  });

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
      <Table align={"left"} verticalSpacing="md">
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
        {data.data.map((device: Device) => {
          return (
              <Table.Tr key={device.id}>
                <Table.Td><ColorSwatch size={24} color={"green"} /></Table.Td>
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
                    <Button variant={"subtle"}>
                      <Pencil size={20}/>
                    </Button>
                  </Group>
                </Table.Td>
              </Table.Tr>
          )
        })}
      </Table>
  );
}

export default MyDevicesTable;