package payment

import (
	"fmt"
	"lms-api/internal/abstraction"
	"lms-api/internal/dto"
	"os"
	"strconv"
	"time"

	"github.com/veritrans/go-midtrans"
)

var (
	MIDTRANS_CLIENT_KEY = os.Getenv("MIDTRANS_CLIENT_KEY")
	MIDTRANS_SERVER_KEY = os.Getenv("MIDTRANS_SERVER_KEY")
)

type Service interface {
	GetPaymentURL(ctx *abstraction.Context, orderID *int, course *dto.CourseGetByIDResponse) (string, error)
}

func generateOrderID(orderID int) string {
	rand := strconv.FormatInt(time.Now().UnixNano(), 10)
	path := fmt.Sprintf("%d-%s", orderID, rand)
	return path
}

func GetPaymentURL(ctx *abstraction.Context, orderID int, course *dto.CourseGetByIDResponse) (string, error) {
	fmt.Println(MIDTRANS_CLIENT_KEY)
	fmt.Println(MIDTRANS_SERVER_KEY)
	midclient := midtrans.NewClient()
	midclient.ServerKey = MIDTRANS_SERVER_KEY
	midclient.ClientKey = MIDTRANS_CLIENT_KEY
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
			OrderID:  generateOrderID(1),
			GrossAmt: int64(course.Price),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
