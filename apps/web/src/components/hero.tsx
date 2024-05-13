import React, { Fragment, useState } from "react";
import { Box, Button, CircularProgress, Typography } from "@mui/material";
import { CollectionCard, Mint_nft, NFTCard } from "@/components";
import { Dialog, Transition } from "@headlessui/react";

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
          COLLECTION
        </Typography>
        <Typography className="text-4xl xl:text-8xl font-extrabold	bg-gradient-to-r from-[#3F09F9] via-[#CB6BE0] to-[#CE2CE5] inline-block text-transparent bg-clip-text">
          BY ARTIST
        </Typography>
      </Box>
      {/* 這裡請寫好IPFS資訊在@/content/ipfs中，再導入map */}
      <Box className="flex flex-col items-center justify-center gap-x-8 px-12 sm:flex-row sm:px-12 md:px-24 lg:px-32 xl:px-48">
        <Box className="flex flex-wrap items-center justify-center">
          <Button onClick={handleCardClick} className="rounded-3xl">
            <CollectionCard />
          </Button>
          <Button onClick={handleCardClick} className="rounded-3xl">
            <CollectionCard />
          </Button>
          <Button onClick={handleCardClick} className="rounded-3xl">
            <CollectionCard />
          </Button>
        </Box>
      </Box>
      <Transition appear show={showBox} as={Fragment}>
        <Dialog as="div" className="relative z-10" onClose={handleCloseNFTCard}>
            <Transition.Child
                as={Fragment}
                enter="ease-out duration-300"
                enterFrom="opacity-0"
                enterTo="opacity-100"
                leave="ease-in duration-200"
                leaveFrom="opacity-100"
                leaveTo="opacity-0"
            >
                <div className="fixed inset-0 bg-black bg-opacity-25" />
            </Transition.Child>
            <div className="fixed inset-0 overflow-y-auto">
                <div className="flex min-h-full items-center justify-center p-4 text-center">
                    <Transition.Child
                        as={Fragment}
                        enter="ease-out duration-300"
                        enterFrom="opacity-0 scale-95"
                        enterTo="opacity-100 scale-100"
                        leave="ease-in duration-200"
                        leaveFrom="opacity-100 scale-100"
                        leaveTo="opacity-0 scale-95"
                    >
                        <Dialog.Panel className="w-full max-w-6xl transform overflow-hidden rounded-2xl bg-white text-left align-middle shadow-xl transition-all">
                          <div className="py-2 px-2">
                            <div className="flex justify-end mt-2">
                              <svg onClick={handleCloseNFTCard} className='cursor-pointer' xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="40" height="40" viewBox="0 0 50 50">
                              <path d="M 25 2 C 12.309534 2 2 12.309534 2 25 C 2 37.690466 12.309534 48 25 48 C 37.690466 48 48 37.690466 48 25 C 48 12.309534 37.690466 2 25 2 z M 25 4 C 36.609534 4 46 13.390466 46 25 C 46 36.609534 36.609534 46 25 46 C 13.390466 46 4 36.609534 4 25 C 4 13.390466 13.390466 4 25 4 z M 32.990234 15.986328 A 1.0001 1.0001 0 0 0 32.292969 16.292969 L 25 23.585938 L 17.707031 16.292969 A 1.0001 1.0001 0 0 0 16.990234 15.990234 A 1.0001 1.0001 0 0 0 16.292969 17.707031 L 23.585938 25 L 16.292969 32.292969 A 1.0001 1.0001 0 1 0 17.707031 33.707031 L 25 26.414062 L 32.292969 33.707031 A 1.0001 1.0001 0 1 0 33.707031 32.292969 L 26.414062 25 L 33.707031 17.707031 A 1.0001 1.0001 0 0 0 32.990234 15.986328 z"></path>
                              </svg>
                            </div>
                            <div className="flex pb-4">
                              <Mint_nft />
                            </div>
                          </div>
                        </Dialog.Panel>
                    </Transition.Child>
                </div>
            </div>
        </Dialog>
    </Transition>
    </Box>
  );
}
