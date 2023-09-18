import React from "react";

interface State {
  menu: string[];
  disabled: string;
  location: "left-0" | "right-0";
  size: "small" | "large";
}
const SelectContext = React.createContext<State | undefined>(undefined);

function SelectProvider({
  children,
  value,
}: {
  children: React.ReactNode;
  value: State | undefined;
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
