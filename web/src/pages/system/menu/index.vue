<template>
  <base-page title="菜单管理" desc="菜单管理">
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
              row-key="code"
              :tree-props="{ children: 'child', hasChildren: 'hasChildren' }"
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
              <el-form-item label="编码" prop="code">
                <el-input
                  v-model="formData.code"
                  :disabled="operationType === 1"
                  placeholder="请输入编码"
                />
              </el-form-item>
              <el-form-item label="名称" prop="name">
                <el-input v-model="formData.name" placeholder="请输入名称" />
              </el-form-item>
              <el-form-item label="路径" prop="frontpath">
                <el-input
                  v-model="formData.frontpath"
                  placeholder="请输入前端路径"
                />
              </el-form-item>
              <el-form-item label="菜单级别" prop="level">
                <el-select
                  v-model="formData.level"
                  placeholder="请选择菜单级别"
                >
                  <el-option :value="0" label="一级菜单" />
                  <el-option :value="1" label="二级菜单" />
                </el-select>
              </el-form-item>
              <el-form-item
                label="父级菜单"
                prop="parentCode"
                v-if="formData.level === 1"
              >
                <el-select
                  v-model="formData.parent_code"
                  placeholder="请选择父级菜单"
                >
                  <el-option
                    v-for="(item, index) in CatalogData"
                    :value="item.code"
                    :label="item.name"
                    :key="index"
                  />
                </el-select>
              </el-form-item>
              <el-form-item label="图标" prop="icon">
                <el-icon-picker v-model="formData.icon" width="520px" />
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
  apiGetMenuList,
  apiAddMenu,
  apiUpdateMenu,
  apiDeleteMenu,
  apiGetCatalog,
} from "~/api/menu";

import { ElMessageBox, ElMessage } from "element-plus";

import { ref, reactive, computed, onMounted } from "vue";

const queryFormRef = ref();
const dataFormRef = ref();

const loading = ref(false);
const operationType = ref(0); // 0:新增 1:编辑

/**
 * 表格信息
 */
const columns = [
  { prop: "code", label: "编码", attrs: { width: 100 } },
  { prop: "name", label: "名称", attrs: { width: 300 } },
  { prop: "icon", label: "图标", attrs: { width: 200 } },
  { prop: "frontpath", label: "路径", attrs: { minWidth: 200 } },
  {
    prop: "operation",
    label: "操作",
    slot: true,
    attrs: { width: 120, fixed: "right" },
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
 * 获取菜单列表
 */
const CatalogData = ref([]);

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
  code: "",
  name: "",
  icon: "",
  frontpath: "",
  desc: "",
  parent_code: "",
  level: 0,
  order: 0,
});

/**
 * 表单验证规则
 */
const computedRules = computed(() => {
  const rules = {
    code: [{ required: true, message: "请输入菜单编码", trigger: "blur" }],
    name: [{ required: true, message: "请输入菜单名称", trigger: "blur" }],
    frontpath: [{ required: true, message: "请输入菜单路径", trigger: "blur" }],
    order: [{ required: true, message: "请输入排序", trigger: "blur" }],
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
    const data = await apiGetMenuList({
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
  dialog.title = "添加菜单";
  operationType.value = 0;
  resetValue();
};

/**
 * 编辑
 * @param value 信息
 */
const handleEditClick = (value) => {
  dialog.visible = true;
  dialog.title = "修改菜单";
  operationType.value = 1;
  setValue(value);
};

/**
 * 设置值
 * @param value 信息
 */
const setValue = (value) => {
  formData.id = value.id;
  formData.code = value.code;
  formData.name = value.name;
  formData.icon = value.icon;
  formData.frontpath = value.frontpath;
  formData.order = value.order;
  formData.parent_code = value.parent_code;
};

/**
 * 重置值
 */
const resetValue = () => {
  formData.id = undefined;
  formData.code = "";
  formData.name = "";
  formData.icon = "";
  formData.frontpath = "";
  formData.order = 0;
  formData.parent_code = "";
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
    await apiAddMenu(formData);
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
    await apiUpdateMenu(formData);
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

const handleOpenDialog = async () => {
  const data = await apiGetCatalog();
  console.log(data);
  CatalogData.value = data;
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
      await apiDeleteMenu(value.id);
      handleQuery();
    },
    () => {
      ElMessage.info("已取消删除");
    }
  );
};
</script>
