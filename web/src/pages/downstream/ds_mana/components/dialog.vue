<template>
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
        <el-form-item label="端口号" prop="port">
          <el-input
            v-model="formData.port"
            :disabled="true"
            placeholder="请输入端口号"
          />
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
        <el-button type="primary" @click="handleSubmitClick">确 定</el-button>
        <el-button @click="handleCloseDialog">取 消</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { apiAddPort, apiUpdatePort } from "~/api/port";

import { reactive, ref, computed } from "vue";

const props = defineProps({
  operationType: Boolean,
});

const emits = defineEmits(["close"]);

const dataFormRef = ref(null);

/**
 * 弹窗数据
 */
const dialog = reactive({
  visible: false,
  title: "",
});

/**
 * 表单
 */
const formData = reactive({
  id: 0, // ID
  port: 0, // 端口号
  use_gzip: false, // 使用 GZIP
  remark: "", // 备注
});

/**
 * 表单验证规则
 */
const computedRules = computed(() => {
  const rules = {
    port: [{ required: true, message: "请输入端口号", trigger: "blur" }],
  };
  return rules;
});

/**
 * 设置值
 * @param value 信息
 */
const setValue = (value) => {
  formData.id = value.id;
  formData.port = value.port;
  formData.use_gzip = value.use_gzip;
  formData.remark = value.remark;
};

/**
 * 重置值
 */
const resetValue = () => {
  formData.id = undefined;
  formData.port = 0;
  formData.use_gzip = false;
  formData.remark = "";
};

/**
 * 关闭弹窗
 */
const handleCloseDialog = () => {};

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
    await apiAddPort(formData);
    dialog.visible = false;
    toast("新增端口号信息成功");
  } catch (error) {}
};

/**
 * 修改数据
 */
const editData = async () => {
  try {
    loading.value = true;
    // 接口调用
    await apiUpdatePort(formData);
    dialog.visible = false;
    toast("修改端口号信息成功");
  } catch (error) {}
};
</script>