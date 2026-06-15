package database

import (
	"database/sql"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/database/migrations"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/server/model"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

type Store struct {
	DB *sql.DB
}

func NewStore(dbURL string) *Store {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal("Unable to connect to the database :", err)
	}
	return &Store{DB: db}
}

func (s *Store) RunMigration() error {
	err := migrations.CreateServiceTable(s.DB)
	if err != nil {
		return err
	}
	err = migrations.CreateUserTable(s.DB)
	if err != nil {
		return err
	}
	return nil
}


func (s *Store) Save(service model.Service) error {
	query := `INSERT INTO services (user_id,name,url,healthy,statusCode,checked_at,response_time) VALUES($1,$2,$3,$4,$5,$6,$7)`
	_, err := s.DB.Exec(query, service.UserID, service.Name, service.URL, service.Healthy, service.StatusCode, service.Checked_at, service.Response_time)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetAllServices(user_id int) ([]model.Service, error) {
	query := `SELECT * FROM services WHERE user_id =$1`
	rows, err := s.DB.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	var services []model.Service
	for rows.Next() {
		var service model.Service
		err = rows.Scan(&service.ID, &service.UserID, &service.Name, &service.URL, &service.Healthy, &service.StatusCode, &service.Checked_at, &service.Response_time)
		if err != nil {
			return nil, err
		}
		services = append(services, service)

	}
	return services, nil
}
func (s *Store) GetAndRunAllServices() ([]model.Service, error) {
	query := `SELECT * FROM services`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var services []model.Service
	for rows.Next() {
		var service model.Service
		err = rows.Scan(&service.ID, &service.UserID, &service.Name, &service.URL, &service.Healthy, &service.StatusCode, &service.Checked_at, &service.Response_time)
		if err != nil {
			return nil, err
		}
		services = append(services, service)

	}
	return services, nil
}

func (s *Store) UpdateServiceStatus(service model.Service) error {
	query := `UPDATE services SET name=$1,url=$2,healthy=$3,statusCode=$4,checked_at=$5,response_time=$6 WHERE id=$7`
	_, err := s.DB.Exec(query, service.Name, service.URL, service.Healthy, service.StatusCode, service.Checked_at, service.Response_time, service.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateService(service model.Service) error {
	query := `UPDATE services SET name=$1,url=$2,healthy=$3,statusCode=$4,checked_at=$5,response_time=$6 WHERE id=$7`
	_, err := s.DB.Exec(query, service.Name, service.URL, service.Healthy, service.StatusCode, service.Checked_at, service.Response_time, service.ID)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) DeleteService(service_id int) error {
	query := `DELETE FROM services WHERE id=$1`
	_, err := s.DB.Exec(query, service_id)
	if err != nil {
		return err
	}
	return nil
}
