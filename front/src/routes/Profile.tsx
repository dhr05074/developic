import ResultList from "@/component/List/Result.List";
import { motion } from "framer-motion";
import styled from "styled-components";
import Rectangle from "@/assets/images/Rectangle.png";

const ComplexityBox = styled.article`
    background: #353346;
    border-radius: 15px;
    width: 100%;
    padding: 2rem 2rem;
    display: flex;
    flex-direction: column;
    gap: 20px;
`;
const Title = styled.h3`
    text-align: left;
`;
const Value = styled.p`
    text-align: right;
    color: #b9ff47;
`;
const ScoreSection = styled.section`
    background: #353346;
    border-radius: 15px;
    width: 100%;
    padding: 2rem 1.8rem;
    display: flex;
    flex-direction: column;
    gap: 20px;
`;
const LineButton = styled.button`
    border-radius: 30px;
    border: solid 1px #1f1c32;
    background: #b9ff47;
    color: black;

    width: 60%;
    &:hover {
        border: solid 1px #b9ff47;
        background: #1f1c32;
        color: #b9ff47;
    }
`;

const gap = 4;

const listSample = [
    { name: "프라블럼", score: 100 },
    { name: "프라블럼", score: 100 },
    { name: "프라블럼", score: 100 },
];

function Profile() {
    return (
        <motion.div
            className="h-full w-full"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <div className=" flex h-full  w-full  flex-row justify-center bg-Navy-900 text-white">
                <div className={`flex h-full w-[400px] flex-col justify-center gap-${gap}`}>
                    <section id="result_title" className="flex flex-col items-center gap-3">
                        <h2 className="pretendard_extrabold_32">이름 넣기</h2>
                        <img src={Rectangle} className="h-2 w-2" />
                        <Value className="pretendard_bold_24">mmr</Value>
                    </section>
                    <ScoreSection id="result_score" className="pretendard_bold_20 flex flex-col items-center ">
                        <article className="flex w-full flex-row justify-between">
                            <Title>Latest Score</Title>
                        </article>
                        <ResultList list={listSample} />
                        <section className="pretendard_medium_20 w-full">
                            <LineButton>More</LineButton>
                        </section>
                    </ScoreSection>
                </div>
            </div>
        </motion.div>
    );
}

export default Profile;
