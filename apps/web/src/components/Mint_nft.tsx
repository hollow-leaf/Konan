"use client";
import { Box, Button, IconButton, Typography } from "@mui/material";
import React, { Suspense, useState, useEffect } from "react";
import {
  Flag as FlagIcon,
  Directions as DirectionsIcon,
  Sms as SmsIcon,
} from "@mui/icons-material";
import { useAccount } from "wagmi";
import { Loading } from "@/components";

export function Mint_nft() {
  const [isClient, setIsClient] = useState(false);

  const { address, isConnecting, isDisconnected, isConnected } = useAccount();
  useEffect(() => {
    if (isConnected === true) {
      setIsClient(true);
    } else {
      setIsClient(false);
    }
  }, [isConnected]);
  return (
    <>
      {isClient ? (
        <Suspense fallback={<Loading />}>
          <Box className="flex w-full overflow-hidden justify-center items-center mx-auto">
            <Box className="flex flex-col w-full gap-10">
              <Box className="w-full flex flex-col xl:flex-row gap-5 justify-center items-center px-16">
                <Box className="xl:w-1/2 flex justify-center">
                  <Box className="mx-10 flex min-w-80 min-h-80 xl:min-w-120 xl:min-h-28 rounded-lg bg-center bg-cover bg-[url('/konan-hero.webp')] object-contain py-64" />
                </Box>
                <Box className="xl:w-1/2 gap-5 flex flex-col px-10 xl:px-0">
                  <Typography className="text-4xl text-black font-bold">
                    NFT Mint
                  </Typography>
                  <Typography className="text-xl text-black ">
                    You can also use variant modifiers to target media queries
                    like responsive breakpoints, dark mode,
                    prefers-reduced-motion, and more. For example, use
                    md:scroll-m-0 to apply the scroll-m-0 utility at only medium
                    screen sizes and above.
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
            </Box>
          </Box>
        </Suspense>
      ) : (
        <Box className="flex flex-col justify-center items-center h-screen">
          <Loading />
          <Typography variant="h5">Connect wallet to mint nft</Typography>
        </Box>
      )}
    </>
  );
}
