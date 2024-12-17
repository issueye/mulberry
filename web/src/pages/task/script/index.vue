<template>
  <base-page :title="task.name" desc="脚本编辑器">
    <template #actions>
      <!-- <el-button type="primary" @click="handleBackClick">返回</el-button> -->
    </template>
    <template #content>
      <div class="flex h-[99%]">
        <div ref="editorBox" class="w-full h-full"></div>
      </div>
    </template>
  </base-page>
</template>

<script setup>
import { apiSaveCode } from "~/api/task";

import { ElMessageBox, ElMessage } from "element-plus";

import { ref, toRaw, onMounted, nextTick } from "vue";
import * as monaco from "monaco-editor";

import { useTaskStore } from "~/store/task";
import { storeToRefs } from "pinia";

const editorBox = ref(null);
const taskStore = useTaskStore();

const { task, editor } = storeToRefs(taskStore);

/**
 * 初始化数据
 */
onMounted(() => {
  // 在 DOM 加载完成后执行
  nextTick(() => {
    editor.value = getEditor(task.value.script_content, editorBox.value, false);
    editor.value.addAction({
      id: "save",
      label: "保存",
      keybindings: [monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyS],
      run: () => {
        let code = toRaw(editor.value).getValue();
        // console.log("code", code);

        let data = {
          id: task.value.id,
          code: code,
        };

        apiSaveCode(data).then(() => {
          ElMessage.success("保存成功");
        });
      },
    });
  });
});

const getEditor = (code, container, readOnly) => {
  return monaco.editor.create(container, {
    value: code,
    language: "javascript",
    folding: true, // 是否折叠
    foldingHighlight: true, // 折叠等高线
    foldingStrategy: "indentation", // 折叠方式  auto | indentation
    showFoldingControls: "always", // 是否一直显示折叠 always | mouseover
    disableLayerHinting: true, // 等宽优化
    emptySelectionClipboard: false, // 空选择剪切板
    selectionClipboard: false, // 选择剪切板
    // automaticLayout: true, // 自动布局
    codeLens: false, // 代码镜头
    scrollBeyondLastLine: false, // 滚动完最后一行后再滚动一屏幕
    colorDecorators: true, // 颜色装饰器
    accessibilitySupport: "off", // 辅助功能支持  "auto" | "off" | "on"
    lineNumbers: "on", // 行号 取值： "on" | "off" | "relative" | "interval" | function
    lineNumbersMinChars: 5, // 行号最小字符   number
    readOnly: readOnly, //是否只读  取值 true | false
    theme: "vs-dark",
    minimap: {
      enabled: true, // 是否启用预览图
    },
  });
};

const handleBackClick = () => {};
</script>
