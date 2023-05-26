import React from "react";
import { Link } from "react-router-dom";
import { motion } from "framer-motion";

export default function HomePage() {
    return (
        <motion.div
            className=""
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <div id="home-page">
                <section className="bg-gray-700 bg-center bg-no-repeat bg-blend-multiply ">
                    <div className="mx-auto max-w-screen-xl px-4 py-24 text-center lg:py-56">
                        <h1 className="mb-4 text-4xl font-extrabold leading-none tracking-tight text-white md:text-5xl lg:text-6xl">
                            리팩토링 챌린지
                        </h1>
                        <p className="mb-8 text-lg font-normal text-gray-300 sm:px-16 lg:px-48 lg:text-xl">
                            Here at Flowbite we focus on markets where technology, innovation, and capital can unlock
                            long-term value and drive economic growth.
                        </p>
                        <div className="flex flex-col space-y-4 sm:flex-row sm:justify-center sm:space-x-4 sm:space-y-0">
                            <Link
                                to="/select"
                                className="motion_basic inline-flex items-center justify-center rounded-lg bg-coco-green_500 px-5 py-3 text-center text-base font-medium text-black  hover:bg-Navy-700 hover:text-white
                             "
                            >
                                참가
                                <svg
                                    aria-hidden="true"
                                    className="-mr-1 ml-2 h-4 w-4"
                                    fill="currentColor"
                                    viewBox="0 0 20 20"
                                    xmlns="http://www.w3.org/2000/svg"
                                >
                                    <path
                                        fillRule="evenodd"
                                        d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z"
                                        clipRule="evenodd"
                                    />
                                </svg>
                            </Link>
                            <a
                                href="#"
                                className="motion_basic inline-flex items-center justify-center rounded-lg border border-white px-5 py-3 text-center text-base font-medium text-white hover:bg-gray-100 hover:text-gray-900 focus:ring-4 focus:ring-gray-400"
                            >
                                Learn more
                            </a>
                        </div>
                    </div>
                </section>
            </div>
        </motion.div>
    );
}
