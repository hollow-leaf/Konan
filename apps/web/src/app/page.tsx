import { Navbar, Footer, Hero } from "@/components";
import { Business, Mint_nft, Team  } from "@/app/components";

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
