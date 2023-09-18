import React, { useEffect, useMemo, useState } from "react";

import { SelectProvider } from "./useSelectContext";
import { Polygon, Wrapper } from "./component";

function Select({ children }: { children: React.ReactNode }) {
  return (
    <SelectProvider value={undefined}>
      <section>{children}</section>
    </SelectProvider>
  );
}

Select.Polygon = Polygon;
Select.Wrapper = Wrapper;

export { Select };
