import React, { useRef, useState, useEffect } from "react";

// recoil로 수정.
let runnerRef: React.RefObject<HTMLDivElement> | null = null;
let bodyRef: React.RefObject<HTMLSelectElement> | null = null;

const useResizable = () => {
    const [runnerWidth, setRunnerWidth] = useState<number>(document.body.clientWidth / 2);
    const [isResizing, setIsResizing] = useState<boolean>(false);

    const getRef = (runner: React.RefObject<HTMLDivElement>, body: React.RefObject<HTMLSelectElement>) => {
        runnerRef = runner;
        bodyRef = body;

        // runnerRef.current.style.width = 300;
        // if (runnerRef.current) runnerRef.current.style.width = "600px";
    };

    const handleMouseMove = (event: React.MouseEvent<HTMLElement>) => {
        if (bodyRef.current.clientWidth < 300) {
            console.log("???");
        }
        if (isResizing && runnerRef?.current) {
            console.log(event);
            const current = bodyRef?.current as HTMLSelectElement;
            const diff = current.clientWidth - event.clientX + 10;
            setRunnerWidth(diff);
        }
    };
    function handleMouseDown() {
        setIsResizing(true);
        // document.addEventListener("mousemove", handleMouseMove);
    }
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
