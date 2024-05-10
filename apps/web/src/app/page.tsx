import { Navbar, Footer, Hero, Mint_nft } from "@/components";
import { Business, Team } from "@/app/components";

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
