import { useQuery } from "@tanstack/react-query";
import { createFileRoute, Outlet, useRouter } from "@tanstack/react-router";
import { Loader } from "lucide-react";
import { useEffect } from "react";

import { verifyToken } from "~/actions/auth/verify-token";
import { setUser } from "~/global-state/persistant-storage/token";

export const Route = createFileRoute("/_authenticated")({
  component: RouteComponent,
});

function RouteComponent() {
  const router = useRouter();

  const {
    data: user,
    isPending,
    isError,
  } = useQuery({
    queryKey: ["verify-token-app-layout"],
    queryFn: verifyToken,
  });

  useEffect(() => {
    if (user) setUser(user);
  }, [user]);

  if (isPending) {
    return (
      <div className="flex items-center justify-center w-screen h-screen overflow-hidden">
        <Loader className="animate-spin" />
      </div>
    );
  }

  console.log(isError);
  if (isError) {
    router.navigate({
      replace: true,
      to: "/sign-in",
      search: { redirect: location.href },
    });
    return;
  }

  return <Outlet />;
}
