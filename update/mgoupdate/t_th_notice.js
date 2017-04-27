conn = new Mongo();
db = conn.getDB("test");

db.t_th_notice.remove({});

db.t_th_notice.insert({
    "id": 1,
    "noticetype": 1,
    "noticetitle": "来一圈游戏,邀约好友玩牌的利器",
    "noticecontent": "来一圈游戏, 邀约好友切磋牌技娱乐, 切不可赌博! 若有疑问, 请联系客服微信lyqyouxi",
    "noticememo": "",
    "noticefileds": []
});

db.t_th_notice.insert({
    "_id": ObjectId("57aae42b2e7eeff0be9514ab"),
    "id": 2, "noticetype": 2,
    "noticetitle": "来一圈棋牌，你的熟人约局利器",
    "noticecontent": "创建房间后,点击微信邀请,发送邀请链接给好友,好友安装App后可一键点击分享的房号进入房间. 赶快去试试吧",
    "noticememo": "",
    "noticefileds": ["购买钻石请联系微信", "lyqyouxi", "代理加盟请联系微信", "lyqyouxi", "投诉建议请联系客服QQ", "3298896830", "【湖南麻将】客服微信号", "229459253", "【湖南麻将】QQ", "229459253"]
});


db.t_th_notice.insert({
    "id": 3, "noticetype": 3,
    "noticetitle": "来一圈棋牌，你的熟人约局利器",
    "noticecontent": "如有购买房卡疑,请联系来一圈棋牌:\n\n客服微信号:lyqyouxi \n客户QQ号: 3298896830 \n关注公众号 :来一圈棋牌\n\n轻松加盟代理,赚取丰厚利润!",
    "noticememo": "",
    "noticefileds": []
});

db.t_th_notice.insert({
    "id": 4,
    "noticetype": 4,
    "noticetitle": "活动公告",
    "noticecontent": "来一圈游戏诚招各地代理！申请代理门槛极低，利润丰厚，轻松月入万元。欢迎联系洽谈！",
    "noticememo": null,
    "fileds": []
});
