# CHESS

## 介绍

这是一个用于创建和加入国际象棋房间的API，可以让用户在房间中进行下棋和观战。我开发这个项目的动机是为了学习如何使用WebSocket来实现实时通讯。这个项目解决了如何在不同的客户端之间同步棋盘状态的问题。

## 目录

- [安装](#安装)
- [使用](#使用)
- [接口](#接口)
- [加分项](#加分项)

## 安装

要运行这个项目，你需要修改此项目中dao层的dao.go，使其能正确连接上你的MySQL数据库。

之后，你需要在你的MySQL数据库中创建名为“chess”的database。

这样，你就可以正确地运行这个项目了。


## 使用

要使用这个API，你需要有一个客户端应用，比如一个网页或者一个移动应用。你可以使用任何支持WebSocket的技术来创建客户端应用。

要创建或加入一个房间，你需要向服务器发送HTTP请求。要在房间中进行下棋或观战，你需要与服务器建立WebSocket连接。

该项目已经部署到`120.79.27.213:9099`,你可以在2023年6月30日之前通过将下面接口的`localhost:9099`替换为该地址来访问已经部署的示例。

## 接口

以下接口采用RESTful风格

### POST http://localhost:9099/user/register

需要提交表单username、telephone、password，其中telephone要等于11位数，password要大于等于6位字符

该接口返回JSON字符串以表示是否注册成功，例如：

```json
{
  "code": 422,
  "message": "用户已存在"
}
```

### POST http://localhost:9099/user/login

需要提交表单telephone、password，其中telephone要等于11位数，password要大于等于6位字符

该接口返回JSON字符串以表示是否登录成功，如果登录成功，会在JSON字符串中返回token，例如：

```json
{
    "status": 10000,
    "info": "success",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MywiZXhwIjoxNjg2NTcyMDE0fQ.sYOJsETtwRv7RdMhVlfcY6TMydvl_AdCgOtkjNF8c2I"
}
```

### POST http://localhost:9099/chess/create

需要提交表单board_name

需要在请求头中用Authorization携带token

该接口用以创建房间，创建者会自动加入房间，并使用白棋

返回JSON字符串，创建成功会在其中带有房间号boardid，例如：

```json
{
  "board_id": 3,
  "message": "create board success"
}
```

### POST http://localhost:9099/chess/join

需要提交表单boardid，用以加入指定房间

需要在请求头中用Authorization携带token

第一个加入者会使用黑棋，之后的加入者会进行观战。

返回JSON字符串，加入成功会在其中带有房间名board_name，例如：

```json
{
  "board_name": "aliang's room2",
  "message": "join board success"
}
```

### ws://localhost:9099/chess/board/:boardid

boardid参数为房间号，用以在指定房间中进行下棋和观战

需要在请求头中用Authorization携带token

断线后可以通过再次使用此链接建立连接以重连

观展者不需要传输任何内容

下棋的双方需要用JSON字符串来与服务器进行通讯，首先需要传输

```json
{
"Ready": true
}
```

的数据，用以表示准备完成，同样可以传输

```json
{
"Ready": false
}
```

取消准备。当双方同时准备完成时，游戏开始，此时可以按国际象棋规则传输"From"和“To”用以表示移动棋子，例如：

```json
{
"From": "a2",
"To": "a4"
}
```

服务器会返回更新后的表示棋盘状态的JSON字符串，解析此字符串可以获得棋盘的信息。

使用Base64解码其中的“State”字段以获取当前棋盘。

如果用户的操作违反规则，那么此时的棋盘状态不会发生变化。下面是一个返回的JSON字符串的例子：

```json
{
    "ID": 2,
    "CreatedAt": "2023-06-11T18:34:56.221+08:00",
    "UpdatedAt": "2023-06-11T20:16:42.827+08:00",
    "DeletedAt": null,
    "Name": "aliang's room2",
    "WhiteID": 1,
    "BlackID": 2,
    "WhiteReady": false,
    "BlackReady": false,
    "Playing": true,
    "AudienceID": "{\"3\":true}",
    "State": "W1siIiwiUk9PSyIsIkJJU0hPUCIsIlFVRUVOIiwiS0lORyIsIkJJU0hPUCIsIktOSUdIVCIsIlJPT0siXSxbIiIsIlBBV04iLCJQQVdOIiwiUEFXTiIsIlBBV04iLCIiLCJQQVdOIiwiUEFXTiJdLFsiUEFXTiIsIiIsIiIsIiIsIiIsIlBBV04iLCIiLCIiXSxbIiIsIiIsIiIsIiIsIiIsIiIsIiIsIiJdLFsicGF3biIsIiIsInBhd24iLCIiLCIiLCJwYXduIiwiIiwiIl0sWyIiLCIiLCIiLCIiLCIiLCIiLCIiLCIiXSxbIiIsInBhd24iLCIiLCJwYXduIiwicGF3biIsIiIsInBhd24iLCJwYXduIl0sWyJyb29rIiwia25pZ2h0IiwiYmlzaG9wIiwicXVlZW4iLCJraW5nIiwiYmlzaG9wIiwia25pZ2h0Iiwicm9vayJdXQ==",
    "Turn": true,
    "Winner": -1
}
```

- " Turn"字段表示轮次，true表示轮到白棋，false表示轮到黑棋。

- "Winner" 字段表示胜利者，1表示白棋胜利，0表示黑棋胜利，-1表示未决出胜负。

- "State"字段使用Base64解码结果为：
```
[["","ROOK","BISHOP","QUEEN","KING","BISHOP","KNIGHT","ROOK"],["","PAWN","PAWN","PAWN","PAWN","","PAWN","PAWN"],["PAWN","","","","","PAWN","",""],["","","","","","","",""],["pawn","","pawn","","","pawn","",""],["","","","","","","",""],["","pawn","","pawn","pawn","","pawn","pawn"],["rook","knight","bishop","queen","king","bishop","knight","rook"]]
```
其中，大写字母表示白棋，小写字母表示黑棋。

该结果为测试中的棋盘布局，并非最终成品所能得到的布局。

## 加分项

- 观战，建立一对多关系，将观众以JSON字符串储存到board的AudienceID中，同时对建立的连接按房间分组，广播时仅对房间内的连接广播。
- 断线重连，通过board储存WhiteID和BlackID来确认重新连接的用户应该操纵哪个棋子。