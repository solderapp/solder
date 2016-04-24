package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder-api/model"
	"github.com/solderapp/solder-api/router/middleware/session"
	"github.com/solderapp/solder-api/store"
)

// GetBuilds retrieves all available builds.
func GetBuilds(c *gin.Context) {
	pack := session.Pack(c)

	records, err := store.GetBuilds(
		c,
		pack.ID,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch builds",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// GetBuild retrieves a specific build.
func GetBuild(c *gin.Context) {
	record := session.Build(c)

	c.JSON(
		http.StatusOK,
		record,
	)
}

// DeleteBuild removes a specific build.
func DeleteBuild(c *gin.Context) {
	pack := session.Pack(c)
	record := session.Build(c)

	err := store.DeleteBuild(
		c,
		pack.ID,
		record,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully deleted build",
		},
	)
}

// PatchBuild updates an existing build.
func PatchBuild(c *gin.Context) {
	pack := session.Pack(c)
	record := session.Build(c)

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind build data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind build data",
			},
		)

		c.Abort()
		return
	}

	err := store.UpdateBuild(
		c,
		pack.ID,
		record,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// PostBuild creates a new build.
func PostBuild(c *gin.Context) {
	pack := session.Pack(c)
	record := &model.Build{}

	if err := c.BindJSON(&record); err != nil {
		logrus.Warn("Failed to bind build data")
		logrus.Warn(err)

		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Failed to bind build data",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateBuild(
		c,
		pack.ID,
		record,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		record,
	)
}

// GetBuildVersions retrieves all versions related to a build.
func GetBuildVersions(c *gin.Context) {
	records, err := store.GetBuildVersions(
		c,
		&model.BuildVersionParams{
			Pack:  c.Param("pack"),
			Build: c.Param("build"),
		},
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch versions",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		records,
	)
}

// PatchBuildVersion appends a version to a build.
func PatchBuildVersion(c *gin.Context) {
	assigned := store.GetBuildHasVersion(
		c,
		&model.BuildVersionParams{
			Pack:    c.Param("pack"),
			Build:   c.Param("build"),
			Mod:     c.Param("mod"),
			Version: c.Param("version"),
		},
	)

	if assigned == true {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Version is already appended",
			},
		)

		c.Abort()
		return
	}

	err := store.CreateBuildVersion(
		c,
		&model.BuildVersionParams{
			Pack:    c.Param("pack"),
			Build:   c.Param("build"),
			Mod:     c.Param("mod"),
			Version: c.Param("version"),
		},
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to append version",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully appended version",
		},
	)
}

// DeleteBuildVersion deleted a version from a build
func DeleteBuildVersion(c *gin.Context) {
	assigned := store.GetBuildHasVersion(
		c,
		&model.BuildVersionParams{
			Pack:    c.Param("pack"),
			Build:   c.Param("build"),
			Mod:     c.Param("mod"),
			Version: c.Param("version"),
		},
	)

	if assigned == false {
		c.JSON(
			http.StatusPreconditionFailed,
			gin.H{
				"status":  http.StatusPreconditionFailed,
				"message": "Version is not assigned",
			},
		)

		c.Abort()
		return
	}

	err := store.DeleteBuildVersion(
		c,
		&model.BuildVersionParams{
			Pack:    c.Param("pack"),
			Build:   c.Param("build"),
			Mod:     c.Param("mod"),
			Version: c.Param("version"),
		},
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to unlink version",
			},
		)

		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Successfully unlinked version",
		},
	)
}
