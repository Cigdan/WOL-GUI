type User = {
  username: string;
  password: string;
}

type Device = {
  id: string;
  name: string;
  mac_address: string;
  ip_address: string;
  last_online: string;
}

export type { User, Device };