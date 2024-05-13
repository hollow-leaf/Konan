"use client";
import { Navbar, Footer, Hero, Mint_nft, Business, Team } from "@/components";
import { Box } from "@mui/material";

export default function Profile() {
  return (
    <div className="" style={{"backgroundColor": "#f2c1f2", "backgroundImage": "linear-gradient(160deg, #f2c1f2 0%, #ffffff 16%, #ffffff 99%)"}}>
      <Navbar />
      <Hero />
      <Business />
      <Team />
      <Footer />
    </div>
  );
}
