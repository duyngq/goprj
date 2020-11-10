package services




import (
	"reward-point/dto"
	"reward-point/models"
	"reward-point/services/go-points/repositories"

	"time"
)
type PointService interface {
	GetPoint(id string)  (dto.InquiryResponseDto,error)
	UpdatePoint(customer *models.Customer,point int,transType int) error

	GetCustomer(id string) (models.Customer,error)
	GetAllCustomers() ([]models.Customer, error)
	InsertCustomer(customer *models.Customer) error
	DeleteCustomer(id string) error
	UpdateCustomer( newCustomer *models.Customer) error
}
type pointService struct {}
var(
	pointRepository repositories.PointRepository
)

func NewPointService(repository repositories.PointRepository) PointService {
	pointRepository = repository
	return &pointService{}
}

func (*pointService) GetPoint(id string)  (dto.InquiryResponseDto,error){
	customer,err := pointRepository.GetCustomer(id)
	if err != nil {
		return dto.InquiryResponseDto{},err
	}
	totalPoint :=0
	var inquiryPoints []dto.InquiryPoint
	for _,v :=range customer.AvailablePoints{
		if v.ExpiredAt.After(time.Now()) {
			totalPoint += v.Amount

			inquiryPoints = append(inquiryPoints,dto.InquiryPoint{
				Point:     v.Amount,
				ExpiredAt: v.ExpiredAt,
			} )
		}
	}
	inquiry := dto.InquiryResponseDto{
		CisId:           customer.CisId,
		CreatedAt:       customer.CreatedAt,
		UpdatedAt:       customer.UpdatedAt,
		PointTotal:      totalPoint,
		AvailablePoints: inquiryPoints,
	}
	return inquiry,nil
}


func (*pointService) GetCustomer(id string) (models.Customer,error){
	customer,err := pointRepository.GetCustomer(id)
	return customer,err
}
func (*pointService) GetAllCustomers() ([]models.Customer, error)  {
	customers,err := pointRepository.GetAllCustomers()
	return customers,err
}
func (*pointService) InsertCustomer(customer *models.Customer) error{
	err :=	pointRepository.InsertCustomer(customer)
	return err
}
func (*pointService) DeleteCustomer(id string) error{
	err:=	pointRepository.DeleteCustomer(id)
	return err
}
func (*pointService) UpdateCustomer( newCustomer *models.Customer) error{
	err := pointRepository.UpdateCustomer( newCustomer)
	return err
}
func (*pointService) UpdatePoint(customer *models.Customer,point int,transType int) error{
	if transType ==0{

	}else {
		for i, availablePoint := range customer.AvailablePoints {
			if availablePoint.Amount > point && point !=0 && availablePoint.Amount !=0 {
				availablePoint.Amount -= point
				point = 0
				availablePoint.UpdatedAt = time.Now()
				customer.AvailablePoints[i] = availablePoint
			}else if availablePoint.Amount < point && point!=0 && availablePoint.Amount !=0{
				point -= availablePoint.Amount
				availablePoint.Amount = 0
				availablePoint.UpdatedAt = time.Now()
				customer.AvailablePoints[i] = availablePoint
			}
		}
	}
	err := pointRepository.UpdateCustomer(customer)
	return err
}

