import { Navbar, Footer } from "@/components";
import { Hero } from "@/components/profile";

export default function Profile() {
  return (
    <div className="bg-white">
      <Navbar />
      <Hero />
      <Footer />
    </div>
  );
}
