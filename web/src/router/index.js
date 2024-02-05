import Vue from 'vue'
import VueRouter from 'vue-router'

// 进度条
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

import store from '@/store/index'
import util from '@/libs/util.js'
import layoutHeaderAside from '@/layout/header-aside'
// 由于懒加载页面太多的话会造成webpack热更新太慢，所以开发环境不使用懒加载，只有生产环境使用懒加载
const _import = require('@/libs/util.import.' + process.env.NODE_ENV)

// fix vue-router NavigationDuplicated
const VueRouterPush = VueRouter.prototype.push
VueRouter.prototype.push = function push (location) {
  return VueRouterPush.call(this, location).catch(err => err)
}
const VueRouterReplace = VueRouter.prototype.replace
VueRouter.prototype.replace = function replace (location) {
  return VueRouterReplace.call(this, location).catch(err => err)
}

Vue.use(VueRouter)

export function buildRouter (menuItem) {
  return {
    path: menuItem.routerPath,
    name: menuItem.routerName,
    meta: {
      title: menuItem.menuTitle,
      auth: menuItem.routerAuth === 1,
      hidden: menuItem.routerHidden === 1,
      cache: menuItem.routerHidden === 1
    },
    component: _import(menuItem.routerComponentPath)
  }
}

export function buildRouterList (menuTree) {
  // 权限路由, 全是一级路由
  const routerData = []
  for (const menuItem of menuTree) {
    const children = menuItem.children
    if (children && children.length > 0) {
      for (const childMenuItem of children) {
        if (childMenuItem.routerPath && childMenuItem.routerComponentPath) {
          routerData.push(buildRouter(childMenuItem))
        }
      }
    }
  }
  return routerData
}

/**
 * @description 创建在 layout 中显示的路由设置
 * @param {Array} routes 动态路由设置
 */
export function createRoutesInLayout (routes = []) {
  return [
    {
      path: '/',
      redirect: { name: 'index' },
      component: layoutHeaderAside,
      children: [
        { path: 'index', name: 'index', meta: { title: '首页', auth: true }, component: _import('system/index') },
        { path: 'log', name: 'log', meta: { title: '前端日志', auth: true }, component: _import('system/log') },
        ...routes
      ]
    }
  ]
}

// 在 layout 之外显示的路由
export const routesOutLayout = [
  // 刷新页面 必须保留
  { path: '/refresh', name: 'refresh', component: _import('system/function/refresh'), hidden: true },
  // 页面重定向 必须保留
  { path: '/redirect/:route*', name: 'redirect', component: _import('system/function/redirect'), hidden: true },
  // 登陆页面 必须保留
  { path: '/login', name: 'login', component: _import('system/login'), hidden: true },
  { path: '*', name: '404', component: _import('system/error/404'), hidden: true }
]

// 默认的路由
export const constantRoutes = createRoutesInLayout().concat(routesOutLayout)

/**
 * @description 创建路由
 * @param {Array} routes 路由设置
 */
const createRouter = (routes = []) => new VueRouter({
  scrollBehavior: () => ({ y: 0 }),
  routes
})

// 导出路由 在 main.js 里使用
const router = createRouter(constantRoutes)

/**
 * @description 重新设置路由
 * @param {Array} routes 额外追加的路由
 */
export function resetRouter (routes = []) {
  router.matcher = createRouter(routes).matcher
}

/**
 * 路由拦截
 * 权限验证
 */
router.beforeEach(async (to, from, next) => {
  // 进度条
  NProgress.start()
  console.log('即将进入的路由:', to)
  console.log('当前离开的路由:', from)
  // 获取路由参数或查询参数
  console.log('路由参数:', to.params)
  console.log('查询参数:', to.query)
  // 确认已经加载多标签页数据 https://github.com/d2-projects/d2-admin/issues/201
  await store.dispatch('d2admin/page/isLoaded')
  // 确认已经加载组件尺寸设置 https://github.com/d2-projects/d2-admin/issues/198
  await store.dispatch('d2admin/size/isLoaded')
  // 加载动态路由 内部已经做了对登录状态和是否已经加载动态路由的判断
  await store.dispatch('d2admin/router/load', { to: to.fullPath })
  // 关闭搜索面板
  store.commit('d2admin/search/set', false)
  // 验证当前路由所有的匹配中是否需要有登录验证的
  if (to.matched.some(r => r.meta.auth)) {
    // 这里暂时将cookie里是否存有token作为验证是否登录的条件
    // 请根据自身业务需要修改
    const token = util.cookies.get('token')
    if (token && token !== 'undefined') {
      next()
    } else {
      // 没有登录的时候跳转到登录界面
      // 携带上登陆成功之后需要跳转的页面完整路径
      next({
        name: 'login',
        query: {
          redirect: to.fullPath
        }
      })
      // https://github.com/d2-projects/d2-admin/issues/138
      NProgress.done()
    }
  } else {
    // 不需要身份校验 直接通过
    next()
  }
})

router.afterEach(to => {
  // 进度条
  NProgress.done()
  // 多页控制 打开新的页面
  store.dispatch('d2admin/page/open', to)
  // 更改标题
  util.title(to.meta.title)
})

export default router
