package cache

import (
	"image-analysis-tools-aggregator/logging"
	"image-analysis-tools-aggregator/py_scripts"
	"image-analysis-tools-aggregator/utils"
	"io/ioutil"
	"os"
	"time"
)

//TODO
//	temporary solution - goroutines which managing tmp folder (cache) and
//	delete outdated images; in a good way it should be a separate
//	service to manage files

var DebugCacheBlacklist = []string{
	"264835d9-1879-11ec-bf39-18c04d99fd86.jpg",
	"c10d88b8-118f-11ec-a708-00059a3c7a00.jpg",
	"00a1f8e3-12f5-11ec-8364-18c04d99fd86.jpg",
}

var InstanceControlChannels = []chan bool{}

func InitCache(cacheConfig Config) {
	logging.DebugFormat("cache config %+v", cacheConfig)
	for i := 0; i < cacheConfig.CacheManagerInstances; i++ {
		go ManageCache(cacheConfig)
		time.Sleep(time.Second * 5)
	}
}

func ManageCache(config Config){
	quit := make(chan bool)
	InstanceControlChannels = append(InstanceControlChannels, quit)
	logging.Info("cache manager instance start")
	for {
		select {
		case <-quit:
			return
		default:  {
			items, _ := ioutil.ReadDir(py_scripts.ImageCacheFolderPath)
			for _, item := range items {
				if _, contains := utils.StringInArray(item.Name(), DebugCacheBlacklist);
					!contains &&
						time.Since(item.ModTime()).Milliseconds() >= config.CacheTTL.Milliseconds() {
					err := os.Remove(py_scripts.ImageCacheFolderPath + item.Name())
					if err != nil {
						logging.ErrorFormat("cannot delete obsolete file from cache due to %s:", err)
						continue
					}
					logging.InfoFormat("file %s successfully deleted", item.Name())
				}
			}
		}
		}
		logging.InfoFormat("cache manager instance sleep for %.2f seconds",
			config.CacheCleanInterval.Seconds() )
		time.Sleep(config.CacheCleanInterval)
	}
}