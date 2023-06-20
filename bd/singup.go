package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/livingstone23/provimaduser/models"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DBConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	//sentencia := "INSERT INTO users (Users_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	sentencia := "insert into users(user_uuid, user_email, user_firstname, user_lastname, user_status, user_dateadd, user_dateupg)	Values('121','livingstone2','lcano','test','1','2023-01-01','2023-01-01')"
	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Ejecución Exitosa")
	return nil

}
