import { useEffect } from "react";
import useDropDown from "./useDropDown";
import { DropDown, DropDownValueState } from "./DropDown";

export default function Usage(props: DropDownValueState) {
  //여기서 훅을 사용해야함
  const {
    onClickDropDown,
    setOptions,
    setSelected,
    initDropDownOption,
    selected,
    isClick,
  } = useDropDown();

  useEffect(() => {
    initDropDownOption();
    setSelected(props.disabled);
  }, []);

  return (
    <DropDown value={props}>
      <DropDown.Wrapper clickEvent={onClickDropDown}>
        <DropDown.Label>
          <p>{selected}</p>
          <DropDown.Polygon isInverted={isClick} />
        </DropDown.Label>
        <DropDown.DropDownMenu isClick={isClick}>
          {isClick && setOptions(props.menu)}
        </DropDown.DropDownMenu>
      </DropDown.Wrapper>
    </DropDown>
  );
}
