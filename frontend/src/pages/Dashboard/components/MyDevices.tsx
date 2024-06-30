import {ActionIcon, Group, Paper, Stack, Title} from "@mantine/core";
import MyDevicesTable from "./MyDevicesTable.tsx";
import Filter from "./Filter.tsx";
import { RefreshCcw } from 'lucide-react';
import {useQueryClient} from "@tanstack/react-query";

function MyDevices() {
  const queryClient = useQueryClient();
  return (
      <Paper className={"card"} withBorder p={"lg"}>
        <Stack align="flex-start">
          <Group w={"100%"} justify={"space-between"}>
            <Group gap={"sm"}>
              <Title order={2}>My Devices</Title>
              <ActionIcon onClick={() => queryClient.invalidateQueries('devices')} variant="transparent" aria-label="Refresh">
                <RefreshCcw size={24} />
              </ActionIcon>
            </Group>
            <Filter />
          </Group>
          <MyDevicesTable />
        </Stack>
      </Paper>
  );
}

export default MyDevices;