"use client";
import React, { useState } from "react";
import CloseIcon from "@mui/icons-material/Close";
import { Box, IconButton, Typography } from "@mui/material";
import Image from "next/image";
import { DNFT_Text } from "@/content/AccessoryText";

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

  return (
    <Box
      className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50"
      onClick={handleBackgroundClick}
    >
      <Box
        className="bg-white rounded-3xl py-5 px-5 shadow-lg my-32 w-3/4 h-3/4 flex flex-col overflow-y-auto"
        onClick={(e) => e.stopPropagation()}
      >
        <Box className="flex justify-end items-end">
          <IconButton onClick={handleClose}>
            <CloseIcon className="text-3xl" />
          </IconButton>
        </Box>
        <Box className="flex flex-col w-full xl:h-full justify-center items-center rounded-3xl xl:px-10 xl:py-2 xl:gap-5">
          <Box className="basis-1/2 xl:basis-4/6 w-full h-full flex flex-col xl:flex-row gap-3 mb-20 xl:mb-0">
            <Box className="basis-1/2 bg-white w-full h-full flex justify-center">
              <Image
                src="https://bafybeihuuhqoo5wqbovk4iz6c22wkryogi4wllguxzj3l22472v5btrloi.ipfs.dweb.link/"
                alt="dnft"
                width={400}
                height={300}
              />
            </Box>
            <Box className="basis-1/2 w-full h-full flex flex-col justify-start items-start gap-3">
              <Typography className="text-4xl text-black font-bold">
                Dynamic update your NFT
              </Typography>
              <Box className="text-gray-600">
                {isExpanded ? (
                  <>
                    <Typography className="text-clip">{DNFT_Text}</Typography>
                    <button
                      onClick={toggleText}
                      className="font-bold hover:text-gray-500"
                    >
                      See Less
                    </button>
                  </>
                ) : (
                  <>
                    <Typography className="truncate w-64 xl:w-96">
                      {DNFT_Text}
                    </Typography>
                    <button
                      onClick={toggleText}
                      className="font-bold hover:text-gray-500"
                    >
                      See More
                    </button>
                  </>
                )}
              </Box>
            </Box>
          </Box>
          {isExpanded ? (
            <Box className="basis-1/2 xl:basis-2/6 w-full flex flex-col xl:flex-row gap-3 items-center justify-center">
              <Box
                className="basis-1/4 w-full rounded-3xl flex flex-col text-center justify-center py-32 xl:py-14"
                id="BG"
              >
                <Typography className="text-2xl font-bold text-black">
                  Wish
                </Typography>
                <Typography className="text-2xl font-bold text-black">
                  Background
                </Typography>
              </Box>
              <Box
                className="basis-1/4 w-full rounded-3xl flex flex-col text-center justify-center py-32 xl:py-14"
                id="BG"
              >
                <Typography className="text-2xl font-bold text-black">
                  Wish
                </Typography>
                <Typography className="text-2xl font-bold text-black">
                  Body
                </Typography>
              </Box>
              <Box
                className="basis-1/4 w-full rounded-3xl flex flex-col text-center justify-center py-32 xl:py-14"
                id="BG"
              >
                <Typography className="text-2xl font-bold text-black">
                  Wish
                </Typography>
                <Typography className="text-2xl font-bold text-black">
                  Accessory
                </Typography>
              </Box>
              <Box className="pt-12 h-full w-full xl:hidden"></Box>
            </Box>
          ) : (
            <Box className="basis-1/2 xl:basis-2/6 w-full h-full flex flex-col xl:flex-row gap-3 items-center justify-center">
              <Box
                className="basis-1/4 w-full rounded-3xl flex flex-col text-center justify-center py-32 xl:py-14"
                id="BG"
              >
                <Typography className="text-2xl font-bold text-black">
                  Wish
                </Typography>
                <Typography className="text-2xl font-bold text-black">
                  Background
                </Typography>
              </Box>
              <Box
                className="basis-1/4 w-full rounded-3xl flex flex-col text-center justify-center py-32 xl:py-14"
                id="BG"
              >
                <Typography className="text-2xl font-bold text-black">
                  Wish
                </Typography>
                <Typography className="text-2xl font-bold text-black">
                  Body
                </Typography>
              </Box>
              <Box
                className="basis-1/4 w-full rounded-3xl flex flex-col text-center justify-center py-32 xl:py-14"
                id="BG"
              >
                <Typography className="text-2xl font-bold text-black">
                  Wish
                </Typography>
                <Typography className="text-2xl font-bold text-black">
                  Accessory
                </Typography>
              </Box>
              <Box className="pt-12 h-full w-full xl:hidden"></Box>
            </Box>
          )}
        </Box>
      </Box>
    </Box>
  );
}
