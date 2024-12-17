<template>
  <base-page title="定时任务管理" desc="定时任务管理">
    <template #content>
      <div class="h-full flex flex-col p-2">
        <div class="search-bar">
          <el-form ref="queryFormRef" :model="queryParams" :inline="true">
            <el-form-item label="关键字" prop="keywords">
              <el-input
                v-model="queryParams.keywords"
                placeholder="名称/编码"
                clearable
                @keyup.enter="handleQuery"
              />
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

        <div class="grow flex flex-col">
          <div class="mb-[10px]">
            <el-button type="success" icon="plus" @click="handleAddClick"
              >新增</el-button
            >
          </div>

          <div class="grow">
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
              <template #task_status="{ scope }">
                <el-tag
                  size="small"
                  :type="scope.row.status ? 'primary' : 'danger'"
                >
                  {{ scope.row.status ? "启用" : "关闭" }}
                </el-tag>
              </template>
              <template #operation="{ scope }">
                <el-button
                  type="primary"
                  link
                  @click="handleRunClick(scope.row)"
                  >执行</el-button
                >
                <el-divider direction="vertical" />
                <el-button
                  type="primary"
                  link
                  @click="handleEditClick(scope.row)"
                  >编辑</el-button
                >
                <el-divider direction="vertical" />
                <el-dropdown @command="handleCommand">
                  <span class="flex items-center text-[--el-color-primary]">
                    更多<el-icon class="el-icon--right"><arrow-down /></el-icon>
                  </span>
                  <template #dropdown>
                    <el-dropdown-menu>
                      <el-dropdown-item
                        :command="{ data: scope.row, type: 'script' }"
                        class="text-[--el-color-primary]"
                        >脚本</el-dropdown-item
                      >
                      <el-dropdown-item
                        :command="{ data: scope.row, type: 'switch' }"
                        class="text-[--el-color-primary]"
                        >{{
                          scope.row.status ? "关闭" : "开启"
                        }}</el-dropdown-item
                      >
                      <el-dropdown-item
                        :command="{ data: scope.row, type: 'del' }"
                        class="text-[--el-color-danger]"
                        >删除</el-dropdown-item
                      >
                    </el-dropdown-menu>
                  </template>
                </el-dropdown>
              </template>
            </d-table>
          </div>
        </div>

        <!--弹窗-->
        <el-dialog
          v-model="dialog.visible"
          :title="dialog.title"
          width="500px"
          @close="handleCloseDialog"
          @open="handleOpenDialog"
          :close-on-click-modal="false"
        >
          <el-form
            ref="dataFormRef"
            :model="formData"
            :rules="computedRules"
            label-width="auto"
          >
            <el-card shadow="never">
              <el-form-item label="名称" prop="name">
                <el-input v-model="formData.name" placeholder="请输入名称" />
              </el-form-item>
              <el-form-item label="cron表达式" prop="cron_expression">
                <el-input
                  v-model="formData.cron_expression"
                  placeholder="请输入cron表达式"
                />
              </el-form-item>
              <el-form-item label="客户端" prop="client_auth_id">
                <el-select
                  v-model="formData.client_auth_id"
                  placeholder="请选择客户端"
                >
                  <el-option
                    v-for="(item, index) in clientList"
                    :key="index"
                    :value="item.client_auth_id"
                    :label="item.name"
                  />
                </el-select>
              </el-form-item>
              <el-form-item label="脚本类型" prop="script_language">
                <el-input
                  v-model="formData.script_language"
                  placeholder="请输入脚本类型"
                  :disabled="true"
                />
              </el-form-item>
            </el-card>
          </el-form>

          <template #footer>
            <div class="dialog-footer">
              <el-button type="primary" @click="handleSubmitClick"
                >确 定</el-button
              >
              <el-button @click="handleCloseDialog">取 消</el-button>
            </div>
          </template>
        </el-dialog>
      </div>
    </template>
  </base-page>
</template>

<script setup>
import {
  apiGetTaskList,
  apiAddTask,
  apiUpdateTask,
  apiDeleteTask,
  apiUpdateTaskStatus,
  apiRunTask,
} from "~/api/task";
import { apiGetClientList } from "~/api/client";

import { ElMessageBox, ElMessage } from "element-plus";

import { ref, reactive, computed, onMounted } from "vue";
import { useRouter } from "vue-router";

import { useTaskStore } from "~/store/task";

const queryFormRef = ref();
const dataFormRef = ref();

const loading = ref(false);
const operationType = ref(0); // 0:新增 1:编辑

const router = useRouter();
const taskStore = useTaskStore();

/**
 * 表格信息
 */
const columns = [
  {
    prop: "name",
    label: "名称",
    attrs: { width: 200, showOverflowTooltip: true },
  },
  { prop: "cron_expression", label: "cron表达式", attrs: { width: 150 } },
  {
    prop: "client_auth_id",
    label: "客户端编码",
    attrs: { minWidth: 250, showOverflowTooltip: true },
  },
  { prop: "script_language", label: "脚本语言", attrs: { width: 150 } },
  {
    prop: "last_execution_time",
    label: "上一次执行时间",
    attrs: { width: 180 },
  },
  {
    prop: "next_execution_time",
    label: "下一次执行时间",
    attrs: { width: 180 },
  },
  {
    prop: "description",
    label: "描述",
    attrs: { width: 300, showOverflowTooltip: true },
  },
  {
    prop: "task_status",
    label: "描述",
    slot: true,
    attrs: { width: 100, fixed: "right", showOverflowTooltip: true },
  },
  {
    prop: "operation",
    label: "操作",
    slot: true,
    attrs: { width: 180, fixed: "right" },
  },
];

/**
 * 分页查询
 */
const pageConfig = reactive({
  pageSize: 10,
  currentPage: 1,
  total: 0,
});

/**
 * 获取表格数据
 */
const tableData = ref([]);

/**
 * 获取客户端列表
 */
const clientList = ref([]);

/**
 * 查询条件
 */
const queryParams = reactive({
  keywords: "",
});

const dialog = reactive({
  title: "",
  visible: false,
});

/**
 * 初始化数据
 */
onMounted(() => {
  handleQuery();
});

/**
 * 表单数据
 */
const formData = reactive({
  id: 0,
  name: "",
  cron_expression: "",
  client_auth_id: "",
  script_language: "",
  script_content: "",
  description: "",
});

/**
 * 表单验证规则
 */
const computedRules = computed(() => {
  const rules = {
    name: [{ required: true, message: "请输入名称", trigger: "blur" }],
    cron_expression: [
      { required: true, message: "请输入cron表达式", trigger: "blur" },
    ],
    client_auth_id: [
      { required: true, message: "请选择客户端", trigger: "blur" },
    ],
  };
  return rules;
});

/**
 * 查询数据
 */
const handleQuery = () => {
  getData();
};

/**
 * 获取数据
 */
const getData = async () => {
  loading.value = true;
  try {
    const data = await apiGetTaskList({
      pageNum: pageConfig.currentPage,
      pageSize: pageConfig.pageSize,
      condition: queryParams,
    });
    loading.value = false;
    tableData.value = data.list;
    pageConfig.total = data.total;
  } catch (error) {
    loading.value = false;
  }
};

/**
 * 重置查询
 */
const handleResetQuery = () => {
  queryFormRef.value.resetFields();
  queryParams.pageNum = 1;
  handleQuery();
};

const handleCommand = (value) => {
  console.log("value", value.data, value.type);

  switch (value.type) {
    case "del": {
      handleDeleteClick(value.data);
      break;
    }
    case "script": {
      taskStore.setTask(value.data);
      // console.log("taskStore", taskStore.task);
      router.push({ name: "script" });
      break;
    }
    case "switch": {
      handleSwitchClick(value.data);
      break;
    }
  }
};

/**
 * 新增
 */
const handleAddClick = () => {
  dialog.visible = true;
  dialog.title = "添加定时任务";
  operationType.value = 0;
  resetValue();
};

/**
 * 编辑
 * @param value 信息
 */
const handleEditClick = (value) => {
  dialog.visible = true;
  dialog.title = "修改定时任务";
  operationType.value = 1;
  setValue(value);
};

/**
 * 运行任务
 * @param value
 */
const handleRunClick = (value) => {
  ElMessageBox.confirm("确认执行该任务?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(
    async () => {
      // 调用接口
      await apiRunTask(value.id);
      handleQuery();
    },
    () => {
      ElMessage.info("已取消执行");
    }
  );
};

const handleSwitchClick = (value) => {
  ElMessageBox.confirm("确认修改该任务状态?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(
    async () => {
      // 调用接口
      await apiUpdateTaskStatus(value.id);
      handleQuery();
    },
    () => {
      ElMessage.info("已取消修改");
    }
  );
};

/**
 * 设置值
 * @param value 信息
 */
const setValue = (value) => {
  formData.id = value.id;
  formData.name = value.name;
  formData.cron_expression = value.cron_expression;
  formData.client_auth_id = value.client_auth_id;
  formData.script_language = value.script_language;
  formData.script_content = value.script_content;
  formData.description = value.description;
};

/**
 * 重置值
 */
const resetValue = () => {
  formData.id = undefined;
  formData.name = "";
  formData.cron_expression = "";
  formData.client_auth_id = "";
  formData.script_language = "javascript";
  formData.script_content = "";
  formData.description = "";
};

/**
 * 提交表单按钮
 */
const handleSubmitClick = () => {
  dataFormRef.value.validate(async (isValid) => {
    if (isValid) {
      loading.value = true;
      switch (operationType.value) {
        case 0:
          await addData();
          break;
        case 1:
          await editData();
          break;
      }
      loading.value = false;
    }
  });
};

/**
 * 新增数据
 */
const addData = async () => {
  try {
    loading.value = true;
    // 接口调用
    await apiAddTask(formData);
    loading.value = false;
    dialog.visible = false;
    handleQuery();
  } catch (error) {
    loading.value = false;
  }
};

/**
 * 修改数据
 */
const editData = async () => {
  try {
    loading.value = true;
    // 接口调用
    await apiUpdateTask(formData);
    loading.value = false;
    dialog.visible = false;
    handleQuery();
  } catch (error) {
    loading.value = false;
  }
};

// 关闭弹窗
const handleCloseDialog = () => {
  dialog.visible = false;

  dataFormRef.value.resetFields();
  dataFormRef.value.clearValidate();

  formData.id = undefined;
};

const handleOpenDialog = () => {
  getClientList();
};

const getClientList = async () => {
  const data = await apiGetClientList();
  clientList.value = data.list;
};

/**
 * 删除数据
 * @param value 信息
 */
const handleDeleteClick = (value) => {
  ElMessageBox.confirm("确认删除已选中的数据项?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(
    async () => {
      // 调用接口
      await apiDeleteTask(value.id);
      handleQuery();
    },
    () => {
      ElMessage.info("已取消删除");
    }
  );
};
</script>