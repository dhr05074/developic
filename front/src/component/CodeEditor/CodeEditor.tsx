import { useEffect, useMemo, useRef } from "react";
import { useCodeMirror } from "@uiw/react-codemirror";
import { javascript } from "@codemirror/lang-javascript";

const code = "console.log('hello world!');\n\n\n";
// Define the extensions outside the component for the best performance.
// If you need dynamic extensions, use React.useMemo to minimize reference changes
// which cause costly re-renders.
const extensions = [javascript()];

export default function CodeEditor() {
    const editor = useRef();
    const { setContainer } = useCodeMirror({
        container: editor.current,
        extensions,
        value: code,
    });

    useEffect(() => {
        if (editor.current) {
            setContainer(editor.current);
        }
    }, [editor.current]);

    return <div ref={editor} />;
}
