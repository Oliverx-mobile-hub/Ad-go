const Layout = () => import('../view/layout/Layout.vue')
// 广告资源路由
export const displayRoutes = 
    {
        path: "/display",
        component: Layout,
        redirect: "/display/get",
        children: [
            {
                path: 'get',
                component: ()=> import('../view/display/get.vue')
            },
            {
                path: 'upload',
                component: ()=> import('../view/display/upload.vue')
            },
            {
                path: 'update',
                component: ()=> import('../view/display/update.vue')
            },
            {
                path: 'delete',
                component: ()=> import('../view/display/delete.vue')
            },
            {
                path: 'getall',
                component: ()=> import('../view/display/getall.vue')
            }
            
        ]
    }