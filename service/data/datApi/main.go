package main

import (
	log "github.com/cihub/seelog"
	"encoding/json"
	"github.com/bottos-project/magiccube/service/data/proto"
	"github.com/micro/go-micro"
	//"github.com/micro/go-micro/client"
	//"github.com/micro/go-micro/client"
	api "github.com/micro/micro/api/proto"

	"github.com/bitly/go-simplejson"
	"golang.org/x/net/context"
	"io/ioutil"
	"fmt"
	//"os"

	//log "github.com/cihub/seelog"
	"net/http"
	"strings"
	//"time"
)

type Data struct {
	Client data.DataClient
}

//global map
var mslice = make(map[string]int)
var msliceip = make(map[string]string)

func (d *Data) FileCheck(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info("Start File Check !")
	var req1 data.FileCheckRequest
	json.Unmarshal([]byte(body), &req1)

	sliceHash := req1.Hash

	rep, err := d.Client.FileCheck(ctx, &data.FileCheckRequest{
		Hash: sliceHash,
	})
	if err != nil {
		log.Error(err)
	}
	log.Info(rep)
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"result":           rep.Result,
		"message":          rep.Message,
		"merkle_root_hash": rep.MerkleRootHash,
		"is_exist":         rep.IsExist,
	})
	rsp.Body = string(b)
	log.Info("rsp.Body")
	log.Info(rsp.Body)
	return nil
}
func (d *Data) GetFileUploadURL(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info("Start Get File URL!")
	var req1 data.GetFileUploadURLRequest
	json.Unmarshal([]byte(body), &req1)

	userName := req1.Username
	fileSlice := req1.Slice

	rep, err := d.Client.GetFileUploadURL(ctx, &data.GetFileUploadURLRequest{
		Username: userName,
		Slice:    fileSlice,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"result":  rep.Result,
		"message": rep.Message,
		"url":     rep.Url,
	})
	rsp.Body = string(b)

	return nil
}
func (d *Data) GetUploadProgress(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info("Start Get Upload Progress !")
	var req1 data.GetUploadProgressRequest
	json.Unmarshal([]byte(body), &req1)

	userName := req1.Username
	fileSlice := req1.Slice
	log.Info("userName !")
	log.Info(userName)
	log.Info("fileSlice !")
	log.Info(fileSlice)
	//1 check cache upload status
	uploadCacheResult, err := d.Client.GetUploadProgress(ctx, &data.GetUploadProgressRequest{
		Username: userName,
		Slice:    fileSlice,
	})
	log.Info("uploadCacheResult !")
	log.Info(uploadCacheResult)
	if err != nil {
		log.Error(err)
	}
	m := int(uploadCacheResult.ProgressDone)

	//2.1 get slice IPlist
	nodes, err := d.Client.GetFileStorageNode(ctx, &data.GetFileStorageNodeRequest{
		Username: userName,
		Slice:    fileSlice,
	})
	if err != nil {
		log.Info(err)
		return err
	}
	log.Info("nodes")
	log.Info(nodes)
	//2.2 storage
	storageOK := 0
	sliceIp := []*data.Ip{}
	sliceIp = nil
	for i := 0; i < m; i++ {
		sguid := fileSlice[i].Sguid
		if mslice[sguid] == 0 {
			//2.2.1 get slice ip
			Sip := nodes.Ip[i].SnodeIp
			log.Info("Sip")
			log.Info(Sip)
			//Sip := "127.0.0.1"
			addr := "http://" + Sip + ":8080/rpc"
			//2.2.2 get slice storage url

			params := `service=go.micro.srv.v2.data&method=Data.GetFileStorageURL&request={
					"username":"%s",
					"guid":"%s"}`
			s := fmt.Sprintf(params, userName, sguid)
			resp_body, err := http.Post(addr, "application/x-www-form-urlencoded",
				strings.NewReader(s))
			log.Info("Get Data from remote rpc err:")
			if err != nil {
				log.Info(err)
			}
			defer resp_body.Body.Close()
			body, err := ioutil.ReadAll(resp_body.Body)
			var url string
			if err != nil {
				log.Info(err)
			} else {
				jss, _ := simplejson.NewJson([]byte(body))
				url = jss.Get("url").MustString()

			}
			log.Info("url")
			log.Info(url)
			//2.2.3 storage slice file
			putResult, err := d.Client.PutFile(ctx, &data.PutFileRequest{
				Username: userName,
				Guid:     sguid,
				Url:      url,
			})
			if err != nil {
				log.Info(err)
				return err
			}
			log.Info("putResult")
			log.Info(putResult)
			mslice[sguid] = 1
			msliceip[sguid] = Sip
			nodeTag := &data.Ip{sguid,
			Sip}
			sliceIp = append(sliceIp, nodeTag)
		}else{
			nodeTag := &data.Ip{sguid,
			msliceip[sguid]}
			sliceIp = append(sliceIp, nodeTag)
		}
		storageOK++
	}

	//response
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"result":              uploadCacheResult.Result,
		"message":             uploadCacheResult.Message,
		"slice_progress_done": uploadCacheResult.SliceProgressDone,
		"slice_progressing":   uploadCacheResult.SliceProgressing,
		"storage_done":         storageOK,
		"storage_ip": 		   sliceIp,
	})
	rsp.Body = string(b)
	return nil

}

func (d *Data) GetStorageIP(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info("Start Get Storage IP !")
	var req1 data.GetStorageIPRequest
	json.Unmarshal([]byte(body), &req1)

	guid := req1.Guid

	rep, err := d.Client.GetStorageIP(ctx, &data.GetStorageIPRequest{
		Guid: guid,
	})
	if err != nil {
		log.Error(err)
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"reslut":  rep.Result,
		"message": rep.Message,
		"storage_addr":      rep.StorageAddr,
	})
	rsp.Body = string(b)

	return nil
}
func (d *Data) GetFileDownloadURL(ctx context.Context, req *api.Request, rsp *api.Response) error {
	body := req.Body
	log.Info("Start Get File Download URL!")
	var req1 data.GetFileDownloadURLRequest

	json.Unmarshal([]byte(body), &req1)

	userName := req1.Username
	guid := req1.Guid
	ip := req1.Ip
	n := len(ip)
	if userName == "" || guid == "" || ip == nil {
		rsp.StatusCode = 200
		b, _ := json.Marshal(map[string]interface{}{
			"result":  "404",
			"message": "Missing  request para",
			"url":     nil,
		})
		rsp.Body = string(b)
		return nil
	}
	for i := 0; i < n; i++ {
		downloadOK := 0
		//1.1 get slice ip
		Sip := ip[i].SnodeIp
		//Sip := "127.0.0.1"
		addr := "http://" + Sip + ":8080/rpc"
		//1.2 get slice storage url
		sguid := ip[i].Sguid
		log.Info("sguid:",sguid)
		log.Info("Sip:",Sip)
		log.Info("username:",userName)
		params := `service=go.micro.srv.v2.data&method=Data.GetFileStorageURL&request={
					"username":"%s",
					"guid":"%s"}`
		s := fmt.Sprintf(params, userName, sguid)
		resp_body, err := http.Post(addr, "application/x-www-form-urlencoded",
			strings.NewReader(s))
		if err != nil {
			return err
		}
		defer resp_body.Body.Close()
		dbody, err := ioutil.ReadAll(resp_body.Body)
		var url string
		if err != nil {
			return err
		} else {
			jss, _ := simplejson.NewJson([]byte(dbody))
			url = jss.Get("url").MustString()

		}
		log.Info("url:",url)
		//1.3 storage slice file
		downloadResult, err := d.Client.DownloadFile(ctx, &data.DownloadFileRequest{
			Username: userName,
			Guid:     sguid,
			Url:      url,
		})
		if err != nil {
			log.Error(err)
			return err
		}
		log.Info("downloadResult:",downloadResult)
		downloadOK++

	}
	log.Info("ip:",ip)
	//2.composeFile
	d.Client.ComposeFile(ctx, &data.ComposeFileRequest{
		Username: userName,
		Guid:     guid,
		Ip:       ip,
	})

	//3.get download url
	rep, err := d.Client.GetFileStorageURL(ctx, &data.GetFileStorageURLRequest{
		Username: userName,
		Guid:     guid,
	})
	if err != nil {
		log.Error(err)
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"result":  rep.Result,
		"message": rep.Message,
		"url":     rep.Url,
	})
	rsp.Body = string(b)

	return nil
}

func init() {
	logger, err := log.LoggerFromConfigAsFile("./config/data-log.xml")
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer logger.Flush()
	log.ReplaceLogger(logger)
}

func main() {

	service := micro.NewService(
		micro.Name("go.micro.api.v2.data"),

		//client.RequestTimeout(time.Second*30),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Data{Client: data.NewDataClient("go.micro.srv.v2.data", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Critical(err)
	}

}
