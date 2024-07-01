import {ActionIcon, Group, Pagination, Paper, Stack, Title, Tooltip} from "@mantine/core";
import MyDevicesTable from "./MyDevicesTable.tsx";
import Filter from "./Filter.tsx";
import { RefreshCcw } from 'lucide-react';
import {useQueryClient} from "@tanstack/react-query";
import {useEffect, useState} from "react";
import {useNavigate, useSearch} from "@tanstack/react-router";

function MyDevices() {
  const queryClient = useQueryClient();
  const navigate = useNavigate();
  const {deviceLimit, deviceSearch} = useSearch(useSearch({ deviceLimit: 5, deviceSearch: "" }));
  const [limit, setLimit] = useState(deviceLimit ? parseInt(deviceLimit) : 10)
  const [search, setSearch] = useState(deviceSearch || "")
  const [offset, setOffset] = useState(0)
  const [deviceCount, setDeviceCount] = useState(0)

  const changeLimit = async (newLimit) => {
    if (!newLimit || newLimit === limit) return;
    setLimit(parseInt(newLimit) || 5)
  }

  const changeSearch = async (e, search) => {
    e.preventDefault();
    if (search === deviceSearch) return;
    setSearch(search)
  }

  const changePage = async (newOffset) => {
    if (newOffset - 1 === offset) return;
    setOffset(newOffset - 1)
  }

  useEffect(() => {
    navigate({
      search: (old) => {
        const newSearch = { ...old, deviceLimit: limit, deviceOffset: offset, deviceSearch: null};
        if (search) {
          newSearch.deviceSearch = search;
        } else {
          delete newSearch.deviceSearch;
        }
        return newSearch;
      },
    }).then(() => {
      queryClient.invalidateQueries('devices').then(() => {});
    });
  }, [limit, search, offset]);

  return (
      <Paper className={"card"} withBorder p={"lg"}>
        <Stack align="flex-start">
          <Group w={"100%"} justify={"space-between"}>
            <Group gap={"sm"}>
              <Title order={2}>My Devices</Title>
              <Tooltip label={"Refresh Devices"}>
                <ActionIcon onClick={() => queryClient.invalidateQueries('devices')} variant="transparent" aria-label="Refresh">
                  <RefreshCcw size={24} />
                </ActionIcon>
              </Tooltip>
            </Group>
            <Filter limit={limit} search={search} changeLimit={changeLimit} changeSearch={changeSearch} />
          </Group>
          <MyDevicesTable limit={limit} search={search} offset={offset} setDeviceCount={setDeviceCount} />
          <Pagination onChange={(page) => changePage(page)} value={offset +1} total={Math.floor(deviceCount / limit) || 1} />
        </Stack>
      </Paper>
  );
}

export default MyDevices;