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

db.t_gray_release_user.insert({
    "id": 2,
    "userid": 11977,
    "status": 1,
});

db.t_gray_release_user.insert({
    "id": 3,
    "userid": 11564,
    "status": 1,
});

db.t_gray_release_user.insert({
    "id": 4,
    "userid": 10356,
    "status": 1,
});

db.t_gray_release_user.insert({
    "id": 5,
    "userid": 11533,
    "status": 1,
});

//  11977 QWERTY 邓艳梅
//  11564 小药丸  钟春燕
//  10356 kory  杨荣华
//  11533 一根藤上七个娃娃 iPhone6
