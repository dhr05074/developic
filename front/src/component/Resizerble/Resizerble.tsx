import React, { useState, useRef, MouseEvent } from "react";

interface SidebarProps {
    children: React.ReactNode;
}

function ResizableSidebar({ children }: SidebarProps) {
    const [width, setWidth] = useState<number>(200); // initial width of the sidebar
    const ref = useRef<HTMLDivElement>(null); // reference to the sidebar's DOM node
    const [isResizing, setIsResizing] = useState<boolean>(false); // track whether the user is resizing the sidebar

    // function to update the sidebar's width based on the mouse position
    function handleMouseMove(event: MouseEvent) {
        if (isResizing && ref.current) {
            const diff = event.clientX - ref.current.offsetLeft;
            setWidth(diff);
        }
    }

    // add a mousedown event listener to the sidebar to start resizing
    function handleMouseDown() {
        setIsResizing(true);
    }

    // remove the mousemove event listener when the user releases the mouse button
    function handleMouseUp() {
        setIsResizing(false);
    }

    return (
        <div
            ref={ref}
            style={{ width: `${width}px` }}
            onMouseDown={handleMouseDown}
            onMouseUp={handleMouseUp}
            onMouseMove={handleMouseMove}
            style={{ cursor: isResizing ? "ew-resize" : "default" }} // change the cursor when the user is resizing
        >
            <div>{children}</div>
        </div>
    );
}

export default ResizableSidebar;

{
    /* <div
onMouseDown={handleMouseDown}
onMouseUp={handleMouseUp}
onMouseMove={handleMouseMove}
className="bg-Navy-800 h-full w-4"
></div> */
}
