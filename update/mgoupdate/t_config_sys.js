conn = new Mongo();
db = conn.getDB("test");


//1.删除现有配置
db.getCollection('t_config_sys').drop();

//更新日志:04-26 更新新用户注册，赠送3张房卡
db.t_config_sys.insert({
    id: 1,
    newusercoin: 1000,
    newuserroomcard: 3
});
