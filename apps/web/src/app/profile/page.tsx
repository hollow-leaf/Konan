"use client";
import { Navbar, Footer } from "@/components";
import { Hero } from "@/components/profile";
import { Box } from "@mui/material";

export default function Profile() {
  return (
    <Box className="bg-white">
      <Navbar />
      <Hero />
      <Footer />
    </Box>
  );
}
