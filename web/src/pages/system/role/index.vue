<template>
  <base-page title="角色管理" desc="角色管理">
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
                  @click="handleAuthClick(scope.row)"
                  >权限</el-button
                >
                <el-divider direction="vertical" />
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
                <el-input v-model="formData.code" placeholder="请输入编码" />
              </el-form-item>

              <el-form-item label="名称" prop="name">
                <el-input v-model="formData.name" placeholder="请输入名称" />
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

        <el-dialog v-model="dialog.treeVisible" title="编辑权限" width="500px">
          <el-tree
            show-checkbox
            :data="menuTree"
            :props="defaultProps"
            node-key="code"
            :default-checked-keys="defCheckedKeys"
            @check="handleCheck"
          />
          <template #footer>
            <div class="dialog-footer">
              <el-button type="primary" @click="handleSaveMenuClick"
                >确 定</el-button
              >
              <el-button @click="handleCloseMenuDialog">取 消</el-button>
            </div>
          </template>
        </el-dialog>
      </div>
    </template>
  </base-page>
</template>

<script setup>
import {
  apiGetRoleList,
  apiAddRole,
  apiUpdateRole,
  apiDeleteRole,
} from "~/api/role";
import { apiGetRoleMenuList, apiSaveRoleMenus } from "~/api/menu";

import { flatten } from "~/utils/index";

import { ElMessageBox, ElMessage } from "element-plus";
import { toast } from '~/composables/util'

import { ref, reactive, computed, onMounted } from "vue";
const queryFormRef = ref();
const dataFormRef = ref();

const loading = ref(false);
const operationType = ref(0); // 0:新增 1:编辑

// 表格信息
const columns = [
  { prop: "code", label: "编码", attrs: { width: 150 } },
  { prop: "name", label: "名称", attrs: { width: 230 } },
  { prop: "remark", label: "备注", attrs: { minWidth: 200 } },
  {
    prop: "operation",
    label: "操作",
    slot: true,
    attrs: { width: 170, fixed: "right" },
  },
];

const pageConfig = reactive({
  pageSize: 10,
  currentPage: 1,
  total: 0,
});

const tableData = ref([]);

const checkData = ref([]);

const queryParams = reactive({
  keywords: "",
});

const dialog = reactive({
  title: "",
  visible: false,
});

const formData = reactive({
  code: "",
  name: "",
  remark: "",
});

const computedRules = computed(() => {
  const rules = {
    code: [{ required: true, message: "请输入编码", trigger: "blur" }],
    name: [{ required: true, message: "请输入名称", trigger: "blur" }],
  };
  return rules;
});

/*** 菜单权限  */

/**
 * 获取菜单树
 */
const menuTree = ref([]);

const defaultProps = {
  label: "name",
  children: "child",
};

const defCheckedKeys = ref([]);

const handleAuthClick = async (value) => {
  defCheckedKeys.value = [];
  menuTree.value = [];

  const data = await apiGetRoleMenuList(value.code);

  data.forEach((e) => {
    let val = getChild(e);
    menuTree.value.push(val);
  });

  let fData = flatten(data);
  defCheckedKeys.value = fData.filter((e) => e.is_have).map((e) => e.code);

  setValue(value);

  dialog.treeVisible = true;
};

// 进行递归处理并且返回对象
const getChild = (data) => {
  let rtnData = {
    code: data.code,
    name: data.name,
    child: null,
  };

  if (data.child) {
    rtnData.child = [];
    data.child.forEach((e) => {
      let val = getChild(e);
      rtnData.child.push(val);
    });
  }

  return rtnData;
};

const handleCheck = (data, checked) => {
  checkData.value = [];
  checkData.value = [...checked.checkedKeys, ...checked.halfCheckedKeys];
};

const handleSaveMenuClick = async () => {
  loading.value = true;
  await apiSaveRoleMenus(formData.code, checkData.value);
  loading.value = false;
  dialog.treeVisible = false;

  ElMessage({
    type: "success",
    message: "保存成功",
  });

  getData();
};

const handleCloseMenuDialog = () => {
  dialog.treeVisible = false;
};

// 查询
function handleQuery() {
  getData();
}

const getData = async () => {
  loading.value = true;
  try {
    const data = await apiGetRoleList({
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

// 重置查询
function handleResetQuery() {
  queryFormRef.value.resetFields();
  queryParams.pageNum = 1;
  handleQuery();
}

const setValue = (value) => {
  formData.id = value.id;
  formData.code = value.code;
  formData.name = value.name;
  formData.remark = value.remark;
};

const resetValue = () => {
  formData.id = undefined;
  formData.code = "";
  formData.name = "";
  formData.remark = "";
};

// 新增字典
function handleAddClick() {
  dialog.visible = true;
  dialog.title = "新增角色";
  operationType.value = 0;
  resetValue();
}

/**
 * 编辑字典
 *
 * @param id 字典ID
 */
function handleEditClick(value) {
  dialog.visible = true;
  dialog.title = "修改角色";
  operationType.value = 1;
  setValue(value);
}

// 提交字典表单
function handleSubmitClick() {
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
}

const addData = async () => {
  try {
    loading.value = true;
    await apiAddRole(formData);
    loading.value = false;
    dialog.visible = false;
    toast('新增角色成功')
    handleQuery();
  } catch (error) {
    loading.value = false;
  }
};

const editData = async () => {
  try {
    loading.value = true;
    await apiUpdateRole(formData);
    loading.value = false;
    dialog.visible = false;
    toast('修改角色成功')
    handleQuery();
  } catch (error) {
    loading.value = false;
  }
};

// 关闭字典弹窗
function handleCloseDialog() {
  dialog.visible = false;

  dataFormRef.value.resetFields();
  dataFormRef.value.clearValidate();

  formData.id = undefined;
}
/**
 * 删除角色
 *
 * @param id 角色ID
 */
function handleDeleteClick(value) {
  ElMessageBox.confirm("确认删除已选中的数据项?", "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(
    async () => {
      await apiDeleteRole(value.id);
      toast('删除角色成功')
      handleQuery();
    },
    () => {
      ElMessage.info("已取消删除");
    }
  );
}

onMounted(() => {
  handleQuery();
});
</script>
