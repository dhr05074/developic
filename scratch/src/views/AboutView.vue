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
