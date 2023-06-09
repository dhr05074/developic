import { useState } from "react";
import { Link } from "react-router-dom";
import Icons from "@/component/Icons";

export default function NavBar() {
    const [isSpinner, setSpinner] = useState<boolean>(false);
    return (
        <nav className="h-full bg-Navy-900 py-2">
            <div className="mx-auto flex h-full max-w-screen-xl flex-wrap items-center justify-between px-4">
                {/* logo */}
                <div className=" flex flex-row items-center gap-4">
                    <Link to="/" className=" flex flex-row items-center gap-4">
                        <div className="h-8 w-8 bg-coco-green_500" />
                        <p className=" pretendard_bold_24 whitespace-nowrap text-white">Developic</p>
                    </Link>
                </div>
                <Icons icon="profile" color="#A5A4AD" width={30} />
            </div>
        </nav>
    );
}
