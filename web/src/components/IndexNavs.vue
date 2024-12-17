<template>
  <el-row :gutter="20" class="mt-5">
    <el-col
      :span="3"
      :offset="0"
      v-for="(item, index) in iconNavs"
      :key="index"
    >
      <el-card shadow="hover" @click="goTo(item.path)">
        <div class="flex flex-col items-center justify-center cursor-pointer">
          <el-icon :size="16" :class="item.color">
            <component :is="item.icon" />
          </el-icon>
          <span class="text-sm mt-2">{{ item.title }}</span>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup>
import { useRouter } from "vue-router";

const router = useRouter();

const iconNavs = [
  {
    icon: "user",
    color: "text-light-blue-500",
    title: "用户管理",
    path: "/system/user",
  },
  {
    icon: "message-box",
    color: "text-violet-500",
    title: "客户端管理",
    path: "/task/client_mana",
  },
  {
    icon: "grid",
    color: "text-violet-500",
    title: "数据库管理",
    path: "/task/db_mana",
  },
  {
    icon: "histogram",
    color: "text-fuchsia-500",
    title: "任务管理",
    path: "/task/task_mana",
  },
];

const goTo = (path) => {
  router.push(path);
};
</script>