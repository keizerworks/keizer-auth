import Cookies from "js-cookie";
import { create } from "zustand";
import { persist, StorageValue } from "zustand/middleware";

import { createSelectors } from "../zustand";

interface ActiveAccountInterface {
  data: string | null;
  setData: (data: ActiveAccountInterface["data"]) => void;
}

const _useActiveAccountStore = create<ActiveAccountInterface>()(
  persist(
    (set) => ({
      data: null,
      setData: (data) => set({ data }),
    }),
    {
      name: "active-account",
      storage: {
        getItem: async (name) => {
          try {
            const str = Cookies.get(name);
            if (!str) return null;
            return JSON.parse(str);
          } catch {
            return null;
          }
        },
        setItem: (name, account: StorageValue<string>) => {
          const str = JSON.stringify(account);
          Cookies.set(name, str, { secure: import.meta.env.PROD });
        },
        removeItem: (name) => Cookies.remove(name),
      },
    },
  ),
);

const useActiveAccountStore = createSelectors(_useActiveAccountStore);

export const getActiveAccount = () => _useActiveAccountStore.getState().data;
export const setActiveAccount: ActiveAccountInterface["setData"] = (data) =>
  _useActiveAccountStore.getState().setData(data);

export default useActiveAccountStore;
