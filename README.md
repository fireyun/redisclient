### 实现方式
 - redisclient是对底层redis client的封装，实现接口Client来提供redis常用的一些操作
 - 当前封装的底层redis client是[go-redis](https://github.com/go-redis/redis/tree/v7)，git地址为<https://github.com/go-redis/redis/tree/v7>

### 实现目的
- 对底层redis之上增加一层接口封装，该接口供上层调用，有利于代码的扩展性
- 在底层redis发生替换或变化时，无需对接口协议协议进行变更，只需改变接口方法内的具体实现，增强了代码的可维护性


