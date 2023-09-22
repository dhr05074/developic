import React, { useEffect, useMemo, useState } from "react";

import { DropDownProvider } from "./useDropDownContext";
import { DropDownMenu, Label, Polygon, Wrapper } from "./component";

export type DropDownValueState = {
  menu: string[];
  disabled: string;
  location: "left-0" | "right-0";
  size: "small" | "large";
};

function DropDown({
  children,
  value,
}: {
  children: React.ReactNode;
  value: DropDownValueState;
}) {
  return (
    <DropDownProvider value={value}>
      <section>{children}</section>
    </DropDownProvider>
  );
}

DropDown.Polygon = Polygon;
DropDown.Wrapper = Wrapper;
DropDown.Label = Label;
DropDown.DropDownMenu = DropDownMenu;

export { DropDown };
