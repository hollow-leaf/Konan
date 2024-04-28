import React from "react";

import Image from "next/image";
import Link from "next/link";
import { Box, Typography } from "@mui/material";

export function Hero() {
  return (
    <Box className="relative w-full overflow-hidden bg-black bg-cover bg-center py-6 sm:py-12 md:py-24 lg:py-32 xl:py-48">
      <Box className="flex flex-col items-center justify-between gap-x-8 px-8 sm:flex-row sm:px-12 md:px-24 lg:px-32 xl:px-48">
        <Box className="flex flex-col">
          <Image
            src={"/konan-hero.webp"}
            alt="xuedao logo"
            width={300}
            height={300}
            priority
          />
          <Typography className="max-w-[600px] md:text-2xl/relaxed text-gray-400">
            Our mission is to break down barriers, foster global connections
            among students, ignite a deep passion for learning, and promote a
            vibrant culture of knowledge sharing for both students and lifelong
            learners.
          </Typography>
        </Box>
        <Box className="absolute right-[20px] top-0 hidden sm:right-[-80px] md:right-[-100px] lg:right-[-120px] xl:right-[-150px] 2xl:right-[-180px] 2xl:block xl:mt-10">
          {/* <Image
            src="/puzzle.png"
            width={700}
            height={600}
            alt="puzzle"
            priority
          /> */}
        </Box>
      </Box>
    </Box>
  );
}
