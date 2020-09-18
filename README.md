### 实现方式
 - redisclient是对底层redis client的封装，提供了一个通用的接口RedisClient，提供redis常用的一些操作
 - 当前封装的底层redis client是[go-redis](https://github.com/go-redis/redis/tree/v7)，git地址为<https://github.com/go-redis/redis/tree/v7>

### 实现目的
- 在底层redis发生替换或变化时，能够保持接口方法对外保持不变，只需改变方法内的实现，增强代码的可维护性


