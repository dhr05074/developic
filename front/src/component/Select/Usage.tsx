import React, { useEffect, useState } from "react";
import useSelect from "./useSelect";
import { Select } from "./Select";

interface propsState {
  menu: string[];
  disabled: string;
  location: "left-0" | "right-0";
  size: "small" | "large";
}

export default function Usage(props: propsState) {
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
    <Select>
      <Select.Wrapper
        location={props.location}
        size={props.size}
        clickEvent={onClickSelect}
      >
        <div className="flex flex-row items-center justify-between ">
          <p>{selected}</p>
          <Select.Polygon isInverted={isClick} />
        </div>
        {/* menu */}
        <ul
          className={`list-none text-white transition-all ease-in-out ${
            isClick ? "flex  flex-col gap-4 pb-4 opacity-100" : "-mt-4 "
          }`}
        >
          {isClick && setOptions(props.menu)}
        </ul>
      </Select.Wrapper>
    </Select>
  );
}
