import {Container, Title, Stack} from "@mantine/core";
import MyDevices from "./components/MyDevices.tsx";
import AddDevice from "./components/AddDevice.tsx";

function Dashboard() {
  return (
      <Container p={"md"} fluid>
        <Title my={"md"} order={1}>Dashboard</Title>
        <Stack>
          <MyDevices/>
          <AddDevice />
        </Stack>
      </Container>

  );
}

export default Dashboard;