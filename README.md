<h1 align='center'>xcache</h1>
<div align=center><img src="https://github.com/sugtex/xcache/blob/main/rideGo.jpg"/></div>
<h2 align='center'>A Cache For Go</h2>

## 📖 简介

`xcahe`是一个提供单机缓存池的第三方库，淘汰机制采用LRU(Least recently used,最近最少使用)。

## ⚠️ 注意

- 单机缓存适用场景较为局限，必须可容忍数据不一致。
- 虽然提供监视者`monitor`功能可定时消除脏数据，但定时器间隔期间读取同样会产生误差，即遵循上条规则。

## 🚀 功能

- 提供默认单机缓存池，最大容量为1000，监控间隔为10s。
- 提供自定义缓存容量，到达容量则使用LRU算法淘汰。
- 容量到达最大值执行淘汰机制（LRU）。
- 提供监视者`monitor`开关及自定义间隔，监视者在定时器间隔`gap`会执行用户提供的数据源方法，返回值为nil时删除单机缓存中脏数据。
- 用户打开监视者`monitor`后可自主向监视者提供"犯人"`prisoner`或者移除"犯人"`prisoner`，此过程是并发安全的。

## 🧰 安装
``` powershell
go get -u github.com/sugtex/xcache
```

## 🛠 使用

### 添加缓存键值对
``` 
xcache.Add("a",[]byte("x"))
```

### 实现获取数据源方法
``` 
type GetF func(context.Context) ([]byte, error)

func getData(ctx context.Context)(reply []byte,err error){
	// 获取数据源逻辑
	return
}
``` 

### 获取缓存
参数：ctx(上下文)，键(key)，getF(获取数据源方法)
``` 
reply,err:=xcache.Get(context.Background(),"a",getData)
if err!=nil{
	// 处理异常
	return
}
```

### 删除缓存
```
xcache.Del("a")
```

### 构建"犯人"
"犯人"由key(键)和getF(获取数据源方法)构成
```
xcache.WithPrisoner("a",getData)
```

### 添加"犯人"[支持不定参数]
```
xcache.AddPrisoner(xcache.WithPrisoner("a",getData))
```

### 删除"犯人"[支持不定参数]
```
xcache.RemovePrisoner(xcache.WithPrisoner("a",getData))
```

### 自定义池
```
cache:=xcache.NewXCache(10)
```

### 打开监视者`monitor`[可选]
```
cache:=xcache.NewXCache(10,xcache.WithOpenMonitor(20))
```

### 重置监视者`monitor`监视间隔
```
if ok := cache.ResetMonitor(30); ok {
	// 处理参数错误逻辑
	return
}
```

## 📚 附言

- LRU算法"增删改查"时间复杂度均为O(1)，算法来源`LeetCode`146题或16.25题，进阶。
- 146题：`https://leetcode-cn.com/problems/lru-cache/`
- 16.25题：`https://leetcode-cn.com/problems/lru-cache-lcci/`
