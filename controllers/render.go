package controllers

/**
* 公共渲染方法
**/

type H map[string]interface{}

type Meta struct {
	Name      string `json:"name"`
	HttpEquiv string `json:"http_equiv"`
	Content   string `json:"content"`
}

// Render
func RenderData(h H) H {

	// 设置网页 title
	if h["head_title"] == nil {
		h["head_title"] = "Hi OwO"
	}

	// 设置网页 meta
	h["head_meta"] = getMeta()

	return h
}

// 获取 meta
func getMeta() []interface{} {

	var metaData Meta

	metaData.Name = "keywords"
	metaData.HttpEquiv = ""
	metaData.Content = "hiOwO,hiowo,owo,bloghiowo,hiowoblog,owoblog,owo个人博客,vocaloid"

	var metaList = make([]interface{}, 0)

	metaList = append(metaList, metaData)

	return metaList

}
