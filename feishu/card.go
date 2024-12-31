package feishu

import (
	"encoding/json"
	"fmt"
)

type MessageCard struct {
	MesgType string `json:"msg_type"`
	Card     card   `json:"card"`
}

type card struct {
	Config   config    `json:"config"`
	Header   header    `json:"header"`
	Elements []element `json:"elements"`
}

type config struct {
	WideScreenMode bool `json:"wide_screen_mode"`
}

type header struct {
	Template string `json:"template"`
	Title    title  `json:"title"`
}

type title struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

type element struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

/* func NewMessageCard() feishuMessage {

	switch disk.Status.IsPlugin {
	case 0:
		return onDiskInserted(disk)
	case 1:
		return onDiskPullout(disk)
	default:
		logg.Error("invalid value of plugin")
	}
	return nil
} */

// 插盘消息
func onDiskInserted() feishuMessager {
	text := fmt.Sprintf(`
	**磁盘信息** 
	**磁盘序列号: ** %s
	**系统标签: ** %s
	**插盘时间: ** %s
	**详细描述: **
		磁盘位置:
			区域: %s
			机房: %s
			机柜: %d
			服务器: %d
			盘位(slot): %s
		插盘状态: %s
		挂载状态: %s
		标签状态: %s
		硬盘寿命: %s
	---
	**集群插盘报错，数据无法进行实时拷贝**
	**请硬盘管理员关注**
	`)
	return &MessageCard{
		MesgType: "interactive",
		Card: card{
			Config: config{
				WideScreenMode: true,
			},
			Header: header{
				Template: "red",
				Title: title{
					Tag:     "plain_text",
					Content: "👋集群插盘异常实时报警",
				},
			},
			Elements: []element{
				{
					Tag:     "markdown",
					Content: text,
				},
				{
					Tag: "hr",
				},
				{
					Tag:     "markdown",
					Content: "💡消息来源:DiskMon 😘",
				},
			},
		},
	}
}

// 拔盘消息
/* func onDiskPullout() feishuMessage {
	var (
		pluginState       string
		transmissionState string
		slotNum           string
		diskSize          string
		diskAvail         string
		diskMoveData      string
	)

	pluginState = "已拔盘"

	if disk.DriveInfo.Slot < 0 {
		slotNum = "未定位"
	} else {
		slotNum = fmt.Sprintf("%d", disk.DriveInfo.Slot)
	}

	if disk.Status.IsFinish == 0 {
		transmissionState = "完成"
	} else {
		transmissionState = "未完成"
	}

	diskSize = fmt.Sprintf("%.1f", disk.Sysinfo.Size)
	diskAvail = fmt.Sprintf("%d", disk.Sysinfo.AvailLatest)
	diskMoveData = fmt.Sprintf("%d", disk.Sysinfo.AvailLatest-disk.Sysinfo.Avail)

	text := fmt.Sprintf(`
	**磁盘信息**
	**磁盘序列号: ** %s
	**系统标签: ** %s
	**插盘时间: ** %s
	**信息描述: **
		磁盘位置:
			区域: %s
			机房: %s
			机柜: %d
			服务器: %d
			盘位(slot): %s
		插盘状态: %s
		传输状态: %s
		空间描述:
			磁盘规格: %s
			可用空间: %s
			迁移数据: %s
	---
	**硬盘数据上传未能完成，上车可能出现空间不足**
	**请硬盘管理员关注**
	`, disk.Uevent.SCSI_IDENT_SERIAL,
		disk.Uevent.ID_FS_LABEL,
		disk.UeventTime,
		conf.Config.Location.Region.Name,
		conf.Config.Location.Region.DataCenter.Name,
		conf.Config.Location.Region.DataCenter.Room.Rack.Index,
		conf.Config.Location.Region.DataCenter.Room.Rack.Row.Index,
		slotNum,
		pluginState,
		transmissionState,
		diskSize,
		diskAvail,
		diskMoveData)
	return &MessageCard{
		MesgType: "interactive",
		Card: card{
			Config: config{
				WideScreenMode: true,
			},
			Header: header{
				Template: "red",
				Title: title{
					Tag:     "plain_text",
					Content: "👋集群拔盘异常实时报警",
				},
			},
			Elements: []element{
				{
					Tag:     "markdown",
					Content: text,
				},
				{
					Tag: "hr",
				},
				{
					Tag:     "markdown",
					Content: "💡消息来源:DiskMon 😘",
				},
			},
		},
	}
} */

func (mesg *MessageCard) payload() ([]byte, error) {
	return json.Marshal(mesg)
}

func (mesg *MessageCard) Send() error {
	return sendFeishuMessage(mesg)
}
