package goSam

import (
	"fmt"
)

func (c *Client) StreamConnect(id int32, dest string) (err error) {
	var r *Reply

	r, err = c.sendCmd(fmt.Sprintf("STREAM CONNECT ID=%d DESTINATION=%s\n", id, dest))
	if err != nil {
		return err
	}

	if r.Topic != "STREAM" || r.Type != "STATUS" {
		return fmt.Errorf("Unknown Reply: %+v\n", r)
	}

	result := r.Pairs["RESULT"]
	if result != "OK" {
		err = ReplyError{result, r}
		return
	}

	fmt.Println("StreamConnect OK")
	return nil
}