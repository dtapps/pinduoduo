package pinduoduo

import (
	"fmt"
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
	"go.dtapp.net/gostring"
	"regexp"
	"strconv"
	"strings"
)

// client *dorm.GormClient
type gormClientFun func() *dorm.GormClient

// client *dorm.MongoClient
// databaseName string
type mongoClientFun func() (*dorm.MongoClient, string)

// ClientConfig 实例配置
type ClientConfig struct {
	ClientId       string         // POP分配给应用的client_id
	ClientSecret   string         // POP分配给应用的client_secret
	MediaId        string         // 媒体ID
	Pid            string         // 推广位
	GormClientFun  gormClientFun  // 日志配置
	MongoClientFun mongoClientFun // 日志配置
	Debug          bool           // 日志开关
	ZapLog         *golog.ZapLog  // 日志服务
	CurrentIp      string         // 当前ip
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	zapLog        *golog.ZapLog  // 日志服务
	currentIp     string         // 当前ip
	config        struct {
		clientId     string // POP分配给应用的client_id
		clientSecret string // POP分配给应用的client_secret
		mediaId      string // 媒体ID
		pid          string // 推广位
	}
	log struct {
		gorm           bool              // 日志开关
		gormClient     *dorm.GormClient  // 日志数据库
		logGormClient  *golog.ApiClient  // 日志服务
		mongo          bool              // 日志开关
		mongoClient    *dorm.MongoClient // 日志数据库
		logMongoClient *golog.ApiClient  // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	var err error
	c := &Client{}

	c.zapLog = config.ZapLog

	c.currentIp = config.CurrentIp

	c.config.clientId = config.ClientId
	c.config.clientSecret = config.ClientSecret
	c.config.mediaId = config.MediaId
	c.config.pid = config.Pid

	c.requestClient = gorequest.NewHttp()
	c.requestClient.Uri = apiUrl

	gormClient := config.GormClientFun()
	if gormClient != nil && gormClient.Db != nil {
		c.log.logGormClient, err = golog.NewApiGormClient(&golog.ApiGormClientConfig{
			GormClientFun: func() (*dorm.GormClient, string) {
				return gormClient, logTable
			},
			Debug:     config.Debug,
			ZapLog:    c.zapLog,
			CurrentIp: c.currentIp,
		})
		if err != nil {
			return nil, err
		}
		c.log.gorm = true
		c.log.gormClient = gormClient
	}

	mongoClient, databaseName := config.MongoClientFun()
	if mongoClient != nil && mongoClient.Db != nil {
		c.log.logMongoClient, err = golog.NewApiMongoClient(&golog.ApiMongoClientConfig{
			MongoClientFun: func() (*dorm.MongoClient, string, string) {
				return mongoClient, databaseName, logTable
			},
			Debug:     config.Debug,
			ZapLog:    c.zapLog,
			CurrentIp: c.currentIp,
		})
		if err != nil {
			return nil, err
		}
		c.log.mongo = true
		c.log.mongoClient = mongoClient
	}

	return c, nil
}

type ErrResp struct {
	ErrorResponse struct {
		ErrorMsg  string      `json:"error_msg"`
		SubMsg    string      `json:"sub_msg"`
		SubCode   interface{} `json:"sub_code"`
		ErrorCode int         `json:"error_code"`
		RequestId string      `json:"request_id"`
	} `json:"error_response"`
}

type CustomParametersResult struct {
	Sid string `json:"sid"`
	Uid string `json:"uid"`
}

func (c *Client) SalesTipParseInt64(salesTip string) int64 {
	parseInt, err := strconv.ParseInt(salesTip, 10, 64)
	if err != nil {
		salesTipStr := salesTip
		if strings.Contains(salesTip, "万+") {
			salesTipStr = strings.Replace(salesTip, "万+", "0000", -1)
		} else if strings.Contains(salesTip, "万") {
			salesTipStr = strings.Replace(salesTip, "万", "000", -1)
		}
		re := regexp.MustCompile("[0-9]+")
		SalesTipMap := re.FindAllString(salesTipStr, -1)
		if len(SalesTipMap) == 2 {
			return gostring.ToInt64(fmt.Sprintf("%s%s", SalesTipMap[0], SalesTipMap[1]))
		} else if len(SalesTipMap) == 1 {
			return gostring.ToInt64(SalesTipMap[0])
		} else {
			return 0
		}
	} else {
		return parseInt
	}
}

func (c *Client) CommissionIntegralToInt64(GoodsPrice, CouponProportion int64) int64 {
	return (GoodsPrice * CouponProportion) / 1000
}
