import { router, addRoutes } from "~/router"
import { getToken } from "~/composables/auth"
import {
    toast,
    showFullLoading,
    hideFullLoading
} from "~/composables/util"
import { useUserStore } from '~/store'

let hasGetInfo = false

router.beforeEach(async (to, from, next) => {
    showFullLoading()

    const token = getToken()

    if (!token && to.path !== "/login") {
        toast("请先登录", "error")
        return next({ path: "/login" })
    }

    if (token && to.path === "/login") {
        toast("请勿重复登录", "error")
        return next({ path: from.path ? from.path : "/" })
    }

    const userStore = useUserStore()

    let hasNewRoutes = false
    if (token && !hasGetInfo) {
        await userStore.fetchUserInfo()
        hasGetInfo = true
        hasNewRoutes = addRoutes(userStore.menus)
    }

    let title = (to.meta.title ? to.meta.title : "")
    document.title = title

    hasNewRoutes ? next(to.fullPath) : next()
})

router.afterEach((to, from) => hideFullLoading())
