import { ReactJSXElement } from "@emotion/react/types/jsx-namespace";
import { GetRecords200Response, Record } from "api/api";
import styled from "styled-components";

const ListComponent = styled.div`
    background: #4c495b;
    border-radius: 15px;
    width: 100%;
    padding: 1rem 1.8rem;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
`;
type PropsType = {
    recordList: Record[] | undefined;
};
type ResultListType = {
    name: string;
    score: number;
};
function ResultList(props: PropsType) {
    const result = [] as ReactJSXElement[];

    const setScore = () => {
        props.recordList?.map((l: Record) => {
            result.push(
                <ListComponent className="pretendard_regular_16" key={`${l.id}_${l.problem_id}`}>
                    <div>{l.id}</div>
                    <div>{l.readability}</div>
                </ListComponent>,
            );
        });
        return result;
    };
    return <div className="flex max-h-72 w-full flex-col gap-4 overflow-y-auto">{setScore()}</div>;
}

export default ResultList;
