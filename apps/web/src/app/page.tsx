"use client";
import { Navbar, Footer, Hero, Mint_nft, Business, Team } from "@/components";
import { Box } from "@mui/material";

export default function Profile() {
  return (
    <Box className="bg-gradient-to-br from-purple-200 via-purple-100 to-purple-50">
      <Navbar />
      <Hero />
      <Mint_nft />
      <Business />
      <Team />
      <Footer />
    </Box>
  );
}
