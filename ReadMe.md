## MyGO Pagination

this unity for generator page number for golang

* if you using gorm & gin

```
type MyInput struct {
  Pagging
}

func Query(c *gin.Context) {  
  inputs := MyInput{}
  if err := c.Bing(); err != nil {
    c.JSON(400, err.Error())
  }
  q := db.Table("mytable").Where("id > 10")
  pg := inputs.Pagging(q)
  q.Offset(pg.Offset).Limit(pg.Limit)
}
```

* if you use other go web / orm

```
pagging := Pagging{
  Limit: 10
  Page: 1
}

// get total by your sql code
pg := pagging.PageInfoGenerator(total)

// You will get toal / total_page / offset / limit
```
