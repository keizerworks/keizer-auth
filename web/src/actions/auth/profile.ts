import apiClient from "~/axios";
import { UserInterface } from "~/schema/user";

export const profile = async () =>
  await apiClient.get<UserInterface>("auth/profile").then((res) => res.data);
