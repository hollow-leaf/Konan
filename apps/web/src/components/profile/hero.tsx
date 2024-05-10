"use client";
import React from "react";
import { Box, Typography } from "@mui/material";
import { useAccount } from "wagmi";
import { Loading } from "@/app/components";

export function Hero() {
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
        <Typography variant="h5">
          Connect wallet to display your profile nft
        </Typography>
      </div>
    );
  } else {
    return (
      <Box className="relative w-full overflow-hidden py-6 sm:py-12 md:py-24 lg:py-32 xl:py-12">
        <Box className="container mx-auto flex flex-col w-full h-screen">
          已連接錢包：{address}
        </Box>
      </Box>
    );
  }
}
