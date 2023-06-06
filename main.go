package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	aws "github.com/aws/aws-lambda-go/lambda"
	"github.com/livingstone23/provimaduser/awsgo"
)

func main() {
	fmt.Println("Inicio Proyecto provimadUser")

	aws.Start(EjecutoLambda)

}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parametros debe enviar 'Secretmanager'")
		err := errors.New("error en los parametros debe enviar SecretName")
		return event, err
	}

}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro

}
