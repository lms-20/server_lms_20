package factory

import (
	"lms-api/database"
	"lms-api/internal/repository"

	"gorm.io/gorm"
)

type Factory struct {
	Db                 *gorm.DB
	UserRepository     repository.User
	SampleRepository   repository.Sample
	MentorRepository   repository.Mentor
	CategoryRepository repository.Category
	CourseRepository   repository.Course
	ChapterRepository  repository.Chapter
	LessonRepository   repository.Lesson
	NoteRepository     repository.Note
	ReviewRepository   repository.Review
	MyCourseRepository repository.MyCourse
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection("POSTGRES")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	f.UserRepository = repository.NewUser(f.Db)
	f.CategoryRepository = repository.NewCategory(f.Db)
	f.SampleRepository = repository.NewSample(f.Db)
	f.MentorRepository = repository.NewMentor(f.Db)
	f.CourseRepository = repository.NewCourse(f.Db)
	f.ChapterRepository = repository.NewChapter(f.Db)
	f.LessonRepository = repository.NewLesson(f.Db)
	f.NoteRepository = repository.NewNote(f.Db)
	f.ReviewRepository = repository.NewReview(f.Db)
	f.MyCourseRepository = repository.NewMyCourse(f.Db)

}
