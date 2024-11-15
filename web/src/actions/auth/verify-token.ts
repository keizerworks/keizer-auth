import apiClient from "~/axios";
import { UserInterface } from "~/schema/user";

export const verifyToken = async () =>
  await apiClient
    .get<UserInterface>("auth/verify-token")
    .then((res) => res.data);
