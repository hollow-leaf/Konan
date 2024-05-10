"use client";
import React from "react";
import CloseIcon from '@mui/icons-material/Close';
import { Box, Button, IconButton, Typography } from "@mui/material";
import {
    Flag as FlagIcon,
    Directions as DirectionsIcon,
    Sms as SmsIcon,
  } from "@mui/icons-material";
import Image from "next/image";

interface NFTCardProps {
    openBTN: boolean;
    onClose: () => void;
}
export function NFTCard(props: NFTCardProps) {
    const handleClose = () => {
        props.onClose()
    }
    const handleBackgroundClick = (event: React.MouseEvent<HTMLDivElement>) => {
        if (event.target === event.currentTarget) {
            handleClose();
        }
    }
    return (
        <Box className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50" onClick={handleBackgroundClick}>
            <Box className="bg-white rounded-3xl py-5 px-5 shadow-lg my-32 w-3/4 h-3/4 flex flex-col overflow-y-auto" onClick={(e) => e.stopPropagation()}>
                <Box className="flex justify-end items-end">
                    <IconButton onClick={handleClose}>
                        <CloseIcon className="text-3xl" />
                    </IconButton>
                </Box>
                <Box className="flex flex-col w-full h-full justify-center items-center rounded-3xl mt-120 xl:mt-0 xl:px-10 xl:py-2 xl:gap-5">
                    <Box className="basis-1/2 xl:basis-4/6 w-full h-full flex flex-col xl:flex-row gap-3 mb-20 xl:mb-0">
                        <Box className="basis-1/2 bg-white w-full h-full flex justify-center">
                        <Box className="mx-10 flex min-w-80 min-h-80 xl:min-w-80 xl:min-h-28 rounded-lg bg-center bg-cover bg-[url('/konan-hero.webp')] object-contain"/>
                        </Box>
                        <Box className="basis-1/2 w-full h-full flex flex-col justify-start items-start gap-3">
                            <Typography className="text-4xl text-black font-bold">NFT Mint</Typography>
                            <Typography className="text-xl text-black ">You can also use variant modifiers to target media queries like responsive breakpoints, dark mode, prefers-reduced-motion, and more. For example, use md:scroll-m-0 to apply the scroll-m-0 utility at only medium screen sizes and above.</Typography>
                            <Typography className="text-xl text-black font-bold">Price:</Typography>
                            <Typography className="text-xl text-black font-bold">Date:</Typography>
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
                                <Button className="bg-gray-500 w-3/4 rounded-3xl text-white hover:bg-gray-500 hover:shadow-lg">Mint</Button>
                            </Box>
                        </Box>
                    </Box>
                    <Box className="basis-1/2 xl:basis-2/6 w-full h-full flex flex-col xl:flex-row gap-3 mt-32 xl:mt-0">
                        <Box className="basis-1/4 bg-gray-500 w-full h-full rounded-3xl flex flex-col text-center justify-center py-32 xl:py-0">
                            <Typography className="text-4xl font-bold text-white">0/500</Typography>
                            <Typography className="font-bold text-white px-10">
                                Fair Launch and distribution-all boars have the same cost of 0.04 ETH
                            </Typography>
                        </Box>
                        <Box className="basis-1/4 bg-gray-500 w-full h-full rounded-3xl flex flex-col text-center justify-center py-32 xl:py-0">
                            <Typography className="text-4xl font-bold text-white">0/500</Typography>
                            <Typography className="font-bold text-white px-10">
                                Fair Launch and distribution-all boars have the same cost of 0.04 ETH
                            </Typography>
                        </Box>
                        <Box className="basis-1/4 bg-gray-500 w-full h-full rounded-3xl flex flex-col text-center justify-center py-32 xl:py-0">
                            <Typography className="text-4xl font-bold text-white">0/500</Typography>
                            <Typography className="font-bold text-white px-10">
                                Fair Launch and distribution-all boars have the same cost of 0.04 ETH
                            </Typography>
                        </Box>
                        <Box className="basis-1/4 bg-gray-500 w-full h-full rounded-3xl flex flex-col text-center justify-center py-32 xl:py-0">
                            <Typography className="text-4xl font-bold text-white">0/500</Typography>
                            <Typography className="font-bold text-white px-10">
                                Fair Launch and distribution-all boars have the same cost of 0.04 ETH
                            </Typography>
                        </Box>
                        <Box className="pt-12 h-full w-full xl:hidden"></Box>
                    </Box>
                </Box>
                
            </Box>
        </Box>
    );
}
