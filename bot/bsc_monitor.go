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
	"strconv"
	"strings"
	"time"
)

var lastBscTimestamps = make(map[string]int64)

func MonitorSmartWalletBsc() {
	bsckey := config.HandleYaml().Keys.BscKey
	address := config.HandleYaml().Wallet.Bsc
	pushType := config.HandleYaml().Push.Type
	interval := time.Duration(config.HandleYaml().Common.Interval) * time.Minute
	timeNow := time.Now().Unix()
	var responses []apis.BSC_TXResponse

	for _, addr := range address {
		url := apis.BSC_TRANSACTION + "&address=" + addr + "&apikey=" + bsckey
		req, _ := http.NewRequest("GET", url, nil)
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

		var response apis.BSC_TXResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}
		if response.Message != "OK" && response.Status != "1" {
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
			for _, res := range response.Result {
				lastBscTimestamp := lastBscTimestamps[addr] // 获取该地址的最新时间戳
				t, err := strconv.Atoi(res.TimeStamp)
				if err != nil {
					return
				}
				if lastBscTimestamp == int64(t) {
					fmt.Println("地址: " + addr + "钱包无新交易")
					continue
				} else {
					// 如果时间戳与上次相同，跳过推送
					if timeNow-int64(t) > int64(interval.Seconds()) {
						continue
					}
					// 排除稳定币
					if strings.Contains(res.TokenName, "Tether USD") || strings.Contains(res.TokenName, "USDC") || strings.Contains(res.TokenName, "USDT") {
						continue
					}
					content := fmt.Sprintf(
						"聪明钱链上异动！！！\n\n"+
							"区块高度: %s\n"+
							"买入时间: %s\n"+
							"钱包地址: %s\n"+
							"合约地址: %s\n"+
							"代币名称: %s\n"+
							"区块网络: BSC\n"+
							"当前时间: %s",
						res.BlockNumber,
						util.FormatTime((timeNow-int64(t))/60),
						res.To,
						res.ContractAddress,
						res.TokenName,
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
					case 2:
						msg := tgbotapi.MessageConfig{Text: content}
						err := push.PushTelegram(msg)
						if err != nil {
							return
						}
					case 3:
						err := push.PushBark(content)
						if err != nil {
							return
						}
					}
					lastBscTimestamps[addr] = int64(t)
				}
			}
		}
	}
}
