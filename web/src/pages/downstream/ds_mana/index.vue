<template>
  <base-page title="代理管理" desc="代理管理">
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
              <template #status="{ scope }">
                <el-tag
                  :type="scope.row.status ? 'success' : 'danger'"
                  :disabled="true"
                  size="small"
                >
                  {{ scope.row.status ? "启用" : "停用" }}
                </el-tag>
              </template>
              <template #operation="{ scope }">
                <el-button
                  type="primary"
                  link
                  @click="handleEditClick(scope.row)"
                  >编辑</el-button
                >
                <el-divider direction="vertical" />
                <el-button
                  type="primary"
                  link
                  @click="handleModifyStatusClick(scope.row)"
                  >{{ scope.row.status ? "关闭" : "打开" }}</el-button
                >
                <el-divider direction="vertical" />
                <el-dropdown @command="handleCommand">
                  <span class="flex items-center text-[--el-color-primary]">
                    更多<el-icon class="el-icon--right"><arrow-down /></el-icon>
                  </span>
                  <template #dropdown>
                    <el-dropdown-menu>
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
          :close-on-click-modal="false"
        >
          <el-form
            ref="dataFormRef"
            :model="formData"
            :rules="computedRules"
            label-width="auto"
          >
            <el-card shadow="never">
              <el-form-item label="认证编码" prop="client_auth_id">
                <el-input
                  v-model="formData.client_auth_id"
                  :disabled="true"
                  placeholder="客户端认证ID"
                />
              </el-form-item>
              <el-form-item label="名称" prop="name">
                <el-input v-model="formData.name" placeholder="请输入名称" />
              </el-form-item>
              <el-form-item label="备注">
                <el-input
                  v-model="formData.remark"
                  placeholder="请输入备注"
                  type="textarea"
                  :rows="4"
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
import { apiGetPortList, apiDeletePort } from "~/api/port";

import { toast } from "~/composables/util";

import { ElMessageBox, ElMessage } from "element-plus";

import { ref, reactive, computed, onMounted } from "vue";

const queryFormRef = ref();

const loading = ref(false);
const operationType = ref(0); // 0:新增 1:编辑

/**
 * 表格信息
 */
const columns = [
  {
    prop: "port",
    label: "端口号",
    attrs: { width: 90, showOverflowTooltip: true },
  },
  { prop: "status", label: "状态", attrs: { width: 100 } },
  {
    prop: "remark",
    label: "备注",
    attrs: { minWidth: 200, showOverflowTooltip: true },
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
 * 查询条件
 */
const queryParams = reactive({
  keywords: "",
});

/**
 * 初始化数据
 */
onMounted(() => {
  handleQuery();
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
    const data = await apiGetPortList({
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

/**
 * 新增
 */
const handleAddClick = () => {
  dialog.visible = true;
  dialog.title = "添加客户端";
  operationType.value = 0;
};

/**
 * 编辑
 * @param value 信息
 */
const handleEditClick = (value) => {
  dialog.visible = true;
  dialog.title = "修改客户端";
  operationType.value = 1;
};

// 关闭弹窗
const handleCloseDialog = () => {
  dialog.visible = false;

  dataFormRef.value.resetFields();
  dataFormRef.value.clearValidate();

  formData.id = undefined;
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
      await apiDeleteClient(value.id);
      toast("删除客户端信息成功");
      handleQuery();
    },
    () => {
      ElMessage.info("已取消删除");
    }
  );
};

const handleCommand = (data) => {
  switch (data.type) {
    case "del":
      handleDeleteClick(data.data);
      break;
  }
};
</script>