import { Navbar, Footer } from "@/components";
import { Hero, Vision } from "@/app/components";

export default function Profile() {
  return (
    <div className="bg-white">
      <Navbar />
      <Hero />
      <Vision />
      <Footer />
    </div>
  );
}
