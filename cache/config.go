package cache

import (
	"image-analysis-tools-aggregator/utils"
	"strconv"
	"time"
)

type Config struct {
	CacheTTL              time.Duration
	CacheCleanInterval    time.Duration
	CacheManagerInstances int
	CacheSize             int
}

func GetCacheConfig() Config {
	cacheTTLStr := utils.GetEnv(utils.CacheTTLEnvKey, "5m")
	cacheCleanIntervalStr := utils.GetEnv(utils.CacheCleanIntervalEnvKey, "30s")
	cacheSizeStr := utils.GetEnv(utils.CacheSizeEnvKey, "100")
	cacheInstancesCountStr := utils.GetEnv(utils.CacheInstancesCount, "2")

	cacheTTL, _ := time.ParseDuration(cacheTTLStr)
	cacheCleanInterval, _ := time.ParseDuration(cacheCleanIntervalStr)
	cacheSize, _ := strconv.Atoi(cacheSizeStr)
	cacheInstancesCount, _ := strconv.Atoi(cacheInstancesCountStr)

	return Config{
		CacheTTL:           cacheTTL,
		CacheSize:          cacheSize,
		CacheCleanInterval: cacheCleanInterval,
		CacheManagerInstances: cacheInstancesCount,
	}
}

