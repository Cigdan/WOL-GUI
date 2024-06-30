import {ColorSwatch, Loader, Tooltip} from "@mantine/core";

type DeviceStatusProps = {
  status: number;
  isFetching: boolean;
  isError: boolean;

}

function DeviceStatus(props: DeviceStatusProps) {
  if (props.isFetching) {
    return <Loader size={24}/>
  }
  if (props.isError) {
    return (
        <Tooltip label={"Couldn't get status. \n Did you specify an IP Address?"}>
          <ColorSwatch size={24} color="orange"/>
        </Tooltip>
    )
  }
  if (props.status === 1) {
    return (
        <Tooltip label={"Device is on"}>
          <ColorSwatch size={24} color="green"/>
        </Tooltip>

    )
  } else if (props.status === 0) {
    return (
        <Tooltip label={"Device is off"}>
          <ColorSwatch size={24} color="red"/>
        </Tooltip>
    )
  } else {
    return (
        <Tooltip label={"Couldn't get status. \n Did you specify an IP Address?"}>
          <ColorSwatch size={24} color="orange"/>
        </Tooltip>
    )
  }
}

export default DeviceStatus;