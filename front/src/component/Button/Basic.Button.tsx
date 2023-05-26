import { Link } from "react-router-dom";

type PropsType = {
    link?: {
        pathName: string;
        search?: string;
    };
    name: string;
    clickFunction: () => void;
};

export default function ButtonBasic(props: PropsType) {
    const className =
        "motion_basic inline-flex items-center justify-center rounded-lg bg-coco-green_500 px-5 py-3 text-center text-base font-medium text-black  hover:bg-Navy-700 hover:text-white";

    if (props.link?.pathName) {
        return (
            <Link
                id="button-basic"
                to={{ pathname: props.link.pathName, search: props.link.search }}
                className={className}
            >
                {props.name}
            </Link>
        );
    }
    return (
        <button id="button-basic" className={className} onClick={props.clickFunction}>
            {props.name}
        </button>
    );
}
