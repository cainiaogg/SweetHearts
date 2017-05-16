### 登陆
/login
@param :json
'username' : 'admin'
'password' : 'admin'

@return :json
ErrorNumber=0, 失败; 1, 成功
ErrorString  登陆信息

#获取朋友圈发布id
url = http://118.89.200.173/circle/get_id
@method :Get
@param :null
@return :string (唯一id)

#上传朋友圈图片
url = http://118.89.200.173/circle//circle/upload_picture
@method :Post
@param : json
{"username": "jiba", "id": "1"}
@image: 文件名"image"

#上传朋友圈内容
url = http://118.89.200.173/circle/add_friend_circle
@method :Post
@param :
{"username": "jiba", "id": "1", "content" : "jibada", "ishoney", "0"}  0表示不是 honey
@return : "1" or null

#获取朋友圈列表
url = http://118.89.200.173/circle/get_friend_circle
@method :Get
@param :"username"  "page"


#获得头像
url = http://118.89.200.173/image/head_image/jiba.jpg

#获取聊天记录
url = http://192.168.253.18:8080/chat/get_chat_record?username=admin&friendname=adminfriend&timestamp=123123123
@method :Get
@return :
[
  [
    1,
    {
      "Id": 1,
      "From": "admin",
      "To": "adminfriend",
      "Content": "JIba",
      "TimeStamp": 1494902803
    }
  ],
  [
    1,
    {
      "Id": 2,
      "From": "admin",
      "To": "adminfriend",
      "Content": "dajiba",
      "TimeStamp": 1494902803
    }
  ],
  [
    1,
    {
      "Id": 3,
      "From": "admin",
      "To": "adminfriend",
      "Content": "WebSocket rocks",
      "TimeStamp": 1
    }
  ]
]
