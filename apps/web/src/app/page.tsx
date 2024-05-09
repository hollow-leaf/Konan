import { Navbar, Footer, Hero, Team } from "@/components";
import { Business, Mint_nft  } from "@/app/components";

export default function Profile() {
  return (
    <div className="bg-white bg-cover bg-center">
      <Navbar />
      <Hero />
      <Mint_nft />
      <Business />
      <Team />
      <Footer />
    </div>
  );
}
