package controllers

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.Redirect("chat/room.html", 302)
	return
}
