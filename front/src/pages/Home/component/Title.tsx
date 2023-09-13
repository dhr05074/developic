import React from "react";

type PropType = {
  title: string;
};

function Title({ title }: PropType) {
  return <h3 className="pretendard_extrabold_32 w-full">{title}</h3>;
}

export { Title };
