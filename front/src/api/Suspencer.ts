// // 비동기 api 로직은 여기서 구현한다.

// const wrapPromise = (promise: Promise<unknown>) => {
//     let status = "pending";
//     let result: null | unknown = null;
//     const suspender = promise
//         .then(
//             (r) => {
//                 if (r) {
//                     status = "success";
//                     result = r;
//                 } else {
//                     status = "pending";
//                     result = r;
//                 }
//             },
//             (e) => {
//                 status = "error";
//                 result = e;
//             },
//         )
//         .catch((e) => {
//             console.error("suspense", e);
//         });

//     console.log("wrapPromise", status);
//     const read = () => {
//         switch (status) {
//             case "pending": // 200에 아직 값이 안왔을 경우.
//                 throw suspender;
//             case "error": // 404
//                 throw result;
//             default:
//                 return result;
//         }
//     };

//     return { read };
// };
// export default wrapPromise;
// // const fetchPosts = () => {
// //     return new Promise((resolve) => {
// //             resolve([
// //                 { id: 0, text: "this is first" },
// //                 { id: 1, text: "this is second" },
// //                 { id: 2, text: "this is third" },
// //             ]);
// //     });
// // };

// // suspece를 상요하는 api들.
