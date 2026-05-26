package database

import (
	"database/sql"
	"github.com/Bhakaresuraj/Go_language/devops-healthcheck/model"
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
	query := `
	CREATE TABLE IF NOT EXISTS services(
		id SERIAL PRIMARY KEY,
		user_id INT NOt NULL, 
		name VARCHAR(255) NOT NULL	,
		url VARCHAR(255) NOT NULL,
		healthy BOOLEAN NOT NULL,
		statusCode INT NOT NULL,
		checked_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		response_time BIGINT
	)`
	_, err := s.DB.Exec(query)
	if err != nil {
		// log.Fatal("Unable to create table :",err)
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
		err = rows.Scan(&service.ID, &service.UserID, &service.Name, &service.URL, &service.Healthy, &service.StatusCode, &service.Checked_at,&service.Response_time)
		if err != nil {
			return nil, err
		}
		services = append(services, service)

	}
	return services, nil
}

func (s *Store) UpdateServiceStatus(service model.Service) error {
	query := `UPDATE services SET healthy=$1,statusCode=$2,checked_at=$3,response_time=$4 WHERE id=$5`
	_, err := s.DB.Exec(query, service.Healthy, service.StatusCode, service.Checked_at,service.Response_time,service.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteService(user_id, service_id int) error {
	query := `DELETE FROM services WHERE user_id=$1 AND id=$2`
	_, err := s.DB.Exec(query, user_id, service_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Select() ([]model.Service, error) {
	query := `SELECT * FROM services`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var services []model.Service
	for rows.Next() {
		var service model.Service
		err = rows.Scan(&service.ID, &service.UserID, &service.Name, &service.URL, &service.Healthy, &service.StatusCode, &service.Checked_at)
		if err != nil {
			return nil, err
		}
		services = append(services, service)

	}
	return services, nil
}
