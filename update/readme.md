## 更新及部署说明：
### 2017-03-18:
1. 更新t_config_active_list表
2. 在各个游戏初始化时初始化common/service下的pushService，
例：pushService.PoolInit("192.168.199.155:2801"),建议将HallTcpAddr字段配置到json文件中。
3. 先部署大厅
4. 再部署casino_super、casino_mj、casino_ddz、casino_zjh
### 2017-03-20:
1. 更新t_th_notice表。
### 2017-03-25:
1. 更新t_config_task_info表。
2. 新版红包系统还差房费统计、不同游戏及房间金币兑率配置等。目前暂时兼容老版红包

### 2017-03-29:
1. 更新t_config_active_list.txt表
2. 更新t_th_notice.txt表
3. 更新大厅：
    1. 修复了红包任务领取时总时0元的bug。
    2. 修复了保险箱提示里含有错误码的问题。
    3. 将活动里的"迎新春"改为"清明节"。
