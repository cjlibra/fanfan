package main

import (
    "log"
	"os"
	"net"
	"encoding/hex"
	"time"
	"fmt"
	
)
const (
   url = "192.168.1.8:1200"
  // url = "127.0.0.1:1200"
   manhandle = "aaeecc07010b010000000014"
   autohandle = "aaeecc07010b020000000015"
 //  lasting  = time.Second*60
   lasting  = time.Hour * 24

)
func zhaofanaction(debugLog *log.Logger , manhandle_bin []byte){
    
	for {
	    control_success_flag  := 0
	    nowtime := time.Now()
		t1 := time.Date(nowtime.Year(),nowtime.Month(),nowtime.Day(),8,30,0,0,time.Local)
		t2 := time.Date(nowtime.Year(),nowtime.Month(),nowtime.Day(),13,0,0,0,time.Local)
		if nowtime.Before(t1) == true  || nowtime.Before(t2) != true {
			 
			 err  := control(debugLog   , manhandle_bin , "手动")
			 if err != nil {
				control_success_flag = 0
			 }else{
				control_success_flag  = 1
				 
			 }
		 
			 
					
		}
		if  control_success_flag == 0 {
		    time.Sleep(time.Second*60)	
		}else {
		    t3time := nowtime.Add(time.Hour*24)
			t3 :=  time.Date(t3time.Year(),t3time.Month(),t3time.Day(),0,0,0,0,time.Local)
		    debugLog.Println("zhaofanaction 休眠")
			<-time.After(t3.Sub(time.Now()))
		    debugLog.Println("zhaofanaction 启动")
		}
		
		
	}

}
func zhongfanaction(debugLog *log.Logger, autohandle_bin []byte){
    
	for {
	    control_success_flag  := 0
	    nowtime := time.Now()
		t1 := time.Date(nowtime.Year(),nowtime.Month(),nowtime.Day(),8,30,0,0,time.Local)
		t2 := time.Date(nowtime.Year(),nowtime.Month(),nowtime.Day(),13,0,0,0,time.Local)
		if nowtime.Before(t2) == true  && nowtime.Before(t1) != true{
			 
				 err  := control(debugLog   , autohandle_bin  , "自动")
				 if err != nil {
					control_success_flag = 0
				 }else{
				    control_success_flag = 1 
					 
				 }
			 
			 
					
		}
		if control_success_flag == 0 {
			time.Sleep(time.Second*60)
		} else {
		    t3time := nowtime.Add(time.Hour*24)
			t3 :=  time.Date(t3time.Year(),t3time.Month(),t3time.Day(),8,30,0,0,time.Local)
			debugLog.Println("zhongfanaction 休眠")
		    <-time.After(t3.Sub(time.Now()))
			debugLog.Println("zhongfanaction 启动")
		}
			
	}

}
func main(){
   
    fileName := time.Now().Format("2006-01-02#15-04-05")+".log"
    logFile,err  := os.Create(fileName)
    defer func(){
	    logFile.Write([]byte("close file  at time"+time.Now().String()))
		logFile.Close()
	}()
    if err != nil {
        log.Fatalln("open file error")
		return
    }
    debugLog := log.New(logFile,"[Info]",log.Llongfile | log.LstdFlags)
    debugLog.Println("Start to be Running")
	
	autohandle_bin,err := hex.DecodeString(autohandle)
	if err != nil {
	    debugLog.Println("error autohandle_bin")
		return
	}
	manhandle_bin,err := hex.DecodeString(manhandle)
	if err != nil {
	    debugLog.Println("error manhandle_bin")
		return
	}
	 
    nowtime := time.Now()
	 
	t1 := time.Date(nowtime.Year(),nowtime.Month(),nowtime.Day(),8,30,0,0,time.Local)
	t2 := time.Date(nowtime.Year(),nowtime.Month(),nowtime.Day(),13,0,0,0,time.Local)
	 
	if nowtime.Before(t1) == true {
	    control(debugLog, manhandle_bin , "手动1")
	}else{
		if nowtime.Before(t2) == true {
			control(debugLog, autohandle_bin ,  "自动1")
		}else{
		    control(debugLog, manhandle_bin , "手动1")
		}
	}
	
	go zhaofanaction(debugLog,manhandle_bin)
	go zhongfanaction(debugLog,autohandle_bin)
	b := make(chan int)
	fmt.Println("start running")
	fmt.Println("此软件界面不要关闭，否则食堂消费机不会自动设置模式")
	<-b
	
}

func control(debugLog *log.Logger , bindata []byte , lei string) error {
    service := url
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		debugLog.Println("net.ResolveTCPAddr error" , lei)
	    
		return err
	}
	
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
        debugLog.Println("无法连接消费机，网络出现问题", service, lei)
		 
        return err
    }	
	
	_, err = conn.Write(bindata)
	if err != nil {
	    debugLog.Println("无法传输数据出去，网络出现问题", service , lei)
		 
        return err
	}
	conn.Close()
	debugLog.Println("成功动作", service , lei)
    return nil
    

}