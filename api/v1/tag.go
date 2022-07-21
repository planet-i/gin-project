package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/planet-i/gin-project/pkg/app"
	"github.com/planet-i/gin-project/pkg/e"
	"github.com/planet-i/gin-project/pkg/export"
	"github.com/planet-i/gin-project/pkg/setting"
	"github.com/planet-i/gin-project/pkg/util"
	"github.com/planet-i/gin-project/service/tag_service"
	"github.com/unknwon/com"
)

// GetTags        godoc
// @Summary       获取标签列表
// @Produce       json
// @Param         name query string true "Name"
// @Param         state query int false "State"
// @Success       200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router        /api/v1/tags [get]
func GetTags(c *gin.Context) { // *gin.Context允许我们在中间件之间传递变量、管理流、验证请求的 JSON 和呈现 JSON 响应
	// c.Query可用于获取?name=test&state=1这类 URL 参数
	appG := app.Gin{C: c}
	name := c.Query("name")
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagService := tag_service.Tag{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	tags, err := tagService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}

	count, err := tagService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": tags,
		"total": count,
	})
}

type AddTagForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

// AddTag         godoc
// @Summary       添加标签
// @Produce       json
// @Param         name query string true "Name"
// @Param         state query int false "State"
// @Param         created_by query int false "CreatedBy"
// @Success       200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router        /api/v1/tags [post]
func AddTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddTagForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := tag_service.Tag{
		Name:      form.Name,
		CreatedBy: form.CreatedBy,
		State:     form.State,
	}
	exists, err := tagService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}

	err = tagService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditTagForm struct {
	ID         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

// EditTag       godoc
// @Summary      修改指定标签
// @Produce      json
// @Param        id path int true "ID"
// @Param        name query string true "Name"
// @Param        state query int false "State"
// @Param        modified_by query string true "ModifiedBy"
// @Success      200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router       /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	appG := app.Gin{C: c}
	form := EditTagForm{ID: com.StrTo(c.Param("id")).MustInt()}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := tag_service.Tag{
		ID:         form.ID,
		Name:       form.Name,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}

	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = tagService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// DeleteTag     godoc
// @Summary      删除指定标签
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router       /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	tagService := tag_service.Tag{ID: id}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	if err := tagService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// ExportTag     godoc
// @Summary      导出标签到excel文件
// @Produce      json
// @Param        name body string false "Name"
// @Param        state body int false "State"
// @Success      200 {object} app.Response
// @Failure      500 {object} app.Response
// @Router       /api/v1/tags/export [post]
func ExportTag(c *gin.Context) {
	appG := app.Gin{C: c}

	name := c.PostForm("name")
	state := -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagService := tag_service.Tag{
		Name:  name,
		State: state,
	}

	filename, err := tagService.Export()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_EXPORT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"export_url":      export.GetExcelFullUrl(filename),
		"export_save_url": export.GetExcelPath() + filename,
	})
}

// ImportTag     godoc
// @Summary      从excel文件中导入标签
// @Produce      json
// @Param        file body file true "Excel File"
// @Success      200 {object} app.Response
// @Failure      500 {object} app.Response
// @Router       /api/v1/tags/import [post]
func ImportTag(c *gin.Context) {
	appG := app.Gin{C: c}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	tagService := tag_service.Tag{}
	err = tagService.Import(file)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_IMPORT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
