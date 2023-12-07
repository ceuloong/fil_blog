package apis

import (
	"blog/blockchain"
	"blog/httpcorrect"
	"blog/httputils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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

	//httputils.UpdateAddresses("")
	//
	httputils.UpdateNodes(node, timeTag)
	//
	//httputils.UpdateAddressesBalance()
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
