import axios from "axios";

const instance = axios.create({
    baseURL: "http://localhost:8080",
    // timeout: 1000,
    // headers: {'X-Custom-Header': 'foobar'}
});
const headers = { "Content-Type": `application/json` };

const apiErrorHandler = (error) => {
    if (error.response) {
        // 요청이 전송되었고, 서버는 2xx 외의 상태 코드로 응답했습니다.
        console.log(error.response.data);
        console.log(error.response.status);
        console.log(error.response.headers);
    } else if (error.request) {
        // 요청이 전송되었지만, 응답이 수신되지 않았습니다.
        // 'error.request'는 브라우저에서 XMLHtpRequest 인스턴스이고,
        // node.js에서는 http.ClientRequest 인스턴스입니다.
        console.log(error.request);
    } else {
        // 오류가 발생한 요청을 설정하는 동안 문제가 발생했습니다.
        console.log("Error", error.message);
    }
    console.log(error.config);
};

export const analyze = async (data: string) => {
    const utf8Str = unescape(encodeURIComponent(data));
    const base64 = btoa(utf8Str);
    try {
        const response = await instance.post("/analyze", { base64 }, { headers });
    } catch (error) {
        apiErrorHandler(error);
    }
};
