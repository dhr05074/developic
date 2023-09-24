import { ReactNode, useDebugValue, useState } from "react";
import { FunnelProvider } from "./useFunnelContext";

type Props = {
  children: ReactNode;
  value: undefined;
};

function Funnel({ children, value }: Props) {
  const [step, setStep] = useState("Home");
  useDebugValue(step);
  return <FunnelProvider>{children}</FunnelProvider>;
}

export default Funnel;
