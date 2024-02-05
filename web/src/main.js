// Vue
import Vue from 'vue'
import i18n from './i18n'
import App from './App'
// 核心插件
import d2Admin from '@/plugin/d2admin'
// store
import store from '@/store/index'

// 菜单和路由设置
import router, { constantRoutes } from './router'
import api from '@/api'
import { flushMenus } from '@/menu'

// 核心插件
Vue.use(d2Admin)

new Vue({
  router,
  store,
  i18n,
  render: h => h(App),
  methods: {
    async loadMenuRouter () {
      // 设置路由菜单
      const menuTree = await api.queryAllMenus({})
      const menuData = flushMenus(menuTree) || []
      // 处理路由 得到每一级的路由设置
      this.$store.commit('d2admin/page/init', constantRoutes)
      // 设置顶栏菜单
      // this.$store.commit('d2admin/menu/headerSet', menuData)
      // 设置侧边栏菜单
      this.$store.commit('d2admin/menu/asideSet', menuData)
      // 初始化菜单搜索功能
      this.$store.commit('d2admin/search/init', menuData)
    }
  },
  created () {
    this.loadMenuRouter()
  },
  mounted () {
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
