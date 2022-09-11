package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/znkisoft/watchdog/schema"
)

func NewUserModel(u *schema.Userver) *UserverModel {
	return &UserverModel{
		schema: u,
		db:     db,
	}
}

type UserverModel struct {
	db     *sql.DB
	schema *schema.Userver
}

func (m *UserverModel) GetId() string {
	// TODO implement me
	panic("implement me")
}

func (m *UserverModel) SetId(id string) {

}

func (m *UserverModel) HasId() bool {
	// TODO implement me
	panic("implement me")
}

func (m *UserverModel) Save() error {
	stm, err := m.db.Prepare("insert into userver (id, name, ip, port, protocol, check_interval, timeout) values (?,?,?,?,?,?,?);")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stm.Close()

	id := uuid.NewString()
	if _, err = stm.Exec(id, m.schema.Name, m.schema.Ip, m.schema.Port, m.schema.Protocol, m.schema.CheckInterval, m.schema.Timeout); err != nil {
		log.Println(err)
		return err
	}
	log.Println(fmt.Sprintf("Successfully inserted new userver with id %s", id))
	return nil
}

func (m *UserverModel) Update(u *schema.Userver) {

}

func (m *UserverModel) Delete(id string) error {
	stm, err := m.db.Prepare("delete from userver where id = ?;")
	if err != nil {
		log.Println(err)
		return err
	}
	r, err := stm.Exec(id)
	if err != nil {
		log.Println(err)
		return err
	}

	affected, err := r.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	if affected == 0 {
		// TODO wrap error and use NotFound error
		return fmt.Errorf("no userver with id %s found", id)
	}

	return nil
}

func (m *UserverModel) Get(p *schema.Pagination) ([]*schema.Userver, error) {
	stm, err := m.db.Prepare("select id, name, ip, port, protocol, check_interval, timeout from userver limit ? offset ?;")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, err := stm.Query(p.Limit, p.Skip)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var result []*schema.Userver
	for rows.Next() {
		var (
			id            string
			name          string
			ip            string
			port          string
			protocol      string
			checkInterval int32
			timeout       int32
		)
		if err := rows.Scan(&id, &name, &ip, &port, &protocol, &checkInterval, &timeout); err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, &schema.Userver{
			Id:            id,
			Name:          name,
			Ip:            ip,
			Port:          port,
			Protocol:      protocol,
			CheckInterval: checkInterval,
			Timeout:       timeout,
		})
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}
