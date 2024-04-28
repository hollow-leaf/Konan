import React from "react";

import { AccessTime, People, Settings } from "@mui/icons-material";
import { Box, Typography } from "@mui/material";

export function Vision() {
  return (
    <Box
      id="vision"
      className="w-full bg-black bg-cover bg-center py-12 md:py-24 lg:py-32"
      component="section"
    >
      <Box className="container mx-auto px-4 md:px-6">
        <Box className="flex flex-col items-center justify-center space-y-4 text-center">
          <Box className="space-y-2">
            <Typography
              variant="h3"
              className="text-gray-400 text-3xl font-bold sm:text-6xl"
            >
              How do we use celestia?
            </Typography>
            <Typography
              variant="body1"
              className="max-w-[900px] text-gray-400 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed"
            >
              As a member of XueDAO, you'll have access to a wide range of
              resources and benefits designed to help you succeed in the world
              of Web3.
            </Typography>
          </Box>
        </Box>
        <Box className="mt-12 grid grid-cols-1 gap-2 sm:grid-cols-3 lg:mt-20 ">
          <Box
            component="a"
            className="group relative z-40 m-2 bg-black p-4 transition-all duration-1000 hover:bg-gray-900 hover:shadow-xl"
          >
            <Box className="absolute left-0 top-0 z-30 h-1 w-24 bg-purple-600/50 transition-all duration-200 group-hover:w-1/2 group-hover:text-gray-400" />
            <Box className="relative px-9 py-2">
              <AccessTime className="h-16 w-16 fill-gray-400 group-hover:text-gray-400" />
              <Typography className="mt-8 text-lg font-semibold text-gray-400 group-hover:text-gray-400">
                Storage
              </Typography>
              <Typography className="mt-4 text-base text-gray-600 group-hover:text-gray-400">
                Access to educational resources and workshops on blockchain and
                Web3 technologies.
              </Typography>
            </Box>
          </Box>
          <Box
            component="a"
            className="group relative z-40 m-2 bg-black p-4 transition-all duration-1000 hover:bg-gray-900 hover:shadow-xl"
          >
            <Box className="absolute left-0 top-0 z-30 h-1 w-24 bg-purple-600/50 transition-all duration-200 group-hover:w-1/2 group-hover:text-gray-400" />
            <Box className="relative px-9 py-2">
              <People className="h-16 w-16 fill-gray-400 group-hover:text-gray-400" />
              <Typography className="mt-8 text-lg font-semibold text-gray-400 group-hover:text-gray-400">
                Prove
              </Typography>
              <Typography className="mt-4 text-base text-gray-600 group-hover:text-gray-400">
                Connect with other students and professionals in the Web3 space.
              </Typography>
            </Box>
          </Box>

          <Box
            component="a"
            className="group relative z-40 m-2 bg-black p-4 transition-all duration-1000 hover:bg-gray-900 hover:shadow-xl"
          >
            <Box className="absolute left-0 top-0 z-30 h-1 w-24 bg-purple-600/50 transition-all duration-200 group-hover:w-1/2 group-hover:text-gray-400" />
            <Box className="relative px-9 py-2">
              <Settings className="h-16 w-16 fill-gray-400 group-hover:fill-white" />
              <Typography className="mt-8 text-lg font-semibold text-gray-400 group-hover:text-gray-400">
                Dynamic NFT
              </Typography>
              <Typography className="mt-4 text-base text-gray-600 group-hover:text-gray-400">
                Participate in the decision-making process of Xue DAO.
              </Typography>
            </Box>
          </Box>
        </Box>
      </Box>
    </Box>
  );
}
