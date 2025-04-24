package data

import (
	model "customer-management/models"
)

type CustomerRelationRepo struct {
	data *Data
}

func NewCustomerRelationRepo(data *Data) *CustomerRelationRepo {
	return &CustomerRelationRepo{data: data}
}

func (r *CustomerRelationRepo) CreateCustomerRelation(info *model.RelationTable) error {

	return r.data.DB.Create(info).Error
}

func (r *CustomerRelationRepo) GetCustomerRelation(info *model.RelationTable) error {

	return r.data.DB.Create(info).Error
}

func (r *CustomerRelationRepo) UpdateCustomerRelation(info *model.RelationTable) error {

	return r.data.DB.Create(info).Error
}

func (r *CustomerRelationRepo) DelationCustomerRelation(info *model.RelationTable) error {

	return r.data.DB.Create(info).Error
}
