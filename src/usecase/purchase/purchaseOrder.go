package purchaseUsecase

import (
	"projectsprintw4/src/constants"
)

func (uc *sPurchaseUsecase) PurchaseOrder(estimationId string) (orderId string, err error) {
	orderId, err = uc.purchaseRepo.FindOrderByEstimationId(estimationId)
	if orderId != "" {
		err = uc.purchaseRepo.UpdateOrderStatus(orderId, constants.PURCHASED)
	}

	return orderId, err
}
