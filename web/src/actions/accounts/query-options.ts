import { queryOptions } from "@tanstack/react-query";

import { getAccounts } from "./index";

export const getAccountQueryOption = queryOptions({
  queryKey: ["get-accounts"],
  queryFn: getAccounts,
});
