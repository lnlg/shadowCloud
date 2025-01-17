package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"path"
	"shadowCloud/app/models"
	"shadowCloud/internal/global"
	"shadowCloud/internal/tool"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
	"golang.org/x/exp/rand"
)

type videoService struct {
	baseService // nolint:unused
}

// 获取视频列表
func (v *videoService) GetVideoList(typeId int, page int, limit int) interface{} {
	data, _ := models.GetVideoList(typeId, page, limit)
	count, _ := models.GetVideoListCount(typeId)
	return map[string]interface{}{
		"data":  data,
		"total": count,
		"page":  page,
		"limit": limit,
	}
}

// 根据id获取视频信息
func (v *videoService) GetVideoInfo(id int) interface{} {
	data, _ := models.GetVideoById(uint(id))
	return data
}

// 重新下载视频
func (v *videoService) DownloadVideo(id int, downloadState int) interface{} {
	// 根据ID获取视频信息
	video, _ := models.GetVideoById(uint(id))
	if video.Id == 0 {
		return nil
	}
	//解析OrgDownloadUrl
	u, err := url.Parse(video.OrgDownloadUrl)
	if err != nil {
		return nil
	}

	mapData := make(map[string]interface{})
	url := v.GetRandDowUrl()
	mapData["id"] = video.Id
	mapData["down_url"] = url + u.Path
	mapData["title"] = video.VideoName + path.Ext(path.Base(u.Path)) //文件名
	mapData["status"] = video.DownloadState                          //下载状态
	local_storage := models.GetVideoSetting("local_storage")
	mapData["local_storage"] = local_storage + "/" + strconv.Itoa(video.TypeId) + "/" //本地存储绝对路径
	// jsontext, _ := json.Marshal(mapData)
	// dao.Rdb.LPush(context.Background(), "video:download_video:list", jsontext)
	//修改状态下载中-1
	models.UpdateVideoDownloadState(video.Id, downloadState)
	//延迟队列2小时未成功下载，恢复状态-0
	r_score := tool.GetDateTime() + 7200
	global.Rdb.ZAdd(context.Background(), "video:download_video:delay", redis.Z{Score: float64(r_score), Member: video.Id}).Err()
	return mapData
}

// 获取随机真实下载地址
func (v *videoService) GetRandDowUrl() string {
	//https://992kp-f.pppp826.xyz/js/u.js
	var items []string
	jsonString := models.GetVideoSetting("dow_url")
	if jsonString == "" {
		return ""
	}
	err := json.Unmarshal([]byte(jsonString), &items)
	if err == nil {
		// 随机取一个元素
		randomIndex := rand.Intn(len(items))
		return items[randomIndex]
	}
	return ""
}

// 获取未下载视频信息
func (v *videoService) GetNoDownloadVideo() interface{} {
	data, _ := models.GetRandomVideoInfo()
	if data.Id == 0 {
		return nil
	}
	down_data := v.DownloadVideo(int(data.Id), 1)
	return down_data
}

// 下载完成处理
func (v *videoService) MarkDownloadSuccess(id int) string {
	data, _ := models.GetVideoById(uint(id))
	if data.Id == 0 {
		return "视频不存在"
	}
	//状态2下载完成不处理
	if data.DownloadState == 2 {
		return "视频已下载完成"
	}
	if data.DownloadState == 0 {
		return "暂不处理"
	}
	//检查文件是否存在
	local_storage := models.GetVideoSetting("domain_path")
	path := local_storage + data.VideoUrl
	fmt.Println(path)
	if tool.FileExists(path) {
		// 下载完成, 修改状态设置成2
		models.UpdateVideoDownloadState(data.Id, 2)
		//删除缓存
		cacheKey := "video:download_video:delay"
		//延迟队列2小时未成功下载，删除
		global.Rdb.ZRem(context.Background(), cacheKey, data.Id)
		return "处理完毕"
	} else {
		// 修改状态下载失败,恢复状态-0
		models.UpdateVideoDownloadState(data.Id, 0)
		return "文件不存在"
	}
}

// 把缓存数据存入数据库
func (v *videoService) CacheVideoTODB() interface{} {
	type Data struct {
		Typeid         string `json:"type_id"`
		VideoName      string `json:"video_name"`
		OrgDownloadUrl string `json:"org_download_url"`
	}
	returnStr := ""
	cacheKey := "video:download_video_one_list"
	whileNum := 0
	for whileNum < 10 {
		jsonData := global.Rdb.RPop(context.Background(), cacheKey).Val()
		if jsonData == "" {
			returnStr += "暂无数据"
			break
		}
		var data Data
		err := json.Unmarshal([]byte(jsonData), &data)
		if err == nil {
			typeId, _ := strconv.Atoi(data.Typeid)
			last_id, err := v.MakeVideoInfo(typeId, data.VideoName, data.OrgDownloadUrl, 0)
			if err == nil {
				returnStr += data.VideoName + ":" + strconv.Itoa(last_id) + ","
			}
		}
	}
	v.ClearRedisHtmlDetails() //清除缓存 video:html_details:*
	return returnStr
}

// 清除缓存 video:html_details:*
func (v *videoService) ClearRedisHtmlDetails() bool {
	cacheKey := "video:html_details:*"
	keys := global.Rdb.Keys(context.Background(), cacheKey).Val()
	for _, key := range keys {
		global.Rdb.Del(context.Background(), key)
	}
	return true
}

// 添加视频信息
func (v *videoService) MakeVideoInfo(typeid int, videoName string, orgDownloadUrl string, downloadState int) (int, error) {
	//判断video_download_url表是否添加过
	if !models.CheckDownloadUrl(orgDownloadUrl, tool.EncryMd5(videoName)) {
		return 0, errors.New("视频已添加")
	}
	var param models.VideoList
	param.VideoName = videoName
	param.TypeId = typeid
	u, _ := url.Parse(orgDownloadUrl)
	filename := path.Base(u.Path)
	video_path := models.GetVideoSetting("video_path")
	param.VideoUrl = video_path + strconv.Itoa(typeid) + "/" + videoName + path.Ext(filename)
	image_path := models.GetVideoSetting("image_path")
	param.ImageUrl = image_path + strconv.Itoa(typeid) + "/" + videoName + ".png"
	param.DownloadState = downloadState
	param.Md5File = tool.EncryMd5(videoName)
	param.OrgDownloadUrl = orgDownloadUrl
	last_id, err := models.AddVideo(&param)
	return int(last_id), err
}

// 定时删除大于当前时间的下载中状态的文件
func (v *videoService) DeleteDownloadVideo() bool {
	cacheKey := "video:download_video:delay"
	opts := redis.ZRangeBy{
		Min:    "-inf",
		Max:    strconv.Itoa(tool.GetDateTime()),
		Offset: 0,
		Count:  int64(100),
	}
	list, err := global.Rdb.ZRangeByScore(context.Background(), cacheKey, &opts).Result()
	if err == nil {
		for _, key := range list {
			println("过期id:", key)
			//状态标记成未下载 0
			id, _ := strconv.ParseUint(key, 10, 64)
			models.UpdateVideoDownloadState(uint(id), 0)
			global.Rdb.ZRem(context.Background(), cacheKey, key)
		}
	}
	return true
}
func (v *videoService) ScanDirToDB(dir string) interface{} {
	//扫描目录
	files := tool.ScanDirectory(dir)
	returnStr := ""
	for _, file := range files {
		//判断是否是文件
		if !tool.FileExists(file) {
			continue
		}

		file_new := strings.Replace(file, dir, "", -1)
		typeIdStr := path.Dir(file_new)
		type_id_str := strings.Trim(typeIdStr, "/")
		video_name := path.Base(file_new)
		if video_name == ".DS_Store" || video_name == "." {
			continue
		}
		video_name = strings.Replace(video_name, path.Ext(file_new), "", -1)
		org_download_url := "http://my.video.cn/public/video/" + typeIdStr + "/" + video_name + path.Ext(file_new)

		type_id, _ := strconv.Atoi(type_id_str)
		last_id, _ := v.MakeVideoInfo(type_id, video_name, org_download_url, 2)
		returnStr += video_name + ":" + strconv.Itoa(last_id) + ","
	}
	return returnStr
}
