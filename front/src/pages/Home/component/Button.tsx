import React from "react";
import ButtonLink from "@/component/Button/Link.Button";

type PropType = {
  name: string;
};

function Button({ name }: PropType) {
  // {optionLength && optionDifficulty ? (
  //     <ButtonLink link={buttonOption} name="Start" />
  //   ) : (
  //     <ButtonLink name="Start" />
  //   )}
  return <ButtonLink name={name} />;
}

export { Button };
