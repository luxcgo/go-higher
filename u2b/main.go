package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os/exec"
	"time"
)

const (
	Name = "ffmpeg"

	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36"
)

func main1() {
	// url, _ := url.Parse("https://www.youtube.com/watch?v=w_Ma8oQLmSM")
	// url, _ := url.Parse("https://www.huya.com/910429")
	cmd := exec.Command(
		"youtube-dl",
		"https://www.youtube.com/watch?v=54t_kLLcKVM",

		// "-nostats",
		// "-progress", "-",
		// "-y", "-re",
		// "-user_agent", userAgent,
		// "-referer", url.String(),
		// "-timeout", "60000000",
		// // "-i", "https://bd.flv.huya.com/src/957063158-957063158-4110554963816480768-1914249772-10057-A-0-1.flv?wsSecret=7fca42eec235db3cd2b6bad6afbf8cd7&wsTime=620a7e8e&u=0&seqid=16447685266223472&fm=RFdxOEJjSjNoNkRKdDZUWV8kMF8kMV8kMl8kMw%3D%3D&ctype=tars_mobile&txyp=o%3Acs6%3B&fs=bgct&&sphdcdn=al_7-tx_3-js_3-ws_7-bd_2-hw_2&sphdDC=huya&sphd=264_*-265_*&t=103",
		// "-c", "copy",
		// "-bsf:a", "aac_adtstoasc",
		// "-f", "flv",
		// "tt.flv",
	)
	log.Println(*cmd)
	cmd.Run()
	time.Sleep(100 * time.Second)
}

func main2() {
	parsed := "https://d1--cn-gotcha103.bilivideo.com/live-bvc/650890/live_500508001_7889020.m3u8?expires=1644949664&len=0&oi=1876795529&pt=h5&qn=10000&trid=10038ede16d5f1594d0aac543d346acdcfc2&sigparams=cdn,expires,len,oi,pt,qn,trid&cdn=cn-gotcha03&sign=04f7f74c99eeb47fb3a3719f91b999d6&sk=a78282cbfc3d64196de65af305b58a0f&p2p_type=0&src=5&sl=2&free_type=0&flowtype=1&machinezone=jd&pp=srt&slot=9&source=onetier&order=1&site=c25a13b0d8c579a3037467b19ea4833e"
	cmd := exec.Command(
		"ffmpeg",
		"-y", "-re",
		"-i", parsed,
		"-c", "copy",
		"-bsf:a", "aac_adtstoasc",
		"-f", "flv",
		"cs.flv",
	)
	cmd.Run()
}

func main() {
	log.Println("h")
	referer := "https://www.huya.com/391226"
	parsed := "https://d1--cn-gotcha103.bilivideo.com/live-bvc/650890/live_500508001_7889020.m3u8?expires=1644949664&len=0&oi=1876795529&pt=h5&qn=10000&trid=10038ede16d5f1594d0aac543d346acdcfc2&sigparams=cdn,expires,len,oi,pt,qn,trid&cdn=cn-gotcha03&sign=04f7f74c99eeb47fb3a3719f91b999d6&sk=a78282cbfc3d64196de65af305b58a0f&p2p_type=0&src=5&sl=2&free_type=0&flowtype=1&machinezone=jd&pp=srt&slot=9&source=onetier&order=1&site=c25a13b0d8c579a3037467b19ea4833e"
	cmd := exec.Command(
		"ffmpeg",
		"-nostats",
		"-progress", "-",
		"-y", "-re",
		"-user_agent", userAgent,
		"-referer", referer,
		// "-timeout", "60000000",
		// "-timeout", "100000000000",
		"-i", parsed,
		"-c", "copy",
		"-bsf:a", "aac_adtstoasc",
		"-f", "flv",
		"d.flv",
	)
	// cmdStdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")

	// go scheduler(cmdStdout)
	err := cmd.Wait()
	log.Printf("Command finished with error: %v", err)

	// cmd.Run()
}

func scanFFmpegStatus(cmdStdout io.ReadCloser) <-chan []byte {
	ch := make(chan []byte)
	br := bufio.NewScanner(cmdStdout)
	br.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if idx := bytes.Index(data, []byte("progress=continue\n")); idx >= 0 {
			return idx + 1, data[0:idx], nil
		}

		return 0, nil, nil
	})
	go func() {
		defer close(ch)
		for br.Scan() {
			ch <- br.Bytes()
		}
	}()
	return ch
}

func scheduler(cmdStdout io.ReadCloser) {
	statusCh := scanFFmpegStatus(cmdStdout)
	for {
		select {
		default:
			if _, ok := <-statusCh; !ok {
				return
			}
		}
	}
}
