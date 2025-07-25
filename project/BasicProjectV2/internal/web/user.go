package web

import (
	"fmt"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/service"
	"github.com/basicprojectv2/pkg/jwt"
	"github.com/basicprojectv2/pkg/snowflake"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type UserHandler struct {
	svc     service.UserService
	codeSvc service.CodeService
}

const (
	//emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	//passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	bizLogin = "login"
)

func NewUserHandler(svc service.UserService, codeSvc service.CodeService) *UserHandler {
	return &UserHandler{svc: svc, codeSvc: codeSvc}
}

// 注册路由
func (h *UserHandler) RegisterRoutes(server *gin.Engine, i18n, loginCheck gin.HandlerFunc) {
	ug := server.Group("/v2/users/")
	ug.GET("/hi", h.Hi) //loginCheck, i18n,
	ug.POST("/signin", h.SignIn)
	ug.POST("/login", h.Login)
	ug.POST("/loginsms/code/send", h.SendSMS)
	ug.POST("/loginsms", h.VerifySMS)
	ug.POST("/userList", h.HandleUserList)
	ug.POST("/updateUser", h.updateUser)
	//用户上传图片
	ug.POST("/uploadAvatar", loginCheck, h.HandleUploadAvatar)
	ug.POST("/uploadFile", loginCheck, h.HandleUploadFile)
	ug.POST("/profile", loginCheck, h.HandleUerProfile)

	// 用户下载文件
	ug.POST("/downloadFile", loginCheck, h.HandlerUserDownloadFile)

	// 分片上传
	ug.POST("/upload_chunk", loginCheck, h.HandleUploadChunk)
	ug.POST("/merge_chunk", loginCheck, h.HandleMergeChunk)
}

func (h *UserHandler) updateUser(ctx *gin.Context) {
	var form domain.User
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "bad request"})
	}

	//交给service层
	if err := h.svc.UpdateUser(ctx, form); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "系统错误"})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

// 处理获取用户列表
func (h *UserHandler) HandleUserList(ctx *gin.Context) {
	// 验证请求参数是否正确
	var form domain.UserListRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "bad request"})
	}

	// 交给service层
	data, err := h.svc.GetUserList(ctx, form)
	if err != nil {
		log.Println("user service get user list error", err)
	}
	log.Println(data)
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

// 处理用户上传图片
func (h *UserHandler) HandleUploadAvatar(ctx *gin.Context) {
	userid, exists := ctx.Get("userid")
	if !exists {
		ctx.JSON(400, gin.H{
			"msg": "用户未登录",
		})
		return
	}
	strUserid := userid.(string)
	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "bad request"})
		return
	}
	// 打开文件
	openedFile, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to open file"})
		return
	}
	defer openedFile.Close()

	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to read file"})
		return
	}
	req := domain.UserAvatar{UserID: strUserid, AvatarFile: fileBytes}

	if err := h.svc.UploadUserAvatar(ctx, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Upload Avatar ERROR"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "upload successful",
	})

}

func (h *UserHandler) HandlerUserDownloadFile(ctx *gin.Context) {
	userid, exists := ctx.Get("userid")
	if !exists {
		ctx.JSON(400, gin.H{
			"msg": "用户未登录",
		})
		return
	}
	fileName := ctx.PostForm("fileName")

	// 获取文件
	req := domain.DownloadFileReq{
		UserID:   userid.(string),
		FileName: fileName,
	}
	file, err := h.svc.GetUserFile(ctx, req)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Download File Error",
		})
		return
	}

	//ctx.Header("Content-Type", "image/png")
	//ctx.Data(http.StatusOK, "image/jpeg", file.File)

	ctx.JSON(http.StatusOK, gin.H{
		"pic": file.Base64,
	})
}

// 处理上传分片的接口
func (h *UserHandler) HandleUploadChunk(ctx *gin.Context) {
	userid, exists := ctx.Get("userid")
	if !exists {
		ctx.JSON(401, gin.H{
			"msg": "用户未登录",
		})
		return
	}
	log.Println(userid)
	// todo  生成分片的临时文件夹
	const chunkDir = "C:\\Users\\hzf19\\Desktop\\chunk"

	// todo 生成文件存储路径
	const finalDir = "C:\\Users\\hzf19\\Desktop\\1821841651400708096"

	// 检查分片存储目录是否存在，不存在则创建
	if _, err := os.Stat(chunkDir); os.IsNotExist(err) {
		err := os.MkdirAll(chunkDir, os.ModePerm) // 创建目录
		if err != nil {
			log.Println("创建分片目录失败:", err)
			ctx.JSON(500, gin.H{"error": "创建分片目录失败"})
			return
		}
	}

	file, err := ctx.FormFile("chunk")
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": "文件上传失败"})
		return
	}

	// 获取文件的其他信息
	// 获取文件名，文件名从表单字段获取
	fileName := ctx.DefaultPostForm("fileName", "") // 获取文件名
	if fileName == "" {
		fileName = file.Filename // 如果没有从表单中获取到文件名，则使用上传的文件名
	}
	index, err := strconv.Atoi(ctx.DefaultPostForm("index", "0")) // 分片索引
	log.Println(index)
	if err != nil {
		log.Println(err, "无效的分片索引")
		ctx.JSON(400, gin.H{"error": "无效的分片索引"})
		return
	}
	chunkCount, err := strconv.Atoi(ctx.DefaultPostForm("chunkCount", "0")) // 总分片数
	if err != nil {
		ctx.JSON(400, gin.H{"error": "无效的总分片数"})
		return
	}

	// 创建存储分片的临时文件
	chunkFilePath := filepath.Join(chunkDir, fmt.Sprintf("%s-chunk-%d", fileName, index))
	// 保存分片到临时文件
	if err := ctx.SaveUploadedFile(file, chunkFilePath); err != nil {
		ctx.JSON(500, gin.H{"error": "保存分片文件失败"})
		return
	}

	// 返回上传成功的响应
	ctx.JSON(200, gin.H{
		"message":    "分片上传成功",
		"chunkIndex": index,
		"chunkCount": chunkCount,
	})

}

// 合并分片的接口
func (h *UserHandler) HandleMergeChunk(ctx *gin.Context) {
	userid, exists := ctx.Get("userid")
	if !exists {
		ctx.JSON(400, gin.H{
			"msg": "用户未登录",
		})
		return
	}
	strUserid := userid.(string)
	log.Println(strUserid)

	fileName := ctx.PostForm("fileName")
	totalChunk, err := strconv.Atoi(ctx.PostForm("totalChunk"))
	if err != nil || totalChunk <= 0 {
		ctx.JSON(400, gin.H{"error": "无效的总分片数"})
		return
	}

	// 分片的临时文件夹
	const chunkDir = "C:\\Users\\hzf19\\Desktop\\chunk"

	// 文件存储路径
	mergePath := fmt.Sprintf("C:\\Users\\hzf19\\Desktop\\1821841651400708096\\%s", fileName)

	// 确保临时文件夹存在
	if _, err := os.Stat(chunkDir); os.IsNotExist(err) {
		ctx.JSON(500, gin.H{"error": "临时文件不存在"})
		return
	}

	// 创建合并文件
	mergeFile, err := os.Create(mergePath)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "创建合并文件失败"})
		return
	}
	defer mergeFile.Close()

	// 按索引顺序合并分片 顺序很重要
	for i := 0; i < totalChunk; i++ {
		chunkPath := filepath.Join(chunkDir, fmt.Sprintf("%s-chunk-%d", fileName, i))
		chunkData, err := os.ReadFile(chunkPath)
		if err != nil {
			log.Printf("读取分片失败: %v", err)
			ctx.JSON(500, gin.H{"error": fmt.Sprintf("读取分片 %d 失败: %v", i, err)})
			return
		}
		if _, err := mergeFile.Write(chunkData); err != nil {
			log.Printf("写入分片失败: %v", err)
			ctx.JSON(500, gin.H{"error": fmt.Sprintf("写入分片 %d 失败: %v", i, err)})
			return
		}
	}

	ctx.JSON(200, gin.H{
		"data": "merge chunk success",
	})
}

// 用户获取自己信息
func (h *UserHandler) HandleUerProfile(ctx *gin.Context) {
	userid, exists := ctx.Get("userid")
	if !exists {
		ctx.JSON(400, gin.H{
			"msg": "用户未登录",
		})
		return
	}
	strUserid := userid.(string)
	log.Println(strUserid)
}

func (h *UserHandler) HandleUploadFile(ctx *gin.Context) {
	userid, exists := ctx.Get("userid")
	if !exists {
		ctx.JSON(400, gin.H{
			"msg": "用户未登录",
		})
		return
	}
	strUserid := userid.(string)
	fileName := ctx.PostForm("fileName")

	//获取文件
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Println("upload file error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Failed to get file"})
		return
	}

	// 获取文件类型
	fileType := file.Header.Get("Content-Type")
	log.Println(fileType)

	// 检查文件大小
	if file.Size > 200*1024*1024 {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "file size too big"})
		return
	}

	// 打开文件
	openedFile, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to open file"})
		return
	}
	defer openedFile.Close()

	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to read file"})
		return
	}

	req := domain.UploadFileReq{UserID: strUserid, FileName: fileName, File: fileBytes, FileType: fileType}

	if err := h.svc.UploadUserFile(ctx, req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Upload File ERROR"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "upload file successful",
	})

}

func (h *UserHandler) Hi(ctx *gin.Context) {
	//localizer, _ := ctx.Get("localizer")
	//localizeConfig := &i18n.LocalizeConfig{MessageID: "welcome_msg"}
	//welcomeMsg, err := localizer.(*i18n.Localizer).Localize(localizeConfig)
	//if err != nil {
	//	log.Println("localize welcome error", err)
	//	welcomeMsg = "Hi 这里是BasicProjectV2 单体~ "
	//}
	welcomeMsg := domain.HiResponse{Msg: "Hi 这里是BasicProjectV2 单体~ "}
	ctx.JSON(200, welcomeMsg)
}

// 处理注册请求
func (h *UserHandler) SignIn(ctx *gin.Context) {

	// 验证请求参数是否正确
	var form domain.SignInRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
	}
	// 检验两次密码输入是否正确
	if form.Password != form.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "两次密码输入不一致",
		})
	}

	err := h.svc.Signup(ctx, domain.User{Email: form.Email, Password: form.Password})
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, "msg:注册成功")
	default:
		ctx.JSON(http.StatusNotFound, "msg:注册失败")
	}

}

// 处理登录请求
func (h *UserHandler) Login(ctx *gin.Context) {
	var form domain.LoginRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, "参数错误")
	}
	u, err := h.svc.Login(ctx, form.Email, form.Password)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "msg: 登录失败")
		return
	}
	// 登录成功
	log.Println(u)
	token, err := jwt.GenToken(u.ID)
	if err != nil {
		// 生成token失败
		ctx.JSON(http.StatusInternalServerError, "系统错误")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "登录成功",
		"token": token,
	})
}

// 处理sms请求
func (h *UserHandler) SendSMS(ctx *gin.Context) {
	var form domain.SMSRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求参数错误",
		})
		return
	}

	if form.Phone == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号为空",
		})
		return
	}

	// 开始交给code svc
	if err := h.codeSvc.SendCode(ctx, bizLogin, form.Phone); err != nil {
		log.Println("发送短信失败", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})

}

// 处理验证sms登录
func (h *UserHandler) VerifySMS(ctx *gin.Context) {
	var form domain.SMSLogin
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求参数错误",
		})
		// 可以交给code svc了
	}
	ok, err := h.codeSvc.VerifyCode(ctx, bizLogin, form.Phone, form.Code)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "系统错误",
		})
	}
	if ok != true {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "验证失败",
		})
	} else {
		/*
			数据库中查询是否有该手机号，如果没有的话帮用户注册
		*/

		// 返回一个token
		id := snowflake.GenId()
		u, err := h.svc.FindOrCreate(ctx, form.Phone, strconv.Itoa(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"msg": "系统错误",
			})
			return
		}
		token, err := jwt.GenToken(u.ID)

		ctx.JSON(http.StatusOK, gin.H{
			"msg":   "登录成功",
			"token": token,
		})
	}

}
