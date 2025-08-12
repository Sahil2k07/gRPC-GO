package service

import (
	interfaces "github.com/Sahil2k07/gRPC-GO/internal/interface"
	"github.com/Sahil2k07/gRPC-GO/internal/model"
	"github.com/Sahil2k07/gRPC-GO/internal/util"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
)

type inventoryGroupService struct {
	repo interfaces.InventoryGroupRepository
}

func NewInventoryGroupService(r interfaces.InventoryGroupRepository) interfaces.InventoryGroupService {
	return &inventoryGroupService{repo: r}
}

func (s *inventoryGroupService) GetInventoryGroup(id string) (view.InventoryGroupView, error) {
	uid, err := util.StringToUint(id)
	if err != nil {
		return view.InventoryGroupView{}, err
	}

	group, err := s.repo.GetInventoryGroup(uid)
	if err != nil {
		return view.InventoryGroupView{}, err
	}

	return view.NewInventoryGroupView(group), nil
}

func (s *inventoryGroupService) AddInventoryGroup(req view.AddInventoryGroup) error {
	if err := s.repo.ValidateExistingCode(0, req); err != nil {
		return err
	}

	return s.repo.AddInventoryGroup(req)
}

func (s *inventoryGroupService) UpdateInventoryGroup(req view.InventoryGroupView) error {
	if err := s.repo.ValidateExistingCode(req.ID, req.AddInventoryGroup); err != nil {
		return err
	}

	return s.repo.UpdateInventoryGroup(req)
}

func (s *inventoryGroupService) DeleteInventoryGroup(id string) error {
	uid, err := util.StringToUint(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteInventoryGroup(uid)
}

func (s *inventoryGroupService) ListInventoryGroup(req view.ListInventoryGroup) (view.ListResponse, error) {
	list, count, err := s.repo.ListInventoryGroup(req)
	if err != nil {
		return view.ListResponse{}, err
	}

	records := util.Map(list, func(m model.InventoryGroup) view.InventoryGroupView {
		return view.NewInventoryGroupView(m)
	})

	response := view.ListResponse{
		TotalRecords: count,
		Data:         records,
	}

	return response, nil
}
