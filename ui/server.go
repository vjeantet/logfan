package ui

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/k0kubun/pp"
	uuid "github.com/nu7hatch/gouuid"

	"github.com/vjeantet/bitfan/api"
	eztemplate "github.com/vjeantet/ez-gin-template"
)

var db *gorm.DB
var apiClient *api.RestClient

func Handler(assetsPath, path string, dbpath string, apiBaseUrl string) http.Handler {
	var err error
	apiClient = api.NewRestClient(apiBaseUrl)
	db, err = gorm.Open("sqlite3", filepath.Join(dbpath, "bitfan.db"))
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Pipeline{}, &Asset{})

	r := gin.New()
	render := eztemplate.New()
	render.TemplatesDir = assetsPath + "/views/" // default
	render.Ext = ".html"                         // default
	render.Debug = true                          // default
	render.TemplateFuncMap = template.FuncMap{
		"dateFormat": (*templateFunctions)(nil).dateFormat,
		"ago":        (*templateFunctions)(nil).dateAgo,
		"string":     (*templateFunctions)(nil).toString,
		"b64":        (*templateFunctions)(nil).toBase64,
		"int":        (*templateFunctions)(nil).toInt,
		"time":       (*templateFunctions)(nil).asTime,
		"now":        (*templateFunctions)(nil).now,
		"isset":      (*templateFunctions)(nil).isSet,

		"numFmt": (*templateFunctions)(nil).numFmt,

		"safeHTML":     (*templateFunctions)(nil).safeHtml,
		"hTMLUnescape": (*templateFunctions)(nil).htmlUnescape,
		"hTMLEscape":   (*templateFunctions)(nil).htmlEscape,
		"lower":        (*templateFunctions)(nil).lower,
		"upper":        (*templateFunctions)(nil).upper,
		"trim":         (*templateFunctions)(nil).trim,
		"trimPrefix":   (*templateFunctions)(nil).trimPrefix,
		"hasPrefix":    (*templateFunctions)(nil).hasPrefix,
		"replace":      (*templateFunctions)(nil).replace,
		"markdown":     (*templateFunctions)(nil).toMarkdown,
	}

	r.HTMLRender = render.Init()
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store), gin.Recovery())
	// r.Use(gin.Recovery())
	g := r.Group(path)
	{
		g.StaticFS("/public", http.Dir(assetsPath+"/public"))
		g.GET("/", getPipelines)

		// list pipelines
		g.GET("/pipelines", getPipelines)
		// New pipeline
		g.GET("/pipelines/:id/new", newPipeline)
		// Create pipeline
		g.POST("/pipelines", createPipeline)

		// Show pipeline
		g.GET("/pipelines/:id", editPipeline)
		// Save pipeline
		g.POST("/pipelines/:id", updatePipeline)

		// Show asset
		g.GET("/pipelines/:id/assets/:assetID", showAsset)
		// Create asset
		g.POST("/pipelines/:id/assets", createAsset)
		// Update asset
		g.POST("/pipelines/:id/assets/:assetID", updateAsset)
		// Replace asset
		g.PUT("/pipelines/:id/assets/:assetID", replaceAsset)
		// Download asset
		g.GET("/pipelines/:id/assets/:assetID/download", downloadAsset)
		// Delete asset
		g.GET("/pipelines/:id/assets/:assetID/delete", deleteAsset)

		// Start pipeline
		g.GET("/pipelines/:id/start", startPipeline)
		// Stop pipeline
		g.GET("/pipelines/:id/stop", stopPipeline)
		// Restart pipeline
		g.GET("/pipelines/:id/restart", restartPipeline)
	}

	return r
}

func startPipeline(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Abort()
		return
	}

	//Récupére le pipeline en base
	var p = Pipeline{ID: (id)}
	db.Preload("Assets").First(&p)

	//Construit l'objet ApiClientPipeline
	appl := api.Pipeline{}
	appl.Uuid = p.Uuid
	appl.Label = p.Label
	appl.Assets = []api.Asset{}

	for _, a := range p.Assets {
		appl.Assets = append(appl.Assets, api.Asset{
			Path:    a.Name,
			Content: base64.StdEncoding.EncodeToString(a.Value),
		})
		if a.Type == "entrypoint" {
			appl.ConfigLocation = a.Name
		}
	}
	//Envoie

	pipeline, err := apiClient.AddPipeline(&appl)
	if err != nil {
		c.JSON(500, err.Error())
		log.Printf("error : %v\n", err)
	} else {
		c.JSON(200, appl)
		log.Printf("Started (UUID:%s) - %s\n", pipeline.Uuid, pipeline.Label)
	}

	//OK

	// c.String(200, "pipelines/new")
}
func restartPipeline(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Abort()
		return
	}

	var p = Pipeline{ID: (id)}
	db.Preload("Assets").First(&p)

	err = apiClient.StopPipeline(p.Uuid)
	if err != nil {
		c.JSON(500, err.Error())
	}

	//Construit l'objet ApiClientPipeline
	appl := api.Pipeline{}
	appl.Uuid = p.Uuid
	appl.Label = p.Label
	appl.Assets = []api.Asset{}

	for _, a := range p.Assets {
		appl.Assets = append(appl.Assets, api.Asset{
			Path:    a.Name,
			Content: base64.StdEncoding.EncodeToString(a.Value),
		})
		if a.Type == "entrypoint" {
			appl.ConfigLocation = a.Name
		}
	}
	//Envoie

	pipeline, err := apiClient.AddPipeline(&appl)
	if err != nil {
		c.JSON(500, err.Error())
		log.Printf("error : %v\n", err)
	} else {
		c.JSON(200, appl)
		log.Printf("Restarted (UUID:%s) - %s\n", pipeline.Uuid, pipeline.Label)
	}
}

func stopPipeline(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Abort()
		return
	}

	var p = Pipeline{ID: (id)}
	db.First(&p)

	err = apiClient.StopPipeline(p.Uuid)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		c.JSON(200, "stopped")
	}

}

func newPipeline(c *gin.Context) {

	c.HTML(200, "pipelines/new", gin.H{})
}

func deleteAsset(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		pp.Println("c-->", c.Param("id"))
		c.Abort()
		return
	}

	assetID, err := strconv.Atoi(c.Param("assetID"))
	if err != nil {
		pp.Println("c-->", c.Param("assetID"))
		c.Abort()
		return
	}

	var p = Pipeline{ID: id}
	var a = Asset{ID: assetID}
	db.First(&p)
	db.Model(&p).Association("Assets").Delete(&a)
	db.Delete(&a)

	c.Redirect(302, fmt.Sprintf("/ui/pipelines/%d", p.ID))
}

func updateAsset(c *gin.Context) {
	c.Request.ParseForm()
	pp.Println("c.Request.PostForm-->", c.Request.PostForm)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		pp.Println("c-->", c.Param("id"))
		c.Abort()
		return
	}

	assetID, err := strconv.Atoi(c.Param("assetID"))
	if err != nil {
		pp.Println("c-->", c.Param("assetID"))
		c.Abort()
		return
	}

	var p = Pipeline{ID: id}
	db.First(&p)
	var a = Asset{ID: assetID}
	db.First(&a)

	a.Name = c.Request.PostFormValue("name")
	a.Value = []byte(c.Request.PostFormValue("content"))
	db.Save(&a)

	session := sessions.Default(c)
	session.AddFlash(fmt.Sprintf("Asset %s saved", a.Name))
	session.Save()

	c.Redirect(302, fmt.Sprintf("/ui/pipelines/%d/assets/%d", p.ID, a.ID))
}

func replaceAsset(c *gin.Context) {
	assetID, err := strconv.Atoi(c.Param("assetID"))
	if err != nil {
		pp.Println("c-->", c.Param("assetID"))
		c.Abort()
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	filename := header.Filename
	size := header.Size

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		c.Abort()
		return
	}
	// Reset the read pointer if necessary.
	file.Seek(0, 0)
	contentType := http.DetectContentType(buffer[:n])

	pp.Println(header)

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Fatal(err)
	}

	var a = Asset{ID: assetID}
	db.First(&a)

	a.Name = filename
	a.ContentType = contentType
	a.Value = buf.Bytes()
	a.Size = int(size)

	db.Save(&a)
	c.String(200, "ok")
}

func createAsset(c *gin.Context) {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	filename := header.Filename
	size := header.Size

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		c.Abort()
		return
	}
	// Reset the read pointer if necessary.
	file.Seek(0, 0)
	contentType := http.DetectContentType(buffer[:n])

	pp.Println(header)

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Fatal(err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Abort()
		return
	}

	var p = Pipeline{ID: id}
	db.First(&p)

	var a = Asset{
		Name:        filename,
		ContentType: contentType,
		Value:       buf.Bytes(),
		Size:        int(size),
	}
	p.Assets = append(p.Assets, a)
	db.Save(&p)
	c.String(200, "ok")
}

func downloadAsset(c *gin.Context) {
	assetID, err := strconv.Atoi(c.Param("assetID"))
	if err != nil {
		pp.Println("c-->", c.Param("assetID"))
		c.Abort()
		return
	}

	var a = Asset{ID: (assetID)}
	db.First(&a)
	pp.Println("a.ContentType-->", a.ContentType)
	c.Header("Content-Disposition", "attachment; filename=\""+a.Name+"\"")
	c.Data(200, "application/octet-stream", a.Value)

}

func showAsset(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		pp.Println("c-->", c.Param("id"))
		c.Abort()
		return
	}

	assetID, err := strconv.Atoi(c.Param("assetID"))
	if err != nil {
		pp.Println("c-->", c.Param("assetID"))
		c.Abort()
		return
	}

	var a = Asset{ID: (assetID)}
	db.First(&a)

	var p = Pipeline{ID: id}
	db.Preload("Assets").First(&p)

	runningPipelines, _ := apiClient.ListPipelines()

	if pup, ok := runningPipelines[p.Uuid]; ok {
		p.IsUP = true
		p.LocationPath = pup.ConfigLocation
		p.StartedAt = pup.StartedAt

	}

	flashes := []string{}
	for _, m := range sessions.Default(c).Flashes() {
		flashes = append(flashes, m.(string))
	}
	sessions.Default(c).Save()

	c.HTML(200, "assets/edit", gin.H{
		"asset":    a,
		"pipeline": p,
		"flashes":  flashes,
	})
}

func createPipeline(c *gin.Context) {
	c.Request.ParseForm()

	uid, _ := uuid.NewV4()
	defaultValue := []byte("input{ }\n\nfilter{ }\n\noutput{ }")
	var p = Pipeline{
		Label:       c.Request.PostFormValue("label"),
		Description: c.Request.PostFormValue("description"),
		Uuid:        uid.String(),
		Assets: []Asset{{
			Name:        "bitfan.conf",
			Type:        "entrypoint",
			ContentType: "text/plain",
			Value:       defaultValue,
			Size:        len(defaultValue),
		}},
	}

	db.Create(&p)
	session := sessions.Default(c)
	session.AddFlash(fmt.Sprintf("Pipeline %s created", p.Label))
	session.Save()

	c.Redirect(302, fmt.Sprintf("/ui/pipelines/%d", p.ID))
}

func updatePipeline(c *gin.Context) {
	c.Request.ParseForm()
	pp.Println("c.Request.PostForm-->", c.Request.PostForm)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Abort()
		return
	}

	var p = Pipeline{ID: id}
	db.First(&p)

	p.Label = c.Request.PostFormValue("label")
	p.Description = c.Request.PostFormValue("description")
	p.UpdatedAt = time.Now()
	db.Save(&p)

	session := sessions.Default(c)
	session.AddFlash(fmt.Sprintf("Pipeline %s saved", p.Label))
	session.Save()

	c.Redirect(302, fmt.Sprintf("/ui/pipelines/%d", p.ID))
}

func editPipeline(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Abort()
		return
	}

	var p = Pipeline{ID: (idInt)}
	db.Preload("Assets").First(&p)

	runningPipelines, _ := apiClient.ListPipelines()

	if pup, ok := runningPipelines[p.Uuid]; ok {
		p.IsUP = true
		p.LocationPath = pup.ConfigLocation
		p.StartedAt = pup.StartedAt

	}

	flashes := []string{}
	for _, m := range sessions.Default(c).Flashes() {
		flashes = append(flashes, m.(string))
	}
	sessions.Default(c).Save()

	c.HTML(200, "pipelines/edit", gin.H{
		"pipeline": p,
		"flashes":  flashes,
	})

}

func getPipelines(c *gin.Context) {

	var pipelines []Pipeline

	db.Find(&pipelines)

	runningPipelines, _ := apiClient.ListPipelines()

	for i, p := range pipelines {
		if pup, ok := runningPipelines[p.Uuid]; ok {
			pipelines[i].IsUP = true
			pipelines[i].LocationPath = pup.ConfigLocation
			pipelines[i].StartedAt = pup.StartedAt

		}
	}

	c.HTML(200, "pipelines/index", gin.H{
		"pipelines": pipelines,
	})
}
