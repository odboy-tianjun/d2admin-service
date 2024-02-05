import router, { buildRouterList, createRoutesInLayout, resetRouter, routesOutLayout } from '@/router'
import api from '@/api'
import { flushMenus, supplementPath } from '@/menu'

export default {
  namespaced: true,
  state: {
    // 是否已经加载
    isLoaded: false,
    // 用户权限
    permissions: []
  },
  actions: {
    /**
     * @description 用户登录后从持久化数据加载一系列的设置
     * @param {Object} context
     */
    async load ({ state, rootState, commit, dispatch }, { focus = false, to = '', data }) {
      // 取消请求 - 没有登录
      if (!data && !rootState.d2admin.account.isLogged) return
      // 取消请求 - 已经加载过动态路由
      if (!focus && state.isLoaded) return
      // 获取接口原始数据
      const menuTree = await api.queryAllMenus({})
      console.log('=================== loadXXXXXX')
      // [ 菜单 ] 计算菜单
      const menuData = flushMenus(menuTree)
      const menus = supplementPath(menuData)
      // [ 菜单 ] 设置侧栏菜单
      commit('d2admin/menu/asideSet', menus, { root: true })
      // [ 路由 ] 计算路由
      const routes = createRoutesInLayout(buildRouterList(menuTree)).concat(routesOutLayout)
      // [ 路由 ] 重新设置路由
      resetRouter(routes)
      // [ 路由 ] 重新设置多标签页池
      commit('d2admin/page/init', routes, { root: true })
      // [ 标签页 ] 重新计算多标签页数据
      dispatch('d2admin/page/openedLoad', { filter: true }, { root: true })
      // [ 搜索 ] 初始化搜索数据
      commit('d2admin/search/init', menus, { root: true })
      // [ 路由 ] 重新访问
      if (to) router.replace(to)
      // 标记已经加载过动态路由
      commit('isLoadedSet', true)
    }
  },
  mutations: {
    /**
     * @description 设置动态路由加载状态
     * @param {Object} state state
     * @param {Boolean} value 是否已经加载动态路由
     */
    isLoadedSet (state, value) {
      state.isLoaded = value
    }
  }
}
