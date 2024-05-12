"use client";
import React, { useState, Fragment } from "react";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Typography from "@mui/material/Typography";
import { CardActionArea } from "@mui/material";
import { nftIPFS } from "@/content/ipfs/ipfs-nft";
import { Dialog, Transition } from '@headlessui/react'
import { nftOwn } from "@/types/types";


export function NFTCard(props: { nft: nftOwn }) {
  const [isOpen, sstIsOpen] = useState<boolean>(false)
  
  function closeModal() {
    sstIsOpen(false)
  }

  return (
    <div className="p-5 rounded-xl w-2/12	shadow-lg m-2" onClick={()=>{sstIsOpen(true)}}>
      <div className="rounded-xl cursor-pointer">
        <img className="ms-auto max-h-xs rounded-xl" src={nftIPFS[1].src} alt="image description" />
        <p className="p-2">{nftIPFS[1].name}</p>
      </div>
      <Transition appear show={isOpen} as={Fragment}>
        <Dialog as="div" className="relative z-10" onClose={closeModal}>
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
                        <Dialog.Panel className="w-full max-w-4xl transform overflow-hidden rounded-2xl bg-white text-left align-middle shadow-xl transition-all">
                          <div className="py-2 px-2" style={{"background": "linear-gradient(90deg, transparent 445px, #fff 0) top left, url('/background1.png')"}}>
                            <div className="flex justify-end mt-2">
                              <svg onClick={closeModal} className='cursor-pointer' xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="40" height="40" viewBox="0 0 50 50">
                              <path d="M 25 2 C 12.309534 2 2 12.309534 2 25 C 2 37.690466 12.309534 48 25 48 C 37.690466 48 48 37.690466 48 25 C 48 12.309534 37.690466 2 25 2 z M 25 4 C 36.609534 4 46 13.390466 46 25 C 46 36.609534 36.609534 46 25 46 C 13.390466 46 4 36.609534 4 25 C 4 13.390466 13.390466 4 25 4 z M 32.990234 15.986328 A 1.0001 1.0001 0 0 0 32.292969 16.292969 L 25 23.585938 L 17.707031 16.292969 A 1.0001 1.0001 0 0 0 16.990234 15.990234 A 1.0001 1.0001 0 0 0 16.292969 17.707031 L 23.585938 25 L 16.292969 32.292969 A 1.0001 1.0001 0 1 0 17.707031 33.707031 L 25 26.414062 L 32.292969 33.707031 A 1.0001 1.0001 0 1 0 33.707031 32.292969 L 26.414062 25 L 33.707031 17.707031 A 1.0001 1.0001 0 0 0 32.990234 15.986328 z"></path>
                              </svg>
                            </div>
                            <div className="flex px-16 pb-4">
                              <div className="rounded-lg shadow-2xl p-2 max-h-[500px]">
                                <img className="ms-auto max-w-xs" src="/main1.png" alt="image description" />
                                <div className="absolute bottom-5 left-8 rounded-lg bg-orange-100 py-6 px-16"><p className="freeman-regular text-3xl font-medium">Dinosaur #1</p></div>
                              </div>
                              <div className="py-6 pl-14 w-full">
                                <div className="px-2 min-h-48">
                                  <p className="text-2xl font-medium font-sans">Dinosaur #1</p>
                                  <p className="text-lg font-sans text-orange-400">Created by SoloLin</p>
                                  <p className="mt-1 text-xs font-sans">arbitrum sepolia:</p>
                                  <p className="text-xs font-sans">{"0xf3419771c2551f88a91Db61cB874347f05640172"}</p>
                                </div>
                                <div>
                                  <p className="p-2 text-2xl font-medium font-sans">Accessories</p>
                                </div>
                                <div className="flex flex-wrap justify-around">
                                  <div className="flex justify-center items-center">
                                    <div className="flex absolute w-40 h-40 z-50 opacity-0 text-2xl font-medium cursor-pointer hover:opacity-90 rounded-xl items-center justify-center bg-black/40">UnLock!</div>
                                    <div className="flex absolute w-36 h-36 z-40 opacity-30 text-2xl font-medium cursor-pointer rounded-xl items-center justify-center bg-center bg-cover bg-[url('/sunglasses.png')] bg-black/10"></div>
                                    <img className="rounded-lg border-2 w-40 h-40 opacity-40 p-2" src="https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z" alt="image description" />
                                  </div>
                                  <div className="flex justify-center items-center">
                                    <div className="flex absolute w-40 h-40 z-50 opacity-0 text-2xl font-medium cursor-pointer hover:opacity-90 rounded-xl items-center justify-center bg-black/40">UnLock!</div>
                                    <div className="flex absolute w-36 h-36 z-40 opacity-30 text-2xl font-medium cursor-pointer rounded-xl items-center justify-center bg-center bg-cover bg-[url('/background.png')] bg-black/10"></div>
                                    <img className="rounded-lg border-2 w-40 h-40 opacity-40 p-2" src="https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z" alt="image description" />
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                        </Dialog.Panel>
                    </Transition.Child>
                </div>
            </div>
        </Dialog>
    </Transition>
    </div>
    
  );
}
