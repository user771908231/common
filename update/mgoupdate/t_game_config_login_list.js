var loginIp = "182.92.225.121";
var prod1 = "182.92.225.121";
var prod2 = "47.93.150.22";
var prod3 = "47.93.150.194";
var prod4 = "182.92.179.230"; //测试服
var tableName = "t_game_config_login_list_shenzhen";
var defaultDownloadUrl = "http://d.tondeen.com/sjtexas.html";

conn = new Mongo();
db = conn.getDB("test");

db.getCollection(tableName).remove({});

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
    "IP": prod4,
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
    "IP": prod4,
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
    "IP": prod4,
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
    "IP": prod4,
    "PORT": 3801,
    "STATUS": 1.0
});

db.getCollection(tableName).insert({
    "GameId": 7.0,
    "name": "跑得快",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0.0,
    "MaintainMsg": "跑得快",
    "ReleaseTag": 1.0,
    "DownloadUrl": defaultDownloadUrl,
    "LatestClientVersion": 1.0,
    "IP": prod4,
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
    "IP": prod4,
    "PORT": 3803,
    "STATUS": 1.0
});

db.getCollection(tableName).insert({
    "GameId": 9.0,
    "name": "白山麻将",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0,
    "MaintainMsg": "白山麻将",
    "ReleaseTag": 1.0,
    "DownloadUrl": defaultDownloadUrl,
    "LatestClientVersion": 1.0,
    "IP": prod4,
    "PORT": 3804,
    "STATUS": 1.0
});


db.getCollection(tableName).insert({
    "GameId": 10,
    "name": "长春麻将",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0,
    "MaintainMsg": "长春麻将",
    "ReleaseTag": 1.0,
    "DownloadUrl": defaultDownloadUrl,
    "LatestClientVersion": 1.0,
    "IP": prod4,
    "PORT": 3805,
    "STATUS": 1.0
});

db.getCollection(tableName).insert({
    "GameId": 11,
    "name": "牛牛",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0,
    "MaintainMsg": "牛牛",
    "ReleaseTag": 1.0,
    "DownloadUrl": defaultDownloadUrl,
    "LatestClientVersion": 1.0,
    "IP": prod4,
    "PORT": 3807,
    "STATUS": 1.0
});

db.getCollection(tableName).insert({
    "GameId": 12,
    "name": "搬坨子",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0,
    "MaintainMsg": "搬坨子",
    "ReleaseTag": 1.0,
    "DownloadUrl": defaultDownloadUrl,
    "LatestClientVersion": 1.0,
    "IP": prod4,
    "PORT": 3806,
    "STATUS": 1.0
});

db.getCollection(tableName).insert({
    "GameId": 22,
    "name": "转转麻将",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0,
    "MaintainMsg": "转转麻将",
    "ReleaseTag": 1.0,
    "DownloadUrl": defaultDownloadUrl,
    "LatestClientVersion": 1.0,
    "IP": prod4,
    "PORT": 3818,
    "STATUS": 1.0
});

db.getCollection(tableName).insert({
    "GameId": 16,
    "name": "松江河麻将",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0,
    "MaintainMsg": "松江河麻将",
    "ReleaseTag": 1.0,
    "DownloadUrl": defaultDownloadUrl,
    "LatestClientVersion": 1.0,
    "IP": prod4,
    "PORT": 3815,
    "STATUS": 1.0
});

db.getCollection("t_game_config_login_list_shenzhen").insert({
    "GameId": 24,
    "name": "海南麻将",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0,
    "MaintainMsg": "海南麻将",
    "ReleaseTag": 1.0,
    "DownloadUrl": "http://d.tondeen.com/sjtexas.html",
    "LatestClientVersion": 1.0,
    "IP": "182.92.179.230",
    "PORT": 3820,
    "STATUS": 1.0
});
db.getCollection("t_game_config_login_list").insert({
    "GameId": 24,
    "name": "海南麻将",
    "CurVersion": 1.0,
    "IsUpdate": 0.0,
    "IsMaintain": 0,
    "MaintainMsg": "海南麻将",
    "ReleaseTag": 1.0,
    "DownloadUrl": "http://d.tondeen.com/sjtexas.html",
    "LatestClientVersion": 1.0,
    "IP": "182.92.179.230",
    "PORT": 3820,
    "STATUS": 1.0
});