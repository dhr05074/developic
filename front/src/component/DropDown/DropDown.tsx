import React, { useEffect, useMemo, useState } from "react";

import { DropDownProvider } from "./useDropDownContext";
import {
  MenuWrapper,
  Label,
  Polygon,
  Wrapper,
  Menu,
  Selected,
} from "./component";
import { useRecoilState } from "recoil";
import { SelectDifficulty, SelectLang } from "@/recoil/problem.recoil";

function DropDown({
  children,
  value,
}: {
  children: React.ReactNode;
  value: DropDownProps;
}) {
  const [isClick, setIsClick] = useState(false);
  const [selected, setSelected] = useState(""); // props.disabled

  const [optionLength, setOptionLength] = useRecoilState(SelectLang);
  const [optionDifficulty, setOptionDifficulty] =
    useRecoilState(SelectDifficulty);

  const { menu, disabled, location, size } = value;
  // 다음페이지로 넘겨야하기때문에 전역으로 관리했음.
  const initDropDownOption = () => {
    setLang("");
    setDifficulty("");
  };

  const setLang = (value: LanguageType) => {
    setOptionLength(value);
  };
  const setDifficulty = (value: difficultyType) => {
    setOptionDifficulty(value);
  };

  const onClickDropDown = () => {
    setIsClick(!isClick);
  };
  const onClickMenu = (e: React.MouseEvent<HTMLLIElement>) => {
    const target = e.target as HTMLLIElement;
    setSelected(target.innerHTML);
    // props?.value.callback(target.innerHTML);
    setIsClick(!isClick);
  };

  useEffect(() => {
    initDropDownOption();
    setSelected(disabled);
  }, []);

  return (
    <DropDownProvider
      value={{
        menu,
        disabled,
        location,
        size,
        onClickDropDown,
        onClickMenu,
        selected,
        isClick,
      }}
    >
      <section>{children}</section>
    </DropDownProvider>
  );
}

DropDown.Polygon = Polygon;
DropDown.Wrapper = Wrapper;
DropDown.Label = Label;
DropDown.MenuWrapper = MenuWrapper;
DropDown.Selected = Selected;
DropDown.Menu = Menu;

export { DropDown };
