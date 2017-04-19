package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"bufio"
	"strings"
	"net/http"
)

//#Config Parameters
var theServer string
var thePort string
var the404Page string
var theSSH bool
var theGZip bool
var theSSHCertification string
var theSSHKey string

func init() {
  // open a file
  if file, err := os.Open("config.conf"); err == nil {

    // make sure it gets closed
    defer file.Close()

    // create a new scanner and read the file line by line
    scanner := bufio.NewScanner(file)
    var theLine string
    for scanner.Scan() {
    	theLine = scanner.Text()
    	if strings.Contains(theLine, "Server") {
    		theServer = strings.Replace(theLine, "Server = ", "", -1)
    	}

  		if strings.Contains(theLine, "Port") {
    		thePort = strings.Replace(theLine, "Port = ", "", -1)
    	}	

  		if strings.Contains(theLine, "404") {
    		the404Page = strings.Replace(theLine, "404_Page = ", "", -1)
    	}	

  		if strings.Contains(theLine, "SSH_Module") {
  			if !strings.Contains(theLine, "#"){
				theSSH = true;
  			}
    	}

    	if theSSH {
			if strings.Contains(theLine, "SSH_Certification") {
				theSSHCertification = theLine[strings.Index(theLine, "=")+2:]
			}

			if strings.Contains(theLine, "SSH_Key") {
				theSSHKey = theLine[strings.Index(theLine, "=")+2:]
			}
    	}

  		if strings.Contains(theLine, "GZip") {
  			if !strings.Contains(theLine, "#"){
				theGZip = true;
  			}
    	}
        
    }

    // check for errors
    if err = scanner.Err(); err != nil {
      log.Fatal(err)
    }

  } else {
    log.Fatal(err)
  }

	if theServer == "" {
		theServer = "localhost"
	}

	if thePort == "" {
		thePort = "8080"
	}

	if the404Page == "" {
		the404Page = "public/404.html"
	}


	theServer 			= strings.TrimSpace(theServer)
	thePort 			= strings.TrimSpace(thePort)	
	the404Page 			= strings.TrimSpace(the404Page)
	theSSHCertification = strings.TrimSpace(theSSHCertification)
	theSSHKey 			= strings.TrimSpace(theSSHKey)

}

func main() {
	preLog()

	http.Handle("/", new(TheHandler))

	if theSSH {
		http.ListenAndServeTLS(theServer+":"+thePort, theSSHCertification, theSSHKey, nil)
	} else {
		err := http.ListenAndServe(strings.TrimSpace(theServer)+":"+thePort, nil)
		log.Fatal(err)
	}
}


func preLog(){
	fmt.Println("\x1b[31;1mVX-Server - V 0.0.1\x1b[0m\n");
	fmt.Println("\x1b[36;1mTime:   ["+time.Now().Format(time.RFC850)+"]\x1b[0m\x1b[33;1m");
	fmt.Println("\x1b[32;1mServer: [localhost | 127.0.0.1 ON PORT "+thePort+"]\x1b[0m")
	if pwd, err := os.Getwd();err == nil {
		fmt.Println("\x1b[33;1mRoot:   ["+ pwd+"]\x1b[0m\n----------------------------------------------------")
	}

}

type TheHandler struct {
	http.Handler
}
 
//The default Handler 
func (this *TheHandler) ServeHTTP(responseWriter http.ResponseWriter, req *http.Request){

	fmt.Println("\x1b[36;1m["+time.Now().Format(time.RFC850)+"]\x1b[0m"+"\t\x1b[32;1m"+(strings.Split(req.RemoteAddr,":")[0])+"\x1b[0m\t\x1b[31;1m"+(theServer+":"+thePort+req.URL.Path)+"\x1b[0m")


	path := "public" + req.URL.Path
	file, err := os.Open(path)

	if err == nil {
		bufferedReader := bufio.NewReader(file)
	
		responseWriter.Header().Add("Content Type", getContentType(path))
		bufferedReader.WriteTo(responseWriter)
	} else {
		w.WriteHeader(404)
		file, err := os.Open(the404Page)
		if err == nil {
			bufferedReader := bufio.NewReader(file)
			w.Header().Add("Content Type", "text/html")
			bufferedReader.WriteTo(responseWriter)
		} else {
			responseWriter.WriteHeader(404)
        	responseWriter.Write([]byte("<b>404</b> - " + http.StatusText(404)))	
		}

 	}

  // Get required file Content Type
 func getContentType(string path) string{

 		if strings.HasSuffix(path, "html") {
			return "text/html"
 		} else if strings.HasSuffix(path, ".htm") {
 			return "text/html"
 		} else if strings.HasSuffix(path, ".aac") {
 			return "audio/aac"
 		} else if strings.HasSuffix(path, ".abw") {
 			return "application/x-abiword"
 		} else if strings.HasSuffix(path, ".arc") {
 			return "application/octet-stream"
 		} else if strings.HasSuffix(path, ".avi") {
 			return "video/x-msvideo"
 		} else if strings.HasSuffix(path, ".zaw") {
 			return "application/vnd.amazon.ebook"
 		} else if strings.HasSuffix(path, ".bin") {
 			return "application/octet-stream"
 		} else if strings.HasSuffix(path, ".bz") {
 			return "application/x-bzip"
 		} else if strings.HasSuffix(path, ".bz2") {
 			return "application/x-bzip2"
 		} else if strings.HasSuffix(path, ".csh") {
 			return "application/x-csh"
 		} else if strings.HasSuffix(path, ".css") {
 			return "text/css"
 		} else if strings.HasSuffix(path, ".csv") {
 			return "text/csv"
 		} else if strings.HasSuffix(path, ".doc") {
 			return "application/msword"
 		} else if strings.HasSuffix(path, ".epub") {
 			return "application/epub+zip"
 		} else if strings.HasSuffix(path, ".gif") {
 			return "image/gif"
 		} else if strings.HasSuffix(path, ".ico") {
 			return "image/x-icon"
 		} else if strings.HasSuffix(path, ".ics") {
 			return "text/calender"
 		} else if strings.HasSuffix(path, ".jar") {
 			return "application/java-archive"
 		} else if strings.HasSuffix(path, ".jpeg") {
 			return "image/jpeg"
 		} else if strings.HasSuffix(path, ".jpg") {
 			return "image/jpeg"
 		} else if strings.HasSuffix(path, ".js") {
 			return "application/javascript"
 		} else if strings.HasSuffix(path, ".json") {
 			return "application/json"
 		} else if strings.HasSuffix(path, ".mid") {
 			return "audio/midi"
 		} else if strings.HasSuffix(path, ".midi") {
 			return "audio/midi"
 		} else if strings.HasSuffix(path, ".mpeg") {
 			return "video/mpeg"
 		} else if strings.HasSuffix(path, ".mpkg") {
 			return "application/vnd.apple.installer+xml"
 		} else if strings.HasSuffix(path, ".odp") {
 			return "application/vnd.oasis.opendocument.presentation"
 		} else if strings.HasSuffix(path, ".ods") {
 			return "application/vnd.oasis.opendocument.spreadsheet"
 		} else if strings.HasSuffix(path, ".odt") {
 			return "application/vnd.oasis.opendocument.text"
 		} else if strings.HasSuffix(path, ".oga") {
 			return "audio/ogg"
 		} else if strings.HasSuffix(path, ".ogv") {
 			return "video/ogg"
 		} else if strings.HasSuffix(path, ".ogx") {
 			return "application/ogg"
 		} else if strings.HasSuffix(path, ".pdf") {
 			return "application/pdf"
 		} else if strings.HasSuffix(path, ".ppt") {
 			return "application/vnd.ms-powerpoint"
 		} else if strings.HasSuffix(path, ".rar") {
 			return "application/x-rar-compressed"
 		} else if strings.HasSuffix(path, ".rtf") {
 			return "application/rtf"
 		} else if strings.HasSuffix(path, ".sh") {
 			return "application/x-sh"
 		} else if strings.HasSuffix(path, ".svg") {
 			return "image/svg+xml"
 		} else if strings.HasSuffix(path, ".swf") {
 			return "application/x-shockwave-flash"
 		} else if strings.HasSuffix(path, ".tar") {
 			return "application/x-tar"
 		} else if strings.HasSuffix(path, ".tif") {
 			return "image/tiff"
 		} else if strings.HasSuffix(path, ".tiff") {
 			return "image/tiff"
 		} else if strings.HasSuffix(path, ".ttf") {
 			return "font/ttf"
 		} else if strings.HasSuffix(path, ".vsd") {
 			return "application/vnd.visio"
 		} else if strings.HasSuffix(path, ".wav") {
 			return "audio/x-wav"
 		} else if strings.HasSuffix(path, ".weba") {
 			return "audio/webm"
 		} else if strings.HasSuffix(path, ".webm") {
 			return "video/webm"
 		} else if strings.HasSuffix(path, ".webp") {
 			return "image/webp"
 		} else if strings.HasSuffix(path, ".woff") {
 			return "font/woff"
 		} else if strings.HasSuffix(path, ".woff2") {
 			return "font/woff2"
 		} else if strings.HasSuffix(path, ".xhtml") {
 			return "application/xhtml+xml"
 		} else if strings.HasSuffix(path, ".xls") {
 			return "application/vnd.ms-excel"
 		} else if strings.HasSuffix(path, ".xml") {
 			return "application/xml"
 		} else if strings.HasSuffix(path, ".xul") {
 			return "application/vnd.mozilla.xul"
 		} else if strings.HasSuffix(path, ".zip") {
 			return "application/zip"
 		} else if strings.HasSuffix(path, ".3gp") {
 			return "video/3gpp"
 		} else if strings.HasSuffix(path, ".3g2") {
 			return "video/3gpp2"
 		} else if strings.HasSuffix(path, ".7z") {
 			return "application/x-7z-compressed"
 		} else{
 			return "text/plain"
 		}
}