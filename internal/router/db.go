package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"svc/proxy-service/internal/data"
	"svc/proxy-service/internal/utils"
)

func addDbRoutes(rg *gin.RouterGroup) {
	customer := rg.Group("/customer")

	customer.GET("/:id", func(c *gin.Context) {
		handleGetCustomer(c)
	})

	customer.GET("/", func(c *gin.Context) {
		handleGetCustomers(c)
	})

	site := rg.Group("/sites")
	site.GET("/:id", func(c *gin.Context) {
		handleGetSite(c)
	})
	site.GET("/", func(c *gin.Context) {
		handleGetSites(c)
	})

	siteConfig := rg.Group("/sites-config")

	siteConfig.GET("/:id", func(c *gin.Context) {
		handleGetSiteConfig(c)
	})

	siteConfig.GET("/", func(c *gin.Context) {
		handleGetSiteConfigs(c)
	})
}

func handleGetCustomer(c *gin.Context) {
	id := utils.ParseId(c, "id")
	customer := data.FindCustomer(&data.Customer{CustomerUniqueID: id})
	c.JSON(http.StatusOK, customer)
}

func handleGetCustomers(c *gin.Context) {
	customer := data.FindCustomer(nil)
	c.JSON(http.StatusOK, customer)
}

func handleGetSite(c *gin.Context) {
	id := utils.ParseId(c, "id")
	site := data.FindSite(&data.Site{SiteID: id})
	c.JSON(http.StatusOK, site)
}

func handleGetSites(c *gin.Context) {
	sites := data.FindSites(nil)
	c.JSON(http.StatusOK, sites)
}

func handleGetSiteConfig(c *gin.Context) {
	id := utils.ParseId(c, "id")
	site := data.FindSiteConfig(&data.SiteConfig{Id: id})
	c.JSON(http.StatusOK, site)
}

func handleGetSiteConfigs(c *gin.Context) {
	site := data.FindSiteConfigs(nil)
	c.JSON(http.StatusOK, site)
}
