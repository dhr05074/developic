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
    //didMount 나중에 넣기.
    // onChangeSelect();

    return (
        <div>
            <label htmlFor="countries" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                Select an option
            </label>
            <select
                id="countries"
                className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                onChange={onChangeSelect}
            >
                {setOptions()};
            </select>
        </div>
    );
}
