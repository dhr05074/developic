import React from "react";
import { Icon } from "@iconify/react";

type PropsType = {
    icon: IconType;
    color: IconColorType;
    width: number;
};

function Icons(props: PropsType) {
    // if ()
    return props.icon === "profile" ? (
        <Icon icon="mdi:account-circle-outline" width={props.width} color={props.color} />
    ) : null;
}

export default Icons;
