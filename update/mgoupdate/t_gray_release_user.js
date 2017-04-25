var tableName = "t_gray_release_user";

//得到db
conn = new Mongo();
db = conn.getDB("test");

//清空数据库
db.t_gray_release_user.remove({});

//插入数据
db.t_gray_release_user.insert({
    "id": 1,
    "userid": 11205,
    "status": 1,
});