package controllers

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	a := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		"jd",
		18,
	}
	this.View(a)
}
