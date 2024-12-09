package goframework_videosys

import (
	videocollection "github.com/kordar/video-collection"
)

type CollectionWrapper struct {
	Collection    videocollection.Collection
	Configuration *videocollection.Configuration
	Retry         videocollection.Retry
}

func (w CollectionWrapper) Start() error {
	return w.Collection.Run(w.Configuration, w.Retry)
}

func (w CollectionWrapper) Stop() {
	w.Collection.Exit(w.Configuration)
}

func (w CollectionWrapper) Reload() error {
	return w.Collection.Reload(w.Configuration, w.Retry)
}

func (w CollectionWrapper) Configs() *videocollection.ConfigurationVO {
	vo := videocollection.ConfigurationVO{}
	vo.Load(w.Configuration)
	return &vo
}

// --------------------------------------

type VideoCollectionIns struct {
	name string
	ins  CollectionWrapper
}

func NewVideoCollectionIns(name string, collection videocollection.Collection, configuration *videocollection.Configuration, retry videocollection.Retry) *VideoCollectionIns {
	return &VideoCollectionIns{name: name, ins: CollectionWrapper{
		Collection:    collection,
		Configuration: configuration,
		Retry:         retry,
	}}
}

func (c VideoCollectionIns) GetName() string {
	return c.name
}

func (c VideoCollectionIns) GetInstance() interface{} {
	return c.ins
}

func (c VideoCollectionIns) Close() error {
	return nil
}
