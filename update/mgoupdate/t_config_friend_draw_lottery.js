conn = new Mongo();
db = conn.getDB("test");

db.getCollection('t_config_friend_draw_lottery').drop();

//谢谢参与
db.t_config_friend_draw_lottery.insert(
{
    "sort" : 1,                 //排序
    "type" : 401,               //奖品类型
    "name" : "谢谢参与",         //名称
    "number" : 1,               //奖品数量
    "percent" : 10,             //中奖概率
    "version" : 1
}
);

//1张房卡
db.t_config_friend_draw_lottery.insert(
{
    "sort" : 2,
    "type" : 101,
    "name" : "1张",
    "number" : 1,
    "percent" : 80,
    "version" : 1
}
);

//10元话费
db.t_config_friend_draw_lottery.insert(
{
    "sort" : 3,
    "type" : 301,
    "name" : "10元",
    "number" : 1,
    "percent" : 0,
    "version" : 1
}
);

//再来一次
db.t_config_friend_draw_lottery.insert(
{
    "sort" : 4,
    "type" : 402,
    "name" : "再来一次",
    "number" : 1,
    "percent" : 20,
    "version" : 1
}
);

//2张房卡
db.t_config_friend_draw_lottery.insert(
{
    "sort" : 5.0,
    "type" : 101,
    "name" : "2张",
    "number" : 2,
    "percent" : 5,
    "version" : 1
}
);

//30元话费
db.t_config_friend_draw_lottery.insert(
{
    "sort" : 6.0,
    "type" : 306,
    "name" : "30元",
    "number" : 1,
    "percent" : 0,
    "version" : 1
}
);

//3张房卡
db.t_config_friend_draw_lottery.insert(
{
    "sort" : 7.0,
    "type" : 101,
    "name" : "3张",
    "number" : 3,
    "percent" : 2,
    "version" : 1
}
);

//50元话费
db.t_config_friend_draw_lottery.insert(
{
    "sort" : 8.0,
    "type" : 302,
    "name" : "50元",
    "number" : 1,
    "percent" : 0,
    "version" : 1
}
);
