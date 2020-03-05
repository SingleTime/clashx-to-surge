# clashx-to-surge
# 介绍
    本程序仅用于将clashx 的订阅节点提取出来并转换成surge 节点格式, 目前仅提供mac 版本 其他版本请自行编译
    
    
# 使用
    下载[执行文件](https://github.com/SingleTime/clashx-to-surge/releases/download/1.0/ClashxToSurge)
    
   
   ```shell script
      ##授予执行权限  
      chmod a+x ClashxToSurge

      ## 查看help
      ./ClashxToSurge --help 


      ## 进行转换
      ./ClashxToSurge -url=https://subscription.kyapi.xyz/modules/servers/V2raySocks/clash.php

      ## 进行转换 并通过关键词黑名单过滤
      ./ClashxToSurge -url=https://subscription.kyapi.xyz/modules/servers/V2raySocks/clash.php -blackKeys="回国"


   ```