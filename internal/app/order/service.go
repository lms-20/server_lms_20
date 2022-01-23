package order

import (
	"encoding/json"
	"fmt"
	"lms-api/internal/abstraction"
	"lms-api/internal/dto"
	"lms-api/internal/factory"
	"lms-api/internal/model"
	"lms-api/internal/repository"
	res "lms-api/pkg/util/response"
	"lms-api/pkg/util/trxmanager"
	"os"
	"strconv"
	"time"

	"github.com/veritrans/go-midtrans"
	"gorm.io/gorm"
)

var err error

type Service interface {
	Create(ctx *abstraction.Context, payload *dto.OrderCreateRequest) (*dto.OrderCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.OrderUpdateRequest, course *dto.CourseGetByIDResponse) (*dto.OrderUpdateResponse, error)
}

type service struct {
	Repository repository.Order
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.OrderRepository
	db := f.Db
	return &service{Repository: repository, Db: db}
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.OrderCreateRequest) (*dto.OrderCreateResponse, error) {
	var result *dto.OrderCreateResponse
	var data *model.OrderEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &payload.CourseID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.OrderCreateResponse{
		OrderEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.OrderUpdateRequest, course *dto.CourseGetByIDResponse) (*dto.OrderUpdateResponse, error) {
	var result *dto.OrderUpdateResponse
	var data *model.OrderEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		metadata := model.MetadataEntity{
			CourseID:        course.ID,
			CoursePrice:     course.Price,
			CourseName:      course.Name,
			CourseThumbnail: course.Thumbnail,
			CourseLevel:     string(course.Level),
		}

		metadataToJson, _ := json.Marshal(metadata)
		paymentUrl, err := s.GetPaymentURL(ctx, payload.ID, course)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}
		payload = &dto.OrderUpdateRequest{
			ID: payload.ID,
			OrderEntity: model.OrderEntity{
				Status:   "pending",
				CourseID: course.ID,
				SnapURL:  paymentUrl,
				Metadata: metadataToJson,
			},
		}

		data, err = s.Repository.Update(ctx, &payload.ID, &payload.OrderEntity)
		if err != nil {
			fmt.Println(err.Error())
			return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.OrderUpdateResponse{
		OrderEntityModel: *data,
	}
	return result, nil

}

func generateOrderID(orderID int) string {
	rand := strconv.FormatInt(time.Now().UnixNano(), 10)
	path := fmt.Sprintf("%d-%s", orderID, rand)
	return path
}

func (s *service) GetPaymentURL(ctx *abstraction.Context, orderID int, course *dto.CourseGetByIDResponse) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midclient.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: ctx.Auth.Email,
			FName: ctx.Auth.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  generateOrderID(orderID),
			GrossAmt: int64(course.Price),
		},
		Items: &[]midtrans.ItemDetail{
			midtrans.ItemDetail{
				Price:    int64(course.Price),
				Name:     course.Name,
				Qty:      1,
				Category: course.Category.Name,
			},
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
