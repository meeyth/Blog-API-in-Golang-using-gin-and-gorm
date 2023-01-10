# A Basic Social Media API in Golang with sqlite database support using [GIN](https://github.com/gin-gonic/gin) & [GORM](https://gorm.io/docs/index.html)

## To start just run the command _`go mod download`_ to download all used packages then run _`go run main.go`_ in the root directory

## **Available routes**

| Method   |      Routes      |  Controllers |
|:----------|:-------------|:------|
| POST |  /api/auth/signup | social-media/controllers.Signup |
| POST |  /api/auth/login | social-media/controllers.Login |
| POST |  /api/auth/logut | social-media/controllers.Logout |
| GET |  /api/post/posts | social-media/controllers.GetAllPosts |
| GET |  /api/post/post/:title | social-media/controllers.GetPostByTitle |
| POST |  /api/post/account/post | social-media/controllers.PostAPost |
| PUT |  /api/post/account/post/:title | social-media/controllers.UpdateAPostByTitle |
| DELETE |  /api/post/account/post/:title | social-media/controllers.DeleteAPostByTitle |
| GET |  /api/user/users | social-media/controllers.GetAllUsers  |
| GET |  /api/user/users/:username | social-media/controllers.GetUsersByUsername  |
| GET |  /api/user/users/:username/posts | social-media/controllers.GetAUsersPost  |
| GET |  /api/user/account | social-media/controllers.GetAccountDetails  |
| PUT |  /api/user/account | social-media/controllers.UpdateAccount  |
| DELETE |  /api/user/account | social-media/controllers.DeleteAccount  |
| GET |  /api/user/account/posts | social-media/controllers.CurrentUsersPosts  |

## **Enjoy Coding**
