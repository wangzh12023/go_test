
package main

import (
    "fmt"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

func main() {
    // 创建一个 gorm.DB 类型的变量
    var db *gorm.DB
    // 调用 Open 方法，传入驱动名和连接字符串
    db, err := gorm.Open(mysql.Open("root:2wsx$RFV246@tcp(127.0.0.1:3306)/test"), &gorm.Config{})
    // 检查是否有错误
    if err != nil {
        fmt.Println("连接数据库失败：", err)
        return
    }
    // 打印成功信息
    fmt.Println("连接数据库成功")
    // 关闭数据库连接

    
    defer func() {
        dbInstance, _ := db.DB()

        _ = dbInstance.Close()
    }()
}
