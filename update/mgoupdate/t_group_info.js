conn = new Mongo();
db = conn.getDB("test");

db.getCollection('t_group_info').drop();

db.getCollection('t_group_info').insert({
    "id" : 18882387,
    "name" : "麻将测试群",
    "info" : "测试群组开房。。",
    "owner" : 10192,
    "members" : [
        {
            "uid" : 10192,
            "nickname" : "秋风",
            "remark" : "",
            "openid" : "ojrylw6FHlBUkcvcn6uVhTzgJCY0",
            "headimg" : "http://wx.qlogo.cn/mmopen/A8u9BjAYLHr0ibsdaqF3aNNQ5EAHjI6hiaqXjDF44RZujvx0hkicKf6EofCKtbBmuRQtbFXMq5Vt7ibP3sI7oUzZViae43CGzdeFV/0"
        },
        {
            "uid" : 14357,
            "nickname" : "游客26675",
            "remark" : "",
            "openid" : "",
            "headimg" : ""
        },
        {
            "uid" : 14632,
            "nickname" : "游客19783",
            "remark" : "",
            "openid" : "",
            "headimg" : ""
        },
        {
            "uid" : 14633,
            "nickname" : "游客92779",
            "remark" : "",
            "openid" : "",
            "headimg" : ""
        },
        {
            "uid" : 14634,
            "nickname" : "游客76634",
            "remark" : "",
            "openid" : "",
            "headimg" : ""
        },
        {
            "uid" : 14636,
            "nickname" : "游客99051",
            "remark" : "",
            "openid" : "",
            "headimg" : ""
        },
        {
            "uid" : 14647,
            "nickname" : "游客36661",
            "remark" : "",
            "openid" : "",
            "headimg" : ""
        }
    ],
    "gameopts" : [
        {
            "id" : 0,
            "gameid" : 7,
            "remark" : "跑得快",
            "option" : "经典跑得快 20局 2人 首出黑桃3 不抓鸟 显示余牌"
        },
        {
            "id" : 0,
            "gameid" : 22,
            "remark" : "红中麻将抓鸟群",
            "option" : "红中麻将 16局 4人 鸟不翻倍 抓6鸟"
        }
    ],
    "syncid" : 9
});
