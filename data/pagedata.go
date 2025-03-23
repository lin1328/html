package data

import (
	"html/template"
	"strings"
	"time"
)

type IconLink struct {
	Ico    string `json:"ico"`
	IcoURL string `json:"icourl"`
}

type SiteLink struct {
	Text string `json:"text"`
	JJ   string `json:"jj"`
	URL  string `json:"url"`
	Ico  string `json:"ico"`
}

type PageData struct {
	HTML       bool       `json:"html"`
	Name       string     `json:"name"`
	Title      string     `json:"title"`
	Background string     `json:"background"`
	Yan        string     `json:"yan"`
	ImgURL     string     `json:"imgurl"`
	Img        string     `json:"img"`
	Dome       string     `json:"dome"`
	ICPLink    string     `json:"icplink"`
	ICPText    string     `json:"icptext"`
	Qu         string     `json:"qu"`
	Arr2       []IconLink `json:"arr2"`
	Arrl       []SiteLink `json:"arrl"`
}

type TemplateData struct {
	Title      string
	ImgURL     string
	Img        string
	YanHTML    template.HTML
	Background string
	IconLinks  []IconLink
	SiteLinks  []SiteLink
	Dome       string
	Qu         string
	ICPLink    string
	ICPText    string
	Year       int
	Name       string
}

func GetPageData() PageData {
	return PageData{
		HTML:       true,
		Name:       "Linqi",
		Title:      "一个发现美的网站!",
		Background: "//imgapi.cn/api.php?zd=pc&fl=suiji&gs=images",
		Yan:        "我们会跑，无论往哪里都靠自己的腿去跑！\n男人总以为挣很多钱，就可以留住女人，其实事实就是这样。",
		ImgURL:     "#",
		Img:        "http://q.qlogo.cn/headimg_dl?dst_uin=3151178322&spec=640&img_type=png",
		Dome:       "12",
		ICPLink:    "https://beian.miit.gov.cn/",
		ICPText:    "浙ICP备2020******号",
		Qu:         "本站资源均来自互联网收集，仅供用于学习和交流，请勿用于商业用途。如有侵权、请联系网站管理并出示版权证明以便删除！",
		Arr2: []IconLink{
			{Ico: "iconfont icon-zuanshi1", IcoURL: ""},
			{Ico: "iconfont icon-cnblogs-grey", IcoURL: ""},
			{Ico: "iconfont icon-xinlangweibo", IcoURL: ""},
			{Ico: "iconfont icon-bilibili-fill", IcoURL: ""},
			{Ico: "iconfont icon-QQ", IcoURL: ""},
			{Ico: "iconfont icon-xinxi", IcoURL: "mailto:linqi0201@qq.com"},
		},
		Arrl: []SiteLink{
			{Text: "实验室", JJ: "理论输出一个视频", URL: "http://api.qemao.com/api/douyin/", Ico: "https://www.dmoe.cc/random.php"},
			{Text: "github", JJ: "全球最大的托管网站之一", URL: "https://github.com", Ico: "https://github.com/fluidicon.png"},
			{Text: "哔哩哔哩", JJ: "国内知名的视频弹幕网站", URL: "https://www.bilibili.com/", Ico: "https://static.hdslb.com/mobile/img/512.png"},
			{Text: "腾讯云", JJ: "提供安全稳定的全方位云服务和各行业解决方案", URL: "https://cloud.tencent.com", Ico: "https://cloudcache.tencent-cloud.com/qcloud/favicon.ico"},
			{Text: "网易云音乐", JJ: "一款专注于发现与分享的音乐产品", URL: "https://music.163.com", Ico: "https://s1.music.126.net/style/favicon.ico"},
		},
	}
}

func PrepareTemplateData(data PageData) TemplateData {
	return TemplateData{
		Title:      data.Title,
		ImgURL:     data.ImgURL,
		Img:        data.Img,
		YanHTML:    template.HTML(strings.Replace(data.Yan, "\n", "<br>", -1)),
		Background: data.Background,
		IconLinks:  data.Arr2,
		SiteLinks:  data.Arrl,
		Dome:       data.Dome,
		Qu:         data.Qu,
		ICPLink:    data.ICPLink,
		ICPText:    data.ICPText,
		Year:       time.Now().Year(),
		Name:       data.Name,
	}
}
