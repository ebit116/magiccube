package service

//"fmt"

import (
	"context"
	"errors"
	"fmt"

	"github.com/bottos-project/magiccube/service/storage/proto"
	"github.com/bottos-project/magiccube/service/storage/util"
)

func (c *StorageService) GetUserRequirementList(ctx context.Context, request *storage.UserRequireListRequest, response *storage.UserRequireListResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetUserRequirementList")
	requires, err := c.mgoRepo.CallGetUserRequirementList(request.Username)
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	response.RequireList = []*storage.Requirement{}
	for _, require := range requires {
		requiresTag := &storage.Requirement{require.RequirementId,
			require.Username,
			require.RequirementName,
			require.FeatureTag,
			require.SamplePath,
			require.SampleHash,
			require.ExpireTime,
			require.Price,
			require.Description,
			require.PublishDate}
		response.RequireList = append(response.RequireList, requiresTag)
	}
	response.Code = 1
	return nil
}
func (c *StorageService) GetRequirementListByFeature(ctx context.Context, request *storage.FeatureRequireListRequest, response *storage.FeatureRequireListResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetRequirementListByFeature")
	requires, err := c.mgoRepo.CallGetRequirementListByFeature(request.FeatureTag)
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	response.RequireList = []*storage.Requirement{}
	for _, require := range requires {
		requiresTag := &storage.Requirement{require.RequirementId,
			require.Username,
			require.RequirementName,
			require.FeatureTag,
			require.SamplePath,
			require.SampleHash,
			require.ExpireTime,
			require.Price,
			require.Description,
			require.PublishDate}
		response.RequireList = append(response.RequireList, requiresTag)
	}
	response.Code = 1
	return nil
}
func (c *StorageService) GetRequirementNumByDay(ctx context.Context, request *storage.AllRequest, response *storage.DayRequirementNumResponse) error {

	response.DayRequirementNum = 200
	response.Code = 1
	return nil
}

func (c *StorageService) GetRequirementNumByWeek(ctx context.Context, request *storage.AllRequest, response *storage.WeekRequirementNumResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetRequirementNumByWeek")
	response.WeekRequirementNum = make([]uint64, 1, 7)
	days := util.WeekDur()
	for _, day := range days {
		require_num, err := c.mgoRepo.CallGetRequirementNumByDay(day.Begin, day.End)
		if err != nil {
			response.Code = 0
			fmt.Println(err)
			return errors.New("Failed CallGetAssetNumByDay")
		}
		response.WeekRequirementNum = append(response.WeekRequirementNum, require_num)
	}
	response.Code = 1
	return nil
}
func (c *StorageService) GetAllRequirementList(ctx context.Context, request *storage.AllRequest, response *storage.AllRequireListResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetRequirementListByFeature")
	requires, err := c.mgoRepo.CallGetAllRequirementList()
	if err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed get put url")

	}
	response.RequireList = []*storage.Requirement{}
	for _, require := range requires {
		dbTag := &storage.Requirement{require.RequirementId,
			require.Username,
			require.RequirementName,
			require.FeatureTag,
			require.SamplePath,
			require.SampleHash,
			require.ExpireTime,
			require.Price,
			require.Description,
			require.PublishDate}
		response.RequireList = append(response.RequireList, dbTag)
	}
	response.Code = 1
	return nil
}
