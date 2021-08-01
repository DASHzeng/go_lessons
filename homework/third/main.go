package third

import (
	"context"
	"fmt"
	"go_lessons/homework/third/httpUtil"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

//
func main() {
	c1 := context.Background()
	c2, cancel := context.WithCancel(c1)
	group, errC := errgroup.WithContext(c2)
	srv := &http.Server{Addr: ":2222"}
	group.Go(func() error {
		return httpUtil.StartServer(srv)
	})
	group.Go(func() error {
		<-errC.Done()
		return srv.Shutdown(errC)
	})
	chanel_flag := make(chan os.Signal, 1)
	signal.Notify(chanel_flag)
	group.Go(func() error {
		for {
			select {
			case <-errC.Done():
				return errC.Err()
			case <-chanel_flag:
				cancel()
			}
		}
		return nil
	})
	if err := group.Wait(); err != nil {
		fmt.Println("unexpected: ", err)
	}
	fmt.Println("end")

}
