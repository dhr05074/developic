import { useDebugValue, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

function useFunnel() {
  //   const [step, setStep] = useState<"Home" | "Loading" | "Problem">("Home");
  const navigate = useNavigate();
  //   useDebugValue(step);

  const setStep = (navigation: "Home" | "Loading" | "Problem") => {
    navigate(navigation);
  };

  return { setStep };
}

export default useFunnel;
