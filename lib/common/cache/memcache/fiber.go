package memcache

import "github.com/gofiber/fiber/v2"

const MEMORY_CACHE_LOCAL = "memoryCache"

func GetMemoryCache(c *fiber.Ctx) *MemoryCache {
	return c.Locals(MEMORY_CACHE_LOCAL).(*MemoryCache)
}

func MemoryCacheMiddleware(mem *MemoryCache) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Locals(MEMORY_CACHE_LOCAL, mem)
		return c.Next()
	}
}
