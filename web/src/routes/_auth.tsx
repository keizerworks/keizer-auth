import { createFileRoute, Outlet } from "@tanstack/react-router";

export const Route = createFileRoute("/_auth")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div className="container relative min-h-dvh flex-col items-center justify-center grid lg:max-w-none lg:grid-cols-2 lg:px-0">
      <div className="relative hidden h-full flex-col bg-muted p-10 text-white dark:border-r lg:flex">
        <div className="absolute inset-0 bg-zinc-900" />

        <div className="relative z-20 flex items-center text-lg font-medium">
          <img src="/assets/logo/logo-full.svg" />
        </div>

        <div className="relative z-20 mt-auto">
          <blockquote className="space-y-2">
            <p className="text-lg">
              &ldquo;Keizer Auth has saved me countless hours of work and helped
              me deliver quick prototypes to my clients faster than ever
              before.&rdquo;
            </p>
            <footer className="text-sm">Sudarsh</footer>
          </blockquote>
        </div>
      </div>

      <div className="lg:p-8">
        <Outlet />
      </div>
    </div>
  );
}
