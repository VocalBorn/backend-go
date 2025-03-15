# 後端 Backend Go 語言
## 資料夾架構
models: 定義資料庫模型 <br>
routes: 定義路由 <br>
handlers : 定義路由處理邏輯 <br>
services: 定義服務邏輯 <br>
## 如何執行
```bash
go run cmd/main.go
```
## 如何讓Swagger自動生成或更新API文件
```bash
swag init -g cmd/main.go
```
## 參考資料
[資料結構參考](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091)
[Golang 教學全集](https://vocus.cc/salon/658d1061fd89780001dec7e8/room/golanglab)
[資料庫 Transaction](https://oldmo860617.medium.com/database-transaction-acid-156a3b75845e)