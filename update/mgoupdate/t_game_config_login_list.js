conn = new Mongo();
db = conn.getDB("test");

db.getCollection('t_game_config_login_list').remove();

var loginIp = "192.168.199.120";

db.getCollection('t_game_config_login_list').insert({
    "GameId": 2.0,
    "name": "麻将",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0.0,
    "MaintainMsg": "测试斗地主的维护信息",
    "ReleaseTag": 1.0,
    "DownloadUrl": "http://d.tondeen.com/sjtexas.html",
    "LatestClientVersion": 1.0,
    "IP": loginIp,
    "PORT": 3802.0,
    "STATUS": 1.0
});
db.getCollection('t_game_config_login_list').insert({
    "GameId": 3.0,
    "name": "斗地主",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0.0,
    "MaintainMsg": "测试麻将的维护信息",
    "ReleaseTag": 1.0,
    "DownloadUrl": "http://d.tondeen.com/sjtexas.html",
    "LatestClientVersion": 1.0,
    "IP": loginIp,
    "PORT": 3799.0,
    "STATUS": 1.0
});
db.getCollection('t_game_config_login_list').insert({
    "GameId": 4.0,
    "name": "炸金花",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0.0,
    "MaintainMsg": "测试炸金花的维护信息",
    "ReleaseTag": 1.0,
    "DownloadUrl": "http://d.tondeen.com/sjtexas.html",
    "LatestClientVersion": 1.0,
    "IP": "192.168.199.120",
    "PORT": 3799.0,
    "STATUS": 2.0
});
db.getCollection('t_game_config_login_list').insert({
    "GameId": 5.0,
    "name": "大厅",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0.0,
    "MaintainMsg": "大厅的维护信息",
    "ReleaseTag": 1.0,
    "DownloadUrl": "http://d.tondeen.com/sjtexas.html",
    "LatestClientVersion": 1.0,
    "IP": loginIp,
    "PORT": 3801.0,
    "STATUS": 1.0
});
db.getCollection('t_game_config_login_list').insert({
    "GameId": 5.0,
    "name": "大厅",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0.0,
    "MaintainMsg": "大厅的维护信息",
    "ReleaseTag": 1.0,
    "DownloadUrl": "http://d.tondeen.com/sjtexas.html",
    "LatestClientVersion": 1.0,
    "IP": loginIp,
    "PORT": 3801.0,
    "STATUS": 1.0
});
db.getCollection('t_game_config_login_list').insert({
    "GameId": 8.0,
    "name": "抓瞎子",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0.0,
    "MaintainMsg": "测试抓虾子的维护信息",
    "ReleaseTag": 1.0,
    "DownloadUrl": "http://d.tondeen.com/sjtexas.html",
    "LatestClientVersion": 1.0,
    "IP": loginIp,
    "PORT": 3803.0,
    "STATUS": 1.0
});
