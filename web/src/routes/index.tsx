import { createFileRoute } from "@tanstack/react-router";

import { verifyTokenQueryFn } from "~/actions/auth/verify-token";

export const Route = createFileRoute("/")({
  component: RouteComponent,
});

function RouteComponent() {
  verifyTokenQueryFn();
  return "Hello /!";
}
