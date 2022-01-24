package order

import (
	"encoding/json"
	"errors"
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
	"strings"
	"time"

	"github.com/veritrans/go-midtrans"
	"gorm.io/gorm"
)

type Service interface {
	PremiumAccess(ctx *abstraction.Context, userID *int, courseID *int) error
	ProcessOrder(ctx *abstraction.Context, payload *dto.TransactionNotificationRequest) (*dto.TransactionNotificationResponse, error)
	FindByUserID(ctx *abstraction.Context, ID *int) (*dto.OrderGetByUserIDResponse, error)
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

func (s *service) ProcessOrder(ctx *abstraction.Context, payload *dto.TransactionNotificationRequest) (*dto.TransactionNotificationResponse, error) {
	var result *dto.TransactionNotificationResponse
	realOrderID := strings.Split(payload.OrderID, "-")
	orderID, _ := strconv.Atoi(realOrderID[0])

	order, err := s.Repository.FindByID(ctx, &orderID)
	if err != nil {
		return result, err
	}

	if payload.TransactionStatus == "capture" && payload.FraudStatus == "accept" {
		order.Status = "success"
	} else if payload.TransactionStatus == "settlement" {
		order.Status = "success"
	} else if payload.TransactionStatus == "deny" || payload.TransactionStatus == "expire" || payload.TransactionStatus == "cancel" {
		order.Status = "cancelled"
	}

	_, err = s.Repository.Update(ctx, &order.ID, &order.OrderEntity)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
	}

	result = &dto.TransactionNotificationResponse{
		UserID:   order.UserID,
		CourseID: order.CourseID,
	}

	return result, nil

}

func (s *service) PremiumAccess(ctx *abstraction.Context, userID *int, courseID *int) error {

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		err = s.Repository.CreateMyCourse(ctx, userID, courseID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return err

	}

	return nil
}

func (s *service) FindByUserID(ctx *abstraction.Context, ID *int) (*dto.OrderGetByUserIDResponse, error) {
	var result *dto.OrderGetByUserIDResponse

	data, err := s.Repository.FindByUserID(ctx, ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.OrderGetByUserIDResponse{
		Datas: *data,
	}
	return result, err

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
