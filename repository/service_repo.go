package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/raynaaliyya/enigma-laundry/entity"
)

type ServiceRepo interface {
	CreateService(service entity.Service)
	GetAllService() any
	GetServiceById(id string) any
	UpdateService(service entity.Service)
	DeleteService(id string)
}

type serviceRepo struct {
	db *sql.DB
}

func (r *serviceRepo) CreateService(service entity.Service) {
	query := "INSERT INTO mst_service(id, pelayanan, harga) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, service.ID, service.Pelayanan, service.Harga)

	if err != nil {
		log.Println(err)
		fmt.Println("failed to create new service")
	} else {
		fmt.Println("new service created successfully")
	}
}

func (r *serviceRepo) GetAllService() any {
	var services []entity.Service

	query := "SELECT * FROM mst_service"

	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	// scan service
	for rows.Next() {
		var service entity.Service

		if err := rows.Scan(&service.ID, &service.Pelayanan, &service.Harga); err != nil {
			log.Println(err)
		}

		services = append(services, service)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(services) == 0 {
		return "no data"
	}

	return services
}

func (r *serviceRepo) GetServiceById(id string) any {
	query := "SELECT * FROM mst_service WHERE id = $1"

	service := entity.Service{}
	err := r.db.QueryRow(query, id).Scan(&service.ID, &service.Pelayanan, &service.Harga)

	if err != nil {
		log.Println(err)
	}

	if service.ID == "" {
		return "service not found"
	}

	return service
}

func (r *serviceRepo) UpdateService(service entity.Service) {
	query := "UPDATE mst_service SET pelayanan = $2, harga = $3 WHERE id = $1;"

	_, err := r.db.Exec(query, service.ID, service.Pelayanan, service.Harga)

	if err != nil {
		log.Println(err)
		fmt.Println("failed to update service")
	} else {
		fmt.Println("service updated successfully")
	}

}

func (r *serviceRepo) DeleteService(id string) {
	query := "DELETE FROM mst_service WHERE id = $1"
	_, err := r.db.Exec(query, id)

	if err != nil {
		log.Println(err)
		fmt.Println("failed to delete service")
	} else {
		fmt.Println("service deleted successfully")
	}

}

func NewServiceRepo(db *sql.DB) ServiceRepo {
	repo := new(serviceRepo)
	repo.db = db

	return repo
}
