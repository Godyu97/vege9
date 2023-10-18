package main

//gorm gen demo
import (
	"github.com/Godyu97/vege9/vegedsn"
	"github.com/Godyu97/vege9/vegedsn/gen/self"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

var dns = vegedsn.NewDefaultMysqlDsn(
	vegedsn.DefaultParams,
	vegedsn.WithAuth("root", "passwd"),
	vegedsn.WithAddress("host:port"),
	vegedsn.WithDatabase("db"),
).String()

// go run . 生成dal模型代码
func main() {
	//Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:           "../dal/query",
		ModelPkgPath:      "../dal/model",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldWithTypeTag:  true,
		FieldWithIndexTag: true,
	})
	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf和thrift
	dataMap := map[string]func(gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)
	// Initialize a *gorm.DB instance
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)
	utOp := gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
		tag.Set("", "autoUpdateTime")
		return tag
	})
	ctOp := gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
		tag.Set("", "autoCreateTime")
		return tag
	})
	Member := g.GenerateModelAs("member_info", "Member", utOp, ctOp)
	Subscriber := g.GenerateModelAs("subscriber_info", "Subscriber", utOp, ctOp, gen.WithMethod(self.Subscriber{}))
	Records := g.GenerateModelAs("browse_records", "BrowseRecords", utOp, ctOp)
	Order := g.GenerateModelAs("order_records", "OrderRecords", utOp, ctOp)
	Collection := g.GenerateModelAs("collection_info", "CollectionInfo", utOp, ctOp)
	// Execute the generator
	g.ApplyBasic(
		Member,
		Subscriber,
		Records,
		Order,
		Collection,
	)
	g.Execute()
}
