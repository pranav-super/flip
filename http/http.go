package http

import (
	"fmt"
	"mime"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/eric-lindau/flip/config"
	"github.com/eric-lindau/flip/core"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type response struct {
	Key string
}

type handler func(env *config.Env, c echo.Context) (error, int)

func Init(env *config.Env) {
	h := echo.New()
	h.Use(middleware.BodyLimit(strconv.Itoa(env.MaxData) + "M"))

	h.GET("/data/:key", handle(env, getFile))
	h.GET("/data/:key/:part", handle(env, getFile))
	h.GET("/data/:key/meta", handle(env, metadata))

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

// TODO: Bufferless proxy?
func getFile(env *config.Env, c echo.Context) (error, int) {
	key := c.Param("key")
	idx := 0  // Default to first file in key

    // Allow other index to be specified
	part := c.Param("part")
	if part != "" {
		idx, err := strconv.ParseInt(part, 10, 32)
		if err != nil {
			return err, http.StatusBadRequest
		}
	}

	// tempKey := core.NewS3Key("s3.flip.io", key)
	obj := core.Flip.Objects(key) // TODO: Cache
	if idx >= len(obj) {
		return nil, http.StatusNotFound
	}
	// buf := core.GetData(env.DataStore, tempKey.Extend(obj[idx]))
	buf := core.Flip.Get(tempKey.Extend(obj[idx]))

	c.Response().Header().Add("Content-Disposition", "attachment; filename=\""+key+"\"")
	c.Response().Write(buf) // TODO: Error check
	return nil, 0
}

// Hand client metadata corresp. to request's key.
// Metadata consists of a list of named objects stored under the key.
func metadata(env *config.Env, c echo.Context) (error, int) {
	key := c.Param("key")
	meta := core.Flip.Objects(key)

	if err := c.JSON(http.StatusOK, meta); err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, 0
}

func postFiles(env *config.Env, c echo.Context) (error, int) {
	// TODO: Take Key Options in request
	// TODO: Delegate key generation
	key, err := core.Flip.GenerateKey(env.KeyFunc, &core.KeyOptions{TTL: 5})
	if err != nil {
		return err, http.StatusInternalServerError
	}

	c.Response().Header().Add("Content-Type", "application/json")
	if err := c.JSON(http.StatusOK, response{key.Token()}); err != nil {
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
        
        core.Flip.Put(key.Extend(chk.FileName()), chk)
		// core.PutData(env.DataStore, chk, chk.FileName(), key)
	}

	return nil, 0
}
