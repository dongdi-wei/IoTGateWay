#基于设备状态的物联网网关设计

主要功能：提供依据设备状态信息进行异常检测的框架

##主要模块：
状态信息收集模块：主要用socket编程，用与和局域网内的物联网设备交互收集数据
web服务器模块：基于gin框架，用于响应和用户的交互
规则解析模块：将用户输入的中缀表达式解析为后缀表达式并进行检测结果的计算
数据库交互模块：基于gorm框架，实现系统和数据库交互的统一接口
检测模块：根据规则解析模块解析出的检测规则列表并行计算每种检测规则下的检测结果并送给规则解析模块进行最终结果的计算
数据可视化模块：根据检测结果，绘制检测结果动图
账号管理模块：生成全局唯一账号（包括用户登录账号和检测订单号）

##待实现：
登陆管理模块：目前处于测试阶段，没有账号认证，正式发布时是需要登录管理的
流量统计模块：需要统计每个检测订单的检测的数据量与调用次数，进而才能进行计费管理
设备基本信息收集模块：需要用户填写检测设备的基本信息，比如产品批次，型号等等，方便云端大数据检测
人工标记模块：用js实现，将原始数据展示在web界面上，人工标记一些疑似异常数据作为训练源
设备管理模块：主要是探针的下发与设备隔离，探针的下发就是一个远程登陆模块，设备隔离暂时考虑使用小米路由器开发版提供的防火墙接口，但是设备隔离应该是有一套完整的隔离规则与防护级别的，目前暂时没有设计。
北向接口模块：用于接收SDS系统的指令，目前没有设计
实时报警模块：用于发现异常时进行报警（发送邮件）

###后续计划：
本项目主要是演示demo，主要流程基本开发完成，后续不定时更新
