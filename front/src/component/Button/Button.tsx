import { ButtonStyle } from "./styles";

type PropsType = {
  name: string;
  click: () => void;
};

function Button(props: PropsType) {
  return (
    <ButtonStyle id="button-basic" onClick={props.click}>
      {props.name}
    </ButtonStyle>
  );
}
export { Button };
