﻿package main

import (
	log "github.com/cihub/seelog"
	"encoding/json"
	"github.com/bottos-project/magiccube/service/requirement/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	errcode "github.com/bottos-project/magiccube/error"
	sign "github.com/bottos-project/magiccube/service/common/signature"
	"os"
)

type Requirement struct {
	Client requirement.RequirementClient
}

func (s *Requirement) Publish(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200

	//验签
	is_true, err := sign.PushVerifySign(req.Body)
	if !is_true {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	var publishRequest requirement.PublishRequest
	err = json.Unmarshal([]byte(req.Body), &publishRequest)
	if err != nil {
		log.Error(err)
		return err
	}

	response, err := s.Client.Publish(ctx, &publishRequest)
	if err != nil {
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

func (s *Requirement) Query(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	body := req.Body
	log.Info(body)
	var requirementQuery requirement.QueryRequest
	err := json.Unmarshal([]byte(body), &requirementQuery)
	if err != nil {
		log.Error(err)
		return err
	}
	requirementQuery.Username = ""
	response, err := s.Client.Query(ctx, &requirementQuery)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}

func (s *Requirement) QueryByUsername(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.StatusCode = 200
	body := req.Body
	var requirementQuery requirement.QueryRequest
	err := json.Unmarshal([]byte(body), &requirementQuery)
	if err != nil {
		log.Error(err)
		return err
	}

	//验签
	is_true, err := sign.QueryVerifySign(req.Body)
	if !is_true {
		rsp.Body = errcode.ReturnError(1000, err)
		return nil
	}

	response, err := s.Client.Query(ctx, &requirementQuery)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Body = errcode.Return(response)
	return nil
}


func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/req-log.xml")
	if err != nil{
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.v3.requirement"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Requirement{Client: requirement.NewRequirementClient("go.micro.srv.v3.requirement", service.Client())},
		),
	)
	if err := service.Run(); err != nil {
		os.Exit(1)
	}
}
