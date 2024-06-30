import {ColorSwatch, Loader} from "@mantine/core";

type DeviceStatusProps = {
  status: number;
  isFetching: boolean;
  isError: boolean;

}

function DeviceStatus(props : DeviceStatusProps) {
  if (props.isFetching) {
    return <Loader size={24}/>
  }
  if (props.isError) {
    return <ColorSwatch size={24} color="orange"/>
  }
  if (props.status === 1) {
    return <ColorSwatch size={24} color="green"/>
  }
  else if (props.status === 0) {
    return <ColorSwatch size={24} color="red"/>
  }
  else {
    return <ColorSwatch size={24} color="orange"/>
  }
}

export default DeviceStatus;