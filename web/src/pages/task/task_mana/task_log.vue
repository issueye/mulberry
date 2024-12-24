<template>
  <base-page title="执行记录查询" desc="执行记录查询">
    <template #content>
      <div class="h-full flex flex-col p-2">
        <div class="search-bar">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true">
            <el-form-item label="关键字" prop="keywords">
              <el-input
                v-model="queryParams.keywords"
                placeholder="名称/编码"
                clearable
                class="w-[200px]"
              />
            </el-form-item>
            <el-form-item label="客户端" prop="client_id">
              <el-select
                v-model="queryParams.client_id"
                placeholder="客户端"
                clearable
                class="w-[200px]"
              >
                <el-option
                  v-for="(item, index) in clientList"
                  :key="index"
                  :label="item.name"
                  :value="item.id"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="任务" prop="task_id">
              <el-select
                v-model="queryParams.task_id"
                placeholder="任务"
                clearable
                class="w-[200px]"
              >
                <el-option
                  v-for="(item, index) in taskList"
                  :key="index"
                  :label="item.name"
                  :value="item.id"
                />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" icon="search" @click="handleQuery"
                >搜索</el-button
              >
              <el-button icon="refresh" @click="handleResetQuery"
                >重置</el-button
              >
            </el-form-item>
          </el-form>
        </div>
        <div class="grow flex flex-col overflow-y-auto">
          <d-table
            :columns="columns"
            :table-data="tableData"
            :page-config="pageConfig"
            usePagination
            highlight-current-row
            stripe
            :loading="loading"
            empty-text="暂无数据"
          >
            <template #status="{ scope }">
              <el-tag :type="scope.row.status === 0 ? 'primary' : 'danger'">{{
                scope.row.status === 0 ? "成功" : "失败"
              }}</el-tag>
            </template>
          </d-table>
        </div>
      </div>
    </template>
  </base-page>
</template>

<script setup>
import { apiGetTaskLogList } from "~/api/task";
import { apiGetTaskList } from "~/api/task";
import { apiGetClientList } from "~/api/client";
import { useRoute, useRouter } from "vue-router";

import { onMounted, reactive, ref } from "vue";
/**
 * 查询条件
 */
const queryParams = reactive({
  keywords: "",
  client_id: null,
  task_id: null,
});

const route = useRoute();
const router = useRouter();

const { query } = route;
if (query["task_id"]) {
  queryParams.task_id = parseInt(query.task_id);
}

const loading = ref(false);

const columns = [
  {
    prop: "task_name",
    label: "任务",
    attrs: { width: 150, fixed: "left", showOverflowTooltip: true },
  },
  {
    prop: "client_name",
    label: "客户端",
    attrs: { width: 150, fixed: "left", showOverflowTooltip: true },
  },
  {
    prop: "execution_id",
    label: "执行ID",
    attrs: { width: 300, showOverflowTooltip: true },
  },
  {
    prop: "start_time",
    label: "开始时间",
    attrs: { width: 180, showOverflowTooltip: true },
  },
  {
    prop: "next_execution_time",
    label: "下次执行时间",
    attrs: { width: 180, showOverflowTooltip: true },
  },
  {
    prop: "result",
    label: "执行结果",
    attrs: { minWidth: 200, showOverflowTooltip: true },
  },
  {
    prop: "status",
    slot: true,
    label: "状态",
    attrs: { width: 80, fixed: "right", showOverflowTooltip: true },
  },
  {
    prop: "end_time",
    label: "结束时间",
    attrs: { width: 180, fixed: "right", showOverflowTooltip: true },
  },
];

const tableData = ref([]);
const taskList = ref([]);
const clientList = ref([]);

/**
 * 分页查询
 */
const pageConfig = reactive({
  pageSize: 10,
  currentPage: 1,
  total: 0,
  handleCurrentChange: (val) => {
    pageConfig.currentPage = val;
    getData();
  },
});

const handleQuery = () => {
  getData();
};

const handleResetQuery = () => {
  queryParams.keywords = "";
  queryParams.client_id = null;
  queryParams.task_id = null;
};

const getData = async () => {
  let task_id = queryParams.task_id ? queryParams.task_id : 0;
  let client_id = queryParams.client_id ? queryParams.client_id : 0;
  let params = {
    pageNum: pageConfig.currentPage,
    pageSize: pageConfig.pageSize,
    condition: {
      keywords: queryParams.keywords,
      client_id: parseInt(client_id),
      task_id: parseInt(task_id),
    },
  };
  let res = await apiGetTaskLogList(params);

  tableData.value = res.list;
  pageConfig.total = res.total;
};

const getTaskList = async () => {
  let res = await apiGetTaskList({});
  taskList.value = res.list;
};

const getClientList = async () => {
  let res = await apiGetClientList({});
  clientList.value = res.list;
};

onMounted(() => {
  getData();
  getTaskList();
  getClientList();
});
</script>
<style scoped lang='scss'></style>