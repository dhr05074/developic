import { useState } from "react";
import { Link } from "react-router-dom";
import Icons from "@/component/Icons";
import Logo from "@/assets/images/type-logo.svg";

export default function NavBar() {
    return (
        <nav className="h-full bg-Navy-900 py-2">
            <div className="mx-auto flex h-full max-w-screen-xl flex-wrap items-center justify-between px-4">
                {/* logo */}
                <div className=" flex flex-row items-center gap-4">
                    <Link to="/" className=" flex flex-row items-center gap-4">
                        <img src={Logo} />
                    </Link>
                </div>
                <Icons icon="profile" width={30} />
            </div>
        </nav>
    );
}
