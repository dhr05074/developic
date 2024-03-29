import ResultList from "@/component/List/Result.List";
import styled from "styled-components";
import Rectangle from "@/assets/images/Rectangle.png";
import useProfile from "../hook/Profile.hook";
import { useEffect } from "react";

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
const Button = styled.button`
    border-radius: 30px;
    border: solid 1px #1f1c32;
    background: ${(props) => (props.disabled ? "gray" : "#b9ff47")};
    color: black;
    cursor: ${(props) => (props.disabled ? "not-allowed" : "pointer")};
    width: 60%;
    &: hover
        ${(props) =>
            props.disabled
                ? null
                : `  {
        border: solid 1px #b9ff47;
        background: #1f1c32;
        color: #b9ff47;
    }`};
`;

const gap = 4;

function Profile() {
    const { profile, records, getProfile, getRecords } = useProfile();
    useEffect(() => {
        getProfile();
        getRecords();
    }, []);
    return (
            <div className=" flex h-full  w-full  flex-row justify-center bg-Navy-900 text-white">
                <div className={`flex h-full w-[400px] flex-col justify-center gap-${gap}`}>
                    <section id="result_title" className="flex flex-col items-center gap-3">
                        <h2 className="pretendard_extrabold_32">{profile.nickname}</h2>
                        <img src={Rectangle} className="h-2 w-2" />
                        <Value className="pretendard_bold_24">{profile.elo_score}</Value>
                    </section>
                    <ScoreSection id="result_score" className="pretendard_bold_20 flex flex-col items-center ">
                        <article className="flex w-full flex-row justify-between">
                            <Title>Latest Score</Title>
                        </article>
                        {records ? <ResultList recordList={records} /> : "없음"}

                        <section className="pretendard_medium_20 w-full">
                            <Button id="profile_More" disabled={true}>More</Button>
                        </section>
                    </ScoreSection>
                </div>
            </div>
    );
}

export default Profile;
