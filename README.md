## LineBot use gin & mongodb
package
1. go get -u github.com/gin-gonic/gin
2. go get go.mongodb.org/mongo-driver/mongo
3. go get -u github.com/line/line-bot-sdk-go/v7/linebot

TestImg

1. Insert User and Message to MongoDB
  ![insertjpg](https://user-images.githubusercontent.com/89484381/197419383-b8041b0d-624b-4137-b06d-b2fc8f05dc40.jpg)
  
2. Get User & UserMessage (Postman)
  ![getuser](https://user-images.githubusercontent.com/89484381/197419420-56b59772-550f-4e24-b293-b8b5bcad0291.jpg)


3. main.exe request & response
![cmd](https://user-images.githubusercontent.com/89484381/197419428-415e4711-c16a-4455-8889-6b5a050489cc.jpg)

待處理項目
1. 環境變數設置
2. Mes struct重新設計
3. 時間格式跑掉
