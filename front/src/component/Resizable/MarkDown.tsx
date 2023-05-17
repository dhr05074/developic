import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

type Props = {
    markdown: string;
};
export default function MarkDown({ markdown }: Props) {
    return (
        <ReactMarkdown
            className=" problem-content prose prose-slate"
            // react-markdown에서 지원해 주지 않는 문법들을 추가로 변형시키는데 사용하는 플러그인 ( table, link 등 )
            remarkPlugins={[remarkGfm]}
            // 아래처럼 사용하면 추가적으로 특정 태그에 원하는 스타일링을 부여할 수 있음
            components={{
                h2: ({ node, ...props }) => <h2 {...props} className="text-white" />,
                p: ({ node, ...props }) => <p {...props} className="text-white" />,
                // code: ({ node, ...props }) => <p {...props} className="text-white bg-Navy-500" />,
                a: ({ node, ...props }) => <a {...props} className="text-indigo-500 no-underline" />,
            }}
            children={markdown}
        />
    );
}
