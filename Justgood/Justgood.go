package Justgood

import (
	"fmt"
	"./requests"
	"encoding/json"
)

var (
	DataObj *Data
	host = "https://api.imjustgood.com"
	print = fmt.Println
	session = requests.Requests()
)

type Imjustgood struct {
	apikey  string
	headers requests.Header
}

type Data struct {
	Message string 					`json:'message'`
	Creator string 					`json:'creator'`
	Result  map[string]interface{}  `json:'result'`
	Status  int 					`json:'status'`
}

type Callback struct {
	pin string
	token string
}

func Justgood(apikey string) *Imjustgood {
	res := new(Imjustgood)
	res.apikey = apikey
	res.headers = requests.Header{"User-Agent": "Justgood/5.0", "apikey": apikey}
	return res
}

func (c *Imjustgood) Get(andPoint string) []byte {
	resp,_ := session.Get(host+andPoint, c.headers)
	return resp.Content()
}

func (c *Imjustgood) Status() *Data {
	rest := fmt.Sprintf("/status?apikey=%s",c.apikey)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Youtube(query string) *Data {
	rest := fmt.Sprintf("/youtube=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Youtubedl(url string) *Data {
	rest := fmt.Sprintf("/youtubedl=%s",url)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Joox(query string) *Data {
	rest := fmt.Sprintf("/joox=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Chord(query string) *Data {
	rest := fmt.Sprintf("/chord=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Lyric(query string) *Data {
	rest := fmt.Sprintf("/lyric=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Smule(userId string) *Data {
	rest := fmt.Sprintf("/smule=%s",userId)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Smuledl(url string) *Data {
	rest := fmt.Sprintf("/smuledl=%s",url)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Tiktok(userId string) *Data {
	rest := fmt.Sprintf("/tiktok=%s",userId)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Tiktokdl(url string) *Data {
	rest := fmt.Sprintf("/tiktokdl=%s",url)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Instagram(userId string) *Data {
	rest := fmt.Sprintf("/instagram=%s",userId)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Instapost(url string) *Data {
	rest := fmt.Sprintf("/instapost=%s",url)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Instastory(userId string) *Data {
	rest := fmt.Sprintf("/instastory=%s",userId)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Twitter(userId string) *Data {
	rest := fmt.Sprintf("/twitter=%s",userId)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Twitterdl(url string) *Data {
	rest := fmt.Sprintf("/twitter/video?url=%s",url)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Facebookdl(url string) *Data {
	rest := fmt.Sprintf("/facebook/video?url=%s",url)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Pinterest(url string) *Data {
	rest := fmt.Sprintf("/pinterest?url=%s",url)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Github(userId string) *Data {
	rest := fmt.Sprintf("/github=%s",userId)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Playstore(query string) *Data {
	rest := fmt.Sprintf("/playstore=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Translate(lang, text string) *Data {
	rest := fmt.Sprintf("/translate/%s=%s",lang, text)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Image(query string) *Data {
	rest := fmt.Sprintf("/image=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Wallpaper(query string) *Data {
	rest := fmt.Sprintf("/wallpaper=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Porn(query string) *Data {
	rest := fmt.Sprintf("/porn=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Pornstar() *Data {
	rest := fmt.Sprintf("/pornstar")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Hentai() *Data {
	rest := fmt.Sprintf("/hentai")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Kamasutra() *Data {
	rest := fmt.Sprintf("/kamasutra")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Dick() *Data {
	rest := fmt.Sprintf("/dick")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Tits() *Data {
	rest := fmt.Sprintf("/tits")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Vagina() *Data {
	rest := fmt.Sprintf("/vagina")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Meme(text1, text2, imageUrl string) *Data {
	rest := fmt.Sprintf("/meme/%s/%s/url=%s",text1,text2,imageUrl)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Movie(query string) *Data {
	rest := fmt.Sprintf("/movie=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Movie_quotes() *Data {
	rest := fmt.Sprintf("/movie/quotes")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Cinema(city string) *Data {
	rest := fmt.Sprintf("/cinema=%s",city)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Tinyurl(url string) *Data {
	rest := fmt.Sprintf("/tinyurl=%s",url)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Bitly(url string) *Data {
	rest := fmt.Sprintf("/bitly=%s",url)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Kbbi(query string) *Data {
	rest := fmt.Sprintf("/kbbi=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Topnews() *Data {
	rest := fmt.Sprintf("/topnews")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Wikipedia(query string) *Data {
	rest := fmt.Sprintf("/wikipedia=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Urban(query string) *Data {
	rest := fmt.Sprintf("/urban=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Zodiac(sign string) *Data {
	rest := fmt.Sprintf("/zodiac=%s",sign)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Alquran() *Data {
	rest := fmt.Sprintf("/alquran=list")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) AlquranQS(query string) *Data {
	rest := fmt.Sprintf("/alquran=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Bible() *Data {
	rest := fmt.Sprintf("/bible")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Adzan(city string) *Data {
	rest := fmt.Sprintf("/adzan=%s",city)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Cuaca(city string) *Data {
	rest := fmt.Sprintf("/cuaca=%s",city)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Bmkg() *Data {
	rest := fmt.Sprintf("/bmkg")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Corona() *Data {
	rest := fmt.Sprintf("/corona")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Karir() *Data {
	rest := fmt.Sprintf("/karir")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Cellular(query string) *Data {
	rest := fmt.Sprintf("/cell=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Lahir(date string) *Data {
	rest := fmt.Sprintf("/lahir=%s",date)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Jadian(date string) *Data {
	rest := fmt.Sprintf("/jadian=%s",date)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Nama(name string) *Data {
	rest := fmt.Sprintf("/nama=%s",name)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Mimpi(query string) *Data {
	rest := fmt.Sprintf("/mimpi=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Acaratv() *Data {
	rest := fmt.Sprintf("/acaratv/now")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Acaratv_channel(channel string) *Data {
	rest := fmt.Sprintf("/acaratv=%s",channel)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Cctv_code() *Data {
	rest := fmt.Sprintf("/cctv/code")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) CctvSearch(code string) *Data {
	rest := fmt.Sprintf("/cctv/search/id=%s",code)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) MangaSearch(query string) *Data {
	rest := fmt.Sprintf("/manga/search=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) MangaChapter(chapterId string) *Data {
	rest := fmt.Sprintf("/manga/chapter=%s",chapterId)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Timeline(url string) *Data {
	rest := fmt.Sprintf("/timeline=%s",url)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Resi(query, code string) *Data {
	rest := fmt.Sprintf("/resi/%s=%s",query,code)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Screenshot(url string) *Data {
	rest := fmt.Sprintf("/screenshot?url=%s",url)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Gif(query string) *Data {
	rest := fmt.Sprintf("/gif=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Search(query string) *Data {
	rest := fmt.Sprintf("/search=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Calc(query string) *Data {
	rest := fmt.Sprintf("/calc=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Language() *Data {
	rest := fmt.Sprintf("/language/code")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Lineapp() *Data {
	rest := fmt.Sprintf("/line")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Lineqr(appName string, sysName string, cert string) *Data {
	c.headers = requests.Header{
		"User-Agent": "Justgood/5.0", 
		"apikey": c.apikey,
		"appName": appName, 
		"sysName": sysName, 
		"cert": cert,
	}
	rest := fmt.Sprintf("/lineqr")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) LineqrGetPin(path string) *Data {
	rest := fmt.Sprintf("/pin%s",path[30:])
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) LineqrGetToken(path string) *Data {
	rest := fmt.Sprintf("/token%s",path[32:])
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Check_ip(query string) *Data {
	rest := fmt.Sprintf("/ip=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) BinaryEncode(query string) *Data {
	rest := fmt.Sprintf("/binary/text?q=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) BinaryDecode(query string) *Data {
	rest := fmt.Sprintf("/binary/digit?q=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) B64Encode(query string) *Data {
	rest := fmt.Sprintf("/base64/digit?q=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) B64Decode(query string) *Data {
	rest := fmt.Sprintf("/base64/code?q=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Fancy(query string) *Data {
	rest := fmt.Sprintf("/fancy?text=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Simisimi(query string) *Data {
	rest := fmt.Sprintf("/simisimi?text=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Imagetext(query string) *Data {
	rest := fmt.Sprintf("/imgtext?text=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Stamp(url string, num int) *Data {
	rest := fmt.Sprintf("/stamp?url=%s&num=%v",url, num)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Stamplist() *Data {
	rest := fmt.Sprintf("/stamplist")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Fansign(query string, num int) *Data {
	rest := fmt.Sprintf("/fansign?text=%s&num=%v",query,num)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Ascii(query string) *Data {
	rest := fmt.Sprintf("/ascii=%s",query)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Customlink(label, url string) *Data {
	c.headers = requests.Header{
		"User-Agent": "Justgood/5.0", 
		"apikey": c.apikey,
		"label": label, 
		"url": url, 
	}
	rest := fmt.Sprintf("/custom/make")
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Watermark_image(imageUrl, iconUrl string) *Data {
	rest := fmt.Sprintf("/watermark/image?image=%s&icon=%s",imageUrl,iconUrl)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}

func (c *Imjustgood) Watermark_text(imageUrl, text string) *Data {
	rest := fmt.Sprintf("/watermark/text?image=%s&text=%s",imageUrl,text)
	resp := c.Get(rest)
	json.Unmarshal(resp, &DataObj)
	return DataObj
}