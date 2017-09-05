conn = new Mongo();
db = conn.getDB("test");

//跑得快战绩索引
db.getCollection('t_pdk_desk_round_all').ensureIndex({userids:1,deskid:1});
db.getCollection('t_pdk_desk_round').ensureIndex({userids:1,deskid:1,gamenumber:1});
db.getCollection('t_pdk_desk_round_playback').ensureIndex({gamenumber:1});


//转转麻将战绩索引
db.getCollection('t_mj_zhzh_desk_round_all').ensureIndex({userids:1,friendplay:1,deskid:1});
db.getCollection('t_mj_zhzh_desk_round').ensureIndex({userids:1,friendplay:1,deskid:1,gamenumber:1});

db.getCollection('t_mj_hzh_desk_round_all').ensureIndex({userids:1,friendplay:1,deskid:1});
db.getCollection('t_mj_hzh_desk_round').ensureIndex({userids:1,friendplay:1,deskid:1,gamenumber:1});


//牛牛战绩索引
db.getCollection('t_niuniu_desk_round_all').ensureIndex({userids:1,endtime:1});
db.getCollection('t_niuniu_desk_round_one').ensureIndex({userids:1,deskid:1,gamenumber:1});


