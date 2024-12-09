package goframework_videosys

import (
	"errors"
	"github.com/kordar/godb"
	videocollection "github.com/kordar/video-collection"
)

var (
	streampool = godb.NewDbPool()
	keys       = map[string]bool{}
)

// GetStreamWrapper 获取Stream
func GetStreamWrapper(name string) CollectionWrapper {
	return streampool.Handle(name).(CollectionWrapper)
}

// GetStream 获取Stream
func GetStream(name string) videocollection.Collection {
	if HasStreamInstance(name) {
		streamWrapper := streampool.Handle(name).(CollectionWrapper)
		return streamWrapper.Collection
	}
	return nil
}

// AddStreamInstance 添加stream
func AddStreamInstance(name string, collection videocollection.Collection, configuration *videocollection.Configuration, retry videocollection.Retry) error {
	ins := NewVideoCollectionIns(name, collection, configuration, retry)
	if err := streampool.Add(ins); err == nil {
		keys[name] = true
		return nil
	} else {
		return err
	}
}

// RemoveStreamInstance 移除stream
func RemoveStreamInstance(name string) {
	if HasStreamInstance(name) {
		GetStreamWrapper(name).Stop()
		streampool.Remove(name)
		delete(keys, name)
	}
}

// HasStreamInstance stream句柄是否存在
func HasStreamInstance(name string) bool {
	return streampool != nil && streampool.Has(name)
}

// Start 开始实例
func Start(name string) error {
	if HasStreamInstance(name) {
		return GetStreamWrapper(name).Start()
	}
	return errors.New("non-valid instances found")
}

// Stop 停止实例
func Stop(name string) {
	if HasStreamInstance(name) {
		GetStreamWrapper(name).Stop()
	}
}

// Reload 重新加载
func Reload(name string) error {
	if HasStreamInstance(name) {
		return GetStreamWrapper(name).Reload()
	}
	return errors.New("non-valid instances found")
}

// Config 配置信息
func Config(name string) *videocollection.ConfigurationVO {
	if HasStreamInstance(name) {
		return GetStreamWrapper(name).Configs()
	}
	return nil
}

// ConfigList 配置信息
func ConfigList() []*videocollection.ConfigurationVO {
	data := make([]*videocollection.ConfigurationVO, 0)
	for key := range keys {
		configs := GetStreamWrapper(key).Configs()
		if configs != nil {
			data = append(data, configs)
		}
	}
	return data
}
