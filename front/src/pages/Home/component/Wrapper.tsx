import React from "react";

function Wrapper({ children }: { children: React.ReactNode }) {
  return (
    <article className=" h-full w-1/2 gap-4">
      <div className="flex h-full  w-[440px] flex-col items-start justify-center  gap-4 text-left text-white">
        {children}
      </div>
    </article>
  );
}

export { Wrapper };
