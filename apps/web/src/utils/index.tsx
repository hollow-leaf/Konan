"use client";

import React, { ReactNode } from "react";
import { config, projectId } from "@/utils/config/wagmi.config";

import { createWeb3Modal } from "@web3modal/wagmi/react";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import { State, WagmiProvider } from "wagmi";
import { Layout } from "@/components";
// Setup queryClient
const queryClient = new QueryClient();

if (!projectId) throw new Error("Project ID is not defined");

// Create modal
createWeb3Modal({
  wagmiConfig: config,
  projectId,
  enableAnalytics: true, // Optional - defaults to your Cloud configuration
  enableOnramp: true, // Optional - false as default
  themeVariables: {
    "--w3m-accent": "#6BE4FF",
    "--w3m-color-mix": "#FFABD8",
    "--w3m-color-mix-strength": 40,
    "--w3m-z-index": 5,
  },
});

export default function Web3ModalProvider({
  children,
  initialState,
}: {
  children: ReactNode;
  initialState?: State;
}) {
  return (
    <WagmiProvider config={config} initialState={initialState}>
      <QueryClientProvider client={queryClient}>
        <Layout>{children}</Layout>
      </QueryClientProvider>
    </WagmiProvider>
  );
}
