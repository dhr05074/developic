import React from "react";

type PropType = {
  description: string;
};

function Description({ description }: PropType) {
  return <p className="pretendard_medium_20 w-11/12">{description}</p>;
}

export { Description };
