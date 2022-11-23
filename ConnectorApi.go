package main

type Connector struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type ConnectorApi interface {
	GetConnector() ([]byte, error)
	DeleteConnector(name string) ([]byte, error)
}

func NewConnector(act ConnectorApi) ConnectorApi {
	return act
}

/*GET /connectors –返回所有正在运行的connector名。

POST /connectors – 新建一个connector; 请求体必须是json格式并且需要包含name字段和config字段，name是connector的名字，config是json格式，必须包含你的connector的配置信息。

GET /connectors/{name} – 获取指定connetor的信息。

GET /connectors/{name}/config – 获取指定connector的配置信息。

PUT /connectors/{name}/config – 更新指定connector的配置信息。

GET /connectors/{name}/status – 获取指定connector的状态，包括它是否在运行、停止、或者失败，如果发生错误，还会列出错误的具体信息。

GET /connectors/{name}/tasks – 获取指定connector正在运行的task。

GET /connectors/{name}/tasks/{taskid}/status – 获取指定connector的task的状态信息。

PUT /connectors/{name}/pause – 暂停connector和它的task，停止数据处理知道它被恢复。

PUT /connectors/{name}/resume – 恢复一个被暂停的connector。

POST /connectors/{name}/restart – 重启一个connector，尤其是在一个connector运行失败的情况下比较常用

POST /connectors/{name}/tasks/{taskId}/restart – 重启一个task，一般是因为它运行失败才这样做。

DELETE /connectors/{name} – 删除一个connector，停止它的所有task并删除配置。*/