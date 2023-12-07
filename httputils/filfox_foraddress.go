package httputils

import (
	"blog/models"
	"blog/services"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
	// 通过SetPrefix设置Logger结构体里的prefix属性
	log.SetPrefix("INFO:")
}

func StartAddress(timeTag int64, addrParam string) {
	// f17225euatkghukeyu6exm6j72fw5aneiiyapotbq  f1emjr5wnycomu7dgbt6aiaiir6rn53ggioiou56a
	//f01807413 f01822659 f01845913  f01874748  f01834260  f01899511
	//nodes := []string{"f01807413", "f01822659", "f01845913", "f01874748", "f01834260", "f01899511"}
	nodes := services.FindAllAddress(addrParam)
	pg := 5
	page := pageCount(len(nodes), pg)
	for p := 1; p <= page; p++ {
		pageAddress(nodes, p, pg, timeTag)
		time.Sleep(5 * time.Second)
	}

}

func pageAddress(nodes []models.FilAddresses, p int, pg int, timeTag int64) {
	s := services.LuckyBlock{}
	var newNodes []SpiderNode
	var updateNodes []SpiderNode
	//var m = make(map[string]int)
	t := pg * p
	if t > len(nodes) {
		t = len(nodes)
	}

	for i := (p - 1) * pg; i < t; i++ {
		node := SpiderNode{
			Node:     nodes[i].AccountId,
			PidNode:  nodes[i].Node,
			Category: "address",
			TimeTag:  timeTag,
		}

		var count int64
		count = nodes[i].TransferCount
		if count == 0 {
			s.CountByNode(node.Node, &count)
		}

		if count == 0 && len(newNodes) < 5 {
			newNodes = append(newNodes, node)
			if len(newNodes) >= 5 {
				break
			}
		} else {
			if len(newNodes) > 0 {
				continue
			}
			total := getTotalNum(node.Node) // 获取全部记录数量
			services.UpdateAddrRealCount(total, node.Node)
			log.Printf("node:%s的real_count:%d, transfer_count:%d\n", node.Node, total, count)
			page := pageCount(total-int(count), pageSize)
			if page == 0 {
				continue
			}
			node.Page = page
			updateNodes = append(updateNodes, node)
		}
	}

	if len(newNodes) > 0 { // 新节点首次抓数据
		log.Printf("新地址首次抓数据: %d\n", len(newNodes))
		newNodeSave(newNodes)
	} else if len(updateNodes) > 0 {
		// 已经完整抓取过，实时更新就行
		log.Println("更新节点: len(m)= ", len(updateNodes))
		updateNodeNew(updateNodes)
	} else {
		log.Println("no data need to update. next page: ", p+1)
	}
}
