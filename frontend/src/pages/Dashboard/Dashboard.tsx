import {Title, Stack} from "@mantine/core";
import MyDevices from "./components/MyDevices.tsx";
import AddDevice from "./components/AddDevice.tsx";

function Dashboard() {
  return (
        <Stack>
          <Title order={1}>Dashboard</Title>
          <Stack>
            <MyDevices/>
            <AddDevice />
          </Stack>
        </Stack>
  );
}

export default Dashboard;