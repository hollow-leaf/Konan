"use client";
import { Avatar, Box, Button, IconButton, Typography } from "@mui/material";
import React, { Suspense, useState, useEffect } from "react";
import {
  Flag as FlagIcon,
  Directions as DirectionsIcon,
  Sms as SmsIcon,
} from "@mui/icons-material";
import { useAccount } from "wagmi";
import { Loading } from "@/components";
import { MintNftButton } from "./MintNftButton";

export function Mint_nft() {
  const { address, isConnecting, isDisconnected, isConnected } = useAccount();

  return (
    <div className="w-full flex flex-col items-center bg-white rounded-3xl m-4 bg-opacity-30 text-black">
      <div className="flex w-full px-5">
        <div className="m-2">
          <div className="rounded-full border-2 border-orage-200 overflow-hidden">
            <Avatar
              src={"/main1.png"}
              alt="Konan logo"
              sx={{ width: 80, height: 80 }}
            />
          </div>
        </div>
        <div className="m-2 content-center">
          <Typography className="text-3xl font-bold">#Dinosaur</Typography>
          <Typography className="text-xl">@By Sophia</Typography>
        </div>
      </div>
      <div className="w-full flex flex-col items-center mb-10">
        <img className="object-contain w-7/12 rounded-xl" src={"https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z"} alt="image description" />
      </div>
      <MintNftButton price={1000000000000}/>
    </div>
  );
}
