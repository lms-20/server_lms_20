package migration

import (
	"fmt"
	"lms-api/database"
	"lms-api/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Migration interface {
	AutoMigrate()
	SetDb(*gorm.DB)
}

type migration struct {
	Db            *gorm.DB
	DbModels      *[]interface{}
	IsAutoMigrate bool
}

func Init() {
	mgConfigurations := map[string]Migration{
		"POSTGRES": &migration{
			DbModels: &[]interface{}{
				&model.UserEntityModel{},
				&model.MentorEntityModel{},
				&model.CategoryEntityModel{},
				&model.CourseEntityModel{},
				&model.ChapterEntityModel{},
				&model.LessonEntityModel{},
				&model.NoteEntityModel{},
				&model.ReviewEntityModel{},
				&model.MyCourseEntityModel{},
				&model.OrderEntityModel{},
				&model.PaymentLogEntityModel{},
				&model.ReqCounselingEntityModel{},
				&model.ReqCourseEntityModel{},
			},
			IsAutoMigrate: true,
		},
	}

	for k, v := range mgConfigurations {
		dbConnection, err := database.Connection(k)
		if err != nil {
			logrus.Error(fmt.Sprintf("Failed to run migration, database not found %s", k))
		} else {
			v.SetDb(dbConnection)
			v.AutoMigrate()
			logrus.Info(fmt.Sprintf("Successfully run migration for database %s", k))
		}
	}

}

func (m *migration) AutoMigrate() {
	if m.IsAutoMigrate {
		m.Db.AutoMigrate(*m.DbModels...)
	}
}

func (m *migration) SetDb(db *gorm.DB) {
	m.Db = db
}
