import React from "react";
import Image from "next/image";
import { Box, IconButton, Typography } from "@mui/material";
import {
  Flag as FlagIcon,
  Directions as DirectionsIcon,
  Sms as SmsIcon,
} from "@mui/icons-material";

export function Hero() {
  return (
    <Box className="flex flex-col w-full h-full md:py-32 bg-white border border-gray-200 rounded-3xl shadow p-4 m-4">
      <Box className="container mx-auto flex flex-col w-full h-full text-black">
        <Box className="basis-4/6 sm:basis-2/3 mx-16 p-10 flex flex-col lg:flex-row">
          <Box className="basis-3/4 lg:basis-2/5">
            <Image src="/konan-hero.webp" width={400} height={700} alt="nft" />
          </Box>
          <Box className="basis-1/4 lg:basis-3/5 mt-5 md:mt-0 md:m-10 md:p-5 flex flex-col">
            <Box className="basis-1/6">
              <Typography className="font-bold text-lg sm:text-3xl md:text-4xl">
                NFT Mint Section
              </Typography>
            </Box>
            <Box className="basis-3/6">
              <Typography>NFT detail</Typography>
              <br />
            </Box>
            <Box className="basis-1/6">
              <Typography className="font-bold flex flex-row">
                Price:
                <Typography className="pl-5 font-bold text-lg sm:text-3xl md:text-4xl flex">
                  29.000ETH
                </Typography>
              </Typography>
              <Typography className="font-bold flex flex-row">
                Date:
                <Typography className="pl-5 font-bold text-lg sm:text-3xl md:text-4xl">
                  4.5.2024
                </Typography>
              </Typography>
            </Box>
            <Box className="basis-1/6 flex flex-row py-5 gap-3">
              <IconButton className="border bg-gray-500 rounded-full">
                <SmsIcon className="text-white hover:text-black" />
              </IconButton>
              <IconButton className="border bg-gray-500 rounded-full">
                <DirectionsIcon className="text-white hover:text-black" />
              </IconButton>
              <IconButton className="border bg-gray-500 rounded-full">
                <FlagIcon className="text-white hover:text-black" />
              </IconButton>
            </Box>
          </Box>
        </Box>
        <Box className="basis-2/6 sm:basis-1/3 flex flex-col md:flex-row gap-5 m-5">
          <Box className="basis-1/4 bg-gray-400 rounded-2xl flex flex-col justify-center items-center text-center px-5 w-1/2 md:w-full mx-auto py-10 lg:py-20">
            <Typography variant="h4" className="font-bold">
              0/500
            </Typography>
            <Typography className="font-bold">
              Fair Launch and distribution-all boars have the same cost of 0.04
              ETH
            </Typography>
          </Box>
          <Box className="basis-1/4 bg-gray-400 rounded-2xl flex flex-col justify-center items-center text-center px-5 w-1/2 md:w-full mx-auto py-10 lg:py-20">
            <Typography variant="h4" className="font-bold">
              0/500
            </Typography>
            <Typography className="font-bold">
              Fair Launch and distribution-all boars have the same cost of 0.04
              ETH
            </Typography>
          </Box>
          <Box className="basis-1/4 bg-gray-400 rounded-2xl flex flex-col justify-center items-center text-center px-5 w-1/2 md:w-full mx-auto py-10 lg:py-20">
            <Typography variant="h4" className="font-bold">
              0/500
            </Typography>
            <Typography className="font-bold">
              Fair Launch and distribution-all boars have the same cost of 0.04
              ETH
            </Typography>
          </Box>
          <Box className="basis-1/4 bg-gray-400 rounded-2xl flex flex-col justify-center items-center text-center px-5 w-1/2 md:w-full mx-auto py-10 lg:py-20">
            <Typography variant="h4" className="font-bold">
              0/500
            </Typography>
            <Typography className="font-bold">
              Fair Launch and distribution-all boars have the same cost of 0.04
              ETH
            </Typography>
          </Box>
        </Box>
      </Box>
    </Box>
  );
}
