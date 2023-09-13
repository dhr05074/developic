import React from "react";

function SelectWrapper({ children }: { children: React.ReactNode }) {
  return (
    <section
      id="home_selectButtons"
      className="relative mt-10 flex h-14 w-full  flex-row"
    >
      {children}
    </section>
  );
}

export { SelectWrapper };
