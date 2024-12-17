<template>
  <el-container class="h-full">
    <el-header>
      <f-header />
    </el-header>
    <el-container>
      <el-aside :width="userStore.asideWidth">
        <f-menu></f-menu>
      </el-aside>
      <el-main>
        <div class="h-full flex flex-col bg-gray-100">
          <f-tag-list />
          <div class="h-[calc(100%-56px)] m-3">
            <router-view v-slot="{ Component }">
              <transition name="fade">
                <keep-alive :max="10">
                  <component :is="Component"></component>
                </keep-alive>
              </transition>
            </router-view>
          </div>
        </div>
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import FHeader from "./components/FHeader.vue";
import FMenu from "./components/FMenu.vue";
import FTagList from "./components/FTagList.vue";
import { useUserStore } from "~/store"; // 导入 Pinia store

const userStore = useUserStore(); // 使用 Pinia store
</script>

<style scoped>
.el-aside {
  transition: all 0.2s;
}

.el-main {
  padding: 0;
}

.fade-enter-from {
  opacity: 0;
}
.fade-enter-to {
  opacity: 1;
}
.fade-leave-from {
  opacity: 1;
}
.fade-leave-to {
  opacity: 0;
}
.fade-enter-active,
.fade-leave-active {
  transition: all 0.3s;
}
.fade-enter-active {
  transition-delay: 0.3s;
}
</style>