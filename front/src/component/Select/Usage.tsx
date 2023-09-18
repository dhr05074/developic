import React, { useEffect, useState } from "react";
import useSelect from "./useSelect";
import { Select, SelectValueState } from "./Select";

export default function Usage(props: SelectValueState) {
  //여기서 훅을 사용해야함
  const {
    onClickSelect,
    setOptions,
    setSelected,
    initSelectOption,
    selected,
    isClick,
  } = useSelect();

  useEffect(() => {
    initSelectOption();
    setSelected(props.disabled);
  }, []);

  return (
    <Select value={props}>
      <Select.Wrapper
        location={props.location}
        size={props.size}
        clickEvent={onClickSelect}
      >
        <Select.Label>
          <p>{selected}</p>
          <Select.Polygon isInverted={isClick} />
        </Select.Label>
        <Select.DropDownMenu isClick={isClick}>
          {isClick && setOptions(props.menu)}
        </Select.DropDownMenu>
      </Select.Wrapper>
    </Select>
  );
}
