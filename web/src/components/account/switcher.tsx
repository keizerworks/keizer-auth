import { ChevronsUpDown, GalleryVerticalEnd, Plus } from "lucide-react";
import { useEffect, useMemo, useState } from "react";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "~/components/ui/dropdown-menu";
import {
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  useSidebar,
} from "~/components/ui/sidebar";
import useAccountStore from "~/global-state/accounts";
import useActiveAccountStore from "~/global-state/persistant-storage/selected-account";

import { buttonVariants } from "../ui/button";
import { CreateAccount } from "./create";

export function AccountSwitcher() {
  const { isMobile } = useSidebar();
  const accounts = useAccountStore.use.data();
  const { data: activeAccountId, setData: setActiveAccount } =
    useActiveAccountStore();

  const [openCreateAccount, setOpenCreateAccount] = useState(false);

  const activeAccount = useMemo(() => {
    if (!activeAccountId) return null;
    return accounts.find((a) => a.id === activeAccountId);
  }, [accounts, activeAccountId]);

  useEffect(() => {
    if (accounts && accounts.length > 0 && !activeAccountId) {
      setActiveAccount(accounts[0].id);
    }
  }, [accounts, activeAccountId, setActiveAccount]);

  return (
    <SidebarMenu>
      <CreateAccount open={openCreateAccount} setOpen={setOpenCreateAccount} />

      <SidebarMenuItem>
        <DropdownMenu>
          <DropdownMenuTrigger disabled={!activeAccountId} asChild>
            <SidebarMenuButton
              size="lg"
              className={buttonVariants({
                variant: "secondary",
                size: "sm",
                className: "shadow-sm truncate",
              })}
            >
              <GalleryVerticalEnd className="size-4" />
              {activeAccount?.name ?? "-"}
              <ChevronsUpDown className="ml-auto" />
            </SidebarMenuButton>
          </DropdownMenuTrigger>

          <DropdownMenuContent
            className="w-[--radix-dropdown-menu-trigger-width] min-w-56 rounded-lg"
            align="start"
            side={isMobile ? "bottom" : "right"}
            sideOffset={4}
          >
            <DropdownMenuLabel>Accounts</DropdownMenuLabel>
            <DropdownMenuSeparator />

            {accounts.map((account) => (
              <DropdownMenuItem
                key={account.name}
                onClick={() => setActiveAccount(account.id)}
                className="gap-2 p-2"
              >
                {account.name}
              </DropdownMenuItem>
            ))}

            <DropdownMenuItem
              onClick={() => setOpenCreateAccount(true)}
              className="gap-2 p-2"
            >
              <div className="flex size-6 items-center justify-center rounded-md border bg-background">
                <Plus className="size-4" />
              </div>
              <div className="font-medium text-muted-foreground">Add team</div>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </SidebarMenuItem>
    </SidebarMenu>
  );
}
