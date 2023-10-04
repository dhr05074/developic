import { PageHome } from "@/pages/Home/Usage";
import { Loading } from "@/pages/Loading";
import { Funnel } from "./Funnel";
import useFunnel from "./useFunnel";

function UsageFunnel() {
  const { setStep } = useFunnel();
  return (
    <Funnel>
      <Funnel.Step name="Home">
        <PageHome onNext={() => setStep("Loading")} />
      </Funnel.Step>
      <Funnel.Step name="Loading">
        <Loading onNext={() => setStep("Problem")} />
      </Funnel.Step>
    </Funnel>
  );
}

export { UsageFunnel };
