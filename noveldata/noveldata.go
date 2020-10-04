package noveldata

import (
	"fmt"
	"reflect"
	"strings"
)

// NovelData : data entity used for storing novel related data
type NovelData struct {
	Title         string
	Image         string
	LatestChapter string
	LastUpdated   string
}

// NewNovelData : constructor for novel  data
func NewNovelData(title string, image string, latestchapter string, lastupdated string) *NovelData {
	newdata := NovelData{
		Title:         title,
		Image:         image,
		LatestChapter: latestchapter,
		LastUpdated:   lastupdated,
	}

	return &newdata
}

func (data NovelData) toString() string {
	var str strings.Builder

	values := reflect.ValueOf(data)
	typesOfData := values.Type()

	for i := 0; i < values.NumField(); i++ {
		str.WriteString(fmt.Sprintf("%s : %v\n", typesOfData.Field(i).Name, values.Field(i).Interface()))
	}

	return str.String()
}
