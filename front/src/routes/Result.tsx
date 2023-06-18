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
    border: solid 1px #b9ff47;
    background: #1f1c32;
    color: #b9ff47;
    width: 50%;
    &:hover {
        background: #b9ff47;
        color: black;
    }
`;

const gap = 4;

const listSample = [
    { name: "가독성", score: 100 },
    { name: "모듈성", score: 100 },
    { name: "효율성", score: 100 },
    { name: "테스트 가능성", score: 100 },
    { name: "유지 보수 가능성", score: 100 },
];

function Result() {
    return (
        <motion.div
            className="h-full w-full"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <div className=" flex h-full  w-full  flex-row justify-center bg-Navy-900 text-white">
                <div className={`flex h-full w-[464px] flex-col justify-center gap-${gap}`}>
                    <section id="result_title">
                        <h2 className="pretendard_extrabold_32">Result</h2>
                    </section>
                    <section
                        id="result_complexity"
                        className={`pretendard_bold_20 flex flex-row gap-${gap} w-full justify-center`}
                    >
                        <ComplexityBox id="result_complexity_execution">
                            <Title>Execution Time</Title>
                            <Value>value</Value>
                        </ComplexityBox>
                        <ComplexityBox id="result_complexity_memory">
                            <Title>Memory Usage</Title>
                            <Value>value</Value>
                        </ComplexityBox>
                    </section>
                    <ScoreSection id="result_score" className="pretendard_bold_20 flex flex-col items-center ">
                        <article className="flex w-full flex-row justify-between">
                            <Title>Score</Title>
                            <Value>value</Value>
                        </article>
                        <img src={Rectangle} className="h-2 w-2" />
                        {/* <ResultList list={listSample} /> */}
                    </ScoreSection>
                    <section className="pretendard_medium_20 mt-1 w-full">
                        <LineButton>처음으로</LineButton>
                    </section>
                </div>
            </div>
        </motion.div>
    );
}

export default Result;
