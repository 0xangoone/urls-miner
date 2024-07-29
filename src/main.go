package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func get_response(domain string) string{
	var server string
	if strings.HasPrefix(domain,"https") {
		server= fmt.Sprintf("%s:%d",domain,443);
	}else{
		server =  fmt.Sprintf("%s:%d",domain,80)
	}
	var res,err = http.Get(server);
	if err != nil{
		return " ";
	}
	response,err := ioutil.ReadAll(res.Body);
	if err != nil{
		return " ";
	}
	return string(response)
}
func get_urls(html string) []string { 
	var out []string ;
	var is_start bool = false;
	var temp bytes.Buffer
	for i := 0 ; i < len(html) ; i++{
		if html[i] == '"' || html[i] == '\''{
			if is_start{
				is_start = false
				var strTemp string = temp.String()
				if strings.HasPrefix(strTemp,"http"){
					out = append(out, strTemp)
				}
				temp.Reset()
				continue
			}
			is_start = true
			continue
		} 
		if is_start{
			temp.WriteString(string(html[i]))
		}
	}
	return out
}
func mining(out []string,url string){
	var r []string = get_urls(get_response(url))
	if len(r) == 0{
		return
	}
	for i:=0;i<len(r);i++{
		var exist bool = false;
		for j := 0 ;j<len(out);j++{
			if out[j] == r[i]{exist = true}
		}
		if strings.Contains(r[i],":443") || strings.Contains(r[i],":80"){
			continue
		}
		if strings.Contains(r[i],"https://enlighterjs.org"){continue} 
		if !exist{
			fmt.Println(r[i])
			out = append(out,r[i]);
			mining(out,r[i])
		}
	}
}
func main(){
	var out []string
	var entry string = "https://www.youtube.com/"
	mining(out,entry)
}