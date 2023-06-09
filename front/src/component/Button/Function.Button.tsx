type PropsType = {
    func: void;
    name: string;
};

export default function ButtonFunction(props: PropsType) {
    const active = "  button_large coco_button";

    const inactive = " button_large coco_button_disabled  ";

    return <button id="button-basic">{props.name}</button>;
}
