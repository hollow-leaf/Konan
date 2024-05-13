"use client";
import { Box, Typography } from "@mui/material";
import React from "react";
import { ForkRight, Wallet, Brush, Bolt, CallReceived, Storefront, Widgets } from "@mui/icons-material";
export function Business() {
  return (
    <Box className=" flex flex-col w-full overflow-hidden py-6 sm:py-12 md:py-24 lg:py-32 xl:py-12">
      <Box className="max-w-full flex flex-col items-center py-20">
        <Typography className="text-4xl xl:text-8xl font-extrabold	bg-gradient-to-r from-blue-600 via-green-500 to-indigo-400 inline-block text-transparent bg-clip-text">
          Why choose celestia?
        </Typography>
      </Box>
      <Box className="flex flex-col xl:flex-row items-center justify-center gap-x-8 px-12 sm:flex-row sm:px-12 md:px-24 lg:px-32 xl:px-48">
        <Box className="basis-1/3 w-full flex flex-col items-center justify-center gap-5 rounded-2xl">
          <Box className="basis-1/3 py-16 w-full rounded-2xl border shadow-lg flex flex-col xl:flex-row gap-5 xl:gap-0">
            <Box className="border rounded-2xl h-14 w-14 ml-10 bg-gradient-to-tr from-blue-600 via-green-500 to-indigo-400 flex justify-center items-center">
              <Wallet className="text-white" />
            </Box>
            <Box className="basis-3/4 flex flex-col px-10 gap-5">
              <Typography className="text-2xl font-bold">
                CONNECT YOUR WALLET
              </Typography>
              <Typography className="text-lg">
                Use Trust Wallet, Metamask or to connect to the app.
              </Typography>
            </Box>
          </Box>
          <Box className="basis-1/3 py-16 w-full rounded-2xl border shadow-lg flex flex-col xl:flex-row gap-5 xl:gap-0">
            <Box className="border rounded-2xl h-14 w-14 ml-10 bg-gradient-to-tr from-blue-600 via-green-500 to-indigo-400 flex justify-center items-center">
              <Brush className="text-white" />
            </Box>
            <Box className="basis-3/4 flex flex-col px-10 gap-5">
              <Typography className="text-2xl font-bold">
                SELECT YOUR QUANTITY
              </Typography>
              <Typography className="text-lg">
                Upload your NFTs and set a title, description and price.
              </Typography>
            </Box>
          </Box>
          <Box className="basis-1/3 py-16 w-full rounded-2xl border shadow-lg flex flex-col xl:flex-row gap-5 xl:gap-0">
            <Box className="border rounded-2xl h-14 w-14 ml-10 bg-gradient-to-tr from-blue-600 via-green-500 to-indigo-400 flex justify-center items-center">
              <Bolt className="text-white" />
            </Box>
            <Box className="basis-3/4 flex flex-col px-10 gap-5">
              <Typography className="text-2xl font-bold">
              CONFIRM TRANSACTION
              </Typography>
              <Typography className="text-lg">
              Earn ETH all your NFTs that you sell on our marketplace.
              </Typography>
            </Box>
          </Box>
        </Box>

        <Box className="basis-1/3 w-full flex flex-col items-center justify-center gap-5 rounded-2xl">
          <Box className="mt-4 min-w-80 min-h-96 rounded-lg bg-center bg-cover bg-[url('https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z')]"></Box>
        </Box>

        <Box className="basis-1/3 w-full flex flex-col items-center justify-center gap-5 rounded-2xl">
          <Box className="basis-1/3 py-16 w-full rounded-2xl border shadow-lg flex flex-col xl:flex-row gap-5 xl:gap-0">
            <Box className="border rounded-2xl h-14 w-14 ml-10 bg-gradient-to-tr from-blue-600 via-green-500 to-indigo-400 flex justify-center items-center">
              <CallReceived className="text-white" />
            </Box>
            <Box className="basis-3/4 flex flex-col px-10 gap-5">
              <Typography className="text-2xl font-bold">
              RECEIVE YOUR OWN NFTS
              </Typography>
              <Typography className="text-lg">
              Invest and manage all your NFTs at one place on one platform.
              </Typography>
            </Box>
          </Box>
          <Box className="basis-1/3 py-16 w-full rounded-2xl border shadow-lg flex flex-col xl:flex-row gap-5 xl:gap-0">
            <Box className="border rounded-2xl h-14 w-14 ml-10 bg-gradient-to-tr from-blue-600 via-green-500 to-indigo-400 flex justify-center items-center">
              <Storefront className="text-white" />
            </Box>
            <Box className="basis-3/4 flex flex-col px-10 gap-5">
              <Typography className="text-2xl font-bold">
              TAKE A MARKET TO SELL
              </Typography>
              <Typography className="text-lg">
              Discover, collect the right NFT collections to buy or sell.
              </Typography>
            </Box>
          </Box>
          <Box className="basis-1/3 py-16 w-full rounded-2xl border shadow-lg flex flex-col xl:flex-row gap-5 xl:gap-0">
            <Box className="border rounded-2xl h-14 w-14 ml-10 bg-gradient-to-tr from-blue-600 via-green-500 to-indigo-400 flex justify-center items-center">
              <Widgets className="text-white" />
            </Box>
            <Box className="basis-3/4 flex flex-col px-10 gap-5">
              <Typography className="text-2xl font-bold">
              DRIVE YOUR COLLECTION
              </Typography>
              <Typography className="text-lg">
              We make it easy to Discover, Invest and manage.
              </Typography>
            </Box>
          </Box>
        </Box>
      </Box>
    </Box>
  );
}
