import { defaultWagmiConfig } from "@web3modal/wagmi/react/config";

import { cookieStorage, createStorage } from "wagmi";
import { mainnet, arbitrumSepolia } from "wagmi/chains";

// Get projectId at https://cloud.walletconnect.com
export const projectId = process.env.NEXT_PUBLIC_PROJECT_ID;
console.log(projectId);
if (!projectId) throw new Error("Project ID is not defined");

const metadata = {
  name: "Konan",
  description: "Konan is a DApp integrated with Celestia and Arbitrun Sepolia chain.",
  url: "https://konan.pages.dev",
  icons: ["/konan_logo.webp"],
};

// Create wagmiConfig
const chains = [arbitrumSepolia, mainnet] as const;
export const config = defaultWagmiConfig({
  chains,
  projectId,
  metadata,
  ssr: true,
  enableEmail: true,
  storage: createStorage({
    storage: cookieStorage,
  }),
  // ...wagmiOptions // Optional - Override createConfig parameters
});
