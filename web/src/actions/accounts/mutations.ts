import apiClient from "~/axios";
import { AccountInterface, CreateAccountInterface } from "~/schema/account";

export const createAccountFn = async (values: CreateAccountInterface) => {
  const formData = new FormData();
  formData.append("name", values.name);
  if (values.logo) formData.append("logo", values.logo);

  return apiClient
    .post<AccountInterface>("/accounts", formData)
    .then((r) => r.data);
};
