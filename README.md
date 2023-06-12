# DeveloPic
존립을 위협받고 있는 이름이지만 일단 어쩔 수 없으니 유지

## 서버 실행 방법
`CHATGPT_API_KEY` 환경변수를 OpenAI API Key로 설정. Key는 프로젝트 리더에게서 발급받거나 개인 계정 사용.
```bash
export CHATGPT_API_KEY=...
```

다음 명령어로 서버를 실행 (Go 1.20.2 이상의 버전이 설치되어 있어야 함)
```bash
cd backend/gateway/cmd && go run main.go
```

`localhost:3000` 으로 접속하면 API 이용 가능
```bash
curl --request GET -sL \
     --url 'http://localhost:3000/problems/234'
```
