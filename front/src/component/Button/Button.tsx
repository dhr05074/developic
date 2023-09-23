import { ButtonStyle } from "./styles";

type PropsType = {
  name: string;
  click: () => void;
};

function Button(props: PropsType) {
  const active = "  button_large coco_button normal_button";
  const inactive = " button_large coco_button_disabled disabled_button  ";
  return (
    <ButtonStyle
      id="button-basic"
      onClick={props.click}
      // className={props.link?.pathName ? active : inactive}
    >
      {props.name}
    </ButtonStyle>
  );
}
export { Button };
