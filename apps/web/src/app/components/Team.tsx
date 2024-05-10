import { AutoPlaySwipe } from "@/components/AutoPlaySwipe";
import { Box, Typography } from "@mui/material";
import React from "react";

export function Team() {
  return (
    <Box className="flex flex-col justify-center items-center">
      <Box className="flex flex-col xl:flex-row py-10 xl:py-48 gap-5">
        <Typography className="text-4xl xl:text-8xl font-bold text-black">
          MEET
        </Typography>
        <Typography className="text-4xl xl:text-8xl font-bold bg-gradient-to-r from-blue-600 via-green-500 to-indigo-400 inline-block text-transparent bg-clip-text">
          OUR TEAM
        </Typography>
      </Box>
      <Box className="flex flex-col xl:flex-row mx-auto items-center justify-center align-middle">
        <AutoPlaySwipe />
      </Box>
    </Box>
  );
}
