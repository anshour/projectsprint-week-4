package purchaseUsecase

import entity "projectsprintw4/src/entities"

func (uc *sPurchaseUsecase) ListOrder(p *entity.ListOrderParams) (*[]entity.ListOrderResult, error) {
	orders, err := uc.purchaseRepo.FindOrders(p)

	if err != nil {
		empty := make([]entity.ListOrderResult, 0)
		return &empty, err
	}

	return orders, nil
}
