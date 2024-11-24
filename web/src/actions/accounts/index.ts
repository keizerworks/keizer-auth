import apiClient from "~/axios";
import { setAccounts } from "~/global-state/accounts";
import type { AccountInterface } from "~/schema/account";

export const getAccounts = async () => {
  return apiClient.get<AccountInterface[]>("accounts").then((r) => {
    setAccounts(r.data);
    return r.data;
  });
};
