<template>
  <base-page title="用户管理" desc="用户管理">
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
              <template #role="{ scope }">
                <el-tag>
                  {{ getRoleName(scope.row.UserRole.role_code) }}
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
        >
          <el-form
            ref="dataFormRef"
            :model="formData"
            :rules="computedRules"
            label-width="auto"
          >
            <el-card shadow="never">
              <el-form-item label="账户名" prop="username">
                <el-input
                  v-model="formData.username"
                  placeholder="请输入账户名称"
                />
              </el-form-item>

              <el-form-item label="昵称" prop="nick_name">
                <el-input
                  v-model="formData.nick_name"
                  placeholder="请输入昵称"
                />
              </el-form-item>

              <el-form-item label="角色" prop="role_code">
                <el-select
                  v-model="formData.role_code"
                  placeholder="请选择角色"
                >
                  <el-option
                    v-for="(item, index) in roleList"
                    :key="index"
                    :value="item.code"
                    :label="item.name"
                  />
                </el-select>
              </el-form-item>

              <el-form-item label="备注">
                <el-input
                  v-model="formData.remark"
                  type="textarea"
                  placeholder="请输入备注"
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
  apiGetUserList,
  apiAddUser,
  apiUpdateUser,
  apiDeleteUser,
} from "~/api/user";

import { apiGetRoleList } from "~/api/role";

import { ElMessageBox, ElMessage } from "element-plus";
import { toast } from '~/composables/util'

import { ref, reactive, computed, onMounted } from "vue";

const queryFormRef = ref();
const dataFormRef = ref();

const loading = ref(false);
const operationType = ref(0); // 0:新增 1:编辑

/**
 * 表格信息
 */
const columns = [
  { prop: "username", label: "用户名", attrs: { width: 150 } },
  { prop: "nick_name", label: "昵称", attrs: { width: 230 } },
  { prop: "role", label: "角色", slot: true, attrs: { width: 230 } },
  { prop: "remark", label: "备注", attrs: { minWidth: 200 } },
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

const roleList = ref([]);

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
  getRoleList();
});

/**
 * 表单数据
 */
const formData = reactive({
  id: 0,
  username: "",
  nick_name: "",
  role_code: "",
  remark: "",
});

/**
 * 表单验证规则
 */
const computedRules = computed(() => {
  const rules = {
    username: [{ required: true, message: "请输入账户名称", trigger: "blur" }],
    nick_name: [{ required: true, message: "请输入昵称", trigger: "blur" }],
  };
  return rules;
});

const getRoleName = (code) => {
  const role = roleList.value.find((item) => item.code === code);
  return role ? role.name : "";
};

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
    const data = await apiGetUserList({
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
 *
 */
const getRoleList = async () => {
  const data = await apiGetRoleList();
  roleList.value = data.list;
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
  dialog.title = "新增用户";
  operationType.value = 0;
  resetValue();
};

/**
 * 编辑
 * @param value 信息
 */
const handleEditClick = (value) => {
  dialog.visible = true;
  dialog.title = "修改用户";
  operationType.value = 1;
  setValue(value);
};

/**
 * 设置值
 * @param value 信息
 */
const setValue = (value) => {
  formData.id = value.id;
  formData.username = value.username;
  formData.nick_name = value.nick_name;
  formData.remark = value.remark;
  formData.role_code = value.UserRole.role_code;
};

/**
 * 重置值
 */
const resetValue = () => {
  formData.id = undefined;
  formData.username = "";
  formData.nick_name = "";
  formData.remark = "";
  formData.role_code = "";
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
    await apiAddUser(formData);
    loading.value = false;
    dialog.visible = false;
    toast('添加用户成功')
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
    await apiUpdateUser(formData);
    loading.value = false;
    dialog.visible = false;
    toast('修改用户成功')
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
      const data = await apiDeleteUser(value.id);
      // console.log("data", data);
      toast('删除用户成功')
      handleQuery();
    },
    () => {
      ElMessage.info("已取消删除");
    }
  );
};
</script>
