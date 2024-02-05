[D2Admin](https://github.com/d2-projects/d2-admin) 是一个完全 **开源免费** 的企业中后台产品前端集成方案，使用最新的前端技术栈，小于 60kb 的本地首屏 js 加载，已经做好大部分项目前期准备工作，并且带有大量示例代码，助力管理系统敏捷开发。

**中文** | [English](https://github.com/d2-projects/d2-admin-start-kit)

## 预览

![Deploy preview](https://github.com/d2-projects/d2-admin-start-kit/workflows/Deploy%20preview/badge.svg)
[![Netlify Status](https://api.netlify.com/api/v1/badges/08ff8c93-f0a8-497a-a081-440b31fb3aa4/deploy-status)](https://app.netlify.com/sites/d2-admin-start-kit/deploys)

下列访问地址均由最新的 master 分支代码同时构建部署，访问效果完全一致，请根据自身网络情况选择合适的访问链接。

| 位置 | 链接 | 部署位置 |
| --- | --- | --- |
| d2.pub | [preview](https://d2.pub/d2-admin-start-kit/preview) | 中国服务器 |
| cdn.d2.pub | [preview](https://cdn.d2.pub/d2-admin-start-kit/preview) | 七牛云 CDN |
| github | [preview](https://d2-projects.github.io/d2-admin-start-kit) | GitHub pages |
| netlify | [preview](https://d2-admin-start-kit.netlify.com) | Netlify CDN |

## 其它同步仓库

| 位置 | 链接 |
| --- | --- |
| 码云 | [https://gitee.com/d2-projects/d2-admin-start-kit](https://gitee.com/d2-projects/d2-admin-start-kit) |
| coding | [https://d2-projects.coding.net/p/d2-projects/d/d2-admin-start-kit/git](https://d2-projects.coding.net/p/d2-projects/d/d2-admin-start-kit/git) |

# 以下内容为开发实践
## node && npm version
nodejs: v16.18.0
npm: 8.19.2

## install package
npm config set registry https://registry.npmmirror.com
sudo npm install yarn -g
yarn install

## 文档地址
https://d2.pub/doc/d2-admin/

### vue-cli-service serve和npm run serve的区别
```text
在原package.json中:
"scripts": {
    "serve": "vue-cli-service serve --open",
    "dev": "npm run serve",
    "build": "vue-cli-service build",
    "build:preview": "NODE_OPTIONS=--max_old_space_size=4096 vue-cli-service build --mode preview",
    "lint": "vue-cli-service lint --fix",
    "test:unit": "vue-cli-service test:unit"
}
vue-cli-service serve和npm run serve在功能上没有区别。
vue-cli-service serve是直接调用Vue CLI服务工具来启动本地开发服务器，它会监听源码变化并自动热更新浏览器内容，通常用于开发环境
npm run serve是npm包管理器提供的命令，用来执行package.json中指定的脚本，当执行npm run serve时，npm会查找并运行scripts部分定义的"serve"命令，
因此实际上它也是调用vue-cli-service命令。
```

### vue-cli-service不同环境配置
```text
在 Vue CLI 项目中，环境变量配置是通过不同的 .env 文件实现的。这些文件允许你在不同环境下设置不同的变量值，例如开发环境、测试环境和生产环境。
Vue CLI 使用 dotenv-webpack 插件处理环境变量，其中以 VUE_APP_ 开头的变量会被 webpack 的 DefinePlugin 静态嵌入到客户端侧的包中，这样你就可以在代码中直接使用它们，并且在构建时会根据当前环境替换为对应的值。
以下是如何配置不同环境变量的基本结构：
默认环境：所有环境都可读取的基础变量可以放在 .env 文件中。
开发环境：针对开发环境（npm run dev）的变量，放在 .env.development 文件中。
生产环境：针对生产环境（npm run build）的变量，放在 .env.production 文件中。

当你执行 vue-cli-service build 时，默认会使用生产模式（.env.production），如果你需要构建特定模式下的项目，可以通过 --mode 参数指定：
vue-cli-service build --mode staging
在这种情况下，你需要有一个名为 .env.staging 的文件来提供相应环境的变量配置。
```

### 在vuex中，mapActions的作用
```text
在 Vuex 中，mapActions 是一个辅助函数，它来自于 vuex 组件绑定插件（VueX 提供的）。这个函数的作用是将 Vuex store 中定义的 actions 映射为
Vue 组件的本地方法，使得在组件内部可以更方便地调用这些 action，而无需直接通过 this.$store.dispatch()。
例如，在 Vuex store 中有以下 actions 定义：

具体参考'src/views/system/login/page.vue'中的:
import { mapActions } from 'vuex'

// 使用 mapActions 将 store 的 actions 映射为本组件的方法
...mapActions('d2admin/account', [
    'login'
])

// 现在可以在组件内像调用普通方法一样调用它们
// 相当于 this.$store.dispatch('login')
this.login()
```

### 请求带token关键代码
```text
web/src/api/service.js
```

### 加载顺序
```text
router > api
```

### store 例子
```vue
// store/index.js
import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
 state: {
   count: 0,
   // 其他状态...
 },
 mutations: {
   increment(state) {
     state.count++
   },
   // 其他 mutation 函数...
 },
 actions: {
   // 可能包含异步操作的 action 函数...
 },
 getters: {
   getCount: state => state.count,
   // 其他 getter 函数...
 }
})
```
```vue
// 在任何一个 Vue 组件中，你就可以通过 this.$store 访问 Vuex store
// 任意Vue组件内部
export default {
    mounted() {
        console.log(this.$store.state.count) // 访问状态
        this.$store.commit('increment') // 提交 mutation 更新状态
        this.$store.dispatch('someAction') // 分发 action
        const count = this.$store.getters.getCount // 获取 getter 数据
    }
}
```


