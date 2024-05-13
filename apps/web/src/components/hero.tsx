"use client"
import React, { useState } from "react";
import { Box, Button, Typography } from "@mui/material";
import { CollectionCard, NFTCard } from "@/components";

export function Hero() {
  const [showBox, setShowBox] = useState(false);

  const handleCardClick = () => {
    setShowBox(true);
  };

  const handleCloseNFTCard = () => {
    setShowBox(false);
  };

  return (
    
    <Box className="relative w-full overflow-hidden py-6 sm:py-12 md:py-24 lg:py-32 xl:py-12">
      <Box className="max-w-full flex flex-col items-center py-20">
        <Typography className="text-4xl xl:text-8xl font-extrabold	text-black">
          Let's start
        </Typography>
        <Typography className="text-4xl xl:text-8xl font-extrabold	bg-gradient-to-r from-blue-600 via-green-500 to-indigo-400 inline-block text-transparent bg-clip-text">
          Mint Dynamic NFT
        </Typography>
      </Box>
      <Box className="flex flex-col items-center justify-center gap-x-8 px-12 sm:flex-row sm:px-12 md:px-24 lg:px-32 xl:px-48">
        <Box className="flex flex-wrap items-center justify-center">
          <button onClick={handleCardClick} className="rounded-3xl transform transition-transform hover:scale-110">
            <CollectionCard />
          </button>
          <button onClick={handleCardClick} className="rounded-3xl transform transition-transform hover:scale-110">
            <CollectionCard />
          </button>
          <button onClick={handleCardClick} className="rounded-3xl transform transition-transform hover:scale-110">
            <CollectionCard />
          </button>
        </Box>
      </Box>
      {showBox && (
        <Box className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
          <Box className="absolute inset-0 flex items-center justify-center">
            <NFTCard openBTN={showBox} onClose={handleCloseNFTCard} />
          </Box>
        </Box>
      )}
    </Box>
  );
}
