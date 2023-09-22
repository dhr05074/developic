import React from "react";
import { useDropDownContext } from "../useDropDownContext";

function MenuWrapper({ children }: { children: React.ReactNode }) {
  const { isClick } = useDropDownContext();
  return (
    <ul
      className={`list-none text-white transition-all ease-in-out ${
        isClick ? "flex  flex-col gap-4 pb-4 opacity-100" : "-mt-4 "
      }`}
    >
      {children}
    </ul>
  );
}

export { MenuWrapper };
