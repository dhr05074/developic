import { ReactNode, useDebugValue, useState } from "react";
import { FunnelProvider } from "./useFunnelContext";
import { Step } from "./component";

type Props = {
  children: ReactNode;
};

function Funnel({ children }: Props) {
  return (
    <FunnelProvider value={undefined}>
      {/* funnel Style tag */}
      <section className="flex h-full w-full flex-row  bg-Navy-900">
        {children}
      </section>
    </FunnelProvider>
  );
}
Funnel.Step = Step;

export { Funnel };
