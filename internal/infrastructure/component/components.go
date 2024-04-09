package component

import (
	"goTest/config"
	"goTest/internal/infrastructure/responder"

	"goTest/internal/infrastructure/godecoder"

	"go.uber.org/zap"
)

type Components struct {
	Conf      config.AppConf
	Responder responder.Responder
	Decoder   godecoder.Decoder
	Logger    *zap.Logger
}

func NewComponents(conf config.AppConf, responder responder.Responder, decoder godecoder.Decoder, logger *zap.Logger) *Components {
	return &Components{Conf: conf, Responder: responder, Decoder: decoder, Logger: logger}
}
