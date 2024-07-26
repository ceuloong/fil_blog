package apis

import (
	"blog/blockchain"
	"blog/filutils"
	"blog/httpcorrect"
	"blog/httputils"
	"blog/ticker"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetIndexCorrect(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Success!",
	})

	//timeTag := time.Now().Unix()
	//httpcorrect.Start(timeTag)

	httpcorrect.UpdateNodes("f01900855", 0)

	//httpcorrect.UpdateNodeChart("")
	//httpcorrect.UpdateNodeBurnAmount("")
	//httpcorrect.UpdateBurnNodeChart("")
}

// GetIndex
// @Success 200 {string} welcome
//
//	@Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Success!",
	})

	node := c.Query("node")
	println("node:" + node)

	timeTag := time.Now().Unix()
	httputils.Start(timeTag, node)

	//
	filutils.UpdateNodes(node, timeTag)
}

// UpdateFilNodes
// @Success 200 {string} welcome
//
//	@Router /update-nodes [get]
func UpdateFilNodes(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateNodes Success!",
	})
	timeTag := time.Now().Unix()
	httputils.UpdateNodes("f01900855", timeTag)
}

// UpdateAddresses
// @Success 200 {string} welcome
//
//	@Router /update-addresses [get]
func UpdateAddresses(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateAddresses Success!",
	})
	httputils.UpdateAddresses("")

}

// UpdateAddressesBalance
// @Success 200 {string} welcome
//
//	@Router /update-balance [get]
func UpdateAddressesBalance(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateAddresses Success!",
	})

	addr := c.Query("account_id")
	println("addr:" + addr)

	timeTag := time.Now().Unix()
	httputils.StartAddress(timeTag, addr)

	httputils.UpdateAddressesBalance(timeTag, addr)
}

func TronAddress(c *gin.Context) {
	addr := c.Query("addr")
	lStr := c.Query("level")

	if len(addr) == 0 && len(lStr) == 0 {
		c.JSON(200, gin.H{
			"message": "addr not null!",
		})
		return
	}
	println("addr:" + addr)

	if len(lStr) > 0 {
		level, _ := strconv.Atoi(lStr)
		blockchain.StartTron(level)
	} else {
		blockchain.GetHttpHtmlNew(addr)
	}

	c.JSON(200, gin.H{
		"message": "tron Success!",
	})
}

// HandUpdate
// @Success 200 {string} welcome
//
//	@Router /hand-update [get]
func HandUpdate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandUpdate Success!",
	})
	filutils.HandUpdate("")

}

func UpdateBlockStats(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateBlockStats Success!",
	})
	filutils.UpdateBlockStats("")

}

func Ticker(c *gin.Context) {
	ticker.SetTickerToRedis()
	c.JSON(200, gin.H{
		"message": "Ticker to redis Success!",
	})
}

func NodeDetails(c *gin.Context) {
	addr := c.Query("node")
	println("addr:" + addr)
	details := filutils.NodeDetails(addr)
	c.JSON(200, gin.H{
		"message": details,
	})
}

// SavePoolChart
// @Success 200 {string} welcome
//
//	@Router /save-chart [get]
func SavePoolChart(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "SavePoolChart Success!",
	})
	filutils.SavePoolChart()

}
