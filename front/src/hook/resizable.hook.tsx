import { useRef, useState, useEffect } from "react";

let runnerRef = null;
let bodyRef = null;

const useResizable = () => {
    console.log("useResizable");
    const [runnerWidth, setRunnerWidth] = useState<number>(300);
    const [isResizing, setIsResizing] = useState<boolean>(false);

    const getRef = (runner: React.RefObject<HTMLDivElement>, body: React.RefObject<HTMLSelectElement>) => {
        runnerRef = runner;
        bodyRef = body;

        // runnerRef.current.style.width = 300;
        console.log("🚀 ~ file: resizable.hook.tsx:17 ~ getRef ~ runnerRef.current:", bodyRef.current);
        runnerRef.current.style.width = "300";
    };

    const handleMouseMove = (event: MouseEvent) => {
        if (isResizing && runnerRef.current) {
            const current = bodyRef.current as HTMLSelectElement;
            const diff = current.clientWidth - event.clientX + 10;
            setRunnerWidth(diff);
        }
    };
    // add a mousedown event listener to the sidebar to start resizing
    function handleMouseDown() {
        setIsResizing(true);
        // document.addEventListener("mousemove", handleMouseMove);
    }
    // remove the mousemove event listener when the user releases the mouse button
    function handleMouseUp() {
        console.log("🚀 ~ file: App.tsx:77 ~ handleMouseUp ~ handleMouseUp:");
        setIsResizing(false);
        // document.removeEventListener("mousemove", handleMouseMove);
    }
    useEffect(() => {}, []);

    return {
        getRef,
        handleMouseMove,
        handleMouseDown,
        handleMouseUp,
        runnerWidth,
    };
};

export default useResizable;
