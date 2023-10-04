import { ReactNode, createContext, useContext } from "react";

const FunnelContext = createContext<FunnelProps | undefined>(undefined);

function FunnelProvider({
  children,
  value,
}: {
  children: ReactNode;
  value: undefined;
}) {
  return (
    <FunnelContext.Provider value={value}>{children}</FunnelContext.Provider>
  );
}

function useFunnelContext() {
  const context = useContext(FunnelContext);
  if (context === undefined) new Error("useFunnelContext undefined");
  return context;
}

export { useFunnelContext, FunnelProvider };
