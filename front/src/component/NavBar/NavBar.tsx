import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import Icons from "@/component/Icons";
import Logo from "@/assets/images/type-logo.svg";
import { useLocation } from "react-router-dom";

export default function NavBar() {
    const location = useLocation();
    const [profileIcon, setProfileIcon] = useState<IconType>("profile");

    useEffect(() => {
        if (location.pathname === "/profile") {
            setProfileIcon("profile_active");
        } else {
            setProfileIcon("profile");
        }
    }, [location]);
    return (
        // max-w-screen-xl mx-auto : max 1280
        <nav className="h-full bg-Navy-900 py-2">
            <div className=" flex h-full  flex-wrap items-center justify-between px-6">
                {/* logo */}
                <div className=" flex flex-row items-center gap-4">
                    <Link to="/" className=" flex flex-row items-center gap-4">
                        <img src={Logo} alt="developic logo image" />
                    </Link>
                </div>
                <Link to="/profile" className=" flex flex-row items-center gap-4 ">
                    <Icons icon={profileIcon} width={40} />
                </Link>
            </div>
        </nav>
    );
}
