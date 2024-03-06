FORMAT: 1A
HOST: https://polls.apiblueprint.org/

# Andy_Training

Polls is a simple API allowing consumers to view polls and vote in them.

## Questions Collection [/questions]

### 登入 [POST /api/login]

- 使用者可以登入

- Parameters

  - name: "小美" (string, optional) - 使用者帳號
  - password: "12345678" (string, optional) - 使用者密碼

+ Response 200 (application/json)

  登入成功

     - body
     
    {
        "Sessions": 1,
        "User": {
            "Id": 1,
            "Name": "JOHN",
            "Email": "rewwer@gmail.com"
        },
        "message": "Login Successfully"
    }


+ Response 404 (application/json)

  登入失敗

     - body
     
    {
        "message": "帳號密碼輸入錯誤"
    }  


        


### 登出 [GET /api/logout]

- 使用者可以登出


+ Response 200 (application/json)

  成功登出
  
    - body
            {
                "message": "Logout Successfully"
            }
            
+ Response 401 (application/json)

  登出失敗

     - body
           {
                "message": "你還沒有登入喔!"
           }
            
            
            

### 取得當前登入之使用者的分數[GET /api/users/:userid]

- 使用者可以利用該API取得自己的分數

+ Request (application/json)
        
        {
        "User": {
            "Id": 1,
            "Name": "JOHN",
            "Email": "rewwer@gmail.com"
            }
        }
      
+ Response 200 (application/json)

  成功搜尋到分數

  - Body

          {
                  "error": "Scores not found"
          }      
      
        
+ Response 404 (application/json)

  搜尋分數失敗

  - Body

          {
                  "error": "Scores not found"
          }
