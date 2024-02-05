import { uniqueId } from 'lodash'

/**
 * @description 给菜单数据补充上 path 字段
 * @description https://github.com/d2-projects/d2-admin/issues/209
 * @param {Array} menu 原始的菜单数据
 */
export function supplementPath (menu) {
  return menu.map(e => ({
    ...e,
    path: e.path || uniqueId('d2-menu-empty-'),
    ...e.children ? {
      children: supplementPath(e.children)
    } : {}
  }))
}

/**
 * 设置菜单
 * @returns {{path: string, icon: string, title: string}[]}
 */
export function flushMenus (menuTree = []) {
  let menuData = [{ path: '/index', title: '首页', icon: 'home' }]
  try {
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

// const menuData = [
//   { path: '/index', title: '首页', icon: 'home' },
//   {
//     title: '页面',
//     icon: 'folder-o',
//     children: [
//       { path: '/page1', title: '页面 1' },
//       { path: '/page2', title: '页面 2' },
//       { path: '/page3', title: '页面 3' }
//     ]
//   }
// ]
//
// export const menuHeader = supplementPath(menuData)
//
// export const menuAside = supplementPath(menuData)
