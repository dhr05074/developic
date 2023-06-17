import React from "react";
import profileImg from "@/assets/images/profile.png";
import profileActiveImg from "@/assets/images/profile-active.png";
type PropsType = {
    icon: IconType;
    width: number;
};

function Icons(props: PropsType) {
    switch (props.icon) {
        case "profile":
            return <img src={profileImg} width={props.width} />;
        case "profile_active":
            return <img src={profileActiveImg} width={props.width} />;
        default:
            return null;
    }
}

export default Icons;
