# Golang 为什么不能直接将任意类型数组赋值给 []interface{}

> 本篇参考: https://github.com/golang/go/wiki/InterfaceSlice

当我在写 go 代码时，想用 []interface{} 类型来接受一个任意类型的数组，从而进行泛型操作时，发现直接赋值会发生错误，导致 panic。

```go
	dataSlice := []int{1, 2, 3, 4, 5, 6}
	var interfaceSlice []interface{} = dataSlice
```

> Cannot use 'dataSlice' (type []int) as type []interface{}

IDE 会报以上错误。

原本我以为 interface{} 可以接收任何类型，[]interface{} 数组就应该可以接收任何类型的数组才对。

## Why?

先说结论：**[]interface{} 不是 interface{}，它是一个元素类型为 interface{} 的切片。**

主要是因为 []interface{} 有特殊的内存布局。先看一下 interface{} 的内存布局：

![image-20210428153736553](/Users/guowenfeng/Library/Application Support/typora-user-images/image-20210428153736553.png)

如果是 []interface{} 就是 N 个上面的布局。可是普通类型的切片内存布局是：

![image-20210428155048225](/Users/guowenfeng/Library/Application Support/typora-user-images/image-20210428155048225.png)

所以，根本原因就是两个东西的内存布局不同，所以无法直接赋值。

## What can I do instead?

两种情况：

1. 你只是需要用一个容器去接收任意类型的切片，然后在进行索引之前就会将它转成原来的类型你只需要：

   ```go
   var tmp interface{} = dataArr
   ```

2. 你的确需要将任意类型的数组转成 []interface{} ，然后去索引 []interface{} ，你只能遍历赋值：

   ```go
   for index, data := range dataArr {
     interfaceSlice = append(interfaceSlice, data)
     // 这里可以直接在声明时就指定 []interface{} 的长度
     // 用 interfaceSlice[index] = data 来代替会减少系统创建内存时消耗的时间
   }
   ```

