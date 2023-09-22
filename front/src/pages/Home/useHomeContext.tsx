import { createContext, useContext } from "react";

const HomeContext = createContext<undefined>(undefined);

function HomeProvider({
  children,
  value,
}: {
  children: React.ReactNode;
  value: undefined;
}) {
  return <HomeContext.Provider value={value}>{children}</HomeContext.Provider>;
}

function useHomeContext() {
  const context = useContext(HomeContext);
  if (context === undefined) {
    throw new Error("useHomeContext must be used within a CounterProvider");
  }
  return context;
}

export { HomeProvider, useHomeContext };
