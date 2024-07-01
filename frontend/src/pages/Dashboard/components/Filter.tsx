import {ActionIcon, Group, Select, TextInput} from "@mantine/core";
import {Search} from 'lucide-react';
import {useRef} from "react";

type FilterProps = {
  limit: number,
  search: string,
  changeLimit: (newLimit: number) => void,
  changeSearch: (search: string) => void
}

function Filter(props: FilterProps) {
  const {limit, search, changeLimit, changeSearch} = props;

  const searchRef = useRef(null);

  return (
      <Group grow>
        <form onSubmit={(e: Event) => changeSearch(e, searchRef.current.value)}>
          <TextInput placeholder={"Search"}
                     defaultValue={search}
                     ref={searchRef}
                     rightSection={
                       <ActionIcon  type={"submit"} variant="subtle" aria-label="Search">
                         <Search/>
                       </ActionIcon>}/>
        </form>
        <Select
            onChange={changeLimit}
            w={"100"}
            value={limit ? limit.toString() : "5"}
            data={['5', '10', '15', '20']}
        />

      </Group>
  );
}

export default Filter;