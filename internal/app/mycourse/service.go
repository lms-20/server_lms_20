package mycourse

import (
	"errors"
	"lms-api/internal/abstraction"
	"lms-api/internal/app/courses"
	"lms-api/internal/dto"
	"lms-api/internal/factory"
	"lms-api/internal/model"
	"lms-api/internal/repository"
	res "lms-api/pkg/util/response"
	"lms-api/pkg/util/trxmanager"

	"gorm.io/gorm"
)

type Service interface {
	Create(ctx *abstraction.Context, payload *dto.MyCourseCreateRequest) (*dto.MyCourseCreateResponse, error)
	FindByID(ctx *abstraction.Context, id *int) (*dto.MyCourseGetByIDResponse, error)
}

type service struct {
	RepositoryMyCourse repository.MyCourse
	ServiceCourse      courses.Service
	// ServiceOrder       order.Service
	Db *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repositoryMyCourse := f.MyCourseRepository
	serviceCourse := courses.NewService(f)
	// serviceOrder := order.NewService(f)
	db := f.Db
	return &service{
		RepositoryMyCourse: repositoryMyCourse,
		ServiceCourse:      serviceCourse,
		// ServiceOrder:       serviceOrder,
		Db: db,
	}
}

func (s *service) FindByID(ctx *abstraction.Context, id *int) (*dto.MyCourseGetResponse, error) {
	var result *dto.MyCourseGetResponse
	var datas *[]model.MyCourseEntityModel

	datas, err := s.RepositoryMyCourse.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	result = &dto.MyCourseGetResponse{
		Datas: *datas,
	}

	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.MyCourseCreateRequest) (*dto.MyCourseCreateResponse, error) {
	var result *dto.MyCourseCreateResponse
	var data *model.MyCourseEntityModel

	// course, errs := s.ServiceCourse.FindByID(ctx, &payload.MyCourseEntity.CourseID)
	// if errs != nil {
	// 	return result, res.CustomErrorBuilder(404, "not found", "course not found")
	// }

	// if course.Type == "fremium" {

	// 	order, errs := s.ServiceOrder.Create(ctx, &dto.OrderCreateRequest{course.ID})
	// 	if errs != nil {
	// 		return result, errs
	// 	}

	// 	newOrder, errs := s.ServiceOrder.Update(ctx, &dto.OrderUpdateRequest{ID: order.ID, OrderEntity: order.OrderEntity}, &dto.CourseGetByIDResponse{course.CourseEntityModel})
	// 	if errs != nil {
	// 		return result, errs
	// 	}

	// } else {

	// }
	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.RepositoryMyCourse.Create(ctx, &payload.MyCourseEntity)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.MyCourseCreateResponse{
		MyCourseEntityModel: *data,
	}
	return result, nil

}
