import React from "react";
import PolygonImg from "@/assets/images/Polygon.svg";

type PropType = {
  isInverted: boolean;
};

function Polygon({ isInverted }: PropType) {
  return (
    <img
      src={PolygonImg}
      className={`h-2 ${
        isInverted ? "rotate-180" : null
      }  transition-all duration-500`}
      alt="select button icon"
    />
  );
}

export { Polygon };
