package controllers

import (
    "github.com/revel/revel"
    "myapp/app/models"
    "errors"
)

type Post struct {
    *revel.Controller
}

func (c Post) Index() revel.Result {
    posts := []models.Post{}

    result := DB.Order("id desc").Find(&posts);
    err := result.Error
    if err != nil {
        return c.RenderError(errors.New("Record Not Found"))
    }
    return c.Render(posts)
}
func (c Post) Create() revel.Result {
    post := models.Post{
        Body: c.Params.Form.Get("body"),
    }
    ret := DB.Create(&post)
    if ret.Error != nil {
        return c.RenderError(errors.New("Record Create failure." + ret.Error.Error()))
    }
    return c.Redirect("/posts")    
}
func (c Post) Delete() revel.Result {
    id := c.Params.Route.Get("id")
    posts := []models.Post{}
    ret := DB.Delete(&posts, id)
    if ret.Error != nil {
        return c.RenderError(errors.New("Record Delete failure." + ret.Error.Error()))
    }
    return c.Redirect("/posts")    
}

func (c Post) RedirectToPosts() revel.Result {
    return c.Redirect("/posts")    
}