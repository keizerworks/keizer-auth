import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { Dispatch, SetStateAction } from "react";
import { useForm } from "react-hook-form";

import { createAccountFn } from "~/actions/accounts/mutations";
import useAccountStore from "~/global-state/accounts";
import { CreateAccountInterface, createAccountSchema } from "~/schema/account";

import { Button } from "../ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "../ui/dialog";
import { Form, FormField } from "../ui/form";
import { Input } from "../ui/input";

interface Props {
  open: boolean;
  setOpen: Dispatch<SetStateAction<boolean>>;
}

export const CreateAccount = ({ open, setOpen }: Props) => {
  const { data: accounts, setData: setAccounts } = useAccountStore();

  const form = useForm<CreateAccountInterface>({
    resolver: zodResolver(createAccountSchema),
  });

  const { mutate, isPending } = useMutation({
    mutationFn: createAccountFn,
    onSuccess: (data) => {
      setAccounts([...accounts.filter((a) => a.id !== data.id), data]);
      setOpen(false);
    },
  });

  function onOpenChange(value: boolean) {
    if (!value && isPending) return;
    return setOpen(value);
  }

  function onSubmit(values: CreateAccountInterface) {
    mutate(values);
  }

  return (
    <Dialog open={accounts.length === 0 || open} onOpenChange={onOpenChange}>
      <DialogContent hideCloseButton={accounts.length === 0}>
        <DialogHeader>
          <DialogTitle>Create Account</DialogTitle>
        </DialogHeader>

        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="grid gap-y-4">
            <FormField
              control={form.control}
              name="name"
              label="Name"
              render={({ field }) => (
                <Input
                  placeholder="Acme Inc."
                  className="font-medium"
                  {...field}
                />
              )}
            />

            <DialogFooter>
              <Button loading={isPending} type="submit">
                Create
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
};
