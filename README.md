# gosignaler

## 项目介绍
go语言版的CDNBye signal服务器

## TODO
准备开源，所以需要将代码优化下

## 信令协议

### peer ---> server 建立websocket时的查询参数，如：wss://signal.cdnbye.com/wss?id=2f99625ae34a04a1
{
    id： string         //客户端注册信令的时候提交的ID
}

### peer ---> server    发送信令给server
```javastript
{
    action: 'signal'         		
    to_peer_id: string            //对等端的Id，即要向哪个peer发送信令
    data: string                  //需要传送的数据
}
```

### server ---> peer        如果目标节点存在则发送信令给目标节点
```javastript
{
    action: 'signal'           		
    from_peer_id: string          //对等端的Id，即从哪个peer发送过来的信令
    data: string                  //需要传送的数据
```

### server ---> peer        如果目标节点不存在则返回该信息
```javastript
{
    action: 'signal'           		
    from_peer_id: string          //对等端的Id，即从哪个peer发送过来的信令
}
```