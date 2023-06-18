import React from "react";
import profileImg from "@/assets/images/profile.svg";
import profileActiveImg from "@/assets/images/profile_active.svg";
type PropsType = {
    icon: IconType;
    height: number;
};

function Icons(props: PropsType) {
    switch (props.icon) {
        case "profile":
            return <img src={profileImg} className={`h-${props.height}`} />;
        case "profile_active":
            return <img src={profileActiveImg} className={`h-${props.height}`} />;
        default:
            return null;
    }
}

export default Icons;
