import { Icon } from "@iconify/react";

type PropsType = {
    icon: IconType;
    color: IconColorType;
    width: number;
};

function Icons(props: PropsType) {
    if (props.icon === "profile")
        return <Icon icon="mdi:account-circle-outline" width={props.width} color={props.color} />;
}

export default Icons;
