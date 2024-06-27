import {useQuery} from "@tanstack/react-query";
import {checkAuth} from "./api.ts";
import {Loader} from '@mantine/core';
import {Navigate} from "@tanstack/react-router";

const PrivateRoute = ({Component}) => {
  const query = useQuery({queryKey: ['auth'], queryFn: checkAuth, staleTime: Infinity, retry: false})
  if (query.isLoading) {
    return <Loader color="blue" />;
  }
  if (query.isError) {
    return <Navigate to="/login" />
  }
  if (query.isSuccess) {
    return <Component />
  }
}

export default PrivateRoute;