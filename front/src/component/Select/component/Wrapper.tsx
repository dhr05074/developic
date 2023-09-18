import React from "react";
import { useSelectContext } from "../useSelectContext";

function Wrapper({
  children,
  clickEvent,
}: {
  children: React.ReactNode;
  clickEvent: () => void;
}) {
  const { location, size } = useSelectContext();
  const style =
    " border-Navy-600 border bg-Navy-700 text-coco-green_500 rounded-[1.6rem] " +
    location;
  const setSize = `${
    size === "large"
      ? " w-[13rem] px-6 py-4"
      : " w-[6rem] text-xs px-4 py-2 mt-[0.18rem]"
  }`;
  return (
    <section
      onClick={clickEvent}
      className={
        "selectBox absolute flex cursor-pointer flex-col gap-4  " +
        style +
        setSize
      }
    >
      {children}
    </section>
  );
}

export { Wrapper };
