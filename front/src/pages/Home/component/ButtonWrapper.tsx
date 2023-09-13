import React from "react";

function ButtonWrapper({ children }: { children: React.ReactNode }) {
  return (
    <section
      id="home_startButton"
      className="flex w-full flex-row justify-center"
    >
      <div className="mt-3 w-[80%]">{children}</div>
    </section>
  );
}

export { ButtonWrapper };
