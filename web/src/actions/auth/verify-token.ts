import apiClient from "~/axios";

export const verifyTokenQueryFn = async () =>
  await apiClient
    .get("auth/verify-token")
    .then((res) => console.log(res.data))
    .catch(console.log);
