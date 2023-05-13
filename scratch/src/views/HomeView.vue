<template>
  <div class="flex w-screen h-screen">
    <div class="flex flex-col w-[1200px] p-[100px]">
      <div class="flex py-4">
        <div class="mb-4">
          <h2 class="text-4xl mb-4">{{ title }}</h2>
        </div>
        <div class="ml-auto text-right">
          <p class="text-lg mb-0.5">Estimated Time</p>
          <p class="text-4xl mb-0.5">{{ estimatedTime }}</p>
          <p class="text-lg">minutes</p>
        </div>
      </div>
      <p class="text-xl mb-4">{{ background }}</p>
      <div class="mb-4 w-full h-[600px]">
        <textarea
          placeholder="Code"
          class="w-full h-full p-2 border border-gray-400"
          v-model="targetCode"
        ></textarea>
      </div>
      <div class="flex mb-4">
        <button
          class="px-6 py-2 bg-red-500 text-white"
          @click="generateProblem"
          :disabled="generating"
        >
          Generate
        </button>
        <p class="mx-4 my-auto">
          {{ generated ? "Generated" : "Generating..." }}
        </p>
        <button class="ml-auto px-6 py-2 bg-indigo-500 text-white">
          Submit
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { DefaultApi } from "../../api";
import { ref } from "vue";
import axios from "axios";

const requestID = ref<string>("");
const title = ref<string>("Title");
const background = ref<string>("Background");
const targetCode = ref<string>("package main");
const estimatedTime = ref<number>(30);

const generated = ref<boolean>(false);
const generating = ref<boolean>(false);

const instance = axios.create();

const intervalID = ref<number>(0);

const generateProblem = async () => {
  const api = new DefaultApi(undefined, "http://localhost:3000", instance);
  const resp = await api.createProblem({
    language: "Go",
    estimated_time: 10,
  });

  generating.value = true;

  const { request_id } = resp.data;
  if (!request_id) {
    return;
  }

  requestID.value = request_id;

  console.log(request_id);

  intervalID.value = setInterval(async () => {
    try {
      const resp = await api.getProblem(requestID.value);
      console.log(resp);

      console.log(resp.data);

      const {
        title: _title,
        background: _background,
        code,
        estimated_time,
      } = resp.data.problem;
      title.value = _title;
      background.value = _background;
      estimatedTime.value = estimated_time;

      // Decode base64
      targetCode.value = atob(code);

      generated.value = true;
      generating.value = false;

      clearInterval(intervalID.value);
    } catch (e) {
      console.log(e);
    }
  }, 10000);
};
</script>
