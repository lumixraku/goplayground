package mytimetask

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func TimeEntry() {
	//初始化断续器,间隔1s
	var ticker *time.Ticker = time.NewTicker(1 * time.Hour)
	waitc := make(chan float64)
	pt := fmt.Println

	go func() {
		for t := range ticker.C {
			daystr := ""
			now := time.Now()

			yesterdayshift := now.Add(-24 * time.Hour)
			daystr = yesterdayshift.Format("20060102")
			// t.Format("2006-01-02")
			// year := t.Year()
			// month := t.Month()
			// day := t.Day()

			fname := string(daystr) //+ string(month) + string(day)

			fmt.Println("Tick at", fname, t)
			isExists := detectFile(fname)
			if isExists == 0 {
				pt("not exist", string(fname))

				runCmd()

			} else {
				fmt.Println("exist")
			}
		}
	}()
	<-waitc
	// time.Sleep(time.Second * 5) //阻塞，则执行次数为sleep的休眠时间/ticker的时间
	// ticker.Stop()
	fmt.Println("Ticker stopped")
}

func detectFile(fname string) (rs int) {
	folder := "/home/luonan/my-map-reduce/spam/"

	// folder := "/Users/luonan/repos/go_path/src/goplayground/"
	path := folder + "tmpdata" + fname
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return 0
	}
	return 1

}

func runCmd() {
	cmdstr := "cd /home/luonan/my-map-reduce/spam/ && ./run.sh"
	// cmdstr = "ls"
	out, err := exec.Command("bash", "-c", cmdstr).Output()
	if err != nil {
		fmt.Println("failed.", err)
		return
	}
	fmt.Printf("%s", out)
}
