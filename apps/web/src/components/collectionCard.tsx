import { Box, Typography } from "@mui/material";
import React from "react";
import Image from "next/image";

export function CollectionCard(onClick: any) {
  return (
    <Box className="bg-white rounded-3xl shadow p-4 m-4 bg-opacity-30 text-black">
      <Box className="flex mb-2">
        <Box className="m-2">
          <Image
            src={"/konan-hero.webp"}
            alt="xuedao logo"
            width={80}
            height={80}
            className="rounded-full"
          />
        </Box>
        <Box className="m-2 content-center">
          <Typography className="text-xl font-bold	">#Polly Doll</Typography>
          <Typography>@By TheSalvare</Typography>
        </Box>
      </Box>
      <Box className="flex">
        <Box className="mt-2 min-w-40 min-h-56 rounded-lg bg-center bg-cover bg-[url('/konan-hero.webp')]"></Box>
        <Box className="mt-2 ml-4">
          <Box className="min-w-40 min-h-28 rounded-lg bg-center bg-cover bg-[url('/konan-hero.webp')]"></Box>
          <Box className="mt-4 min-w-40 min-h-28 rounded-lg bg-center bg-cover bg-[url('/konan-hero.webp')]"></Box>
        </Box>
      </Box>
    </Box>
  );
}
