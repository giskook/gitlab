# rangelist

# 实现思路说明
* 使用接口方便多种实现，实行了接口和实现的分离
* 目前支持双向链表+hashtable的方式实现rangelist, 双向链表用来遍历，hashtable用来加快查询
* range的判断封装到类型Range中进行

# 工程结构说明
```
├── pkg
│    └─ rangelist // 主目录
│           ├── rangelist.go      // 核心接口
│           ├── rangelist_test.go // 核心接口单元测试
│           ├── range.go          // 核心接口用到的交换结构
│           │
│           ├── base   // 公共基础函数&错误码
│           │    ├── errors.go      // 错误码
│           │    ├── utils.go       // 公共函数
│           │    └── utils_test.go  // 公共函数单元测试文件
│           │
│           └── impl  // 实现 
│               │
│               └── likedlist // 链表实现方式
│                      │
│                      ├── linkedlist.go // 数据结构定义，对外接口实
│                      ├── node.go       // 链表节点定义
│                      ├── add.go        // 链表添加节点封装的内部函数
│                      ├── remove.go     // 链表删除节点封装的内部函数
│                      └── utils.go      // 链表实现使用的业务工具函数
│   
├── cmd
|    └── main.go // 工程入口函数
│   
└── doc 
     └── README.md // it's me

```
