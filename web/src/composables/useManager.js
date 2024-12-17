import { ref, reactive } from 'vue'
import { logout, updatepassword } from "~/api/manager"
import { showModal, toast } from "~/composables/util"
import { useRouter } from "vue-router"
import { useUserStore } from '~/store' // 导入 Pinia store

export function useRepassword() {
    const router = useRouter()
    const userStore = useUserStore() // 使用 Pinia store

    // 修改密码
    const formDrawerRef = ref(null)
    const form = reactive({
        oldpassword: "",
        password: "",
        repassword: ""
    })

    const rules = {
        oldpassword: [
            {
                required: true,
                message: '旧密码不能为空',
                trigger: 'blur'
            },
        ],
        password: [
            {
                required: true,
                message: '新密码不能为空',
                trigger: 'blur'
            },
        ],
        repassword: [
            {
                required: true,
                message: '确认密码不能为空',
                trigger: 'blur'
            },
        ]
    }

    const formRef = ref(null)
    const onSubmit = () => {
        formRef.value.validate((valid) => {
            if (!valid) {
                return false
            }
            formDrawerRef.value.showLoading()
            updatepassword(form)
                .then(res => {
                    toast("修改密码成功，请重新登录")
                    userStore.logoutUser() // 使用 Pinia store 的 action
                    // 跳转回登录页
                    router.push("/login")
                })
                .finally(() => {
                    formDrawerRef.value.hideLoading()
                })

        })
    }

    const openRePasswordForm = () => formDrawerRef.value.open()

    return {
        formDrawerRef,
        form,
        rules,
        formRef,
        onSubmit,
        openRePasswordForm
    }
}

export function useLogout() {
    const router = useRouter()
    const userStore = useUserStore() // 使用 Pinia store

    function handleLogout() {
        showModal("是否要退出登录？").then(res => {
            logout().finally(() => {
                userStore.logoutUser() // 使用 Pinia store 的 action
                // 跳转回登录页
                router.push("/login")
                // 提示退出登录成功
                toast("退出登录成功")
            })
        })
    }

    return {
        handleLogout
    }
}