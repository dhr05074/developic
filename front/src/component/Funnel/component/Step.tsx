import { ReactNode } from "react";

function Step({ children, name }: { children: ReactNode; name: string }) {
  return <>{children}</>;
}

export { Step };
