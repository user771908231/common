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

db.t_gray_release_user.insert({
    "id": 6,
    "userid": 11754,
    "status": 1,
});

db.t_gray_release_user.insert({
    "id": 7,
    "userid": 11317,
    "status": 1,
});

db.t_gray_release_user.insert({
    "id": 8,
    "userid": 12074,
    "status": 1,
});

db.t_gray_release_user.insert({
    "id": 9,
    "userid": 12039,
    "status": 1,
});

