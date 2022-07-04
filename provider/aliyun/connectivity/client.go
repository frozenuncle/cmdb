package connectivity

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	bssopenapi "github.com/alibabacloud-go/bssopenapi-20171214/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
)

// NewAliCloudClient client
func NewAliCloudClient(ak, sk, region string) *AliCloudClient {
	return &AliCloudClient{
		Region:       region,
		AccessKey:    ak,
		AccessSecret: sk,
	}
}

type AliCloudClient struct {
	AccessKey    string `env:"AL_CLOUD_ACCESS_KEY"`
	AccessSecret string `env:"AL_CLOUD_ACCESS_SECRET"`
	Region       string `env:"AL_CLOUD_REGION"`

	ecsConn *ecs.Client
	rdsConn *rds.Client
	bssConn *bssopenapi.Client
}

// EcsClient 客户端
func (c *AliCloudClient) EcsClient() (*ecs.Client, error) {
	if c.ecsConn != nil {
		return c.ecsConn, nil
	}

	client, err := ecs.NewClientWithAccessKey(c.Region, c.AccessKey, c.AccessSecret)
	if err != nil {
		return nil, err
	}

	c.ecsConn = client

	return client, nil
}

func (c *AliCloudClient) BssClient() (*bssopenapi.Client, error) {
	if c.bssConn != nil {
		return c.bssConn, nil
	}

	client, err := bssopenapi.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(c.AccessKey),
		AccessKeySecret: tea.String(c.AccessSecret),
		Endpoint:        tea.String("business.aliyuncs.com"),
	})
	if err != nil {
		return nil, err
	}

	c.bssConn = client

	return client, nil
}

func (c *AliCloudClient) RdsClient() (*rds.Client, error) {
	if c.rdsConn != nil {
		return c.rdsConn, nil
	}

	client, err := rds.NewClientWithAccessKey(c.Region, c.AccessKey, c.AccessSecret)
	if err != nil {
		return nil, err
	}

	c.rdsConn = client

	return client, nil
}

func (c *AliCloudClient) OssClient() (*oss.Client, error) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	return oss.New("yourEndpoint", "yourAccessKeyId", "yourAccessKeySecret")
}

// 获取客户端账号ID
func (c *AliCloudClient) Account() (string, error) {
	args := sts.CreateGetCallerIdentityRequest()

	stsClient, err := sts.NewClientWithAccessKey(c.Region, c.AccessKey, c.AccessSecret)
	stsClient.GetConfig().WithScheme("HTTPS")

	if err != nil {
		return "", fmt.Errorf("unable to initialize the STS client: %#v", err)
	}
	stsClient.AppendUserAgent("Infraboard", "1.0")
	identity, err := stsClient.GetCallerIdentity(args)
	if err != nil {
		return "", err
	}
	if identity == nil {
		return "", fmt.Errorf("caller identity not found")
	}

	return identity.AccountId, nil
}
