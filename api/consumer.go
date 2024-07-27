package api

import (
	"SecKill/api/RobbitMqService"
	"SecKill/api/dbService"
	"fmt"
	"log"
)

func ConsumeReq() {
	//rabbitMQSimple := NewRabbitMQSimple(TestQueueName)
	consumeSimple(RobbitMqService.Rabbitmq)
}

func consumeSimple(r *RobbitMqService.RabbitMQ) {
	// 无论生产还是消费，第一步都是尝试先申请队列
	_, e := r.Channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否会自动删除：当最后一个消费者断开后，是否将队列中的消息清除
		false,
		// 是否具有排他性，意思只有自己可见，其他人不可用
		false,
		// 是否阻塞
		false,
		// 其他额外属性
		nil,
	)
	if e != nil {
		fmt.Println(e)
	}

	// 3.消费者流控。防止爆库
	r.Channel.Qos(
		1,     // 当前消费之一次能接受的最大消费数量
		0,     // 服务器传递的最大容量
		false, // false 该配置只对当前这个 channel 有效。true 则对所有 channel 有效
	)

	msgs, e := r.Channel.Consume(
		r.QueueName,
		//用来区分多个消费者，消费者处理器名称
		"",

		// 是否自动应答通知已收到消息
		// 1.这里改成 false,通过手动应答的方式来自己处理消息通知
		false,
		//是否排他性,非唯一的消费者，其他消费者处理器也可以去竞争这个队列里面的消息任务
		false,
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false,
		//是否阻塞
		false,
		nil,
	)
	if e != nil {
		fmt.Println(e)
	}

	Done := make(chan bool)
	go func() {
		for d := range msgs {
			// 这里接收到消息，实现处理逻辑
			log.Printf("接收到的完整消息：%s", d)
			log.Printf("接收到消息：%s", d.Body)
			message := &RobbitMqService.Message{}
			message.StrToJson(d.Body)
			//// 执行下单事物
			//if err := service.SubNumberOne(message); err != nil {
			//	log.Printf("接收到消息,下单事物处理失败：%s", err.Error())
			//}
			//
			//// 2.手动应答，通知消息队列确认该消息
			//// true 表示确认所有队列中未确认的消息，false 表示确认队列中当前消息。
			//// 如果不调用则队列中该消息依然未被剔除，会被其他队列接受到。这里和 autoAck 配对使用
			//d.Ack(false)
			username := message.Username            //抢购成功的用户的用户名
			sellerName := message.Coupon.Username   //优惠券的商家名
			couponName := message.Coupon.CouponName //优惠券名

			var err error
			err = dbService.UserHasCoupon(username, message.Coupon) //用户优惠券数+1
			if err != nil {
				println("Error when inserting user's coupon. " + err.Error())
			}
			err = dbService.DecreaseOneCouponLeft(sellerName, couponName) //优惠券库存自减1
			if err != nil {
				println("Error when decreasing coupon left. " + err.Error())
			}
			d.Ack(false)
		}
	}()

	log.Printf("消费者已开启，等待消息产生。。。")
	<-Done

	r.Destory()
	log.Printf("消费者关闭。。。")
}
