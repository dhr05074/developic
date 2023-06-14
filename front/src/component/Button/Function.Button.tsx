import styled from "styled-components";

type PropsType = {
    func?: void;
    type: "line" | "full";
    size: "large" | "medium" | "small";
    name: string;
};
const Button = styled.button`
    background: #b9ff47;
    border-radius: 20px;
    border: none;
    border: solid 1px #b9ff47;
`;
const FullButton = styled(Button)`
    color: black;
    &:hover {
        background: #353346;
        color: #b9ff47;
        border: solid 1px #b9ff47;
    }
`;
const LineButton = styled(Button)`
    border: solid 1px #b9ff47;
    background: #353346;
    color: #b9ff47;
    &:hover {
        background: #b9ff47;
        color: black;
    }
`;
export default function ButtonFunction(props: PropsType) {
    const active = `button_${props.size}`;
    if (props.type === "full") {
        return (
            <FullButton id="button_function" className={active}>
                {props.name}
            </FullButton>
        );
    } else if (props.type === "line") {
        return (
            <LineButton id="button_function" className={active}>
                {props.name}
            </LineButton>
        );
    } else {
        return null;
    }
}
