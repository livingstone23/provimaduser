package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	aws "github.com/aws/aws-lambda-go/lambda"
	"github.com/livingstone23/provimaduser/awsgo"
	"github.com/livingstone23/provimaduser/bd"
	"github.com/livingstone23/provimaduser/models"
)

func main() {
	fmt.Println("Inicio Proyecto provimadUser")

	aws.Start(EjecutoLambda)

}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	fmt.Println("Inicializando funcion  EjecutoLambda")

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parametros debe enviar 'Secretmanager'")
		err := errors.New("error en los parametros debe enviar SecretName")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el Secret" + err.Error())
		return event, err
	}

	err = bd.SignUp(datos)
	return event, err
}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro

}
