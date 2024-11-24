import { create } from "zustand";

import { AccountInterface } from "~/schema/account";

import { createSelectors } from "../zustand";

interface AccoutStoreInterface {
  data: AccountInterface[];
  setData: (data: AccoutStoreInterface["data"]) => void;
}

const _useAccountStore = create<AccoutStoreInterface>((set) => ({
  data: [],
  setData: (data) => set({ data }),
}));

const useAccountStore = createSelectors(_useAccountStore);

export const getAccounts = () => _useAccountStore.getState().data;
export const setAccounts: AccoutStoreInterface["setData"] = (data) =>
  _useAccountStore.getState().setData(data);

export default useAccountStore;
