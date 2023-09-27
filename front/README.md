# README.md 작성 문서

# Developic-web

## **Description**

- 해당 서비스 프롬프트 엔지니어링을 활용하여 개발자들에게 언어와 난이도에 맞는 리팩토링 과제를 던져주어 리팩토링 공부를 할 수 있도록 도와주는 서비스입니다.

**Skills**

- React & Recoil
- styled-compoenet + Tailwindcss
- E2E : playwright
- yarn berry(pnp)
- msw
- swc

I**ssue**

- yarn berry pnp mode - **이슈로 인해 중단**
  **사용 이유**
  - zero-install로 인한 CI 속도 향상
  **문제**
  - zero-install 안되는 문제
    - `swc`의 경우 운영체제에 종속되는 부분이 있다보니 커밋에 포함시킬 경우 실행환경에 따라 문제를 일으킬 수 있어 커밋에서 제외되고 있습니다. 따라서 swc를 사용한다면 정상적인 빌드를 위해 최초 1회 설치 명령어가 필요합니다.
  - definition을 찾지못하는 문제.
    - definition을 찾지못하는 문제가 발생하여 개발 도중 import가 전부 빨간 줄이 생겨 마련한 대책은 일단 local에서는 node_modules를 사용하고 레포지토리에는 pnp모드로 올려 zero-install을 사용하고있다.
- React, Styled, Recoil과 같은 기술 경험에 의의를 두어 관심사의 분리, 재사용성을 제대로 고려하지 않은 코드가 생겼다.
- Tailwindcss를 제거하고 emotion 혹은 styled-component를 사용하여 리팩토링을 할 예정

**~~Cumtom hook 패턴 사용~~**

- ~~기존 Presentation Component - Container Component의 Container의 로직만 hooks로 관리하는 방법. - 단일 책임을 지지않고 있기에 리팩토링 예정.~~

**Compound 패턴으로 리팩토링 중**

- 리팩토링 과정은 노션으로 기록 중입니다. - [Link](https://www.notion.so/Developic-Refactoring-4ada551e41004a4eb609f4e92809c4f2?pvs=21)

## **Environment**

### **Installation**

```jsx
yarn && npm i
```

```jsx
yarn dev && npm run dev
```

### **Prerequisite**

node version v18.16.0 사용 중.

### **Files**

이 항목은 내가 추가한 것이다. 중요한 코드 파일들 몇 개를 대상으로 해당 파일이 어떠한 역할을 하는 파일인지를 간단히 설명해주면 전반적인 맥락을 파악하기에 좋을 것 같아 추가하였다.

### **Usage**

- yarn dev로 프로젝트를 실행하면 msw가 같이 실행됩니다. 루트 디렉토리에 .env 파일을 생성하여 API_KEY를 설정해주면 프로젝트가 정상 작동합니다.

.env

```jsx
API_KEY = "https://local.kr";
```
