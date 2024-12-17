<template>
  <base-page title="系统设置" desc="系统设置">
    <template #content>
      <el-tabs v-model="activeName" @tab-click="handleClick">
        <el-tab-pane label="系统设置" name="first">
          <System />
        </el-tab-pane>
        <el-tab-pane label="日志设置" name="second">
          <Logger />
        </el-tab-pane>
      </el-tabs>
    </template>
  </base-page>
</template>

<script setup>
import { apiGetUserList } from "~/api/user";

import { ref, reactive, onMounted } from "vue";
import System from "./components/system.vue";
import Logger from "./components/logger.vue";

const activeName = ref("first");

const getData = async () => {
  try {
    const data = await apiGetUserList({
      pageNum: pageConfig.currentPage,
      pageSize: pageConfig.pageSize,
      condition: queryParams,
    });
    tableData.value = data.list;
    pageConfig.total = data.total;
  } catch (error) {}
};

const handleClick = (tab, event) => {};

onMounted(() => {
  getData();
});
</script>
