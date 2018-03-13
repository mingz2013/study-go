# gate


- gate 需要至少提供给客户端tcp/websocket 两种连接方式

- gate 要连接到agent，代理服务器

---

- client发送userId连接

- 接收client连接
- 验证client是否登陆，如果没有session，返回错误，断开连接

- 保持连接socket，返回连接成功

- 转发消息

- 