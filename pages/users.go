package pages

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	_ "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"html/template"
)

func GetUserTable(ctx *context.Context) (userTable table.Table) {

	// config the table model.
	userTable = table.NewDefaultTable(table.Config{
		Driver:     db.DriverMysql,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := userTable.GetInfo()

	// set id sortable.
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "name", db.Varchar)

	// use FieldDisplay.
	info.AddField("Gender", "gender", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "men"
		}
		if model.Value == "1" {
			return "women"
		}
		return "unknown"
	})

	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("City", "city", db.Varchar)
	info.AddField("CreatedAt", "created_at", db.Timestamp)
	info.AddField("UpdatedAt", "updated_at", db.Timestamp)

	// set the title and description of table page.
	info.SetTable("users").SetTitle("Users").SetDescription("Users").
		SetAction(template.HTML(`<a href="http://google.com"><i class="fa fa-google"></i></a>`)) // custom operation button

	formList := userTable.GetForm()

	// set id editable is false.
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit()
	formList.AddField("Ip", "ip", db.Varchar, form.Text)
	formList.AddField("Name", "name", db.Varchar, form.Text)

	// use FieldOptions.
	formList.AddField("Gender", "gender", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{
				Text:  "male",
				Value: "0",
			}, {
				Text:  "female",
				Value: "1",
			},
		}).FieldDefault("0")
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("City", "city", db.Varchar, form.Text)

	// add a custom field and use FieldPostFilterFn to do more things.
	formList.AddField("Custom Field", "role", db.Varchar, form.Text).
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			fmt.Println("user custom field", value)
			return ""
		})

	formList.AddField("UpdatedAt", "updated_at", db.Timestamp, form.Default).FieldNotAllowAdd()
	formList.AddField("CreatedAt", "created_at", db.Timestamp, form.Default).FieldNotAllowAdd()

	// use SetTabGroups to group a form into tabs.
	formList.SetTabGroups(types.
		NewTabGroups("id", "ip", "name", "gender", "city").
		AddGroup("phone", "role", "created_at", "updated_at")).
		SetTabHeaders("profile1", "profile2")

	// set the title and description of form page.
	formList.SetTable("user").SetTitle("Users").SetDescription("Users")

	// use SetPostHook to add operation when form posted.
	/*formList.SetPostHook(func(values form2.Values) {
		fmt.Println("userTable.GetForm().PostHook", values)
	})*/

	return
}
