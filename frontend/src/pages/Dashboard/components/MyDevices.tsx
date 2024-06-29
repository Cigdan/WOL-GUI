import {Paper, Stack, Title} from "@mantine/core";
import MyDevicesTable from "./MyDevicesTable.tsx";

function MyDevices() {
  return (
      <Paper withBorder p={"lg"}>
        <Stack align="flex-start">
          <Title order={2}>My Devices</Title>
          <MyDevicesTable />
        </Stack>
      </Paper>
  );
}

export default MyDevices;