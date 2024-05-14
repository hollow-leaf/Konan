"use client";
import { WagmiProvider, createConfig, http } from "wagmi";
import { mainnet, arbitrumSepolia } from "wagmi/chains";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ConnectKitProvider, getDefaultConfig } from "connectkit";
import { ReactNode } from "react";

const walletConnectProjectId =
  process.env.NEXT_PUBLIC_WALLETCONNECT_PROJECT_ID || "";
export const config = createConfig(
  getDefaultConfig({
    // Your dApps chains
    chains: [arbitrumSepolia],
    ssr: true,
    transports: {
      // RPC URL for each chain
      [arbitrumSepolia.id]: http(
        `https://eth-mainnet.g.alchemy.com/v2/${process.env.NEXT_PUBLIC_ALCHEMY_ID}`,
      ),
    },

    // Required API Keys
    walletConnectProjectId,

    // Required App Info
    appName: "Konan",

    // Optional App Info
    appDescription: "2024 5 Celestia Hackathon",
    appUrl: "https://konan.pages.dev/", // your app's url
    appIcon: "https://family.co/logo.png", // your app's icon, no bigger than 1024x1024px (max. 1MB)
  }),
);

const queryClient = new QueryClient();

export const Web3Provider = ({ children }: { children: ReactNode }) => {
  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>
        <ConnectKitProvider debugMode>{children}</ConnectKitProvider>
      </QueryClientProvider>
    </WagmiProvider>
  );
};
