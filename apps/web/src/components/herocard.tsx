import React from "react";
import Spline from "@splinetool/react-spline";

function HeroCard() {
    return (
        <div className="container z-10 flex flex-col-reverse lg:flex-row justify-center items-center bg-opacity-0">
            <div className="text-center lg:w-1/2">
                <svg>
                    <text x="50%" y="50%" dy=".35em" textAnchor="middle" className="text-4xl md:text-6xl lg:text-7xl">
                        KonanNote
                    </text>
                </svg>
            </div>
            <div className="flex justify-center items-center w-full lg:w-1/2 opacity-50">
                <Spline
                    className="spline w-full ml-10 md:w-3/4 lg:w-2/3 xl:w-1/2"
                    scene="https://prod.spline.design/qpAI8P8HP5irz6hR/scene.splinecode"
                />
            </div>
        </div>
    )
}
export default HeroCard;