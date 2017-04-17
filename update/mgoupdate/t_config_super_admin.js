conn = new Mongo();
db = conn.getDB("test");

db.getCollection('t_super_admin').drop();

db.getCollection('t_super_admin').insert({
    "_id" : ObjectId("589af0b70e4041708bb8e4d5"),
    "id" : 10007.0,
    "nickname" : "admin",
    "passwd" : "33ced495ab6ce4eb238053ab1e28e176"
});
