# MyImageTool
MyImageTool  (开发中)  
一个可以方便扩展使用AI部署框架，方便模型部署，及在单卡上进行模型其切换，让跑过的模型不再积灰

主要分为网页前端，center，inference_app 三层

网页前端：上传显示图片，选择推理应用

center：管理 各个模型从内存到显存的加载，及图片消息队列维护与推送

inference_app：部署模型进行推理


TODO:  
模型可能会用到多个版本的cuda 或者cudnn 目前的想法是使用nvidia docker 来控制多个版本  
center控制负载均衡，inference_app弹性扩容   
支持单机部署，及分布式部署  
服务之间使用grpc进行通信，使用glog 进行日志记录    
拟用k8s来管理docker  
