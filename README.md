# MyImageTool
MyImageTool  
一个可以方便扩展使用AI集成框架，方便部署模型，及在单卡上进行模型其切换

采用网页作为前端，网页后台用beego，数据库用mysql 或者mogodb
目前正在开发网站后台

模型可能会用到多个版本的cuda 或者cudnn 目前的想法是使用nvidia docker 来控制多个版本

其中
DeOldify 来自:https://github.com/jantic/DeOldify
暂时放入之后会尝试是否能用submodule 形式加入
