import React, { ChangeEvent, MouseEventHandler, useState } from "react";
import Polygon from "@/assets/images/Polygon.svg";

interface propsState {
    value: {
        callback: Function;
        menu: LanguageType[] | difficultyType[];
    };
    disabled: string;
    location: "left-0" | "right-0";
    size: "small" | "large";
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
                <li onClick={onClickMenu} key={i} className="motion_basic hover:text-coco-green_500 ">
                    {menu[i]}
                </li>,
            );
        }
        return result;
    };
    const onClickSelect = () => {
        setIsClick(!isClick);
    };
    const onClickMenu = (e: React.MouseEvent<HTMLLIElement>) => {
        const target = e.target as HTMLLIElement;
        setSelectMenu(target.innerHTML);
        props?.value.callback(target.innerHTML);
        setIsClick(!isClick);
    };
    // didMount 나중에 넣기.
    // onChangeSelect();
    const style = " border-Navy-600 border bg-Navy-700 text-coco-green_500 rounded-[1.6rem] " + props.location;
    const size = `${props.size === "large" ? " w-[13rem] px-6 py-4" : " w-[6rem] text-xs px-4 py-2 mt-[0.18rem]"}`;

    return (
        <div
            className={"selectBox absolute flex cursor-pointer flex-col gap-4  " + style + size}
            onClick={onClickSelect}
        >
            <div className="flex flex-row items-center justify-between ">
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
                className={`list-none text-white transition-all ease-in-out ${
                    isClick ? "flex  flex-col gap-4 pb-4 opacity-100" : "-mt-4 "
                }`}
            >
                {isClick && setOptions()}
            </ul>
        </div>
    );
}
