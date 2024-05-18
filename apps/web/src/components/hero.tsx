"use client"
import React, { Fragment, useState } from "react";
import { Box, Button, CircularProgress, Typography } from "@mui/material";
import { CollectionCard, Mint_nft, NFTCard } from "@/components";
import { Dialog, Transition } from "@headlessui/react";
import { GetMysteryBoxNutton } from "./getMysteryBoxNutton";

export function Hero() {
  const [showBox, setShowBox] = useState<boolean>(false);

  const handleCardClick = () => {
    setShowBox(true);
  };

  const handleCloseNFTCard = () => {
    setShowBox(false);
  };

  return (
    
    <Box className="relative w-full overflow-hidden py-6 sm:py-12 md:py-24 lg:py-32 xl:py-12">
      <div className="max-w-full flex flex-col items-center py-10">
        <p className="text-4xl xl:text-8xl font-extrabold	text-black">COLLECT & DRAW</p>
        <p className="text-4xl xl:text-8xl font-extrabold	bg-gradient-to-r from-[#3F09F9] via-[#CB6BE0] to-[#CE2CE5] inline-block text-transparent bg-clip-text">SUPER RARE NFTS</p>
      </div>
      <div className="max-w-full flex items-center px-48 py-5 Space between">
        <div className="p-2 rounded-xl w-4/12 m-2 flex">
          <div className="p-2 w-1/2 rounded-xl flex items-center	">
            <img className="ms-auto border-2 border-gray-50 rounded-xl" src={"https://gateway.pinata.cloud/ipfs/QmfSUfwXbc739ax7HSo5e6UHUJCh2P9necbULAZK5dBS45"} alt="image description" />
          </div>
          <div className="p-2 w-1/2">
            <div className="p-2 bg-white rounded-xl shadow-lg mb-6">
              <img className="ms-auto border-2 border-gray-50	max-h-xs rounded-xl" src={"https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z"} alt="image description" />
            </div>
            <div className="p-2 bg-white rounded-xl shadow-lg">
              <img className="ms-auto border-2 border-gray-50	max-h-xs rounded-xl" src={"https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z"} alt="image description" />
            </div>
          </div>
        </div>
        <div className="p-4 rounded-xl w-4/12 m-2">
          <div className="p-2 bg-white rounded-xl shadow-lg">
            <img className="ms-auto border-2 border-gray-50	max-h-xs rounded-xl" src={"https://gateway.pinata.cloud/ipfs/QmfSUfwXbc739ax7HSo5e6UHUJCh2P9necbULAZK5dBS45"} alt="image description" />
          </div>
        </div>
        <div className="p-2 rounded-xl w-4/12 m-2 flex">
          <div className="p-2 w-1/2">
            <div className="p-2 bg-white rounded-xl shadow-lg mb-6">
              <img className="ms-auto border-2 border-gray-50	max-h-xs rounded-xl" src={"https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z"} alt="image description" />
            </div>
            <div className="p-2 bg-white rounded-xl shadow-lg">
              <img className="ms-auto border-2 border-gray-50	max-h-xs rounded-xl" src={"https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z"} alt="image description" />
            </div>
          </div>
          <div className="p-2 w-1/2 rounded-xl flex items-center	">
            <img className="ms-auto border-2 border-gray-50 rounded-xl" src={"https://gateway.pinata.cloud/ipfs/QmfSUfwXbc739ax7HSo5e6UHUJCh2P9necbULAZK5dBS45"} alt="image description" />
          </div>
        </div>
      </div>
      <Box className="max-w-full flex flex-col items-center py-20">
        <Typography className="text-4xl xl:text-8xl font-extrabold	text-black">
          Let's start
        </Typography>
        <Typography className="text-4xl xl:text-8xl font-extrabold	bg-gradient-to-r from-[#3F09F9] via-[#CB6BE0] to-[#CE2CE5] inline-block text-transparent bg-clip-text">
          BY ARTIST
        </Typography>
      </Box>
      {/* 這裡請寫好IPFS資訊在@/content/ipfs中，再導入map */}
      <Box className="flex flex-col w-full items-center justify-center gap-x-8 sm:flex-row px-24">
        <Box className="flex w-full justify-center">
            <CollectionCard />
            <CollectionCard />
        </Box>
      </Box>
      <GetMysteryBoxNutton />
    </Box>
  );
}
