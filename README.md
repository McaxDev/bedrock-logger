## APIs
### 响应体结构
```json
{
    "message": "请求成功",
    "data": 可能是任何数据类型
}
```
### GET /tables
* 用于获取一共有哪些数据库表，目前有：place, break, die, interact, chat, session，online，分别是，放置方块，破坏方块，实体死亡，与方块互动，玩家聊天，玩家进退服务器，每分钟在线玩家的坐标
* 响应体的data字段结构：
```json
[
    {
        "name": "break",
        "display": "方块挖掘事件"
    },
    {
        "name": "place",
        "display": "方块放置事件"
    }
]
```
### GET /get/表名
* 此请求路径里的“表名”对应上一个API的响应体的对象数组里的对象的`name`字段，例如`break`或`place`
* 所有的查询字符串参数都是可选的，可以不带。
* 查询字符串参数（所有的路径(表)都能用的）：
  * `start_time`，筛选在给定时间之前的记录，值为`YYYY-MM-DD hh:mm:ss`格式。
  * `end_time`，筛选在给定时间之后的记录，值为`YYYY-MM-DD hh:mm:ss`格式。
  * `order`，指定排序顺序，值为`字段名 顺序`格式，顺序可以是`asc`升序，`desc`降序，例如`player asc`即是根据player字段升序排列。
  * `limit`，指定限制查询的记录数，值为整数，默认为10。
  * `page`，指定查第几页的，需要搭配`limit`使用，值为整数，默认为1。
* 查询字符串参数（只有一部分表能用）：
  * `x`或`y`或`z`，只有place break die interact能用，用于限制事件发生坐标。可以用`lt`（小于）`eq`（等于）`gt`（大于）比较符，格式示例：`x=lt100`（x小于100）`y=eq150`（y等于150）
  * `block_id`，只有place break interact能用，用于指定方块，例如`block_id=minecraft:dirt`就是筛选泥土相关
  * `player`，只有place break interact chat session能用，用于指定玩家，例如`player=Nerakolo`
  * `dimension`，只有place break interact die能用，用于指定事件发生的纬度，例如`dimension=overworld`就是限制主世界
  * `is_join`，只有session能用，为true代表进服，为false代表退服
  * 对于`die`表，除了可以用时间，xyz坐标，dimension维度，还能用`killer_id=minecraft:zombie`（指定杀手种类，这里面为僵尸），`killer_name=Nerakolo`（指定杀手名称），`dead_id`和`dead_name`也能用，与killer类似，不过dead代表被杀者
  * 对于`online`表，只能用时间来筛选
* 响应体里data字段的结构
```json
{
    "count": 100,
    "field": [
        {"key": "player", "title": "玩家名"},
        {"key": "block_id", "title": "方块ID"}
    ],
    "data": [
        {"player": "Nerakolo", "block_id": "minecraft:grass_block"},
        {"player": "Nerakolo", "block_id": "minecraft:dirt"}
    ]
}
```
