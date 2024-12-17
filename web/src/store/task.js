import { defineStore } from 'pinia'
import { ref, toRaw } from 'vue'

export const useTaskStore = defineStore(
    'task',
    () => {
        const task = ref({})
        const editor = ref(null)

        const reset = () => {
            task.value = {}
        }

        const setTask = (value) => {
            task.value = value

            if (editor.value) {
                toRaw(editor.value).setValue(task.value.script_content)
            }
        }

        return {
            task,
            editor,

            setTask,
            reset
        }
    },
    {
        persist: true
    }
)