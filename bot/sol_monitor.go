package bot

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"net/http"
	"pushbot/apis"
	"pushbot/config"
	"pushbot/util"
	"pushbot/util/push"
	"time"
)

var lastSolTimestamps = make(map[string]int64)

func MonitorSmartWalletSol() {
	solkey := config.HandleYaml().Keys.SolKey
	address := config.HandleYaml().Wallet.Sol
	pushType := config.HandleYaml().Push.Type
	interval := time.Duration(config.HandleYaml().Common.Interval) * time.Minute
	timeNow := time.Now().Unix()
	var responses []apis.SOL_TXResponse
	for _, addr := range address {
		url := apis.SOL_TRANSFER + "&address=" + addr
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("token", solkey)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			content := fmt.Sprintf("API请求错误，请稍后再试，错误：%v", err)
			switch pushType {
			case 1:
				message := push.WeChatRobotMsg{
					MsgType: "text",
					Text: &push.TextMsg{
						Content: content,
					},
				}
				err := push.PushWX(message)
				if err != nil {
					return
				}
				continue
			case 2:
				msg := tgbotapi.MessageConfig{Text: content}
				err := push.PushTelegram(msg)
				if err != nil {
					return
				}
				continue
			case 3:
				err := push.PushBark(content)
				if err != nil {
					return
				}
				continue
			}
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var response apis.SOL_TXResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		if response.Success != true {
			content := "API请求错误，请稍后再试"
			switch pushType {
			case 1:
				message := push.WeChatRobotMsg{
					MsgType: "text",
					Text: &push.TextMsg{
						Content: content,
					},
				}
				err := push.PushWX(message)
				if err != nil {
					return
				}
				continue
			case 2:
				msg := tgbotapi.MessageConfig{Text: content}
				err := push.PushTelegram(msg)
				if err != nil {
					return
				}
				continue
			case 3:
				err := push.PushBark(content)
				if err != nil {
					return
				}
				continue
			}
		} else {
			for _, res := range response.Data {
				lastSolTimestamp := lastSolTimestamps[addr] // 获取该地址的最新时间戳
				if lastSolTimestamp == res.BlockTime {
					fmt.Println("地址: " + addr + "钱包无新交易")
					continue
				} else {
					// 如果时间戳与上次相同，跳过推送
					fmt.Println(timeNow, res.BlockTime)
					if timeNow-res.BlockTime > int64(interval.Seconds()) {
						continue
					}
					content := fmt.Sprintf(
						"聪明钱链上异动！！！\n\n"+
							"区块高度: %d\n"+
							"买入时间: %s\n"+
							"钱包地址: %s\n"+
							"合约地址: %s\n"+
							//"代币名称: %s\n"+
							"区块网络: SOL\n"+
							"当前时间: %s",
						res.BlockId,
						util.FormatTime((timeNow-res.BlockTime)/60),
						res.ToAddress,
						res.TokenAddress,
						//res.TokenName,
						time.Now().Format("2006-01-02 15:04:05"))
					responses = append(responses, response)
					switch pushType {
					case 1:
						message := push.WeChatRobotMsg{
							MsgType: "text",
							Text: &push.TextMsg{
								Content: content,
							},
						}
						err := push.PushWX(message)
						if err != nil {
							return
						}
						continue
					case 2:
						msg := tgbotapi.MessageConfig{Text: content}
						err := push.PushTelegram(msg)
						if err != nil {
							return
						}
						continue
					case 3:
						err := push.PushBark(content)
						if err != nil {
							return
						}
						continue
					}
					lastSolTimestamps[addr] = res.BlockTime
				}
			}
		}
	}
}
