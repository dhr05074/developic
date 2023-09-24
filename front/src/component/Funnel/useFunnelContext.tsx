import { ReactNode, createContext, useContext } from "react";

const FunnelContext = createContext(undefined);

function FunnelProvider({ children }: { children: ReactNode }) {
  return <FunnelContext.Provider>{children}</FunnelContext.Provider>;
}

function useFunnelContext() {
  const context = useContext(FunnelContext);
  if (context === undefined) new Error("useFunnelContext undefined");
  return context;
}

export { useFunnelContext, FunnelProvider };
