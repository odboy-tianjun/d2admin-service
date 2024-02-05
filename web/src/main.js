// Vue
import Vue from 'vue'
import i18n from './i18n'
import App from './App'
// 核心插件
import d2Admin from '@/plugin/d2admin'
// store
import store from '@/store/index'

// 菜单和路由设置
import router from './router'
import { frameInRoutes } from '@/router/routes'
import util from '@/libs/util'
import { uniqueId } from 'lodash'
import api from '@/api'

// 核心插件
Vue.use(d2Admin)

/**
 * @description 给菜单数据补充上 path 字段
 * @description https://github.com/d2-projects/d2-admin/issues/209
 * @param {Array} menu 原始的菜单数据
 */
function supplementPath (menu) {
  return menu.map(e => ({
    ...e,
    path: e.path || uniqueId('d2-menu-empty-'),
    ...e.children ? {
      children: supplementPath(e.children)
    } : {}
  }))
}

async function flushMenus () {
  let menuData = [{ path: '/index', title: '首页', icon: 'home' }]
  try {
    console.log('========= 设置菜单 ======== ')
    const menuTree = await api.queryAllMenus({})
    // 权限菜单
    const d2adminMenus = []
    if (menuTree && menuTree.length > 0) {
      for (const menuItem of menuTree) {
        let menuTemp
        // 一级
        if (menuItem.menuPath) {
          menuTemp = { id: menuItem.id, title: menuItem.menuTitle, ico: menuItem.menuIcon, path: menuItem.menuPath }
        } else {
          // path不能为undefined会报错, 详情看源码'src/layout/header-aside/components/libs/util.menu.js'
          menuTemp = { id: menuItem.id, title: menuItem.menuTitle, ico: menuItem.menuIcon, path: '' }
        }
        // 二级
        const children = menuItem.children
        if (children && children.length > 0) {
          for (const childrenMenuItem of children) {
            const childrenMenuTemp = {
              id: menuItem.id,
              title: childrenMenuItem.menuTitle,
              ico: childrenMenuItem.menuIcon,
              path: childrenMenuItem.menuPath
            }
            if (menuTemp.children && menuTemp.children.length > 0) {
              menuTemp.children.push(childrenMenuTemp)
            } else {
              menuTemp.children = []
              menuTemp.children.push(childrenMenuTemp)
            }
          }
        }
        d2adminMenus.push(menuTemp)
      }
    }
    // 拼合菜单
    menuData = menuData.concat(d2adminMenus)
    return menuData
  } catch (e) {
    console.error(e)
    return menuData
  }
}

async function flushRouter(){
  const menuTree = await api.queryAllMenus({})
  // 权限路由
}

new Vue({
  router,
  store,
  i18n,
  render: h => h(App),
  created () {
    console.log('=========================== app created ===========================')
    const _this = this
    flushMenus().then((data) => {
      // 设置顶栏菜单
      _this.$store.commit('d2admin/menu/headerSet', supplementPath([]))
      // 设置侧边栏菜单
      _this.$store.commit('d2admin/menu/asideSet', supplementPath(data))
      // 初始化菜单搜索功能
      _this.$store.commit('d2admin/search/init', supplementPath(data))
    })
    flushRouter().then((data) => {
      // 处理路由 得到每一级的路由设置
      _this.$store.commit('d2admin/page/init', frameInRoutes)
    })
  },
  mounted () {
    console.log('=========================== app mounted ===========================')
    // 展示系统信息
    this.$store.commit('d2admin/releases/versionShow')
    // 用户登录后从数据库加载一系列的设置
    this.$store.dispatch('d2admin/account/load')
    // 获取并记录用户 UA
    this.$store.commit('d2admin/ua/get')
    // 初始化全屏监听
    this.$store.dispatch('d2admin/fullscreen/listen')
  }
}).$mount('#app')
