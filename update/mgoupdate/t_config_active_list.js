conn = new Mongo();
db = conn.getDB("test");

db.getCollection('t_config_active_list').drop();

db.t_config_active_list.insert(
{
    "id" : 0,
    "type" : 1,
    "reward" : 1,
    "richtext" : [
        "在来一圈棋牌任意金币场游戏，只需赢五局就可拆红包，所得红包在游戏商城兑换区进行兑换，登录微信进入“来一圈棋牌”公众号即可一键领取。",
        "banner"
    ],
    "title" : "清明节"
}
);
db.t_config_active_list.insert(
{
    "id" : 1,
    "type" : 1,
    "reward" : 1,
    "richtext" : [
        " 来一圈棋牌为广大玩家提供休闲娱乐的平台，旗下所属的游戏均为模拟竞技类游戏，我们希望每一位玩家在此能感受到竞技的乐趣。游戏中使用的游戏金币为游戏道具，不具有任何的实际价值，只限玩家本人在游戏中使用。",
        ""
    ],
    "title" : "重要公告"
}
);
db.t_config_active_list.insert(
{
    "id" : 2,
    "type" : 1,
    "reward" : 1,
    "richtext" : [
        "邀请好友注册游戏，即可获金币奖励！好友玩游戏还可获得更多奖励，钻石、房卡、奖券送不停！",
        ""
    ],
    "title" : "邀请奖励"
}
);
db.t_config_active_list.insert(
{
    "_id" : ObjectId("5879fc0ca513c0bbc83c1147"),
    "id" : 3,
    "type" : 1,
    "reward" : 1,
    "richtext" : [
        "登录即可抽取幸运转盘大奖，连续登录7天更有惊喜大奖！奖品不限量，人人有份拿！",
        ""
    ],
    "title" : "连续登录"
}
);
