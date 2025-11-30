import {createRouter,createWebHashHistory} from 'vue-router'
import {CONFIG} from '../config/index.js'
// import userRoutes from './user.js'

import {displayRoutes} from './display.js'

// const User = () => import('../views/User.vue')
const Login = () => import('../view/Login.vue')
// const Index = () => import('../view/Index.vue')
const Layout = () => import('../view/layout/Layout.vue')

// 定义用户相关的路由映射
// const userRoutes = []

// 定义路由映射
const routes = [
    // userRoutes,
    displayRoutes,
    {
        path: "/login",
        component: Login,
    },
    {
        path: "/",
        component: Layout
    },
    // 独立的图片获取路由，不需要认证和Layout
    // {
    //     path: "/display/get",
    //     component: () => import('../view/display/get.vue')
    // },
    // 根路径直接访问get页面的快捷方式
    // {
    //     path: "/get",
    //     component: () => import('../view/display/get.vue')
    // }
]

// 创建路由实例
const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

// 定义一个全局的守卫，去判断请求链接有灭有token字段
router.beforeEach(
  (to, from, next) => {
    // 1. 访问的是login，没有携带token ==> next()
    // 2. 访问的还是login, 携带了token ==> next("/")
    // 3. 访问的不是login，但是携带了token ==> next()
    // 4. 访问的不是login，没有携带token ==> next("/login")
    // 5. 访问的是不需要认证的公共路由 ==> next()
    // 拿到访问路径
    const toPath = to.path
    const toLogin = toPath.indexOf("/login") // 0  没有找到-1
    
    // 定义不需要认证的公共路由
    const publicRoutes = ["/display/get", "/get"]
    
    // 判断是否为公共路由
    const isPublicRoute = publicRoutes.includes(toPath)
    
    // 判断本地是否有token
    const tokenStatus = window.localStorage.getItem(CONFIG.TOKEN_NAME)
    
    if (isPublicRoute) {
        // 公共路由，直接放行
        next()
    } else if (toLogin == 0 && tokenStatus) {
        next("/")
    } else if (toLogin == 0 && !tokenStatus) {
        // 放行
        next()
    } else if (tokenStatus) {
        // 放行
        next()
    } else {
        next("/login")
    }
  }
)
// router.push("/xxxx")
export default router