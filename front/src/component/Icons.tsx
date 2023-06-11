import React from "react";
import profileImg from "@/assets/images/profile.png";
type PropsType = {
    icon: IconType;
    width: number;
};

function Icons(props: PropsType) {
    // if ()
    return props.icon === "profile" ? <img src={profileImg} width={props.width} /> : null;
}

export default Icons;
