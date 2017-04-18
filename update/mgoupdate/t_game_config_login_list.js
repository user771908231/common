var loginIp = "182.92.225.121";
var tableName = "t_game_config_login_list";
var defaultDownloadUrl = "http://d.tondeen.com/sjtexas.html";

conn = new Mongo();
db = conn.getDB("test");

db.getCollection('t_game_config_login_list').remove({});


db.getCollection(tableName).insert({
    "GameId": 2.0,
    "name": "麻将",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0.0,
    "MaintainMsg": "麻将",
    "ReleaseTag": 1.0,
    "DownloadUrl": defaultDownloadUrl,
    "LatestClientVersion": 1.0,
    "IP": loginIp,
    "PORT": 3798,
    "STATUS": 1.0
});

db.getCollection(tableName).insert({
    "GameId": 3.0,
    "name": "斗地主",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0.0,
    "MaintainMsg": "斗地主",
    "ReleaseTag": 1.0,
    "DownloadUrl": defaultDownloadUrl,
    "LatestClientVersion": 1.0,
    "IP": loginIp,
    "PORT": 3799,
    "STATUS": 1.0
});

db.getCollection(tableName).insert({
    "GameId": 4.0,
    "name": "炸金花",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0.0,
    "MaintainMsg": "炸金花",
    "ReleaseTag": 1.0,
    "DownloadUrl": defaultDownloadUrl,
    "LatestClientVersion": 1.0,
    "IP": "192.168.199.120",
    "PORT": 3800,
    "STATUS": 2.0
});

db.getCollection(tableName).insert({
    "GameId": 5.0,
    "name": "大厅",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0.0,
    "MaintainMsg": "大厅",
    "ReleaseTag": 1.0,
    "DownloadUrl": defaultDownloadUrl,
    "LatestClientVersion": 1.0,
    "IP": loginIp,
    "PORT": 3801,
    "STATUS": 1.0
});

db.getCollection(tableName).insert({
    "GameId": 8.0,
    "name": "跑得快",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0.0,
    "MaintainMsg": "跑得快",
    "ReleaseTag": 1.0,
    "DownloadUrl": defaultDownloadUrl,
    "LatestClientVersion": 1.0,
    "IP": loginIp,
    "PORT": 3802,
    "STATUS": 1.0
});

db.getCollection(tableName).insert({
    "GameId": 8.0,
    "name": "抓瞎子",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0.0,
    "MaintainMsg": "抓瞎子",
    "ReleaseTag": 1.0,
    "DownloadUrl": defaultDownloadUrl,
    "LatestClientVersion": 1.0,
    "IP": loginIp,
    "PORT": 3803,
    "STATUS": 1.0
});

