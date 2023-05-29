import { useEffect, useRef } from "react";
import { createTheme } from "@uiw/codemirror-themes";
import { useCodeMirror } from "@uiw/react-codemirror";
import { javascript } from "@codemirror/lang-javascript";
import { tags as t } from "@lezer/highlight";
import { useState } from "react";

type PropsType = {
    code: string | undefined;
};

const samplecode = `
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {

};`;
const extensions = [javascript({ jsx: true })];

//gutter : line 번호
//caret : 깜빡이
const myTheme = createTheme({
    theme: "",
    settings: {
        background: "#1F1C32",
        foreground: "#75baff",
        caret: "#6B668E",
        selection: "#1F1C32",
        selectionMatch: "#1F1C32",
        lineHighlight: "#1F1C32",
        gutterBackground: "#1F1C32",
        gutterForeground: "#6B668E",
        fontFamily: "Pretendard",
        gutterBorder: "#1F1C32",
    },
    styles: [
        { tag: t.comment, color: "#B9FF47" }, //주석
        { tag: t.bracket, color: "white" },
        { tag: t.labelName, color: "red" },
        { tag: t.variableName, color: "white" },
        { tag: [t.string, t.special(t.brace)], color: "#5c6166" },
        { tag: t.number, color: "#5c6166" },
        { tag: t.bool, color: "#5c6166" },
        { tag: t.null, color: "#4ec9b0" },
        { tag: t.keyword, color: "#5c6166" },
        { tag: t.operator, color: "white" }, // | &
        { tag: t.className, color: "#4ec9b0" },
        { tag: t.definition(t.typeName), color: "#5c6166" },
        { tag: t.typeName, color: "#4ec9b0" },
        { tag: t.angleBracket, color: "#5c6166" },
        { tag: t.tagName, color: "#5c6166" },
        { tag: t.attributeName, color: "#5c6166" },
        { tag: t.keyword, color: "#3AA0FF" }, // function
    ],
});

export default function CodeEditor(props: PropsType) {
    const [code, setCode] = useState("");

    const editor = useRef<HTMLDivElement>(null);

    const { setContainer } = useCodeMirror({
        container: editor.current,
        extensions,
        value: atob(code),
        theme: myTheme,
    });

    useEffect(() => {
        if (editor.current) {
            setContainer(editor.current);
        }
    }, [editor.current]);

    useEffect(() => {
        if (props.code) {
            setCode(props.code);
        } else {
            setCode(btoa(samplecode));
        }
    }, []);

    if (!editor) {
        return null;
    }

    return <div id="codeEditor" ref={editor} className="flex flex-wrap text-left" />;
}
