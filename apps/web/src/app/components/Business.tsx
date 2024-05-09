import { Box, Typography } from '@mui/material';
import React from 'react';
import { ForkRight } from '@mui/icons-material';
export function Business() {
    return (
        <Box className="relative w-full overflow-hidden py-6 sm:py-12 md:py-24 lg:py-32 xl:py-12">
      <div className="max-w-full flex flex-col items-center py-20">
        <p className="text-4xl xl:text-8xl font-extrabold	bg-gradient-to-r from-blue-600 via-green-500 to-indigo-400 inline-block text-transparent bg-clip-text">
          Why choose celestia?
        </p>
      </div>
      <Box className="flex flex-col xl:flex-row items-center justify-center gap-x-8 px-12 sm:flex-row sm:px-12 md:px-24 lg:px-32 xl:px-48">

        <Box className="basis-1/3 w-full flex flex-col items-center justify-center gap-5 rounded-2xl">
          <Box className="basis-1/3 py-16 w-full rounded-2xl border shadow-lg flex flex-col xl:flex-row gap-5 xl:gap-0">
            <Box className="border rounded-2xl h-14 w-14 ml-10 bg-gradient-to-tr from-blue-600 via-green-500 to-indigo-400 flex justify-center items-center">
                <ForkRight className="text-white"/>
            </Box>
            <Box className="basis-3/4 flex flex-col px-10 gap-5">
                <Typography className="text-2xl font-bold">CONNECT YOUR WALLET</Typography>
                <Typography className="text-lg">Use Trust Wallet, Metamask or to connect to the app.</Typography>
            </Box>
          </Box>
          <Box className="basis-1/3 py-16 w-full rounded-2xl border shadow-lg flex flex-col xl:flex-row gap-5 xl:gap-0">
            <Box className="border rounded-2xl h-14 w-14 ml-10 bg-gradient-to-tr from-blue-600 via-green-500 to-indigo-400 flex justify-center items-center">
                <ForkRight className="text-white"/>
            </Box>
            <Box className="basis-3/4 flex flex-col px-10 gap-5">
                <Typography className="text-2xl font-bold">CONNECT YOUR WALLET</Typography>
                <Typography className="text-lg">Use Trust Wallet, Metamask or to connect to the app.</Typography>
            </Box>
          </Box>
          <Box className="basis-1/3 py-16 w-full rounded-2xl border shadow-lg flex flex-col xl:flex-row gap-5 xl:gap-0">
            <Box className="border rounded-2xl h-14 w-14 ml-10 bg-gradient-to-tr from-blue-600 via-green-500 to-indigo-400 flex justify-center items-center">
                <ForkRight className="text-white"/>
            </Box>
            <Box className="basis-3/4 flex flex-col px-10 gap-5">
                <Typography className="text-2xl font-bold">CONNECT YOUR WALLET</Typography>
                <Typography className="text-lg">Use Trust Wallet, Metamask or to connect to the app.</Typography>
            </Box>
          </Box>
        </Box>

        <Box className="basis-1/3 w-full flex flex-col items-center justify-center gap-5 rounded-2xl">
            <div className="mt-4 min-w-80 min-h-96 rounded-lg bg-center bg-cover bg-[url('/konan-hero.webp')]"></div>
        </Box>

        <Box className="basis-1/3 w-full flex flex-col items-center justify-center gap-5 rounded-2xl">
          <Box className="basis-1/3 py-16 w-full rounded-2xl border shadow-lg flex flex-col xl:flex-row gap-5 xl:gap-0">
            <Box className="border rounded-2xl h-14 w-14 ml-10 bg-gradient-to-tr from-blue-600 via-green-500 to-indigo-400 flex justify-center items-center">
                <ForkRight className="text-white"/>
            </Box>
            <Box className="basis-3/4 flex flex-col px-10 gap-5">
                <Typography className="text-2xl font-bold">CONNECT YOUR WALLET</Typography>
                <Typography className="text-lg">Use Trust Wallet, Metamask or to connect to the app.</Typography>
            </Box>
          </Box>
          <Box className="basis-1/3 py-16 w-full rounded-2xl border shadow-lg flex flex-col xl:flex-row gap-5 xl:gap-0">
            <Box className="border rounded-2xl h-14 w-14 ml-10 bg-gradient-to-tr from-blue-600 via-green-500 to-indigo-400 flex justify-center items-center">
                <ForkRight className="text-white"/>
            </Box>
            <Box className="basis-3/4 flex flex-col px-10 gap-5">
                <Typography className="text-2xl font-bold">CONNECT YOUR WALLET</Typography>
                <Typography className="text-lg">Use Trust Wallet, Metamask or to connect to the app.</Typography>
            </Box>
          </Box>
          <Box className="basis-1/3 py-16 w-full rounded-2xl border shadow-lg flex flex-col xl:flex-row gap-5 xl:gap-0">
            <Box className="border rounded-2xl h-14 w-14 ml-10 bg-gradient-to-tr from-blue-600 via-green-500 to-indigo-400 flex justify-center items-center">
                <ForkRight className="text-white"/>
            </Box>
            <Box className="basis-3/4 flex flex-col px-10 gap-5">
                <Typography className="text-2xl font-bold">CONNECT YOUR WALLET</Typography>
                <Typography className="text-lg">Use Trust Wallet, Metamask or to connect to the app.</Typography>
            </Box>
          </Box>
        </Box>

      </Box>
    </Box>
    );
}
