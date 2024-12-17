<template>
  <div class="h-full flex flex-col">
    <div
      class="h-[40px] p-2 flex justify-between items-center bg-white rounded-md"
      style="border-bottom: 1px solid #e5e7eb"
    >
      <div class="flex items-end">
        <h2 class="text-xl font-bold mr-2">{{ props.title }}</h2>
        <p class="text-gray-300">{{ props.desc }}</p>
      </div>
      <slot name="actions"></slot>
    </div>
    <div class="h-[calc(100%-40px)] p-2 bg-white rounded-md">
      <slot name="content"></slot>
    </div>
  </div>
</template>
<script setup>
import { ref, reactive, onMounted } from "vue";

const props = defineProps({
  title: {
    type: String,
    default: "测试标题",
  },
  // 描述
  desc: {
    type: String,
    default: "测试描述",
  },
});
</script>