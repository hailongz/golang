package top

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (*Top, error) {
	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = fmt.Sprintf("%s%s_", prefix, task.Name)

	v := Top{}

	err = func() error {

		rs, err := db.Query(conn, &v, prefix, " WHERE tid=?", task.Tid)

		if err != nil {
			return err
		}

		defer rs.Close()

		if rs.Next() {

			scaner := db.NewScaner(&v)

			err = scaner.Scan(rs)

			if err != nil {
				return err
			}
		} else {
			return micro.NewError(ERROR_NOT_FOUND, "未找到推荐项")
		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	_, err = db.Delete(conn, &v, prefix)

	if err != nil {
		return nil, err
	}

	{
		// 清除缓存

		redis, prefix, err := app.GetRedis("default")

		if err == nil {
			redis.Del(fmt.Sprintf("%s%s", prefix, task.Name)).Result()
			redis.Del(fmt.Sprintf("%s%s_rank", prefix, task.Name)).Result()
		}
	}

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
