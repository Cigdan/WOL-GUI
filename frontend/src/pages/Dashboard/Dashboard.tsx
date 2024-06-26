import {Button} from "@mantine/core";
import {useMutation} from "@tanstack/react-query";
import {getMyDevices} from "../../misc/api.ts";

function Dashboard() {
  const mutation = useMutation({
    mutationFn: () => getMyDevices(),
  })
  return (
      <div>
        <Button onClick={() => mutation.mutate()}>TEst</Button>
      </div>
  );
}

export default Dashboard;