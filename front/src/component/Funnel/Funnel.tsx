import { ReactNode, useDebugValue, useState } from "react";
import { FunnelProvider } from "./useFunnelContext";

type Props = {
  children: ReactNode;
};

function Funnel({ children }: Props) {
  const [step, setStep] = useState("Home");
  useDebugValue(step);
  return <FunnelProvider value={{ setStep }}>{children}</FunnelProvider>;
}

export { Funnel };
