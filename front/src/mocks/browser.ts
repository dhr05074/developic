import { setupWorker } from "msw";
import handlers from "./handlers.ts";

export const worker = setupWorker(...handlers); // worker: 가짜 서버