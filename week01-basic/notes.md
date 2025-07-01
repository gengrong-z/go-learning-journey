# Week 01 Log

---

## Day 1 Log

🎯 初始化项目结构 + 热身基础语法 ✅  
🧠 收获：
- 熟悉 GoLand 工程流程（go mod / go run）
- 回忆变量声明、函数、for、map、slice 基本用法  
  ❓ 不懂：无，均已掌握

---

## Day 2 Log

🎯 add/list 命令实现 ✅  
🧠 收获：
- struct 定义 + 方法调用熟练
- 切片 append 机制回忆完成
- string.Join 组合输入字符串很实用

❓ 疑问：
- todo_cli目录下的两个文件的package不同,编译会报错
  - **不同的package必须放在不同目录**

## Day 3 Log

🎯 实现 done <编号> 命令 ✅  
🧠 收获：
- strconv.Atoi() 可快速解析数字输入
- 通过 *[]Item 修改原始 slice 中 struct 内容
- 输入校验防止程序崩溃（if err != nil）  
 
❓ 疑问：
- 为什么修改要使用指针`*`,而传入需要使用地址引用`&`
  - 如此理解:`&`是取地址,`*`是访问地址
- 为什么`*todos`需要写成`(*todos)`的形式
  - 因为Go不允许优先级不清的指针引用+索引组合

## Day 4 Log

🎯 实现 delete 命令 + 自动文件保存机制 ✅  
🧠 收获：
- 完成任务删除逻辑的实现
- 在 `done` 和 `delete` 操作后自动调用 SaveList() 提高数据安全
- 程序结构进一步简洁，逻辑清晰，功能完整

❓ 疑问整理
1. Go 的静态常量怎么定义 ✅
    - 使用 `const` 定义不可变基础类型（string/int/bool）
    - 如：`const FileName = "todo.json"`
    - 复杂结构（如 map、slice）用 `var` 或 `init()` 初始化

2. 文件读写路径如何把握 ⏳（明日课题）
    - 相对 vs 绝对路径的选择
    - 动态路径构建（如使用 `os.UserHomeDir()`）
    - 使用 `filepath.Join()` 实现跨平台路径拼接

3. IO 流如何更好地掌握 ⏳（明日课题）
    - 理解 `io.Reader` / `io.Writer` 接口
    - 熟练使用 `os.ReadFile` / `bufio.Scanner` / `io.Copy`
    - 实战练习建议（如日志分析器、文件合并器等）

## Day 5 Log

🎯 完善错误处理 + 添加基础测试 ✅  
🧠 收获：
- 将时间类型从 string 改为 time.Time，格式化输出更灵活
- 实现对核心函数的测试（New, MarkDone, Delete）
- 知道了如何测试标准输出：可以用 io.Writer 或临时接管 os.Stdout
  
❓ 疑问：
- PrintAll() 改为可测试设计更通用