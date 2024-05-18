"use client";
import React, { useState, Fragment, useEffect } from "react";
import { nftIPFS } from "@/content/ipfs/ipfs-nft";
import { Dialog, Transition } from '@headlessui/react'
import { nftOwn } from "@/types/types";
import { Draw } from "./draw";
import { useReadContract } from 'wagmi'
import { useQuery } from "@tanstack/react-query";
import { blackstoneABI } from "@/contracts/blackstone";
import { Loading } from "../Loading";
import { readContract } from '@wagmi/core'
import { config } from "@/config/wagmi.config";
import { getNftImageByIndex } from "@/services/api";



export function NFTCard(props: { nft: nftOwn }) {
  const [isOpen, sstIsOpen] = useState<boolean>(false)
  const [uri, setURI] = useState<string>("https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z")

  function closeModal() {
    sstIsOpen(false)
  }

  async function getNftImage(nftId: number): Promise<string> {
    const result:any = await readContract(config, {
      abi: blackstoneABI,
      address: '0x585a1DDaB9116F483d367bCa6eb64797252051c8',
      functionName: 'tokenURI',
      args: [props.nft.nftId],
    })
    if(result) {
      const index = result.split("/")[3].split(".")[0]
      const uri = await getNftImageByIndex(Number(index))
      return uri
    }
    return ""
  }

  useEffect(() => {
    getNftImage(props.nft.nftId).then(res => {
      if(res) {
        console.log(res)
        setURI("https://gateway.pinata.cloud/ipfs/" + res)
      }
    })
  })

  return (
    <div className="p-5 rounded-xl w-2/12	shadow-lg m-2 min-w-[250px]" onClick={()=>{sstIsOpen(true)}}>
      <div>
      {false || false ?<></>:
        <div className="rounded-xl cursor-pointer flex flex-col content-center">
          <img className="ms-auto max-h-xs rounded-xl" src={uri} alt="image description" />
          {uri != "https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z"
          ?
          <p className="text-2xl font-medium font-sans p-2 text-center">Dinosaur #{props.nft.nftId}</p>
          :
          <p className="text-2xl font-medium font-sans p-2 text-center">Dinosaur</p>
          }
        </div>
      }
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
                              <div className="rounded-lg shadow-2xl max-h-[0px]">
                                <img className="rounded-xl ms-auto max-w-xs" src={uri} alt="image description" />
                              </div>
                              <div className="py-6 pl-14 w-full">
                                <div className="px-2 min-h-48">
                                  <p className="text-2xl font-medium font-sans">Dinosaur #{props.nft.nftId}</p>
                                  <p className="text-lg font-sans text-orange-400">Created by Sophia</p>
                                  <p className="mt-1 text-xs font-sans">arbitrum sepolia:</p>
                                  <p className="text-xs font-sans">{"0xf3419771c2551f88a91Db61cB874347f05640172"}</p>
                                </div>
                                {uri != "https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z"
                                ?
                                <><div>
                                  <p className="p-2 text-2xl font-medium font-sans">Accessories</p>
                                </div>
                                <div className="flex flex-wrap justify-around">
                                  <Draw tokenId={props.nft.nftId} drawId={1} url={"https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z"} type="Background" />
                                  <Draw tokenId={props.nft.nftId} drawId={2} url={"https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z"} type="Glasses"/>
                                </div></>
                                :
                                <div>
                                  <Draw tokenId={props.nft.nftId} drawId={0} url={"https://gateway.pinata.cloud/ipfs/QmeobMV4X8WZNfQJhmUgMFh3gZ6rAsWD63wVuENazZM24Z"} type="Dinosaur" />
                                </div>}
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
