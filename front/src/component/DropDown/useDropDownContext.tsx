import { createContext, useContext } from "react";

const DropDownContext = createContext<ProviderValue | undefined>(undefined);

function DropDownProvider({
  children,
  value,
}: {
  children: React.ReactNode;
  value: ProviderValue | undefined;
}) {
  return (
    <DropDownContext.Provider value={value}>
      {children}
    </DropDownContext.Provider>
  );
}

function useDropDownContext() {
  const context = useContext(DropDownContext);
  if (context === undefined) {
    throw new Error("useDropDownContext must be used within a CounterProvider");
  }
  return context;
}

export { DropDownProvider, useDropDownContext };
