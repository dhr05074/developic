import React from "react";

function DropDownMenu({
  children,
  isClick,
}: {
  children: React.ReactNode;
  isClick: boolean;
}) {
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

export { DropDownMenu };
