import {
    createRouter,
    createWebHashHistory
} from 'vue-router'

import Admin from "~/layouts/admin.vue";
import Index from '~/pages/index.vue'
import Login from '~/pages/login.vue'
import NotFound from '~/pages/404.vue'

import TaskMana from '~/pages/task/task_mana/index.vue';
import ClientMana from '~/pages/task/client_mana/index.vue';
import Script from '~/pages/task/script/index.vue';
import DBMana from '~/pages/task/db_mana/index.vue';
import TaskLog from '~/pages/task/task_mana/task_log.vue';

import User from '~/pages/system/user/index.vue';
import Role from '~/pages/system/role/index.vue';
import Menu from '~/pages/system/menu/index.vue';
import Setting from '~/pages/system/setting/index.vue';


// 默认路由，所有用户共享
const routes = [
    { path: "/", name: "admin", component: Admin, },
    { path: "/login", component: Login, meta: { title: "登录页" } },
    { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound }
]


// 动态路由，用于匹配菜单动态添加路由
const asyncRoutes = [
    { path: "/home", name: "home", component: Index, meta: { title: "后台首页" } },

    // 任务管理
    { path: "/task/task_mana", name: "task_mana", component: TaskMana, meta: { title: "任务管理" } },
    { path: "/task/client_mana", name: "client_mana", component: ClientMana, meta: { title: "客户端管理" } },
    { path: "/task/db_mana", name: "db_mana", component: DBMana, meta: { title: "数据库管理" } },
    { path: "/task/script", name: "script", component: Script, meta: { title: "脚本编辑器" } },
    { path: "/task/task_log", name: "task_log", component: TaskLog, meta: { title: "执行记录查询" } },

    // 系统管理
    { path: "/system/user", name: "user", component: User, meta: { title: "用户管理" } },
    { path: "/system/role", name: "role", component: Role, meta: { title: "角色管理" } },
    { path: "/system/menu", name: "menu", component: Menu, meta: { title: "菜单管理" } },
    { path: "/system/setting", name: "setting", component: Setting, meta: { title: "系统设置" } },
]


// 定义静态路由
export const router = createRouter({
    history: createWebHashHistory(),
    routes
})



// 定义动态路由
// 动态添加路由的方法
// 这个代码的作用是根据后端返回的用户菜单权限动态添加路由。具体来说，它遍历后端返回的菜单权限数组，找到与前端定义的动态路由匹配的项，并将这些路由动态添加到 Vue Router 中。
// 动态添加路由的函数
//这个会把所有的后端传过来的路由变成admin的子路由
// 这个只是返回是否有新的路由被添加了，不会改变路由结构，也不会改变菜单结构
export function addRoutes(menus) {
    // 是否有新的路由被添加的标志，这个标志用于判断是否需要重新加载页面
    // 如果有新的路由被添加，则返回 true，否则返回 false。
    // 如果有新的路由被添加，则需要重新加载页面，否则不需要重新加载页面。
    let hasNewRoutes = false;

    // 定义一个递归函数来查找并添加路由
    const findAndAddRoutesByMenus = (arr) => {
        // 遍历传入的菜单数组，arr.forEach 是 JavaScript 数组的一个内置方法，用于遍历数组中的每个元素，并对每个元素执行指定的回调函数。
        arr.forEach((menuItem) => {
            // 在 asyncRoutes 中查找与当前菜单项 frontpath 匹配的路由，route.path是动态路由定义好的属性和menuItem.frontpath是后端返回的菜单项的frontpath属性的。
            // 如果找到匹配的路由，则将其添加到名为 "admin" 的路由中。
            let matchedRoute = asyncRoutes.find((route) => route.path === menuItem.frontpath);

            // 如果找到匹配的路由且该路由尚未添加到 router 中
            if (matchedRoute && !router.hasRoute(matchedRoute.path)) {
                // 将该路由添加到名为 "admin" 的路由中
                router.addRoute("admin", matchedRoute);
                // 标记有新的路由被添加
                hasNewRoutes = true;
            }

            // 如果当前菜单项有子菜单，递归处理子菜单，这个里面的当前菜单项是后端返回的菜单项，不是前端定义的动态路由。
            if (menuItem.child && menuItem.child.length > 0) {
                findAndAddRoutesByMenus(menuItem.child);
            }
        });
    };

    // 从 menus 开始递归查找并添加路由
    findAndAddRoutesByMenus(menus);

    // 返回是否有新的路由被添加
    return hasNewRoutes;
}