package awsgo

import (
	"context"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	config "github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var cfg aws.Config
var err error

func InicializoAWS() {

	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	if err != nil {
		panic("Error al cargar configuracion de aws/config" + err.Error())
	}

}
