<template>
  <div
    class="h-full f-menu shadow-md fixed bg-light-50"
    :style="{ width: userStore.asideWidth }"
  >
    <el-menu
      :default-active="defaultActive"
      unique-opened
      :collapse="isCollapse"
      class="border-0"
      @select="handleSelect"
      :collapse-transition="false"
    >
      <!-- 菜单 -->
      <!-- 遍历数组，:key 属性用于唯一标识数组中的每个元素，帮助 Vue 更高效地追踪和更新 DOM。 -->
      <template v-for="(item, index) in asideMenus" :key="index">
        <!-- 子菜单 -->
        <el-sub-menu
          v-if="item.child && item.child.length > 0"
          :index="item.name"
        >
          <!-- 子菜单标题，Vue 会将 <component :is="item.icon"></component> 渲染成指定的 icon 组件，比如 el-icon-home 或 el-icon-user 等，item.icon的属性是一个icon的值 -->
          <template #title>
            <el-icon>
              <component :is="item.icon"></component>
            </el-icon>
            <span>{{ item.name }}</span>
          </template>

          <!-- v-for 指令遍历并渲染 item.child 数组，生成多个 el-menu-item 子菜单项。 -->
          <!-- key 是 Vue 的内部属性，用于唯一标识 v-for 循环中的每一个元素。
                    index 是 el-menu-item 组件的属性，并非 Vue 的内部属性。el-menu-item 的 index 属性用于为菜单项设置唯一的标识符，用来在菜单中追踪选中的项目。 -->
          <el-menu-item
            v-for="(item2, index2) in item.child"
            :key="index2"
            :index="item2.frontpath"
          >
            <el-icon>
              <component :is="item2.icon"></component>
            </el-icon>
            <span>{{ item2.name }}</span>
          </el-menu-item>
        </el-sub-menu>

        <el-menu-item v-else :index="item.frontpath">
          <el-icon>
            <component :is="item.icon"></component>
          </el-icon>
          <span>{{ item.name }}</span>
        </el-menu-item>
      </template>
    </el-menu>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useUserStore } from "~/store"; // 导入 Pinia store
import { global } from "~/init/global";

const router = useRouter();
const route = useRoute();
const userStore = useUserStore(); // 使用 Pinia store

// 默认选中
const defaultActive = ref(route.path);

// 是否折叠
const isCollapse = computed(
  () => !(userStore.asideWidth === global.CARAMBOLA_MENU_WIDTH)
);

const asideMenus = computed(() => userStore.asideMenus);

const handleSelect = (e) => {
  router.push(e);
};
</script>

<style>
.f-menu::-webkit-scrollbar {
  width: 0px;
}
</style>