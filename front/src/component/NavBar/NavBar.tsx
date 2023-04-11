import react from "@vitejs/plugin-react-swc";
export default function NavBar() {
    return (
        <nav className="bg-Navy-900">
            <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
                {/* logo */}
                <a href="#" className="flex flex-row items-start gap-4">
                    <div className="bg-coco-green_500 h-6 w-6"></div>
                    <span className="self-center text-2xl  font-semibold whitespace-nowrap text-white">
                        codeconnect
                    </span>
                </a>
            </div>
        </nav>
    );
}
