import React from "react"

import type { Metadata, Viewport } from "next"
import { Roboto } from "next/font/google"
import { Layout } from "@/components"
import "./globals.css"

import { headers } from 'next/headers'

import { cookieToInitialState } from 'wagmi'

import { config } from '@/utils/config/wagmi.config'
import Web3ModalProvider from '@/utils'

const roboto = Roboto({
  subsets: ["latin"],
  weight: ["300", "400", "500", "700", "900"],
  display: "swap",
})

export const metadata: Metadata = {
  manifest: "/manifest.json",
  title: "Konan",
  description: "Konan website for web3 hackathon",
  icons: {
    icon: "/konan_logo.webp",
  },
}

export const viewport: Viewport = {
  themeColor: "#3367D6",
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  const initialState = cookieToInitialState(config, headers().get('cookie'))
  return (
    <html lang="en">
      <head>
        <link rel="shortcut icon" href="/konan_logo.webp" type="image/webp" />
      </head>
      <body className={roboto.className}>
      <Web3ModalProvider initialState={initialState}>
        <Layout>
          {children}
        </Layout>
      </Web3ModalProvider>
      </body>
    </html>
  )
}
