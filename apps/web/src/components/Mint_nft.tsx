"use client";
import { Box, Button, IconButton, Typography } from "@mui/material";
import React from "react";
import {
  Flag as FlagIcon,
  Directions as DirectionsIcon,
  Sms as SmsIcon,
} from "@mui/icons-material";
import { useAccount } from "wagmi";
import { Loading } from "@/app/components";

export function Mint_nft() {
  const { address, isConnecting, isDisconnected } = useAccount();
  if (isConnecting) {
    return (
      <div className="flex justify-center items-center h-screen">
        <Loading />
        <Typography variant="h5">Connecting</Typography>
      </div>
    );
  } else if (isDisconnected) {
    return (
      <div className="flex flex-col justify-center items-center h-screen">
        <Loading />
        <Typography variant="h5">Connect wallet to mint nft</Typography>
      </div>
    );
  } else {
    return (
      <Box className="flex w-full overflow-hidden py-6 sm:py-12 md:py-24 lg:py-32 xl:py-12 justify-center items-center mx-auto">
        <Box className="flex flex-col w-full mx-24 gap-10">
          <Box className="bg-white w-full flex flex-col xl:flex-row py-24 gap-5 xl:px-24 justify-center items-center">
            <Box className="xl:w-1/2 bg-white  flex justify-center">
              <Box className="mx-10 flex min-w-80 min-h-80 xl:min-w-120 xl:min-h-28 rounded-lg bg-center bg-cover bg-[url('/konan-hero.webp')] object-contain py-64" />
            </Box>
            <Box className="xl:w-1/2 bg-white xl:py-32 gap-5 flex flex-col px-10 xl:px-0">
              <Typography className="text-4xl text-black font-bold">
                NFT Mint
              </Typography>
              <Typography className="text-xl text-black ">
                You can also use variant modifiers to target media queries like
                responsive breakpoints, dark mode, prefers-reduced-motion, and
                more. For example, use md:scroll-m-0 to apply the scroll-m-0
                utility at only medium screen sizes and above.
              </Typography>
              <Typography className="text-xl text-black font-bold">
                Price:
              </Typography>
              <Typography className="text-xl text-black font-bold">
                Date:
              </Typography>
              <Box className="flex flex-row gap-2 mt-5">
                <IconButton className="border bg-gray-500 rounded-full">
                  <SmsIcon className="text-white hover:text-black" />
                </IconButton>
                <IconButton className="border bg-gray-500 rounded-full">
                  <DirectionsIcon className="text-white hover:text-black" />
                </IconButton>
                <IconButton className="border bg-gray-500 rounded-full">
                  <FlagIcon className="text-white hover:text-black" />
                </IconButton>
                <Button className="bg-gray-500 w-20 rounded-3xl text-white hover:bg-gray-500 hover:shadow-lg">
                  Mint
                </Button>
              </Box>
            </Box>
          </Box>

          <Box className="w-full h-full flex flex-col xl:flex-row gap-3 px-2">
            <Box className="basis-1/4 bg-gray-500 w-full h-72 rounded-3xl flex flex-col text-center justify-center py-32 xl:py-0 px-10 gap-5">
              <Typography className="text-4xl font-bold text-white">
                0/500
              </Typography>
              <Typography className="font-bold text-white">
                Fair Launch and distribution-all boars have the same cost of
                0.04 ETH
              </Typography>
            </Box>
            <Box className="basis-1/4 bg-gray-500 w-full h-72 rounded-3xl flex flex-col text-center justify-center py-32 xl:py-0 px-10 gap-5">
              <Typography className="text-4xl font-bold text-white">
                0/500
              </Typography>
              <Typography className="font-bold text-white">
                Fair Launch and distribution-all boars have the same cost of
                0.04 ETH
              </Typography>
            </Box>
            <Box className="basis-1/4 bg-gray-500 w-full h-72 rounded-3xl flex flex-col text-center justify-center py-32 xl:py-0 px-10 gap-5">
              <Typography className="text-4xl font-bold text-white">
                0/500
              </Typography>
              <Typography className="font-bold text-white">
                Fair Launch and distribution-all boars have the same cost of
                0.04 ETH
              </Typography>
            </Box>
            <Box className="basis-1/4 bg-gray-500 w-full h-72 rounded-3xl flex flex-col text-center justify-center py-32 xl:py-0 px-10 gap-5">
              <Typography className="text-4xl font-bold text-white">
                0/500
              </Typography>
              <Typography className="font-bold text-white">
                Fair Launch and distribution-all boars have the same cost of
                0.04 ETH
              </Typography>
            </Box>
            <Box className="pt-12 h-full w-full xl:hidden"></Box>
          </Box>
        </Box>
      </Box>
    );
  }
}
