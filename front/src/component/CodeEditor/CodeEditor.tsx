import { useEffect, useMemo, useRef } from "react";
import { createTheme } from "@uiw/codemirror-themes";
import { useCodeMirror } from "@uiw/react-codemirror";
import { javascript } from "@codemirror/lang-javascript";
import { tags as t } from "@lezer/highlight";

const code = "console.log('hello world!');\n\n\n";
// Define the extensions outside the component for the best performance.
// If you need dynamic extensions, use React.useMemo to minimize reference changes
// which cause costly re-renders.
const extensions = [javascript({ jsx: true })];

const myTheme = createTheme({
    theme: "light",
    settings: {
        background: "black",
        foreground: "#75baff",
        caret: "#5d00ff",
        selection: "#036dd626",
        selectionMatch: "#036dd626",
        lineHighlight: "#8a91991a",
        gutterBackground: "#fff",
        gutterForeground: "#8a919966",
    },
    styles: [
        { tag: t.comment, color: "#787b8099" },
        { tag: t.variableName, color: "#0080ff" },
        { tag: [t.string, t.special(t.brace)], color: "#5c6166" },
        { tag: t.number, color: "#5c6166" },
        { tag: t.bool, color: "#5c6166" },
        { tag: t.null, color: "#5c6166" },
        { tag: t.keyword, color: "#5c6166" },
        { tag: t.operator, color: "#5c6166" },
        { tag: t.className, color: "#5c6166" },
        { tag: t.definition(t.typeName), color: "#5c6166" },
        { tag: t.typeName, color: "#5c6166" },
        { tag: t.angleBracket, color: "#5c6166" },
        { tag: t.tagName, color: "#5c6166" },
        { tag: t.attributeName, color: "#5c6166" },
    ],
});

export default function CodeEditor() {
    const editor = useRef();
    const { setContainer } = useCodeMirror({
        container: editor.current,
        extensions,
        value: code,
        theme: myTheme,
    });

    useEffect(() => {
        if (editor.current) {
            setContainer(editor.current);
        }
    }, [editor.current]);

    return <div ref={editor} />;
}
