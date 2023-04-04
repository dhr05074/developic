import { analyze } from "src/api/analyze.ts";
import { useState } from "react";

export default function TextArea() {
    return (
        <form>
            <h4 className="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-white">
                Regain <mark className="px-2 text-white bg-blue-600 rounded dark:bg-blue-500">control</mark> over your
                days
            </h4>
        </form>
    );
}
