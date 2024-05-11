import { Navbar, Footer, Hero, Mint_nft, Business, Team } from "@/components";
import { Box } from "@mui/material";

export default function Profile() {
  return (
    <Box className="bg-white bg-cover bg-center">
      <Navbar />
      <Hero />
      <Mint_nft />
      <Business />
      <Team />
      <Footer />
    </Box>
  );
}
