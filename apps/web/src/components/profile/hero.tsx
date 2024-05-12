"use client";
import React, { Fragment, useEffect, useState } from "react";
import { Dialog, Transition } from '@headlessui/react'
import { useAccount, useBalance } from "wagmi";
import { Loading } from "@/components";
import { useRouter } from "next/navigation";
import { NFTCard } from "./NFTCard";


export function Hero() {
  const router = useRouter()

  const [isOpen, sstIsOpen] = useState<boolean>(false)
  const [loading, setLoading] = useState<boolean>(false)

  const { address, isConnected, chain } = useAccount();

  useEffect(() => {
    if (!isConnected) {
      router.push('/')
    }
  }, [isConnected]);

  function closeModal() {
    sstIsOpen(false)
  }

  return (
    <>
      <div className="w-full min-h-72 bg-center bg-cover bg-[url('/background1.png')]"></div>
      <div className="flex items-end relative w-full h-10 bg-black px-32 bg-white">
        <div className="bg-orange-200 rounded-full p-1">
          <img className="rounded-full w-40 h-40" src="/main1.png" alt="image description" />
        </div>
      </div>
      <div className="px-32 py-6">
       
      </div>
      <div className="flex px-20 ">
      <NFTCard index={1}/>
      <NFTCard index={1}/>
      <NFTCard index={1}/>
      <NFTCard index={1}/>
      <NFTCard index={1}/>
      <NFTCard index={1}/>
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
                        <Dialog.Panel className="w-full max-w-md transform overflow-hidden rounded-2xl bg-white text-left align-middle shadow-xl transition-all">
                    
                        </Dialog.Panel>
                    </Transition.Child>
                </div>
            </div>
        </Dialog>
    </Transition>
    </>
  )
}
