import { generateProblem } from "./problem.api";

// 비동기 api 로직은 여기서 구현한다.

const wrapPromise = (promise: Promise<unknown>) => {
    let status = "pending";
    let result = null;
    const suspender = promise.then(
        (r) => {
            if (r) {
                status = "success";
                result = r;
            } else {
                status = "pending";
                result = r;
            }
        },
        (e) => {
            status = "error";
            result = e;
        },
    );

    console.log("wrapPromise", status);
    const read = () => {
        switch (status) {
            case "pending":
                throw suspender;
            case "error":
                throw result;
            default:
                return result;
        }
    };

    return { read };
};
// const fetchPosts = () => {
//     return new Promise((resolve) => {
//             resolve([
//                 { id: 0, text: "this is first" },
//                 { id: 1, text: "this is second" },
//                 { id: 2, text: "this is third" },
//             ]);
//     });
// };

// suspece를 상요하는 api들.

export default function fetchProblem() {
    return {
        // language: string, difficulty: number
        create: () => wrapPromise(generateProblem().create("go", 50)),
        problem: (requestId: string) => wrapPromise(generateProblem().get(requestId)),
    };
}
