import { installElementIcons, installComponents } from './install_components'

export function install(Vue) {
    // 安装图标
    installElementIcons(Vue)

    // 安装组件
    installComponents(Vue)
}