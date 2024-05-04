import React from "react";

import Image from "next/image";
import Link from "next/link";
import { Box, Typography } from "@mui/material";
import { CollectionCard } from "./collectionCard";

export function Hero() {
  return (
    <Box className="relative w-full overflow-hidden bg-cover bg-center py-6 sm:py-12 md:py-24 lg:py-32 xl:py-48 bg-gradient-to-b">
      <div className="max-w-full flex flex-col items-center py-20">
        <p className="text-8xl font-extrabold	">COLLECTION</p>
        <p className="text-8xl font-extrabold	bg-gradient-to-r from-blue-600 via-green-500 to-indigo-400 inline-block text-transparent bg-clip-text">
          BY ARTIST
        </p>
      </div>
      <Box className="flex flex-col items-center justify-center gap-x-8 px-12 sm:flex-row sm:px-12 md:px-24 lg:px-32 xl:px-48">
        <Box className="flex flex-wrap items-center justify-center">
          <CollectionCard />
          <CollectionCard />
          <CollectionCard />
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
