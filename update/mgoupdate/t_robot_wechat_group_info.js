conn = new Mongo();
db = conn.getDB("test");

//先清空表
db.getCollection('t_robot_wechat_group_info').drop();

//插入数据
db.getCollection('t_robot_wechat_group_info').insert({
    "groupname": "来一圈转转麻将",  //群名称
    "ownerid": 12921,    //群主id
    "gametype": "转转麻将"     //游戏类型
});

db.getCollection('t_robot_wechat_group_info').insert({
    "groupname": "来一圈红中麻将",  //群名称
    "ownerid": 12921,    //群主id
    "gametype": "红中麻将"     //游戏类型
});

db.getCollection('t_robot_wechat_group_info').insert({
    "groupname": "来一圈经典跑得快",  //群名称
    "ownerid": 12921,    //群主id
    "gametype": "经典跑得快"     //游戏类型
});

db.getCollection('t_robot_wechat_group_info').insert({
    "groupname": "来一圈十五张跑得快",  //群名称
    "ownerid": 12921,    //群主id
    "gametype": "十五张跑得快"     //游戏类型
});

