import { useState } from "react";

interface propsState {
    value: {
        callback: () => void;
        menu: Languages;
    };
}

export default function Select(props: propsState) {
    const onChangeSelect = (e) => {
        props.value.callback(e.target.value);
    };
    const setOptions = () => {
        console.log(props.value);
        const languages = props.value.menu;
        const result = [];
        for (let i = 0; i < languages.length; i++) {
            result.push(<option key={i}>{languages[i]}</option>);
        }
        return result;
    };
    // didMount 나중에 넣기.
    // onChangeSelect();

    return (
        <div>
            <select
                id="countries"
                className="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
                onChange={onChangeSelect}
            >
                {setOptions()};
            </select>
        </div>
    );
}
