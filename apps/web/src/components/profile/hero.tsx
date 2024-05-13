"use client";
import React, { Fragment, useEffect, useState } from "react";
import { Dialog, Transition } from '@headlessui/react'
import { useAccount, useBalance } from "wagmi";
import { Loading } from "@/components";
import { useRouter } from "next/navigation";
import { NFTCard } from "./NFTCard";
import { getNftosByAddress } from "@/services/api";
import { nftOwn } from "@/types/types";
import { useQuery } from "@tanstack/react-query";
import { time } from "console";


export function Hero() {
  const router = useRouter()

  const [isOpen, sstIsOpen] = useState<boolean>(false)
  const [loading, setLoading] = useState<boolean>(false)
  const [nfts, setNfts] = useState<nftOwn[]>([])

  const { address, isConnected, chain } = useAccount();

  useEffect(() => {
    if (!isConnected) {
      router.push('/')
    }
  }, [isConnected]);

  const { isLoading, error, data } = useQuery({
    queryKey: ["getNftosByAddress"],
    queryFn: async () => {
      const res = await getNftosByAddress("0x957A49308CAA2bfFEc32F1cA1c75Ce751BBBaAee")
      if(res) {
        console.log(res)
        setNfts(res)
      }
    }
  })

  function closeModal() {
    sstIsOpen(false)
  }

  const nftElements:any = nfts.map((nft, index) => {
    if(index <30) {
      return <NFTCard nft={nft}/>

    }
  })

  return (
    <>
      <div className="w-full min-h-72 bg-center bg-cover bg-[url('/background1.png')]"></div>
      <div className="flex items-end relative w-full h-10 bg-black px-48 bg-white">
        <div className="bg-orange-200 rounded-full p-1">
          <img className="rounded-full w-40 h-40" src="/main1.png" alt="image description" />
        </div>
      </div>
      <div className="px-48 py-6">
        <p>ID</p>
        <p>{address}</p>
      </div>
      <div className="flex px-48 justify-center	">
        <div className="flex flex-wrap w-10/12">
        {nftElements.length > 0 ? (nftElements) : (<><p>fuck uuuu</p></>)}
        </div>
      </div>
    </>
  )
}
