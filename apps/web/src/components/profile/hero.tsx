"use client";
import React, { useEffect, useState } from "react";
import { Box, Button, Typography } from "@mui/material";
import { useAccount, useBalance } from "wagmi";
import { Loading } from "@/components";
import { Avatar } from "connectkit";
import { NestedList } from "./SideList";
import { NftTable } from "./NftTable";
import { ProfileText } from "@/content/ProfileData";

export function Hero() {
  const [isClient, setIsClient] = useState(false);
  const [isExpanded, setIsExpanded] = useState(false);

  const { address, isConnected, chain } = useAccount();

  const isAddress: `0x${string}` | undefined = address

  // data?.value可以取的Address的balance
  const { data, isError, isLoading } = useBalance({
    address: isAddress,
  });
  useEffect(() => {
    if (isConnected === true) {
      setIsClient(true);
    } else {
      setIsClient(false);
    }
  }, [isConnected]);

  const toggleText = () => {
    setIsExpanded(!isExpanded);
  };
  
  return (
    <>
      {isClient ? (
        <Box className="flex w-full justify-center items-center mx-auto">
          <Box className="flex flex-col w-full">
            {/* user profile data */}
            <Box className="bg-gradient-to-b from-blue-500 to-bg-white flex flex-row py-10 px-10 ">
              <Box className="flex flex-col gap-3">
                {/* 這裡放合約頭像、名稱 */}
                <Avatar address={address} size={50} radius={50} />
                <Typography className="text-black font-bold text-xs sm:text-sm md:text-lg xl:text-xl flex-wrap">
                  Konan Dynamic NFTs
                </Typography>
              </Box>
            </Box>
            
            <Box className="flex flex-col p-10 gap-3">
              <Box className="text-gray-600">
                {isExpanded ? (
                  <>
                    <Typography className="text-clip w-full xl:w-1/3">{ProfileText}</Typography>
                    <button
                      onClick={toggleText}
                      className="font-bold hover:text-gray-500"
                    >
                      See Less
                    </button>
                  </>
                ) : (
                  <>
                    <Typography className="truncate w-1/3">{ProfileText}</Typography>
                    <button
                      onClick={toggleText}
                      className="font-bold hover:text-gray-500"
                    >
                      See More
                    </button>
                  </>
                )}
              </Box>
              <Box className="flex flex-col sm:flex-row gap-5 text-gray-600 items-center justify-start">
                <Typography>Created May 2024</Typography>
                <Typography>Chain {chain?.name}</Typography>
              </Box>
              {/* 這裡放Profile NFT，NestedList是屬性sideBar，NftTable放Profile的IPFS NFT  */}  
              <Box className="flex flex-col md:flex-row w-full h-full py-10 gap-10">
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
