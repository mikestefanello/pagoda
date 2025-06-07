import { HeroGeometric } from "@/components/ui/shape-landing-hero";
import PublicLayout from "@/Layouts/public-layout";
import { FlashMessages } from "@/types/global";

type Props = {
  flash: FlashMessages;
};

export default function Home({ flash }: Props) {
  return (
    <PublicLayout flash={flash}>
      <HeroGeometric
        badge="Pagode"
        title1="Elevate Your"
        title2="Fullstack Game"
      />
    </PublicLayout>
  );
}
