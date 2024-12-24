<template>
  <base-page title="数据库管理" desc="数据库管理">
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
              <template #operation="{ scope }">
                <el-button
                  type="primary"
                  link
                  @click="handleEditClick(scope.row)"
                  >编辑</el-button
                >
                <el-divider direction="vertical" />
                <el-button
                  type="danger"
                  link
                  @click="handleDeleteClick(scope.row)"
                  >删除</el-button
                >
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
              <el-form-item label="类型" prop="db_type">
                <el-select
                  v-model="formData.db_type"
                  placeholder="请选择数据库类型"
                >
                  <el-option
                    v-for="(item, index) in db_type_list"
                    :key="index"
                    :value="item.value"
                    :label="item.label"
                  />
                </el-select>
              </el-form-item>
              <el-form-item
                label="数据库"
                prop="database"
                v-if="formData.db_type != 'sqlite'"
              >
                <el-input
                  v-model="formData.database"
                  placeholder="请输入数据库"
                />
              </el-form-item>
              <el-form-item
                label="主机"
                prop="host"
                v-if="formData.db_type != 'sqlite'"
              >
                <el-input v-model="formData.host" placeholder="请输入主机" />
              </el-form-item>
              <el-form-item
                label="端口号"
                prop="port"
                v-if="formData.db_type != 'sqlite'"
              >
                <el-input
                  v-model.number="formData.port"
                  placeholder="请输入端口号"
                />
              </el-form-item>
              <el-form-item
                label="用户名"
                prop="username"
                v-if="formData.db_type != 'sqlite'"
              >
                <el-input
                  v-model="formData.username"
                  placeholder="请输入用户名"
                />
              </el-form-item>
              <el-form-item
                label="密码"
                prop="password"
                v-if="formData.db_type != 'sqlite'"
              >
                <el-input
                  v-model="formData.password"
                  placeholder="请输入密码"
                />
              </el-form-item>
              <el-form-item
                label="模式"
                prop="schema"
                v-if="['postgresql'].indexOf(formData.db_type) > -1"
              >
                <el-input
                  v-model="formData.database"
                  placeholder="请输入数据库"
                />
              </el-form-item>
              <el-form-item
                label="路径"
                prop="path"
                v-if="formData.db_type === 'sqlite'"
              >
                <el-input
                  v-model="formData.path"
                  placeholder="请输入数据库路径"
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
  apiGetDBList,
  apiAddDB,
  apiUpdateDB,
  apiDeleteDB,
} from "~/api/database";

import { ElMessageBox, ElMessage } from "element-plus";
import { toast } from '~/composables/util'

import { ref, reactive, computed, onMounted } from "vue";
import { useRouter } from "vue-router";

import { useTaskStore } from "~/store/task";

const queryFormRef = ref();
const dataFormRef = ref();

const loading = ref(false);
const operationType = ref(0); // 0:新增 1:编辑

const router = useRouter();
const taskStore = useTaskStore();

const db_type_list = [
  { label: "mysql", value: "mysql" },
  { label: "sqlserver", value: "sqlserver" },
  { label: "sqlite", value: "sqlite" },
  { label: "oracle", value: "oracle" },
  { label: "postgresql", value: "postgresql" },
];

/**
 * 表格信息
 */
const columns = [
  {
    prop: "name",
    label: "名称",
    attrs: { width: 200, showOverflowTooltip: true },
  },
  { prop: "db_type", label: "数据库类型", attrs: { width: 150 } },
  {
    prop: "host",
    label: "主机",
    attrs: { minWidth: 250, showOverflowTooltip: true },
  },
  { prop: "port", label: "端口", attrs: { width: 150 } },
  {
    prop: "username",
    label: "用户名",
    attrs: { width: 180 },
  },
  {
    prop: "password",
    label: "密码",
    attrs: { width: 180 },
  },
  {
    prop: "database",
    label: "数据库",
    attrs: { width: 200, showOverflowTooltip: true },
  },
  {
    prop: "path",
    label: "路径",
    attrs: { width: 300, showOverflowTooltip: true },
  },
  {
    prop: "schema",
    label: "模式",
    slot: true,
    attrs: { width: 100, showOverflowTooltip: true },
  },
  {
    prop: "operation",
    label: "操作",
    slot: true,
    attrs: { width: 110, fixed: "right" },
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
  db_type: "sqlserver",
  host: "",
  port: 1433,
  username: "",
  password: "",
  database: "",
  schema: "",
  path: "",
});

/**
 * 表单验证规则
 */
const computedRules = computed(() => {
  const rules = {
    name: [{ required: true, message: "请输入名称", trigger: "blur" }],
    db_type: [{ required: true, message: "请选择数据库类型", trigger: "blur" }],
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
    const data = await apiGetDBList({
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
  dialog.title = "添加数据库信息";
  operationType.value = 0;
  resetValue();
};

/**
 * 编辑
 * @param value 信息
 */
const handleEditClick = (value) => {
  dialog.visible = true;
  dialog.title = "修改数据库信息";
  operationType.value = 1;
  setValue(value);
};

/**
 * 设置值
 * @param value 信息
 */
const setValue = (value) => {
  formData.id = value.id;
  formData.name = value.name;
  formData.db_type = value.db_type;
  formData.host = value.host;
  formData.port = value.port;
  formData.username = value.username;
  formData.password = value.password;
  formData.database = value.database;
  formData.schema = value.schema;
  formData.path = value.path;
};

/**
 * 重置值
 */
const resetValue = () => {
  formData.id = undefined;
  formData.name = "";
  formData.db_type = "sqlserver";
  formData.host = "";
  formData.port = 1433;
  formData.username = "";
  formData.password = "";
  formData.database = "";
  formData.schema = "";
  formData.path = "";
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
    await apiAddDB(formData);
    loading.value = false;
    dialog.visible = false;
    toast('新增数据库信息成功，重启客户端才能使用最新的数据库信息')
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
    await apiUpdateDB(formData);
    loading.value = false;
    dialog.visible = false;
    toast('修改数据库信息成功，重启客户端才能使用最新的数据库信息')
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

const handleOpenDialog = () => {};

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
      await apiDeleteDB(value.id);
      toast('移除数据库信息成功')
      handleQuery();
    },
    () => {
      ElMessage.info("已取消删除");
    }
  );
};
</script>