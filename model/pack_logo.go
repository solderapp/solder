package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/vincent-petithory/dataurl"
)

// PackLogos is simply a collection of pack logo structs.
type PackLogos []*PackLogo

// PackLogo represents a pack logo model definition.
type PackLogo struct {
	ID          int              `json:"id" gorm:"primary_key"`
	PackID      int              `json:"-"`
	Pack        *Pack            `json:"-"`
	ContentType string           `json:"content_type"`
	Path        string           `json:"-" sql:"-"`
	URL         string           `json:"url" sql:"-"`
	MD5         string           `json:"md5"`
	Upload      *dataurl.DataURL `json:"upload,omitempty" sql:"-"`
	CreatedAt   time.Time        `json:"-"`
	UpdatedAt   time.Time        `json:"-"`
}

// BeforeSave invokes required actions before persisting.
func (u *PackLogo) BeforeSave(db *gorm.DB) error {
	if u.Upload != nil {
		check := md5.Sum(
			u.Upload.Data,
		)

		hash := hex.EncodeToString(
			check[:],
		)

		u.MD5 = hash
		u.ContentType = u.Upload.MediaType.String()
	}

	return nil
}

// AfterSave invokes required actions after persisting.
func (u *PackLogo) AfterSave(db *gorm.DB) error {
	if u.Upload != nil {
		absolutePath, err := u.AbsolutePath()

		if err != nil {
			return fmt.Errorf("Missing storage path for logo")
		}

		errDir := os.MkdirAll(
			filepath.Dir(
				absolutePath,
			),
			os.ModePerm,
		)

		if errDir != nil {
			return fmt.Errorf("Failed to create logo directory")
		}

		file, errCreate := os.Create(
			absolutePath,
		)

		if errCreate != nil {
			return fmt.Errorf("Failed to open logo at %s", absolutePath)
		}

		_, errWrite := u.Upload.WriteTo(
			file,
		)

		if errWrite != nil {
			return fmt.Errorf("Failed to write logo at %s", absolutePath)
		}
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *PackLogo) Validate(db *gorm.DB) {
	if u.Upload == nil {
		db.AddError(fmt.Errorf("A logo is required"))
	}

	if isInvalidPackLogoType(u.Upload.MediaType.String()) {
		db.AddError(fmt.Errorf("Invalid logo media type"))
	}
}

// Path generates the absolute path to the logo.
func (u *PackLogo) AbsolutePath() (string, error) {
	if u.Path == "" {
		return "", fmt.Errorf("Missing storage path for logo")
	}

	return path.Join(
		u.Path,
		"logo",
		u.MD5,
	), nil
}

func isInvalidPackLogoType(mediaType string) bool {
	logrus.Debugf("Got %s pack logo media type", mediaType)
	return false
}
