import { Dispatch, SetStateAction } from "react";
import { ButtonStyle } from "./styles";

type PropsType = {
  name: string;
  action: Function;
};

function Button(props: PropsType) {
  return (
    <ButtonStyle
      id="button-basic"
      onClick={() => {
        props.action();
      }}
    >
      {props.name}
    </ButtonStyle>
  );
}
export { Button };
