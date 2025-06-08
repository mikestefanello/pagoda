import { HeroGeometric } from "@/components/ui/shape-landing-hero";
import PublicLayout from "@/Layouts/PublicLayout";

export default function Home() {
  return (
    <PublicLayout>
      <HeroGeometric
        badge="Pagode"
        title1="Elevate Your"
        title2="Fullstack Game"
      />
    </PublicLayout>
  );
}
