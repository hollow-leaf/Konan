import { Navbar, Footer } from "@/components";
import { Hero, Business } from "@/app/components";

export default function Profile() {
  return (
    <div className="bg-[url('/bg.png')] bg-cover bg-center">
      <Navbar />
      <Hero />
      <Business />
      <Footer />
    </div>
  );
}
