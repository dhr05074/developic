import { ChangeEvent, MouseEventHandler, useState } from "react";
import Polygon from "@/assets/Polygon.svg";

interface propsState {
    value: {
        callback: Function;
        menu: Languages | string;
    };
    disabled: string;
    location: "left-0" | "right-0";
}

export default function Select(props: propsState) {
    const [isClick, setIsClick] = useState(false);
    const [selectMenu, setSelectMenu] = useState(props.disabled);

    const setOptions = () => {
        console.log(props?.value);
        const menu = props?.value.menu;
        const result = [];
        for (let i = 0; i < menu.length; i++) {
            result.push(
                <li onClick={onClickMenu} key={i}>
                    {menu[i]}
                </li>,
            );
        }
        return result;
    };
    const onClickSelect = () => {
        setIsClick(!isClick);
    };
    const onClickMenu = (e: MouseEventHandler<HTMLLIElement>) => {
        const target = e.target as HTMLLIElement;
        setSelectMenu(target.innerHTML);
        props?.value.callback(target.innerHTML);
        setIsClick(!isClick);
    };
    // didMount 나중에 넣기.
    // onChangeSelect();
    const style = " border-new-600 bg-new-700 text-coco-green_500 rounded-[1.6rem] " + props.location;

    return (
        <div className={"selectBox absolute flex w-[13rem]  cursor-pointer flex-col gap-4 px-6 pt-4" + style}>
            {/* <select
                className={" coco_select w-full  border px-6 py-3 text-sm  " + style}
                onChange={onChangeSelect}
            >
                <option disabled selected>
                    {props.disabled}
                </option>
                {setOptions()};
            </select>
                <img src={Polygon} alt="select button icon" /> */}
            <div onClick={onClickSelect} className="flex flex-row items-center justify-between">
                <p>{selectMenu}</p>
                {/* image */}
                {isClick ? (
                    <img
                        src={Polygon}
                        className="h-2 rotate-180 transition-all duration-500"
                        alt="select button icon"
                    />
                ) : (
                    <img src={Polygon} className="h-2  transition-all duration-500" alt="select button icon" />
                )}
            </div>
            {/* menu */}
            <ul
                className={`list-none text-white transition-all  ease-in-out ${
                    isClick ? "flex flex-col gap-4 pb-4 opacity-100" : "h-0 opacity-0 "
                }`}
            >
                {isClick && setOptions()}
            </ul>
        </div>
    );
}
// 버튼으로 language + 세모
// 옵션은 안보이게한다.
// 버튼을 클릭했을 경우 옵션을 보이게한다.
// 옵션은 버튼바로 밑이 아닌 조금 올라오게 겹쳐서 보이게한다.
// 세모는 누를때마다 돈다.
