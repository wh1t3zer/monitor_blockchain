package binance

import (
	"context"
	"fmt"
	"log"
)

// 监控币安下架代币计划

func Monitor_delist() {
	c, err := NewBNClient().NewGetSymbolsDelistScheduleService().Do(context.Background())
	if err != nil {
		log.Fatal("创建监控下线币Client失败: ", err)
	}
	fmt.Println("Sdwa: ", &c)
	for _, clist := range c {
		fmt.Println(clist.DelistTime)
		fmt.Println(clist.Symbols)
	}
}
