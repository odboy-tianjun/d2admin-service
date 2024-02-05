export default ({ service, request, tools }) => ({
  /**
   * @description 获取所有菜单
   */
  queryAllMenus (data = {}) {
    // 接口请求
    return request({
      url: '/queryAllMenus',
      method: 'post',
      data
    })
  }
})
