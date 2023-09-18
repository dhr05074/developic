import React, { useEffect, useMemo, useState } from "react";

import { SelectProvider } from "./useSelectContext";
import { DropDownMenu, Label, Polygon, Wrapper } from "./component";

export type SelectValueState = {
  menu: string[];
  disabled: string;
  location: "left-0" | "right-0";
  size: "small" | "large";
};

function Select({
  children,
  value,
}: {
  children: React.ReactNode;
  value: SelectValueState;
}) {
  return (
    <SelectProvider value={value}>
      <section>{children}</section>
    </SelectProvider>
  );
}

Select.Polygon = Polygon;
Select.Wrapper = Wrapper;
Select.Label = Label;
Select.DropDownMenu = DropDownMenu;

export { Select };
