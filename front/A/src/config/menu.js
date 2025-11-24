export const MENU_CONFIG = [
    // 集群管理
    {
        title: "图片管理",
        index: "/imagedisplay",
        icon: "iconfont icon-tiaodu",
        items: [
         // 查看图片
         {
             index: "/display/get",
             title: "查看图片"
         },
        //  上传图片
        {
            index: "/display/upload",
            title: "上传图片"
        },
        //更新图片
        {
            index: "/display/update",
            title: "更新图片"
        },
        //删除图片
        {
            index: "/display/delete",
            title: "删除图片"
        },
        //获取所有图片
        {
            index: "/display/getall",
            title: "获取所有图片"
        },
        ]
    }
]