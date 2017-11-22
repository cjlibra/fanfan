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
   //url = "192.168.1.8:1200"
   url = "127.0.0.1:1200"
   manhandle = "aaeecc07010b010000000014"
   autohandle = "aaeecc07010b020000000015"
 //  lasting  = time.Second*60
   lasting  = time.Hour * 24

)
func main(){
   
    fileName := "Info_First.log"
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
	tzhao := time.Date(nowtime.Year(),nowtime.Month(),nowtime.Day(),7,0,0,0,time.Local)
	tzhong := time.Date(nowtime.Year(),nowtime.Month(),nowtime.Day(),10,45,0,0,time.Local)
	t1 := time.Date(nowtime.Year(),nowtime.Month(),nowtime.Day(),8,30,0,0,time.Local)
	t2 := time.Date(nowtime.Year(),nowtime.Month(),nowtime.Day(),13,0,0,0,time.Local)
	//fmt.Println(tzhao,tzhong,t1,t2)
	if nowtime.Before(t1) == true {
	    control(debugLog, manhandle_bin)
	}else{
		if nowtime.Before(t2) == true {
			control(debugLog, autohandle_bin)
		}else{
		    control(debugLog, manhandle_bin)
		}
	}
	
	if nowtime.Before(tzhao) == true {
	    
		 go func() {
			<-time.After(tzhao.Sub(nowtime))
			for {
				control(debugLog, manhandle_bin)
				time.Sleep(time.Second*10)
				control(debugLog, manhandle_bin)
				time.Sleep(time.Second*10)
				control(debugLog, manhandle_bin)
				time.Sleep(time.Second*10)
				<-time.After(lasting -time.Second*10*3)
			}
		 
		 }()
		  
		go func() {
			 <-time.After(tzhong.Sub(nowtime))
			 for {
				 control(debugLog, autohandle_bin)
				 time.Sleep(time.Second*10)
				 control(debugLog, autohandle_bin)
				 time.Sleep(time.Second*10)
				 control(debugLog, autohandle_bin)
				 time.Sleep(time.Second*10)
				 <-time.After(lasting -time.Second*10*3)
			 
			 }
		
		
		
		}()
	
	}else{
	    if nowtime.Before(tzhong) == true {
		   
			go func() {
			     <-time.After(tzhong.Sub(nowtime))
			     for {
					 control(debugLog, autohandle_bin)
					 time.Sleep(time.Second*10)
					 control(debugLog, autohandle_bin)
					 time.Sleep(time.Second*10)
					 control(debugLog, autohandle_bin)
					 time.Sleep(time.Second*10)
					 <-time.After(lasting -time.Second*10*3)
				 
			     }
			
			
			
			}()
			
			go func() {
				<-time.After(tzhao.Add(time.Hour*24).Sub(nowtime))
				for {
					control(debugLog, manhandle_bin)
					time.Sleep(time.Second*10)
					control(debugLog, manhandle_bin)
					time.Sleep(time.Second*10)
					control(debugLog, manhandle_bin)
					time.Sleep(time.Second*10)
					<-time.After(lasting -time.Second*10*3)
				}
		 
		 }()
		
		}else{
		
			go func() {
					 <-time.After(tzhong.Add(time.Hour*24).Sub(nowtime))
					 for {
						 control(debugLog, autohandle_bin)
						 time.Sleep(time.Second*10)
						 control(debugLog, autohandle_bin)
						 time.Sleep(time.Second*10)
						 control(debugLog, autohandle_bin)
						 time.Sleep(time.Second*10)
						 <-time.After(lasting -time.Second*10*3)
					 
					 }
				
				
				
				}()
				
				go func() {
					<-time.After(tzhao.Add(time.Hour*24).Sub(nowtime))
				     
					for {
						control(debugLog, manhandle_bin)
						time.Sleep(time.Second*10)
		    			control(debugLog, manhandle_bin)
						time.Sleep(time.Second*10)
						control(debugLog, manhandle_bin)
						time.Sleep(time.Second*10)
		    			<-time.After(lasting -time.Second*10*3)
					}
			 
			 }()
		
		
		
		}
	
	
	}
	b := make(chan int)
	fmt.Println("start running")
	<-b
	
}

func control(debugLog *log.Logger , bindata []byte) error {
    service := url
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		debugLog.Println("net.ResolveTCPAddr error")
	    
		return err
	}
	
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
        debugLog.Println("无法连接消费机，网络出现问题", service)
		 
        return err
    }	
	
	_, err = conn.Write(bindata)
	conn.Close()
    return err
    

}