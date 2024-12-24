<template>
  <el-form :model="formData" label-width="auto" class="w-[40%]">
    <el-form-item label="日志保存路径">
      <el-input
        v-model="formData.path"
        :disabled="disabled"
        placeholder="请输入日志保存路径"
      />
    </el-form-item>
    <el-form-item label="日志等级">
      <el-select
        v-model="formData.level"
        :disabled="disabled"
        placeholder="请选择日志等级"
      >
        <el-option label="DEBUG" :value="-1" />
        <el-option label="INFO" :value="0" />
        <el-option label="WARN" :value="1" />
        <el-option label="ERROR" :value="2" />
      </el-select>
    </el-form-item>
    <el-form-item label="压缩日志">
      <el-switch v-model="formData.compress" :disabled="disabled" />
    </el-form-item>
    <el-form-item label="日志保存天数">
      <el-input-number
        v-model="formData.maxAge"
        :disabled="disabled"
        controls-position="right"
      />
    </el-form-item>
    <el-form-item label="日志备份数">
      <el-input-number
        v-model="formData.maxBackups"
        :disabled="disabled"
        controls-position="right"
      />
    </el-form-item>
    <el-form-item label="日志分割大小">
      <el-input-number
        v-model="formData.maxSize"
        :disabled="disabled"
        controls-position="right"
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

import { apiGetLoggerSettings, apiSetLoggerSettings } from "~/api/settings";
import { toast } from '~/composables/util'

const disabled = ref(true);

const formData = reactive({
  path: "",
  level: -1,
  compress: true,
  maxAge: 30,
  maxBackups: 30,
  maxSize: 100,
});

onMounted(() => {
  getData();
});

const getData = async () => {
  const data = await apiGetLoggerSettings();

  formData.path = data.find((e) => e.key == "path").value;
  formData.level = parseInt(data.find((e) => e.key == "level").value);
  formData.compress = data.find((e) => e.key == "compress").value == "true";
  formData.maxAge = parseInt(data.find((e) => e.key == "max-age").value);
  formData.maxBackups = parseInt(
    data.find((e) => e.key == "max-backups").value
  );
  formData.maxSize = parseInt(data.find((e) => e.key == "max-size").value);
};

const handleSubmitClick = async () => {
  disabled.value = true;

  let datas = [
    { group: "log", key: "path", value: formData.path },
    { group: "log", key: "level", value: formData.level.toString() },
    { group: "log", key: "compress", value: formData.compress.toString() },
    { group: "log", key: "max-age", value: formData.maxAge.toString() },
    {
      group: "log",
      key: "max-backups",
      value: formData.maxBackups.toString(),
    },
    { group: "log", key: "max-size", value: formData.maxSize.toString() },
  ];

  await apiSetLoggerSettings(datas);
  toast('保存成功')
};
</script>