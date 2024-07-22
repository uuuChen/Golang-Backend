# Golang Backend

## 預先具備

- Docker
- Docker Compose

## 環境準備

```=bash
$ git clone https://github.com/uuuChen/Golang-Backend.git
$ cd Golang-Backend
$ vim .env
```

### `.env` 範例

```=.env
MYSQL_ROOT_PASSWORD=defaultRootPassword
MYSQL_DATABASE=defaultDatabase
MYSQL_USER=defaultUser
MYSQL_PASSWORD=defaultPassword
REDIS_HOST=redis
REDIS_PORT=6379
JWT_SECRET=defaultSecret
```

## 建構、執行

- `$ sh ./build_and_run.sh`

## 功能

### 註冊

- 以 Email、密碼 進行 User 註冊
  - 檢測 Email 格式
  - 檢測密碼格式：密碼不少於六個字、不多於 16 個字、需要有一個大寫、有一個小寫跟一個特殊符號 ()[]{}<>+-*/?,.:;"'_\|~`!@#$%^&=

#### 使用 API

- `POST /v1/user/register`
- 參數:

    ```json5
    {
        "email": "user@example.com",
        "password": "Password123!"
    }
    ```

- Curl：

    ```=bash
    curl -X POST http://localhost:8080/v1/user/register \
    -H "Content-Type: application/json" \
    -d '{"email": "user@example.com", "password": "Password123!"}'
    ```

- 回應：
    |  Code   | 含義  |
    |  ----  | ----  |
    | 201  | 成功註冊 |
    | 400  | 參數格式錯誤 |
    | 409  | Email 已經被註冊過 |
    | 500  | 伺服器錯誤 |

- 當 Code 是 201 時，會回傳
    ```json5
    {
        "message": string,
        "verification_code": string, // verify email 時會用到（模擬在Email 中的驗證碼）
    }
    ```

### 重發驗證 Email

- 重發「驗證 email」
- Rate Limit: 同個 email 一分鐘最多發一次
- 已驗證過的 email 會回傳 409 error code

#### 使用 API

- `POST /v1/user/send-verification-email`
- 參數:

    ```json5
    {
        "email": "user@example.com",
    }
    ```

- Curl：

    ```=bash
    curl -X POST http://localhost:8080/v1/user/send-verification-email \
    -H "Content-Type: application/json" \
    -d '{"email": "user@example.com"}'
    ```

- 回應：
    |  Code   | 含義  |
    |  ----  | ----  |
    | 200  | 成功重發驗證信 |
    | 400  | 參數格式錯誤 |
    | 409  | Email 已經驗證過 |
    | 429  | 過多請求，稍後重試 |
    | 500  | 伺服器錯誤 |


### 驗證 Email

- 以 Verification Code 驗證 email

#### 使用 API

- `POST /v1/user/verify-email`
- 參數:

    ```json5
    {
        "email": "user@example.com",
        "code": "111111" // 註冊成功時回覆的驗證碼
    }
    ```

- Curl：

    ```=bash
    curl -X POST http://localhost:8080/v1/user/verify-email \
    -H "Content-Type: application/json" \
    -d '{"email": "user@example.com", "code": "056307"}'
    ```

- 回應：
    |  Code   | 含義  |
    |  ----  | ----  |
    | 200  | 成功驗證 |
    | 400  | 參數格式錯誤 |
    | 401  | 驗證碼錯誤 |
    | 500  | 伺服器錯誤 |


### 登入

- 以 Email、Password 登入
- 成功登入的話，會回傳 JWT Token（期限為一天）

#### 使用 API

- `POST /v1/user/login`
- 參數:

    ```json5
    {
        "email": "user@example.com",
        "password": "Password123!", // plain password
    }
    ```

- Curl：

    ```=bash
    curl -X POST http://localhost:8080/v1/user/login \
    -H "Content-Type: application/json" \
    -d '{"email": "user@example.com", "password": "Password123!"}'
    ```

- 回應：
    |  Code   | 含義  |
    |  ----  | ----  |
    | 200  | 成功登入 |
    | 400  | 參數格式錯誤 |
    | 401  | 密碼錯誤 |
    | 500  | 伺服器錯誤 |

- 當 Code 是 201 時，會回傳

  ```json5
  {
      "token": string // JWT token
  }
  ```

### 推薦資料

- 需要透過 JWT Token 才能使用（成功登入才會獲得，會在一天後過期）
- 獲取推薦商品資料

#### 使用 API

- `POST /v1/recommendation`
- Header:

    ```json5
    {
        "Authorization: Bearer {token}",
    }
    ```

- Curl：

    ```=bash
    curl -X GET http://localhost:8080/v1/recommendation \
    -H "Authorization: Bearer {token}"
    ```

- 回應：
    |  Code   | 含義  |
    |  ----  | ----  |
    | 200  | 成功 |
    | 400  | 參數格式錯誤 |
    | 401  | 密碼錯誤 |
    | 500  | 伺服器錯誤 |

- 當 Code 是 200 時，會回傳

  ```json5
  {
      "recommendations": string // JWT token
  }
  ```

