package base


func UrlFor( url string ) (string){
	if AppConfig.String("","runmode") == "dev" {
		port := AppConfig.String("","httpport")
		return "http://localhost:" + port + "/" + url
	}
	return ""
}