import { Link } from "react-router-dom";

type PropsType = {
    link?: {
        pathName: string;
        search?: string;
    };
    name: string;
};

export default function ButtonBasic(props: PropsType) {
    const active = "  button_large coco_button normal_button";

    const inactive = " button_large coco_button_disabled disabled_button  ";

    return (
        <Link
            id="button-basic"
            to={{ pathname: props.link?.pathName, search: props.link?.search }}
            className={props.link?.pathName ? active : inactive}
        >
            {props.name}
        </Link>
    );
}
