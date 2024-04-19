"use client";
import React, { Suspense } from "react";
import HeroCard from "@/components/herocard";

function Hero() {
  return (
    <div className="relative min-h-screen w-full overflow-hidden flex justify-center items-center">
      <video autoPlay muted loop className="absolute inset-0 z-0 object-cover min-h-screen">
        <source src="/snow_moon.mp4" type="video/mp4" />
      </video>
      <div className="absolute inset-0 h-full w-full bg-gray-900/75" />
      <Suspense fallback={<div>Loading...</div>}>
        <HeroCard />
      </Suspense>
    </div>
  );
}
export default Hero;





