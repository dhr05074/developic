import React from "react";
import { DropDownValueState } from "./DropDown";

const DropDownContext = React.createContext<DropDownValueState | undefined>(
  undefined
);

function DropDownProvider({
  children,
  value,
}: {
  children: React.ReactNode;
  value: DropDownValueState | undefined;
}) {
  return (
    <DropDownContext.Provider value={value}>
      {children}
    </DropDownContext.Provider>
  );
}

function useDropDownContext() {
  const context = React.useContext(DropDownContext);
  if (context === undefined) {
    throw new Error("useDropDownContext must be used within a CounterProvider");
  }
  return context;
}

export { DropDownProvider, useDropDownContext };
