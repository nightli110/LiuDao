# MyImageTool
MyImageTool  
一个可以方便扩展使用AI部署框架，方便模型部署，及在单卡上进行模型其切换，让跑过的模型不再积灰

采用网页作为前端，网页后台用beego，数据库用mysql  

由一个center node管理 各个模型从内存到显存的加载，及图片消息队列  
服务之间使用grpc进行通信，使用glog 进行日志记录  
目前正在开发网站后台，数据表设计

模型可能会用到多个版本的cuda 或者cudnn 目前的想法是使用nvidia docker 来控制多个版本
拟用k8s来管理docker(k8s 学习中)

其中
DeOldify 来自:https://github.com/jantic/DeOldify
暂时放入之后会尝试是否能用submodule 形式加入
