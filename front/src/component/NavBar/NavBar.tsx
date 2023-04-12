import { getProblem, postProblem } from "@src/api/problem";
import { Spinner } from "flowbite-react";
import { useState } from "react";

type PropsType = {
    currentLang: LanguageType;
    getProblem: () => void;
};
export default function NavBar(props: PropsType) {
    const [isSpinner, setSpinner] = useState<boolean>(false);
    const post = () => {
        // e.preventDefault();

        if (props.currentLang) {
            return postProblem(props.currentLang as LanguageType, 90);
        }
        return postProblem(props.currentLang as LanguageType, 90);
    };
    const onClickCreateProblem = async () => {
        const base64 =
            "UHJvYmxlbSBTdGF0ZW1lbnQ6CgpZb3UgYXJlIGdpdmVuIGFuIHVuc29ydGVkIGFycmF5IG9mIGludGVnZXJzIGFuZCB5b3UgbmVlZCB0byBmaW5kIHRoZSBsZW5ndGggb2YgdGhlIGxvbmdlc3QgaW5jcmVhc2luZyBzdWJzZXF1ZW5jZS4KCldyaXRlIGEgZnVuY3Rpb24gYGxvbmdlc3RJbmNyZWFzaW5nU3Vic2VxdWVuY2UobnVtczogbnVtYmVyW10pYCB0byBpbXBsZW1lbnQgdGhpcy4KCiMjIEV4YW1wbGUgVXNhZ2UKCmBgYGphdmFzY3JpcHQKbG9uZ2VzdEluY3JlYXNpbmdTdWJzZXF1ZW5jZShbMTAsIDksIDIsIDUsIDMsIDcsIDEwMSwgMThdKTsgLy8gT3V0cHV0OiA0Cmxvbmdlc3RJbmNyZWFzaW5nU3Vic2VxdWVuY2UoWzAsIDEsIDAsIDMsIDIsIDNdKTsgLy8gT3V0cHV0OiA0Cmxvbmdlc3RJbmNyZWFzaW5nU3Vic2VxdWVuY2UoWzcsIDcsIDcsIDcsIDcsIDcsIDddKTsgLy8gT3V0cHV0OiAxCmBgYAoKIyMgQ29uc3RyYWludHM6CgoqIFRoZSBpbnB1dCBhcnJheSBgbnVtc2Agd2lsbCBjb250YWluIGF0IG1vc3QgYDI1MDBgIGludGVnZXJzLgoqIEFsbCBvZiB0aGUgaW50ZWdlcnMgaW4gYG51bXNgIGFyZSB3aXRoaW4gdGhlIHJhbmdlIGBbMSwgMl4zMSAtIDFdYC4KCiMjIEV2YWx1YXRpb24gQ3JpdGVyaWE6CgoqIFlvdXIgaW1wbGVtZW50YXRpb24gc2hvdWxkIGJlIGNvcnJlY3QgYW5kIHNob3VsZCBoYW5kbGUgYWxsIGVkZ2UgY2FzZXMuCgoqIFlvdXIgc29sdXRpb24gc2hvdWxkIGhhdmUgYSB0aW1lIGNvbXBsZXhpdHkgb2YgYE8obiBsb2cgbilgIG9yIGJldHRlci4gCgoqIFlvdXIgc29sdXRpb24gc2hvdWxkIHVzZSBkYXRhIHN0cnVjdHVyZXMgZWZmaWNpZW50bHku";
        props.getProblem(atob(base64));
        return;
        try {
            setSpinner(true);
            const problemId = await post();
            if (problemId.id) {
                const problemInterval = setInterval(async () => {
                    const problem = await getProblem(problemId.id);
                    if (problem) {
                        setSpinner(false);
                        console.log(atob(problem.content));
                        props.getProblem(atob(problem.content));
                        clearInterval(problemInterval);
                    }
                }, 5000);
            }
        } catch (err) {
            console.log("onClickCreateProblem : ", err);
        }
    };
    return (
        <nav className="bg-Navy-900">
            <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
                {/* logo */}
                <a href="#" className="flex flex-row items-start gap-4">
                    <div className="bg-coco-green_500 h-6 w-6" />
                    <span className="self-center text-2xl  font-semibold whitespace-nowrap text-white">
                        codeconnect
                    </span>
                    <button onClick={onClickCreateProblem} className="bg-coco-green_500 flex flex-row gap-2">
                        <span>문제 제출</span>
                        {isSpinner ? <Spinner color="info" /> : null}
                    </button>
                </a>
            </div>
        </nav>
    );
}
