import React from "react";
import SelectComponent from "@/component/Select/Select";
import { useHomeContext } from "../useHomeContext";

type propType = {
  menu: string[];
  disabled: string;
  location: "left-0" | "right-0";
  size: "small" | "large";
};

function Select({ menu, disabled, location, size }: propType) {
  const { setLang } = useHomeContext();
  return (
    <SelectComponent
      value={{ menu }}
      disabled={disabled}
      location={location}
      size={size}
    />
  );
}

export { Select };
