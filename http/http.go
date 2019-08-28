package http

import (
	"net/http"
	//"io/ioutil"
	//"path"
	"github.com/eric-lindau/flip/config"
	"github.com/eric-lindau/flip/core"
	"strconv"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"mime/multipart"
	"mime"
)

const (
	uuidLength = 36
)

type response struct {
	key string
}

type handler func(env *config.Env, c echo.Context) (error, int)

func Init(env *config.Env) {
	h := echo.New()
	h.Use(middleware.BodyLimit(strconv.Itoa(env.MaxData) + "M"))
	h.GET("/data/:key/:part", handle(env, getFile))
	h.GET("/data/:key", handle(env, info))
	h.POST("/data", handle(env, postFiles))

	h.Logger.Fatal(h.Start(":80"))
}

func handle(env *config.Env, h handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		err, code := h(env, c)
		if code > 0 {
			return echo.NewHTTPError(code, "an error occurred")
		}
		if err != nil {
			fmt.Printf("ERROR: %v\n", err.Error())
		}

		return err
	}
}

func getFile(env *config.Env, c echo.Context) (error, int) {
	//key := c.Param("key") // TODO: Verify key

	// TODO: Query S3
	//reg, err := ioutil.ReadDir(path.Join(env.DataPath, i))
	//if err != nil {
	//	return err, http.StatusInternalServerError
	//}

	// TODO: Fetch first file if this param doesn't exist
	//_, err := strconv.ParseInt(c.Param("part"), 10, 32)
	//if err != nil {
	//	return err, http.StatusBadRequest
	//}
	//if int(idx) >= len(reg) {
	//	return nil, http.StatusNotFound
	//}

	// TODO: Bufferless proxy?

	//c.Response().Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	//if err := c.File(path.Join(env.DataPath, i, fileName)); err != nil {
	//	return err, http.StatusInternalServerError
	//}
	return nil, 0
}

// Respond to client with information about given key
func info(env *config.Env, c echo.Context) (error, int) {
	//id := c.Param("key")
	//i, err := env.DB.Get(id).Result()
	//if err != nil || len(i) != uuidLength {
	//	return err, http.StatusNotFound
	//}
	//
	//reg, err := ioutil.ReadDir(path.Join(env.DataPath, i))
	//if err != nil {
	//	return err, http.StatusInternalServerError
	//}
	//
	//inf := make([]string, len(reg))
	//for i, e := range reg {
	//	inf[i] = e.Name()
	//}
	//
	//if err := c.JSON(http.StatusOK, inf); err != nil {
	//	return err, http.StatusInternalServerError
	//}

	return nil, 0
}

func postFiles(env *config.Env, c echo.Context) (error, int) {
	key, err := core.GenerateKey(env)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	c.Response().Header().Add("Content-Type", "application/json")
	if err := c.JSON(http.StatusOK, response{key.Id()}); err != nil {
		return err, http.StatusInternalServerError
	}

	// RFC 2557
	_, params, err := mime.ParseMediaType(c.Request().Header.Get("Content-Type"))
	if err != nil {
		return err, http.StatusInternalServerError
	}

	parts := multipart.NewReader(c.Request().Body, params["boundary"])
	for {
		chk, err := parts.NextPart()
		if err != nil {
			return err, http.StatusInternalServerError
		}

		core.ProcessData(chk, chk.FileName(), key)
	}

	return nil, 0
}