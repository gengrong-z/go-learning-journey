## Day 1 Log

### 🧪 今日成果
1. 使用 interface 定义角色通用行为（Name/Attack/Defend/Status）
2. 实现 Warrior 结构体并绑定方法，自动实现 Character 接口
3. 将接口作为参数（多态调用），如 PrintStatus(Character)
4. 了解组合结构体与初始化函数的实践写法

### ❓ 当前疑问
1. 多个角色如何放入统一的数组中管理？是否用 []Character？
2. 如果后续角色行为各异，该如何优雅设计接口拆分或组合？
3. interface 类型是否有性能影响？

### 🔜 明日预定
1. 实现第二个角色：Mage（与 Warrior 有不同属性与行为）
2. 将多个角色放入切片中统一管理
3. 使用类型断言 / type switch 实现多态分支