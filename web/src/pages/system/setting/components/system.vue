<template>
  <el-form :model="formData" label-width="auto" class="w-[40%]">
    <el-form-item label="运行模式">
      <el-radio-group v-model="formData.mode" :disabled="disabled">
        <el-radio value="debug">调试模式</el-radio>
        <el-radio value="release">生产模式</el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item label="服务端口号">
      <el-input
        v-model="formData.port"
        :disabled="disabled"
        placeholder="请输入服务端口号"
      />
    </el-form-item>
    <el-form-item label="客户端路径">
      <el-input
        v-model="formData.client_path"
        :disabled="disabled"
        placeholder="客户端路径"
      />
    </el-form-item>
  </el-form>
  <div>
    <el-button @click="disabled = true" v-if="!disabled">取消</el-button>
    <el-button type="primary" @click="disabled = false" v-if="disabled"
      >修改</el-button
    >
    <el-button type="primary" @click="handleSubmitClick" v-if="!disabled"
      >保存</el-button
    >
  </div>
</template>
<script setup>
import { ref, reactive, onMounted } from "vue";

import { apiGetSystemSettings, apiSetSystemSettings } from "~/api/settings";
import { toast } from '~/composables/util'

const disabled = ref(true);

const formData = reactive({
  mode: "debug",
  port: "",
  client_path: "",
});

onMounted(() => {
  getData();
});

const getData = async () => {
  const data = await apiGetSystemSettings();

  formData.mode = data.find((e) => e.key == "mode").value;
  formData.port = data.find((e) => e.key == "port").value;
  formData.client_path = data.find((e) => e.key == "client-path").value;
};

const handleSubmitClick = async () => {
  disabled.value = true;

  let datas = [
    { group: "system", key: "mode", value: formData.mode },
    { group: "system", key: "port", value: formData.port },
    { group: "system", key: "client-path", value: formData.client_path },
  ];

  await apiSetSystemSettings(datas);
  toast('保存成功')
};
</script>