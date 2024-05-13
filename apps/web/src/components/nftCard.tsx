"use client";
import React, { useState } from "react";
import { Box, Button, IconButton, Typography } from "@mui/material";
import Image from "next/image";
import { useAccount } from "wagmi";
import { Loading } from "./Loading";
import {
  Flag as FlagIcon,
  Directions as DirectionsIcon,
  Sms as SmsIcon,
} from "@mui/icons-material";

interface NFTCardProps {
  openBTN: boolean;
  onClose: () => void;
}
export function NFTCard(props: NFTCardProps) {
  const [isExpanded, setIsExpanded] = useState(false);

  const toggleText = () => {
    setIsExpanded(!isExpanded);
  };

  const handleClose = () => {
    props.onClose();
  };

  const handleBackgroundClick = (event: React.MouseEvent<HTMLDivElement>) => {
    if (event.target === event.currentTarget) {
      handleClose();
    }
  };

  const { isConnected } = useAccount();
  return (
    <>
      {isConnected ? (
        <Box
          className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50"
          onClick={handleBackgroundClick}
        >
          <Box
            className="bg-white rounded-3xl py-5 px-5 shadow-lg my-32 w-3/4 h-3/4 flex flex-col overflow-y-auto"
            onClick={(e) => e.stopPropagation()}
          >
            <Box className="flex justify-end items-end">
              <div className="flex justify-end mt-2 transform transition-transform hover:scale-110">
                <svg onClick={handleClose} className='cursor-pointer' xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="40" height="40" viewBox="0 0 50 50">
                  <path d="M 25 2 C 12.309534 2 2 12.309534 2 25 C 2 37.690466 12.309534 48 25 48 C 37.690466 48 48 37.690466 48 25 C 48 12.309534 37.690466 2 25 2 z M 25 4 C 36.609534 4 46 13.390466 46 25 C 46 36.609534 36.609534 46 25 46 C 13.390466 46 4 36.609534 4 25 C 4 13.390466 13.390466 4 25 4 z M 32.990234 15.986328 A 1.0001 1.0001 0 0 0 32.292969 16.292969 L 25 23.585938 L 17.707031 16.292969 A 1.0001 1.0001 0 0 0 16.990234 15.990234 A 1.0001 1.0001 0 0 0 16.292969 17.707031 L 23.585938 25 L 16.292969 32.292969 A 1.0001 1.0001 0 1 0 17.707031 33.707031 L 25 26.414062 L 32.292969 33.707031 A 1.0001 1.0001 0 1 0 33.707031 32.292969 L 26.414062 25 L 33.707031 17.707031 A 1.0001 1.0001 0 0 0 32.990234 15.986328 z"></path>
                </svg>
              </div>
            </Box>
            <Box className="flex flex-col w-full xl:h-3/4 justify-center items-center rounded-3xl xl:px-10 xl:py-2 xl:gap-5">
              <Box className="basis-1/2 xl:basis-4/6 w-full h-full flex flex-col xl:flex-row gap-3 mb-20 xl:mb-0 xl:mt-24">
                <Box className="basis-1/2 bg-white w-full h-full flex justify-center">
                  <Image
                    src="https://bafybeihuuhqoo5wqbovk4iz6c22wkryogi4wllguxzj3l22472v5btrloi.ipfs.dweb.link/"
                    alt="dnft"
                    width={400}
                    height={300}
                    className=" transform transition-transform hover:scale-110"
                  />
                </Box>
                <Box className="basis-1/2 w-full h-full flex flex-col justify-start items-start gap-3">
                  <Typography className="text-4xl text-black font-bold">
                    Dynamic update your NFT
                  </Typography>
                  <Typography className="text-xl text-black ">
                  Imagine your rollup as a canvas, whether it be through
                  unique distribution and minting processes, dynamic NFTs
                  that utilise novel tooling, or applications that leverage Celestia’s
                  high throughput and abundant blockspace.
                  </Typography>
                  <Typography className="text-xl text-black font-bold">
                    Price: 0.000001 ETH
                  </Typography>
                  <Typography className="text-xl text-black font-bold">
                    Date: 12 May 2024
                  </Typography>
                  <Box className="flex flex-row gap-2 mt-5">
                    <IconButton className="border bg-gray-500 rounded-full hover:bg-gradient-to-br from-orange-300 to-bg-orange-100 transform transition-transform hover:scale-110">
                      <SmsIcon className="text-white" />
                    </IconButton>
                    <IconButton className="border bg-gray-500 rounded-full hover:bg-gradient-to-br from-orange-300 to-bg-orange-100 transform transition-transform hover:scale-110">
                      <DirectionsIcon className="text-white" />
                    </IconButton>
                    <IconButton className="border bg-gray-500 rounded-full hover:bg-gradient-to-br from-orange-300 to-bg-orange-100 transform transition-transform hover:scale-110">
                      <FlagIcon className="text-white" />
                    </IconButton>
                    <Button className="bg-gray-500 w-20 rounded-3xl text-white hover:bg-gradient-to-br from-orange-300 to-bg-orange-100 transform transition-transform hover:scale-110">
                      Mint
                    </Button>
                  </Box>
                </Box>
              </Box>
              <Box className="w-full h-full flex flex-col xl:flex-row gap-3 px-2">
                <Box className="basis-1/4 bg-gray-500 w-full h-52 rounded-3xl flex flex-col text-center justify-center py-32 xl:py-0 px-10 gap-5 hover:bg-gradient-to-br from-orange-300 to-bg-orange-100 transform transition-transform hover:scale-110">
                  <Typography className="text-4xl font-bold text-white">
                    Coonect Wallet
                  </Typography>
                </Box>
                <Box className="basis-1/4 bg-gray-500 w-full h-52 rounded-3xl flex flex-col text-center justify-center py-32 xl:py-0 px-10 gap-5 hover:bg-gradient-to-br from-orange-300 to-bg-orange-100 transform transition-transform hover:scale-110">
                  <Typography className="text-4xl font-bold text-white">
                    Mint your NFT
                  </Typography>
                </Box>
                <Box className="basis-1/4 bg-gray-500 w-full h-52 rounded-3xl flex flex-col text-center justify-center py-32 xl:py-0 px-10 gap-5 hover:bg-gradient-to-br from-orange-300 to-bg-orange-100 transform transition-transform hover:scale-110">
                  <Typography className="text-4xl font-bold text-white">
                    Checkout to Profile
                  </Typography>
                </Box>
                <Box className="basis-1/4 bg-gray-500 w-full h-52 rounded-3xl flex flex-col text-center justify-center py-32 xl:py-0 px-10 gap-5 hover:bg-gradient-to-br from-orange-300 to-bg-orange-100 transform transition-transform hover:scale-110">
                  <Typography className="text-4xl font-bold text-white">
                    Update your accessory
                  </Typography>
                </Box>
              </Box>
            </Box>
          </Box>
        </Box>
      ) : (
        <Box
          className="flex flex-col justify-center items-center w-full h-screen"
          onClick={handleClose}
        >
          <Loading />
          <Typography variant="h5">Connect wallet to mint nft</Typography>
        </Box>
      )}
    </>
  );
}
