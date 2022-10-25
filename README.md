## LineBot use gin & mongodb
package
1. go get -u github.com/gin-gonic/gin
2. go get go.mongodb.org/mongo-driver/mongo
3. go get -u github.com/line/line-bot-sdk-go/v7/linebot
4. go get github.com/spf13/viper


TestImg

1. Insert User and Message to MongoDB
  ![insertjpg](https://user-images.githubusercontent.com/89484381/197419383-b8041b0d-624b-4137-b06d-b2fc8f05dc40.jpg)
  
2. Get User & UserMessage (Postman)
  ![getuser2](https://user-images.githubusercontent.com/89484381/197776741-2b3ef6f7-4f4a-4420-b25d-26b9af930069.jpg)




3. main.exe request & response
![cmd](https://user-images.githubusercontent.com/89484381/197419428-415e4711-c16a-4455-8889-6b5a050489cc.jpg)

待處理項目
1. Mes struct重新設計

v.1.1
1. 修改GET API格式更符合RestFul設計模式
2. 時間現在能正確顯示
3. app.env環境變數設置完成
