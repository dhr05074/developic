import React from "react";

type State = {
  setLang : (value: LanguageType) => void
  };

const HomeContext = React.createContext<State | undefined>(undefined);

function HomeProvider({ children, value }: { children: React.ReactNode, value:State | undefined }) {
    //  value={value}
  return (
    <HomeContext.Provider value={value}>{children}</HomeContext.Provider>
  );
}

function useHomeContext() {
  const context = React.useContext(HomeContext);
  if (context === undefined) {
    throw new Error("useHomeContext must be used within a CounterProvider");
  }
  return context;
}

export { HomeProvider, useHomeContext };
