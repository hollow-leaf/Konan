"use client";
import React, { useEffect, useState } from "react";
import { Box, Button, Typography } from "@mui/material";
import { useAccount, useBalance } from "wagmi";
import { Loading } from "@/components";
import { Avatar } from "connectkit";
import { NestedList } from "./SideList";
import { NftTable } from "./NftTable";

export function Hero() {
  const [isClient, setIsClient] = useState(false);
  const [isExpanded, setIsExpanded] = useState(false);

  const { address, isConnected, chain } = useAccount();

  const { data, isError, isLoading } = useBalance({
    address: "0x5C16e64Eac8bf0e8CE0d6f6eAb0b73918cfB0a96",
  });

  useEffect(() => {
    if (isConnected === true) {
      setIsClient(true);
      console.log(chain?.name);
    } else {
      setIsClient(false);
    }
  }, [isConnected]);

  const toggleText = () => {
    setIsExpanded(!isExpanded);
  };
  const text =
    "Courtyard.io is a first-of-its-kind marketplace and tokenization service that enables anyone to easily own and trade physical cards on the blockchain.Keep them vaulted at Brink’s, sell with a click, or burn the NFT to have the identical physical item shipped anywhere in the world, whenever you want. Learn more.Tokenize your cards for free on Courtyard.io";
  return (
    <>
      {isClient ? (
        <Box className="flex w-full justify-center">
          <Box className="flex flex-col w-full">
            {/* user profile data */}
            <Box className="bg-blue-500 flex flex-row">
              <Box className="mx-10 mt-48 flex flex-col gap-3">
                <Avatar address={address} size={100} radius={50} />
                <Typography className="text-white font-bold text-xl">
                  {address}
                </Typography>
              </Box>
              <Box className="text-white font-bold text-xl ml-120" />
              <Box className="mx-72 mt-72 flex flex-col">
                <Typography className="text-white font-bold text-xl">
                  {data?.formatted} {data?.symbol}
                </Typography>
                <Typography className="text-white text-sm">
                  Account Balance
                </Typography>
              </Box>
            </Box>
            {/* user NFT table */}
            <Box className="flex flex-col p-10 gap-3">
              <Box className="text-gray-600">
                {isExpanded ? (
                  <>
                    <Typography className="text-clip w-1/3">{text}</Typography>
                    <button
                      onClick={toggleText}
                      className="font-bold hover:text-gray-500"
                    >
                      See Less
                    </button>
                  </>
                ) : (
                  <>
                    <Typography className="truncate w-1/3">{text}</Typography>
                    <button
                      onClick={toggleText}
                      className="font-bold hover:text-gray-500"
                    >
                      See More
                    </button>
                  </>
                )}
              </Box>
              <Box className="flex flex-row gap-5 text-gray-600 items-center justify-start">
                <Typography>Created May 2024</Typography>
                <Typography>Chain {chain?.name}</Typography>
              </Box>

              <Box className="flex flex-row w-full h-full py-10 gap-5">
                <NestedList />
                <NftTable />
              </Box>
            </Box>
          </Box>
        </Box>
      ) : (
        <Box className="flex flex-col justify-center items-center h-screen">
          <Loading />
          <Typography variant="h5">
            Connect wallet to show your collection
          </Typography>
        </Box>
      )}
    </>
  );
}
