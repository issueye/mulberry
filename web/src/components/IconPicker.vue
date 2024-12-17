<template>
  <div class="el-icon-picker_con">
    <el-popover
      v-model:visible="showPopv"
      :width="440"
      placement="bottom-start"
    >
      <template #reference>
        <el-button
          :icon="modelValue"
          class="icon_picker_btn"
          @click="showPopover"
          >{{ modelValue || '请选择图标' }}</el-button
        >
      </template>
      <div class="el-icon-picker">
        <el-input
          v-model="filterVal"
          class="filter_input"
          placeholder="输入图标查询"
          :suffix-icon="Search"
          @input="inputFilter"
        />
        <div class="icon_con">
          <component
            v-for="icon in iconList"
            :key="icon"
            class="icon"
            :class="{ icon, 'icon-active': icon == modelValue }"
            :is="icon"
            @click="chooseIcon(icon)"
          >
          </component>
        </div>
      </div>
    </el-popover>
    <el-icon
      v-if="modelValue"
      class="delete_icon"
      size="16"
      title="清除"
      @click="clearPicker"
      ><Close
    /></el-icon>
  </div>
</template>

<script setup>
import { ref, watch, getCurrentInstance } from 'vue'
import { Search } from '@element-plus/icons-vue'
const globalProperties = getCurrentInstance().appContext.config.globalProperties // 获取全局挂载

const props = defineProps({
  modelValue: {
    type: String
  }
})

let filterVal = ref('')
const emit = defineEmits(['update:modelValue'])
const showPopover = () => {
    filterVal.value = ''
    showPopv.value = true
}

let showPopv = ref(false)
const chooseIcon = (icon) => {
  emit('update:modelValue', icon)
  showPopv.value = false
}

// 图标过滤查询
let iconList = ref([])
const inputFilter = (val) => {
    if(val){
        iconList.value = globalProperties.$icons.filter((item) => item.toLowerCase().indexOf(val.toLowerCase()) !== -1)
    }else{
        iconList.value = [].concat(globalProperties.$icons)
    }
}

// 清除选中
const clearPicker = () => {
  emit('update:modelValue', '')
}

watch(
  () => showPopv.value,
  (newVal, oldVal) => {
    if(newVal){
        filterVal.value = ''
        inputFilter()
    }
  }
)

</script>

<style scoped>
.el-icon-picker {
  height: 256px;
  overflow-y: scroll;
}

.icon {
    display: inline-block;
    width: 20px;
    height: 20px;
    color: #606266;
    font-size: 14px;
    border-radius: 4px;
    cursor: pointer;
    text-align: center;
    line-height: 45px;
    margin: 10px;
  }

  .icon:hover {
    color: var(--el-color-primary);
  }

  .icon-active {
    color: var(--el-color-primary);
  }
  .icon_picker_btn {
    float: left;
  }
  .filter_input {
    width: 100%;
    padding: 0 12px;
    box-sizing: border-box;
    margin-bottom: 3px;
  }
  .icon_con {
    heigth: calc(100% - 39px);
  }


.el-icon-picker_con {
  
}

.delete_icon {
    float: right;
    margin-left: 8px;
    height: 32px;
    cursor: pointer;
    visibility: hidden;
  }
  .delete_icon:hover {
    color: #606266;
  }

.el-icon-picker_con .delete_icon:hover {
  visibility: visible;
}
</style>