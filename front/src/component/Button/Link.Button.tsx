import { Link } from "react-router-dom";

type PropsType = {
    link?: {
        pathName: string;
        search?: string;
    };
    name: string;
};

export default function ButtonBasic(props: PropsType) {
    const active =
        "pretendard_bold_24 motion_basic inline-flex items-center justify-center rounded-full bg-coco-green_500 w-full py-5 text-center text-base font-medium text-black  hover:bg-Navy-700 hover:text-white";

    const inactive =
        "pretendard_bold_24 motion_basic inline-flex items-center justify-center rounded-full bg-Navy-800 w-full py-5 text-center text-base font-medium text-Navy-500 hover:text-Navy-500 cursor-not-allowed ";

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
