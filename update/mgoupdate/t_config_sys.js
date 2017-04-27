conn = new Mongo();
db = conn.getDB("test");


//1.数据库需要增加配置文件


//更新日志:04-26 更新新用户注册，赠送10张房卡
db.t_config_sys.insert({
    id: 1,
    newusercoin: 1000,
    newuserroomcard: 10
});
