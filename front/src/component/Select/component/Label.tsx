import React from "react";

function Label({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex flex-row items-center justify-between ">
      {children}
    </div>
  );
}

export { Label };
