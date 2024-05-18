import { Dialog, Transition } from "@headlessui/react";
import { Avatar, Typography } from "@mui/material";
import React, { Fragment, useState } from "react";
import { Mint_nft } from "./Mint_nft";

export function CollectionCard() {
  const [showBox, setShowBox] = useState<boolean>(false);

  const handleCardClick = () => {
    setShowBox(true);
  };

  const handleCloseNFTCard = () => {
    setShowBox(false);
  };
  
  return (
    <div onClick={handleCardClick} className="w-5/12 max-w-[600px] flex flex-col items-center bg-white rounded-3xl shadow p-4 py-6 m-4 bg-opacity-30 text-black">
      <div className="flex w-full mb-6">
        <div className="m-2">
          <div className="rounded-full border-2 border-orage-200 overflow-hidden">
            <Avatar
              src={"/main1.png"}
              alt="Konan logo"
              sx={{ width: 110, height: 110 }}
            />
          </div>
        </div>
        <div className="m-2 content-center">
          <Typography className="text-4xl font-bold">#Dinosaur</Typography>
          <Typography className="text-xl">@By Sophia</Typography>
        </div>
      </div>
      <div className="flex">
        <div className="mt-2 min-w-60 min-h-72 rounded-lg bg-center bg-cover bg-[url('/main2_bg.gif')]"></div>
        <div className="mt-2 ml-4">
          <div className="min-w-60 min-h-48 rounded-lg bg-center bg-cover bg-[url('/bg7_main9_combined.png')]"></div>
          <div className="mt-4 min-w-60 min-h-48 rounded-lg bg-center bg-cover bg-[url('/bg4_main6_combined.png')]"></div>
        </div>
      </div>
    </div>
  );
}
