import Cookies from "js-cookie";
import { create } from "zustand";
import { persist, StorageValue } from "zustand/middleware";

import { profile } from "~/actions/auth/profile";
import { UserInterface } from "~/schema/user";

import { createSelectors } from "../zustand";

interface UserStoreInterface {
  data: UserInterface | null;
  logout: () => void;
  setData: (data: UserInterface) => void;
}

const _useUserStore = create<UserStoreInterface>()(
  persist(
    (set) => ({
      data: null,
      setData: (data: UserInterface) => set({ data }),
      logout: () => {
        // TODO:
        Cookies.remove("user-storage");
        set({ data: null });
        window.location.reload();
      },
    }),
    {
      name: "user-storage",
      storage: {
        getItem: async (name) => {
          try {
            const str = Cookies.get(name);
            let data: StorageValue<UserInterface>;
            if (!str) {
              data = {
                state: await profile(),
                version: 0,
              };
              setUser(data.state);
            } else data = JSON.parse(str);
            return data;
          } catch {
            logout();
            return null;
          }
        },
        setItem: (name, user: StorageValue<UserInterface>) => {
          const str = JSON.stringify(user);
          Cookies.set(name, str);
        },
        removeItem: (name) => Cookies.remove(name),
      },
    },
  ),
);

const useUserStore = createSelectors(_useUserStore);

export const getUser = () => _useUserStore.getState().data;
export const setUser = (data: UserInterface) =>
  _useUserStore.getState().setData(data);
export const logout = () => _useUserStore.getState().logout();

export default useUserStore;
