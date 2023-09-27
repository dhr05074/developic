import React from "react";
import PolygonImg from "@/assets/images/Polygon.svg";
import { useDropDownContext } from "../useDropDownContext";

function Polygon() {
  const { isClick } = useDropDownContext();
  return (
    <img
      src={PolygonImg}
      className={`h-2 ${
        isClick ? "rotate-180" : null
      }  transition-all duration-500`}
      alt="select button icon"
    />
  );
}

export { Polygon };
