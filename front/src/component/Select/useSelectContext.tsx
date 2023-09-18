import React from "react";
import { SelectValueState } from "./Select";

const SelectContext = React.createContext<SelectValueState | undefined>(
  undefined
);

function SelectProvider({
  children,
  value,
}: {
  children: React.ReactNode;
  value: SelectValueState | undefined;
}) {
  return (
    <SelectContext.Provider value={value}>{children}</SelectContext.Provider>
  );
}

function useSelectContext() {
  const context = React.useContext(SelectContext);
  if (context === undefined) {
    throw new Error("useSelectContext must be used within a CounterProvider");
  }
  return context;
}

export { SelectProvider, useSelectContext };
