conn = new Mongo();
db = conn.getDB("test");

db.getCollection('t_goods_info').drop();

db.getCollection('t_goods_info').insert(
{
    "goodsid" : 1,
    "name" : "3万金币",
    "category" : 1,
    "pricetype" : 2,
    "price" : 30,
    "goodstype" : 1,
    "amount" : 30000,
    "discount" : "10折",
    "image" : "coin_1",
    "isshow" : true,
    "sort" : 1
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 2,
    "name" : "6万金币",
    "category" : 1,
    "pricetype" : 2,
    "price" : 60,
    "goodstype" : 1,
    "amount" : 60000,
    "discount" : "10折",
    "image" : "coin_2",
    "isshow" : true,
    "sort" : 2
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 3,
    "name" : "18万金币",
    "category" : 1,
    "pricetype" : 2,
    "price" : 180,
    "goodstype" : 1,
    "amount" : 180000,
    "discount" : "10折",
    "image" : "coin_3",
    "isshow" : true,
    "sort" : 3
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 4,
    "name" : "30万金币",
    "category" : 1,
    "pricetype" : 2,
    "price" : 300,
    "goodstype" : 1,
    "amount" : 300000,
    "discount" : "10折",
    "image" : "coin_4",
    "isshow" : true,
    "sort" : 4
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 5,
    "name" : "98万金币",
    "category" : 1,
    "pricetype" : 2,
    "price" : 680,
    "goodstype" : 1,
    "amount" : 980000,
    "discount" : "10折",
    "image" : "coin_5",
    "isshow" : true,
    "sort" : 5
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 6,
    "name" : "198万金币",
    "category" : 1,
    "pricetype" : 2,
    "price" : 1680,
    "goodstype" : 1,
    "amount" : 1980000,
    "discount" : "10折",
    "image" : "coin_6",
    "isshow" : true,
    "sort" : 6
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 7,
    "name" : "30钻石",
    "category" : 2,
    "pricetype" : 5,
    "price" : 3,
    "goodstype" : 2,
    "amount" : 30,
    "discount" : "10折",
    "image" : "diamond_1",
    "isshow" : true,
    "sort" : 1
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 8,
    "name" : "60钻石",
    "category" : 2,
    "pricetype" : 5,
    "price" : 6,
    "goodstype" : 2,
    "amount" : 60,
    "discount" : "10折",
    "image" : "diamond_2",
    "isshow" : true,
    "sort" : 2
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 9,
    "name" : "180钻石",
    "category" : 2,
    "pricetype" : 5,
    "price" : 18,
    "goodstype" : 2,
    "amount" : 180,
    "discount" : "10折",
    "image" : "diamond_3",
    "isshow" : true,
    "sort" : 3
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 10,
    "name" : "300钻石",
    "category" : 2,
    "pricetype" : 5,
    "price" : 30,
    "goodstype" : 2,
    "amount" : 300,
    "discount" : "10折",
    "image" : "diamond_4",
    "isshow" : true,
    "sort" : 4
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 11,
    "name" : "680钻石",
    "category" : 2,
    "pricetype" : 5,
    "price" : 68,
    "goodstype" : 2,
    "amount" : 680,
    "discount" : "10折",
    "image" : "diamond_5",
    "isshow" : true,
    "sort" : 5
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 12,
    "name" : "1680钻石",
    "category" : 2,
    "pricetype" : 5,
    "price" : 168,
    "goodstype" : 2,
    "amount" : 1680,
    "discount" : "10折",
    "image" : "diamond_6",
    "isshow" : true,
    "sort" : 6
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 13,
    "name" : "5红包",
    "category" : 3,
    "pricetype" : 3,
    "price" : 5,
    "goodstype" : 3,
    "amount" : 5,
    "discount" : "10折",
    "image" : "coin_6",
    "isshow" : true,
    "sort" : 1
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 14,
    "name" : "10红包",
    "category" : 3,
    "pricetype" : 3,
    "price" : 10,
    "goodstype" : 3,
    "amount" : 10,
    "discount" : "10折",
    "image" : "coin_6",
    "isshow" : true,
    "sort" : 2
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 15,
    "name" : "20红包",
    "category" : 3,
    "pricetype" : 3,
    "price" : 20,
    "goodstype" : 3,
    "amount" : 20,
    "discount" : "10折",
    "image" : "coin_6",
    "isshow" : true,
    "sort" : 3
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 16,
    "name" : "50红包",
    "category" : 3,
    "pricetype" : 3,
    "price" : 50,
    "goodstype" : 3,
    "amount" : 50,
    "discount" : "10折",
    "image" : "coin_6",
    "isshow" : true,
    "sort" : 4
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 17,
    "name" : "100红包",
    "category" : 3,
    "pricetype" : 3,
    "price" : 100,
    "goodstype" : 3,
    "amount" : 100,
    "discount" : "10折",
    "image" : "coin_6",
    "isshow" : true,
    "sort" : 5
}
);

db.getCollection('t_goods_info').insert(
{
    "goodsid" : 19,
    "name" : "房卡",
    "category" : 4,
    "pricetype" : 2,
    "price" : 90,
    "goodstype" : 101,
    "amount" : 3,
    "discount" : "10折",
    "image" : "",
    "isshow" : true,
    "sort" : 1
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 20,
    "name" : "房卡",
    "category" : 4,
    "pricetype" : 2,
    "price" : 180,
    "goodstype" : 101,
    "amount" : 6,
    "discount" : "10折",
    "image" : "",
    "isshow" : true,
    "sort" : 2
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 21,
    "name" : "房卡",
    "category" : 4,
    "pricetype" : 2,
    "price" : 540,
    "goodstype" : 101,
    "amount" : 18,
    "discount" : "10折",
    "image" : "",
    "isshow" : true,
    "sort" : 3
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 22,
    "name" : "房卡",
    "category" : 4,
    "pricetype" : 2,
    "price" : 900,
    "goodstype" : 101,
    "amount" : 30,
    "discount" : "10折",
    "image" : "",
    "isshow" : true,
    "sort" : 4
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 23,
    "name" : "房卡",
    "category" : 4,
    "pricetype" : 2,
    "price" : 2040,
    "goodstype" : 101,
    "amount" : 68,
    "discount" : "10折",
    "image" : "",
    "isshow" : true,
    "sort" : 5
}
);
db.getCollection('t_goods_info').insert(
{
    "goodsid" : 24,
    "name" : "房卡",
    "category" : 4,
    "pricetype" : 2,
    "price" : 5040,
    "goodstype" : 101,
    "amount" : 168,
    "discount" : "10折",
    "image" : "",
    "isshow" : true,
    "sort" : 6
}
);
