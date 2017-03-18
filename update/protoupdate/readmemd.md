
# 修改说明
###  按天增加文件：同一天修改的proto放在同一个文件
###  修改信息包括：修改的文件-协议名称-字段名称-修改原因-影响版本

## 更新及部署说明：
### 2017-03-18:
1. 更新t_config_active_list表
2. 在各个游戏初始化时初始化common/service下的pushService，
例：pushService.PoolInit("192.168.199.155:2801"),建议将HallTcpAddr字段配置到json文件中。
3. 先部署大厅
4. 再部署casino_super、casino_mj、casino_ddz、casino_zjh
