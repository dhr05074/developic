import { SelectDifficulty, SelectLang } from "@/recoil/problem.recoil";
import { useState } from "react";
import { useRecoilState } from "recoil";
const useSelect = () => {
  const [isClick, setIsClick] = useState(false);
  const [selected, setSelected] = useState(""); // props.disabled

  const [optionLength, setOptionLength] = useRecoilState(SelectLang);
  const [optionDifficulty, setOptionDifficulty] =
    useRecoilState(SelectDifficulty);
  // 다음페이지로 넘겨야하기때문에 전역으로 관리했음.
  const initSelectOption = () => {
    setLang("");
    setDifficulty("");
  };

  const setLang = (value: LanguageType) => {
    setOptionLength(value);
  };
  const setDifficulty = (value: difficultyType) => {
    setOptionDifficulty(value);
  };

  const onClickSelect = () => {
    setIsClick(!isClick);
  };
  const onClickMenu = (e: React.MouseEvent<HTMLLIElement>) => {
    const target = e.target as HTMLLIElement;
    setSelected(target.innerHTML);
    // props?.value.callback(target.innerHTML);
    setIsClick(!isClick);
  };
  const setOptions = (menu: string[]) => {
    // 컴포넌트화 시킬 수 없을까?
    const result = [];
    for (let i = 0; i < menu.length; i++) {
      result.push(
        <li
          onClick={onClickMenu}
          key={i}
          className="motion_basic hover:text-coco-green_500 "
        >
          {menu[i]}
        </li>
      );
    }
    return result;
  };

  return {
    optionLength,
    optionDifficulty,
    selected,
    isClick,
    setLang,
    setDifficulty,
    initSelectOption,
    setOptions,
    setSelected,
    onClickSelect,
    onClickMenu,
  };
};

export default useSelect;
