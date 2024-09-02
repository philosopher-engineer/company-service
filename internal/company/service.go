package company

import (
	"github.com/IBM/sarama"
)

type CreateCompanyReq struct {
	Name              string `json:"name" binding:"required,max=15"`
	Description       string `json:"description,omitempty" binding:"max=3000"`
	AmountOfEmployees int    `json:"amount_of_employees" binding:"required"`
	Registered        bool   `json:"registered" binding:"required"`
	Type              string `json:"type" binding:"required,oneof=Corporations NonProfit Cooperative SoleProprietorship"`
}

type UpdateCompanyReq struct {
	Name              string `json:"name" binding:"max=15"`
	Description       string `json:"description,omitempty" binding:"max=3000"`
	AmountOfEmployees int    `json:"amount_of_employees,omitempty"`
	Registered        bool   `json:"registered,omitempty"`
	Type              string `json:"type" binding:"omitempty,oneof=Corporations NonProfit Cooperative SoleProprietorship"`
}

type Service interface {
	CreateCompany(req *CreateCompanyReq) error
	GetCompanyByID(id string) (*Company, error)
	UpdateCompany(id string, req *UpdateCompanyReq) error
	DeleteCompany(id string) error
}

type service struct {
	repo     Repository
	producer sarama.SyncProducer
}

func NewService(repo Repository, producer sarama.SyncProducer) Service {
	return &service{repo, producer}
}

func (s *service) CreateCompany(req *CreateCompanyReq) error {
	company := &Company{
		Name:              req.Name,
		Description:       req.Description,
		AmountOfEmployees: req.AmountOfEmployees,
		Registered:        req.Registered,
		Type:              req.Type,
	}

	err := s.repo.CreateCompany(company)
	if err != nil {
		return err
	}

	// Produce Kafka event
	msg := &sarama.ProducerMessage{
		Topic: "company_created",
		Value: sarama.StringEncoder(company.ID.String()),
	}
	_, _, err = s.producer.SendMessage(msg)
	return err
}

func (s *service) GetCompanyByID(id string) (*Company, error) {
	return s.repo.GetCompanyByID(id)
}

func (s *service) UpdateCompany(id string, req *UpdateCompanyReq) error {
	company, err := s.repo.GetCompanyByID(id)
	if err != nil {
		return err
	}

	company.Name = req.Name
	company.Description = req.Description
	company.AmountOfEmployees = req.AmountOfEmployees
	company.Registered = req.Registered
	company.Type = req.Type

	err = s.repo.UpdateCompany(company)
	if err != nil {
		return err
	}

	// Produce Kafka event
	msg := &sarama.ProducerMessage{
		Topic: "company_updated",
		Value: sarama.StringEncoder(company.ID.String()),
	}
	_, _, err = s.producer.SendMessage(msg)
	return err
}

func (s *service) DeleteCompany(id string) error {
	err := s.repo.DeleteCompany(id)
	if err != nil {
		return err
	}

	// Produce Kafka event
	msg := &sarama.ProducerMessage{
		Topic: "company_deleted",
		Value: sarama.StringEncoder(id),
	}
	_, _, err = s.producer.SendMessage(msg)
	return err
}
