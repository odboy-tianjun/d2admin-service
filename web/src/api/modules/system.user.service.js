export default ({ service, request, tools }) => ({
  /**
   * @description 登录
   * @param {Object} data 登录携带的信息
   */
  login (data = {}) {
    // 接口请求
    return request({
      url: '/login',
      method: 'post',
      data
    })
  }
})
